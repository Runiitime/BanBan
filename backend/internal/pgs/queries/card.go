package queries

import (
	"backend/models"
	"context"

	"github.com/jackc/pgx/v5"
	log "github.com/sirupsen/logrus"
)

type CardQuery struct {
	tableName string
	conn      *pgx.Conn
}

func NewCardQuery(tableName string, conn *pgx.Conn) *CardQuery {
	return &CardQuery{tableName, conn}
}

func (b *CardQuery) Create(ctx context.Context, title string, columnId int) (int, error) {
	var id int

	q := `INSERT INTO ` + b.tableName + ` (title, column_id) VALUES ($1, $2)`
	err := b.conn.QueryRow(ctx, q, title, columnId).Scan(&id)
	if err != nil {
		log.Error("Error inserting into the table: ", err)
		return -1, err
	}

	return id, nil
}

func (b *CardQuery) UpdateCard(ctx context.Context, id int, card *models.Card) error {
	q := `UPDATE ` + b.tableName + ` SET title=$1, column_id=$2 WHERE id = $3`

	_, err := b.conn.Exec(ctx, q, card.Title, card.ColumnID, id)

	if err != nil {
		log.Error("Error updating the table: ", err)
	}

	return err
}

//
//func (b *CardQuery) GetCard(ctx context.Context, id int) (*models.Column, error) {
//	q := `SELECT * FROM ` + b.tableName + `WHERE id = $1`
//
//	rows, err := b.conn.Query(ctx, q, id)
//	if err != nil {
//		log.Error("Error getting the column: ", err)
//		return new(models.Column), err
//	}
//
//	defer rows.Close()
//
//	var column models.Column
//
//	for rows.Next() {
//		if err := rows.Scan(&column.ID, &column.Title, &column.Cards); err != nil {
//			log.Error("Error getting the column: ", err)
//			return &column, err
//		}
//	}
//
//	return &column, nil
//}

func (b *CardQuery) GetAllCards(ctx context.Context, id int) (*[]models.Card, error) {
	q := `SELECT * FROM ` + b.tableName + `WHERE column_id = $1`

	rows, err := b.conn.Query(ctx, q, id)
	if err != nil {
		log.Error("Error getting all cards: ", err)
		return new([]models.Card), err
	}

	defer rows.Close()

	var cards []models.Card

	for rows.Next() {
		var card models.Card
		if err := rows.Scan(&card.ID, &card.Title); err != nil {
			/* TODO OA переделать ошибку*/
			log.Error("Error getting all cards: ", err)
			return &cards, err
		}

		cards = append(cards, card)
	}

	return &cards, nil
}
