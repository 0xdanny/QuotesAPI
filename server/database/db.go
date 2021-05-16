package database

import (
	"log"

	"github.com/0xdanny/QuotesAPI/server/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type QuotesDB interface {
	Open() error
	Close() error
	CreateQuote(q *models.Quote) error
	GetQuotes() ([]*models.Quote, error)
	DeleteQuote(id string) error
	UpdateQuote(q *models.Quote) error
}

type DB struct {
	db *sqlx.DB
}

func (d *DB) Open() error {
	pg, err := sqlx.Open(DB_DRIVER, DB_URL)
	if err != nil {
		return err
	}
	log.Println("Connected to Production Database")

	pg.MustExec(createSchema)

	d.db = pg

	return nil
}

func (d *DB) Close() error {
	return d.db.Close()
}

func (d *DB) CreateQuote(q *models.Quote) error {
	res, err := d.db.Exec(insertQuoteSchema, q.Content, q.Author)
	if err != nil {
		return err
	}

	res.LastInsertId()
	return err
}

func (d *DB) GetQuotes() ([]*models.Quote, error) {
	var quotes []*models.Quote
	err := d.db.Select(&quotes, "SELECT * FROM quotes")
	if err != nil {
		return quotes, err
	}
	return quotes, nil
}

func (d *DB) DeleteQuote(id string) error {
	_, err := d.db.Exec(deleteQuoteSchema, id)
	if err != nil {
		return err
	}
	return err
}

func (d *DB) UpdateQuote(q *models.Quote) error {
	_, err := d.db.Exec(updateQuoteSchema,
		q.Author, q.Content, q.ID)
	if err != nil {
		return err
	}
	return err
}
