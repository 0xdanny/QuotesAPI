package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/0xdanny/QuotesAPI/server/models"
)

// Parse request body
func parse(w http.ResponseWriter, r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}

func mapPOSTtoJSON(q *models.Quote) models.JsonQuote {
	return models.JsonQuote{
		ID:      q.ID,
		Content: q.Content,
		Author:  q.Author,
	}
}

func respondWithJSON(w http.ResponseWriter, code int, data interface{}) {
	if data == nil {
		return
	}

	res, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("COuld not unmarshal data. %v \n", err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(res)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}
