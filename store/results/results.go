package results

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/lotusirous/codescan/core"
)

func New(conn *sql.DB) core.ScanResultStore {
	return &resultStore{db: conn}
}

type resultStore struct {
	db *sql.DB
}

func (s *resultStore) Count(ctx context.Context) (int64, error) {
	query := `SELECT COUNT(*) FROM scan_results`
	var cnt int64
	err := s.db.QueryRowContext(ctx, query).Scan(&cnt)
	if err != nil {
		return 0, err
	}
	return cnt, nil
}

// Create return the status of adding a scan result to datastore.
func (s *resultStore) Create(ctx context.Context, result *core.ScanResult) error {
	query, args, err := squirrel.Insert("scan_results").SetMap(squirrel.Eq{
		"scan_id": result.ScanID,
		"repo_id": result.RepoID,
		"commit":  result.Commit,
		"created": result.Created,
		"updated": result.Updated,
	}).PlaceholderFormat(squirrel.Question).ToSql()
	if err != nil {
		return fmt.Errorf("%w - sql: %s", err, query)
	}

	r, err := s.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	id, err := r.LastInsertId()
	if err != nil {
		return err
	}
	result.ID = id
	return err

}

// Find returns a scan result from datastore.
func (s *resultStore) Find(ctx context.Context, id int64) (*core.ScanResult, error) {
	query := `SELECT scan_result_id, scan_id, repo_id, commit, created, updated
	FROM scan_results
	WHERE scan_result_id = ?
	`
	r := s.db.QueryRowContext(ctx, query, id)
	return scanRow(r)
}

// List returns a list of scan result from datastore.
func (s *resultStore) List(ctx context.Context) ([]*core.ScanResult, error) {
	query, args, err := squirrel.
		Select("scan_result_id, scan_id, repo_id, created, updated").
		From("scan_results").
		PlaceholderFormat(squirrel.Question).
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return scanRows(rows)
}