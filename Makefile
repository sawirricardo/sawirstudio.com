# Builds and deploys sawirstudio.com. See README.
#
#   make site     after editing content/: pages, translations, Markdown copies
#   make deploy   build the server (with the site inside) and put it on ricardosawir1

DEPLOY_HOST ?= ricardosawir1
BIN = bin/sawirstudio-server
DATABASE_URL ?= postgres://postgres@localhost:5432/sawirstudio?sslmode=disable

.PHONY: generate pages translations markdown site serve migrate test check-built dist deploy purge

generate:            ## templ and sqlc code
	go tool templ generate
	sqlc generate

pages: generate      ## the English pages, from content/
	go run ./cmd/build

translations:        ## rebuild /<lang>/ from the cached translations (no claude calls)
	uv run scripts/i18n.py --no-translate

markdown:            ## index.md copies and llms.txt
	uv run scripts/markdown_pages.py

site: pages translations markdown

serve:               ## the repo's site and the contact API on http://127.0.0.1:8098
	DATABASE_URL="$(DATABASE_URL)" go run ./cmd/server -root .

migrate:             ## apply database migrations (dbmate works too)
	DATABASE_URL="$(DATABASE_URL)" go run ./cmd/server migrate

test:                ## Go tests (the API tests need TEST_DATABASE_URL)
	go vet ./...
	go test ./...

check-built:         ## fail if content/ has changes the built pages don't have yet
	@go run ./cmd/build -check || (echo "Pages are out of date: run make site" >&2; exit 1)

# The site files (everything but what .assetsignore and dotfiles name) go into .dist/site,
# which embed.go builds into the binary.
dist: check-built
	rm -rf .dist/site && mkdir -p .dist/site && touch .dist/site/.gitkeep
	rsync -a --exclude-from=.assetsignore --exclude='.*' ./ .dist/site/
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o $(BIN) ./cmd/server
	@ls -lh $(BIN) | awk '{print "built", $$NF, $$5}'

deploy: dist         ## copy the binary to the server, restart, check, purge Cloudflare
	scp -q $(BIN) $(DEPLOY_HOST):/usr/local/bin/sawirstudio-server.new
	scp -q deploy/sawirstudio.service $(DEPLOY_HOST):/etc/systemd/system/sawirstudio.service
	ssh $(DEPLOY_HOST) 'set -e; \
	  chmod 755 /usr/local/bin/sawirstudio-server.new; \
	  mv /usr/local/bin/sawirstudio-server.new /usr/local/bin/sawirstudio-server; \
	  systemctl daemon-reload; systemctl enable --quiet sawirstudio; systemctl restart sawirstudio; \
	  for i in 1 2 3 4 5; do curl -fsS http://127.0.0.1:8098/healthz && exit 0; sleep 1; done; \
	  journalctl -u sawirstudio -n 20 --no-pager; exit 1'
	@$(MAKE) --no-print-directory purge

# Empties Cloudflare's cache for sawirstudio.com (not the zone's other sites), so the new pages show right away. Needs
# CLOUDFLARE_ZONE_ID and CLOUDFLARE_API_TOKEN (a token with Cache Purge permission).
purge:
	@if [ -z "$$CLOUDFLARE_ZONE_ID" ] || [ -z "$$CLOUDFLARE_API_TOKEN" ]; then \
	  echo "Skipped the Cloudflare purge (set CLOUDFLARE_ZONE_ID and CLOUDFLARE_API_TOKEN)"; \
	else \
	  curl -fsS -X POST "https://api.cloudflare.com/client/v4/zones/$$CLOUDFLARE_ZONE_ID/purge_cache" \
	    -H "Authorization: Bearer $$CLOUDFLARE_API_TOKEN" -H "Content-Type: application/json" \
	    --data '{"hosts":["sawirstudio.com","www.sawirstudio.com"]}' >/dev/null && echo "Purged Cloudflare's cache"; \
	fi
