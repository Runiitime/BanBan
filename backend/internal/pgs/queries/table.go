package queries

import (
	"context"

	"github.com/jackc/pgx/v5"
	log "github.com/sirupsen/logrus"
)

func CreateTableBoards(conn *pgx.Conn, ctx context.Context) {
	q := `
	CREATE TABLE IF NOT EXISTS boards (
	    id serial PRIMARY KEY,
	    title VARCHAR(255) NOT NULL,
	    description VARCHAR(255),
	    columns VARCHAR(255)[]
	);
`
	if _, err := conn.Exec(ctx, q); err != nil {
		log.Error("Failed to create board table:", err)
	}
}

func CreateTableColumns(conn *pgx.Conn, ctx context.Context) {
	q := `
	CREATE TABLE IF NOT EXISTS columns (
	    id serial PRIMARY KEY,
	    board_id serial REFERENCES boards (id),
	    title VARCHAR(255) NOT NULL,
	    description VARCHAR(255),
	    cards VARCHAR(255)[]
	);
`
	if _, err := conn.Exec(ctx, q); err != nil {
		log.Error("Failed to create columns table:", err)
	}

}

func CreateTableCards(conn *pgx.Conn, ctx context.Context) {
	q := `
	CREATE TABLE IF NOT EXISTS cards (
	    id serial PRIMARY KEY,
	    column_id serial REFERENCES columns (id),
	    title VARCHAR(255) NOT NULL,
	    description VARCHAR(255),
	    status VARCHAR(255)
	);
`
	if _, err := conn.Exec(ctx, q); err != nil {
		log.Error("Failed to create columns cards:", err)
	}
}
