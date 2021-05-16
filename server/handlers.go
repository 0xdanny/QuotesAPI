package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/0xdanny/QuotesAPI/server/models"
	"github.com/gorilla/mux"
)

func (s *Server) HomeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Quotes API v1")
	}
}

func (s *Server) CreateQuoteHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := models.QuoteRequest{}
		err := parse(w, r, &req)
		if err != nil {
			log.Printf("Cannot parse POST body. %v \n", err)
			respondWithError(w, http.StatusBadRequest, "Invalid POST data")
			return
		}

		// Create a Quote
		q := &models.Quote{
			ID:      0,
			Content: req.Content,
			Author:  req.Author,
		}

		// Save Quote to DB
		err = s.DB.CreateQuote(q)
		if err != nil {
			log.Printf("Could not save quote in DB. %v\n", err)
			respondWithError(w, http.StatusInternalServerError, "Error while saving quote")
		}

		res := mapPOSTtoJSON(q)
		respondWithJSON(w, http.StatusCreated, res)
	}
}

func (s *Server) GetQuotesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		quotes, err := s.DB.GetQuotes()
		if err != nil {
			log.Printf("Could not get quotes. %v\n", err)
			respondWithError(w, http.StatusInternalServerError, "Could not get quotes")
			return
		}

		var res = make([]models.JsonQuote, len(quotes))
		for i, q := range quotes {
			res[i] = mapPOSTtoJSON(q)
		}

		respondWithJSON(w, http.StatusOK, res)
	}
}

func (s *Server) DeleteQuoteHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		err := s.DB.DeleteQuote(id)
		if err != nil {
			log.Fatalf("Could not delete quote. %v", err)
			respondWithError(w, http.StatusInternalServerError, "Could not delete the quote")
			return
		}

		respondWithJSON(w, http.StatusOK, "quote deleted")

	}
}

func (s *Server) UpdateQuoteHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := models.Quote{}
		err := parse(w, r, &req)

		if err != nil {
			log.Printf("Cannot parse POST body. %v \n", err)
			respondWithError(w, http.StatusBadRequest, "Invalid POST data")
			return
		}

		q := &models.Quote{
			ID:      req.ID,
			Content: req.Content,
			Author:  req.Author,
		}

		err = s.DB.UpdateQuote(q)
		if err != nil {
			log.Printf("Could not update the quote. %v\n", err)
			respondWithError(w, http.StatusInternalServerError, "Error while updating the quote")
			return
		}

		respondWithJSON(w, http.StatusCreated, "Successfully updated")
	}
}
