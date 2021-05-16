package models

type Quote struct {
	ID      int64  `db:"id"`
	Content string `db:"content"`
	Author  string `db:"author"`
}

type JsonQuote struct {
	ID      int64  `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type QuoteRequest struct {
	Content string `json:"content"`
	Author  string `json:"author"`
}
