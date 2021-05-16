package database

const createSchema = `
CREATE TABLE IF NOT EXISTS quotes
(
	id SERIAL PRIMARY KEY,
	content TEXT,
	author TEXT
)
`

var insertQuoteSchema = `INSERT INTO quotes(content, author) VALUES($1,$2) RETURNING id`

var deleteQuoteSchema = `DELETE FROM quotes WHERE id=$1`

var updateQuoteSchema = `UPDATE quotes SET author=$1,content=$2 WHERE id=$3`
