// Command server serves sawirstudio.com: the site (built into the binary by `make dist`),
// the contact form's API and the blog newsletter's signups, behind Cloudflare.
//
//	sawirstudio-server                              serve (applies new migrations first, with DATABASE_URL)
//	sawirstudio-server migrate                      only apply migrations
//	sawirstudio-server subscribers                  count the newsletter's subscribers
//	sawirstudio-server subscribers import FILE.csv  add confirmed subscribers (ConvertKit's export)
//	sawirstudio-server newsletter send [-dry-run] [-test ADDRESS] /posts/SLUG/
//	                                                email a post to the subscribers who haven't had it
//
// Settings come from the environment: DATABASE_URL (without it the contact and signup APIs
// answer 503), ADDR (default 127.0.0.1:8098), for email NOTIFY_FROM and NOTIFY_TO with either
// SES_REGION and AWS credentials (Amazon SES) or CLOUDFLARE_ACCOUNT_ID and
// CLOUDFLARE_EMAIL_TOKEN (see notify.FromEnv), NEWSLETTER_FROM (the
// newsletter's sender, default newsletter@ NOTIFY_FROM's domain) and SITE_URL (the start of
// the links in the newsletter's emails, default https://sawirstudio.com). In development,
// serve the repo instead:
//
//	go run ./cmd/server -root .
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	sawirstudio "github.com/sawirricardo/sawirstudio.com"
	"github.com/sawirricardo/sawirstudio.com/internal/migrate"
	"github.com/sawirricardo/sawirstudio.com/internal/notify"
	"github.com/sawirricardo/sawirstudio.com/internal/server"
)

func main() {
	addr := flag.String("addr", envOr("ADDR", "127.0.0.1:8098"), "address to listen on")
	root := flag.String("root", "", "serve this folder instead of the site built into the binary")
	trust := flag.Bool("trust-proxy", os.Getenv("TRUST_PROXY") == "1", "trust Cloudflare's client IP and scheme headers")
	flag.Parse()
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	dbURL := os.Getenv("DATABASE_URL")

	if flag.Arg(0) == "migrate" {
		if dbURL == "" {
			log.Info("DATABASE_URL is not set; no migrations to apply")
			return
		}
		if err := migrate.Up(dbURL, sawirstudio.Migrations); err != nil {
			log.Error("migrate", "err", err)
			os.Exit(1)
		}
		log.Info("migrations applied")
		return
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	var site fs.FS
	if *root != "" {
		site = os.DirFS(*root)
	} else {
		sub, err := fs.Sub(sawirstudio.Site, ".dist/site")
		if err != nil {
			log.Error("embedded site", "err", err)
			os.Exit(1)
		}
		site = sub
	}

	if cmd := flag.Arg(0); cmd == "subscribers" || cmd == "newsletter" {
		if err := listCommand(ctx, log, dbURL, site, flag.Args()); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		return
	}

	cfg := server.Config{Site: site, TrustProxy: *trust, Log: log}
	if dbURL != "" {
		// A database that's down mustn't take the site down: log it, serve, and let the
		// contact API fail (and /healthz report it) until it's back.
		if err := migrate.Up(dbURL, sawirstudio.Migrations); err != nil {
			log.Error("migrate (serving anyway)", "err", err)
		}
		pool, err := pgxpool.New(ctx, dbURL)
		if err != nil {
			log.Error("database", "err", err)
			os.Exit(1)
		}
		defer pool.Close()
		notifier := mailer(log)
		cfg.Inquiries = server.NewInquiries(pool, notifier)
		cfg.Subscribers = newsletter(pool, notifier)
		defer cfg.Inquiries.Wait() // let emails still being sent finish
		go cfg.Inquiries.SendPending(ctx, log)
	} else {
		log.Warn("DATABASE_URL is not set; the contact form API is off")
	}

	h, err := server.New(cfg)
	if err != nil {
		log.Error("setup", "err", err)
		os.Exit(1)
	}
	srv := &http.Server{
		Addr:              *addr,
		Handler:           h,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       120 * time.Second,
	}
	go func() {
		<-ctx.Done()
		shutdown, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		srv.Shutdown(shutdown)
	}()
	log.Info("listening", "addr", *addr, "root", *root)
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Error("serve", "err", err)
		os.Exit(1)
	}
}

// mailer is the Cloudflare mailer from the environment, or nil (emails off).
func mailer(log *slog.Logger) server.Notifier {
	m, missing := notify.FromEnv(context.Background())
	if m == nil {
		log.Warn("emails are off (inquiry emails, newsletter signups)", "missing", missing)
		return nil
	}
	return m
}

func newsletter(pool *pgxpool.Pool, n server.Notifier) *server.Subscribers {
	from := os.Getenv("NEWSLETTER_FROM")
	if from == "" {
		if _, domain, ok := strings.Cut(os.Getenv("NOTIFY_FROM"), "@"); ok {
			from = "newsletter@" + domain
		}
	}
	return server.NewSubscribers(pool, n, from, envOr("SITE_URL", "https://sawirstudio.com"))
}

// listCommand runs `subscribers` and `newsletter send`.
func listCommand(ctx context.Context, log *slog.Logger, dbURL string, site fs.FS, args []string) error {
	usage := errors.New("usage: subscribers | subscribers import FILE.csv | newsletter send [-dry-run] [-test ADDRESS] /posts/SLUG/")
	if dbURL == "" {
		return errors.New("DATABASE_URL is not set")
	}
	if err := migrate.Up(dbURL, sawirstudio.Migrations); err != nil {
		return fmt.Errorf("migrate: %w", err)
	}
	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		return err
	}
	defer pool.Close()

	switch {
	case args[0] == "subscribers" && len(args) == 1:
		st, err := newsletter(pool, nil).Stats(ctx)
		if err != nil {
			return err
		}
		fmt.Printf("%d confirmed, %d waiting to confirm, %d unsubscribed\n", st.Confirmed, st.Pending, st.Unsubscribed)
		return nil

	case args[0] == "subscribers" && len(args) == 3 && args[1] == "import":
		f, err := os.Open(args[2])
		if err != nil {
			return err
		}
		defer f.Close()
		res, err := newsletter(pool, nil).Import(ctx, f)
		fmt.Printf("%d added, %d already on the list, %d skipped (left, or not an email address)\n", res.Added, res.Existing, res.Skipped)
		return err

	case args[0] == "newsletter" && len(args) > 1 && args[1] == "send":
		fl := flag.NewFlagSet("newsletter send", flag.ContinueOnError)
		dry := fl.Bool("dry-run", false, "only count who would get it")
		test := fl.String("test", "", "send it only to this address")
		delay := fl.Duration("delay", 500*time.Millisecond, "wait between emails")
		if err := fl.Parse(args[2:]); err != nil || fl.NArg() != 1 {
			return usage
		}
		s := newsletter(pool, mailer(log))
		n, err := s.SendIssue(ctx, site, fl.Arg(0), server.SendOptions{DryRun: *dry, Test: *test, Delay: *delay, Log: log})
		switch {
		case *dry:
			fmt.Printf("would send %s to %d subscribers\n", fl.Arg(0), n)
		case err == nil:
			fmt.Printf("sent %s to %d subscribers\n", fl.Arg(0), n)
		}
		return err
	}
	return usage
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
