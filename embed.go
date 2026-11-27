// Package sawirstudio holds the files built into the server binary: the database
// migrations, and the built site (staged into .dist/site by `make dist`).
package sawirstudio

import "embed"

//go:embed db/migrations/*.sql
var Migrations embed.FS

// Site is the built site, without the source and scripts (see .assetsignore).
// A plain `go build` without `make dist` embeds only .gitkeep.
//
//go:embed all:.dist/site
var Site embed.FS
