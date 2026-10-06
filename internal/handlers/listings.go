package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Listing struct {
	ID          string
	Title       string
	Description string
	Price       int
	City        string
	CreatedAt   time.Time
}

// wrap the handler to inject the database connection -> return a http.HandlerFunc and handle the request
func ListingsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//TODO: query the database and return the results as JSON
		rows, err := db.Query(
			`SELECT id, title, description, price, city, created_at
			FROM listings
			ORDER BY created_at DESC
			LIMIT 100`)
		if err != nil {
			log.Printf("query: %v", err)
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		//TODO: iterate over the rows and return them as JSON
		var listings []Listing
		for rows.Next() {
			var listing Listing
			if err := rows.Scan(&listing.ID, &listing.Title, &listing.Description, &listing.Price, &listing.City, &listing.CreatedAt); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			listings = append(listings, listing)
		}
		// check for errors from the rows iterator
		if err := rows.Err(); err != nil {
			log.Printf("rows: %v", err)
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}
		// TODO: convert the listings to JSON and write it to the response
		// set the content type to application/json
		w.Header().Set("Content-Type", "application/json")
		// set the status code to 200 OK
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(listings)
	}
}
