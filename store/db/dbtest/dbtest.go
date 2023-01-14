package dbtest

import (
	"context"
	"database/sql"
	_ "embed"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lotusirous/codescan/store/db"
)

// Connect opens a new test database connection.
func Connect() (*sql.DB, error) {
	var datasource string
	datasource = os.Getenv("DATABASE_DATASOURCE")
	return db.Connect(datasource, 0)

}

var noContext = context.TODO()

// Reset resets the database state.
func Reset(d *sql.DB) error {
	tx, err := d.Begin()
	if err != nil {
		return err
	}
	tx.Exec("DELETE FROM repos")
	tx.Exec("DELETE FROM scans")
	return tx.Commit()
}

// Disconnect closes the database connection.
func Disconnect(d *sql.DB) error {
	return d.Close()
}
