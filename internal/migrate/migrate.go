// Package migrate applies the embedded dbmate migrations (db/migrations) at startup, so
// the server binary is all a deploy needs. dbmate itself works on the same database too.
package migrate

import (
	"io/fs"
	"net/url"

	"github.com/amacneil/dbmate/v2/pkg/dbmate"
	_ "github.com/amacneil/dbmate/v2/pkg/driver/postgres"
)

// Up applies every migration in migrations (a filesystem with db/migrations/*.sql) that the
// database at databaseURL hasn't had yet.
func Up(databaseURL string, migrations fs.FS) error {
	u, err := url.Parse(databaseURL)
	if err != nil {
		return err
	}
	db := dbmate.New(u)
	db.FS = migrations
	db.MigrationsDir = []string{"db/migrations"}
	db.AutoDumpSchema = false
	db.Verbose = false
	return db.Migrate()
}
