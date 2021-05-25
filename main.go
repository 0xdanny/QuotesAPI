package main

import (
	"log"
	"net/http"

	"github.com/0xdanny/QuotesAPI/server"
	"github.com/0xdanny/QuotesAPI/server/database"
)

func main() {
	server := server.NewServer()
	server.DB = &database.DB{}
	err := server.DB.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer server.DB.Close()

	http.HandleFunc("/", server.Router.ServeHTTP)

	log.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
