// Command server serves sawirstudio.com: the site (built into the binary by `make dist`)
// and the contact form's API, behind Cloudflare.
//
//	sawirstudio-server            serve (applies new migrations first, with DATABASE_URL)
//	sawirstudio-server migrate    only apply migrations
//
// Settings come from the environment: DATABASE_URL (without it the contact API answers
// 503), ADDR (default 127.0.0.1:8098), and for an email about each new inquiry
// CLOUDFLARE_ACCOUNT_ID, CLOUDFLARE_EMAIL_TOKEN, NOTIFY_FROM and NOTIFY_TO (see notify). In development, serve the repo instead:
//
//	go run ./cmd/server -root .
package main

import (
	"context"
	"errors"
	"flag"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
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
		var notifier server.Notifier
		if m, missing := notify.FromEnv(); m != nil {
			notifier = m
		} else {
			log.Warn("inquiry emails are off", "missing", missing)
		}
		cfg.Inquiries = server.NewInquiries(pool, notifier)
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

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
