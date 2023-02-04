package postgres

import (
	"context"
	"io"
	"os"

	"github.com/jackc/pgx/v5"
)

const (
	tableSchemas           = "migrations/up.sql"
	workspacesTable        = "workspaces"
	usersTable             = "users"
	workspaceRelationTable = "users_workspaces"
	boardTable             = "boards"
	listTable              = "lists"
	cardTable              = "cards"
)

func createTables(db *pgx.Conn, ctx context.Context) error {
	f, err := os.OpenFile(tableSchemas, os.O_RDONLY, 0o755)
	if err != nil {
		return err
	}
	defer f.Close()
	tables, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	_, err = db.Exec(ctx, string(tables))
	if err != nil {
		return err
	}

	return nil
}
