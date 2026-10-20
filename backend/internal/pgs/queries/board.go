package queries

import (
	"backend/models"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	log "github.com/sirupsen/logrus"
)

type BoardQuery struct {
	tableName string
	conn      *pgx.Conn
}

func NewBoardQuery(tableName string, conn *pgx.Conn) *BoardQuery {
	return &BoardQuery{tableName, conn}
}

func (b *BoardQuery) Create(ctx context.Context, title string) (int, error) {
	var id int

	q := `INSERT INTO ` + b.tableName + ` (title) VALUES ($1) RETURNING id;`
	err := b.conn.QueryRow(ctx, q, title).Scan(&id)
	if err != nil {
		log.Error("Error inserting into the table: ", err)
		return -1, err
	}

	return id, nil
}

func (b *BoardQuery) UpdateColumns(ctx context.Context, id int, columns string) error {
	q := `UPDATE ` + b.tableName + ` SET columns = {$1} WHERE id = $2`

	_, err := b.conn.Exec(ctx, q, columns, id)

	if err != nil {
		log.Error("Error updating the table: ", err)
	}

	return err
}

func (b *BoardQuery) GetBoard(ctx context.Context, id int) (*models.Board, error) {
	q := `SELECT id, title, columns FROM ` + b.tableName + `
WHERE id = $1;`

	rows, err := b.conn.Query(ctx, q, id)
	if err != nil {
		log.Error("Error getting the board: ", err)
		return new(models.Board), err
	}

	defer rows.Close()

	var boards []models.Board

	for rows.Next() {
		var board models.Board
		if err := rows.Scan(&board.ID, &board.Title, &board.Columns); err != nil {
			log.Error("Error getting the board: ", err)
			return &board, err
		}

		boards = append(boards, board)
	}

	if len(boards) == 0 {
		return new(models.Board), errors.New("no board found")
	}

	return &boards[0], nil
}

func (b *BoardQuery) GetBoards(ctx context.Context) (*[]models.Board, error) {
	q := `SELECT id, title, columns FROM ` + b.tableName

	var boards []models.Board

	rows, err := b.conn.Query(ctx, q)
	if err != nil {
		log.Error("Error getting the board: ", err)
		return &boards, err
	}

	defer rows.Close()

	for rows.Next() {
		var board models.Board
		if err := rows.Scan(&board.ID, &board.Title, &board.Columns); err != nil {
			log.Error("Error getting the board: ", err)
			return &boards, err
		}

		boards = append(boards, board)
	}

	return &boards, nil
}

func (b *BoardQuery) DeleteBoard(ctx context.Context, id int) error {
	q := `DELETE FROM ` + b.tableName + ` WHERE id = $1;`

	tag, err := b.conn.Exec(ctx, q, id)

	if err != nil {
		log.Error("Error deleting the board: ", err)
	}

	if tag.RowsAffected() == 0 {
		return errors.New("no board found")
	}

	return nil
}
