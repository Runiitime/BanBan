package queries

import (
	"backend/models"
	"context"

	"github.com/jackc/pgx/v5"
	log "github.com/sirupsen/logrus"
)

type ColumnQuery struct {
	tableName string
	conn      *pgx.Conn
}

func NewColumnQuery(tableName string, conn *pgx.Conn) *ColumnQuery {
	return &ColumnQuery{tableName, conn}
}

func (b *ColumnQuery) Create(ctx context.Context, title string, boardId int) (int, error) {
	var id int

	q := `INSERT INTO ` + b.tableName + ` (title, board_id) VALUES ($1, $2)`
	err := b.conn.QueryRow(ctx, q, title, boardId).Scan(&id)
	if err != nil {
		log.Error("Error inserting into the table: ", err)
		return -1, err
	}

	return id, nil
}

func (b *ColumnQuery) UpdateCards(ctx context.Context, id int, boardId int, cards string) error {
	q := `UPDATE ` + b.tableName + ` SET cards = {$1}, board_id=$2 WHERE id = $3`

	_, err := b.conn.Exec(ctx, q, cards, boardId, id)

	if err != nil {
		log.Error("Error updating the table: ", err)
	}

	return err
}

func (b *ColumnQuery) GetColumn(ctx context.Context, id int) (*models.Column, error) {
	q := `SELECT * FROM ` + b.tableName + `WHERE id = $1`

	rows, err := b.conn.Query(ctx, q, id)
	if err != nil {
		log.Error("Error getting the column: ", err)
		return new(models.Column), err
	}

	defer rows.Close()

	var column models.Column

	for rows.Next() {
		if err := rows.Scan(&column.ID, &column.Title, &column.Cards); err != nil {
			log.Error("Error getting the column: ", err)
			return &column, err
		}
	}

	return &column, nil
}

func (b *ColumnQuery) GetAllColumns(ctx context.Context, id int) (*[]models.Column, error) {
	q := `SELECT * FROM ` + b.tableName + `WHERE board_id = $1`

	rows, err := b.conn.Query(ctx, q, id)
	if err != nil {
		log.Error("Error getting the all columns: ", err)
		return new([]models.Column), err
	}

	defer rows.Close()

	var columns []models.Column

	for rows.Next() {
		var column models.Column
		if err := rows.Scan(&column.ID, &column.Title, &column.Cards); err != nil {
			log.Error("Error getting columns: ", err)
			return &columns, err
		}

		columns = append(columns, column)
	}

	return &columns, nil
}
