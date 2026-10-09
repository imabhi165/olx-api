package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"log/slog"
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

type ListingHandlers struct {
	db     *sql.DB
	logger *slog.Logger
}

// NewListingHandlers creates a new ListingHandlers with the given database connection.
func NewListingHandlers(db *sql.DB, logger *slog.Logger) *ListingHandlers {
	return &ListingHandlers{
		db:     db,
		logger: logger,
	}
}

// ListingsHandler handles GET /listings and returns up to 100 listings as JSON.

func (lh *ListingHandlers) List(w http.ResponseWriter, r *http.Request) {
	// Request scope context
	ctx := r.Context()
	rows, err := lh.db.QueryContext(ctx,
		`SELECT id, title, description, price, city, created_at
			FROM listings
			ORDER BY created_at DESC
			LIMIT 100`)
	if err != nil {
		// log.Printf("query: %v", err)
		lh.logger.Error("query error: ", "err", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var listings []Listing
	for rows.Next() {
		var listing Listing
		if err := rows.Scan(&listing.ID, &listing.Title, &listing.Description, &listing.Price, &listing.City, &listing.CreatedAt); err != nil {
			// log.Printf("scan: %v", err)
			lh.logger.Error("listing scan error: ", "err", err)
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}
		listings = append(listings, listing)
	}
	if err := rows.Err(); err != nil {
		// log.Printf("rows: %v", err)
		lh.logger.Error("listings rows error: ", "err", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(listings); err != nil {
		// log.Printf("encode: %v", err)
		lh.logger.Error("listings rows error: ", "err", err)
	}
}

// DeleteListings handles DELETE /listings/{id} and removes the listing from the database.

func (lh *ListingHandlers) Delete(w http.ResponseWriter, r *http.Request) {
	listingID := r.PathValue("id")
	if listingID == "" {
		http.Error(w, "Missing listing ID", http.StatusBadRequest)
		return
	}
	log.Printf("deleting listing ID: %s", listingID)
	ctx := r.Context()
	_, err := lh.db.ExecContext(ctx, `DELETE FROM listing WHERE id = $1`, listingID)
	if err != nil {
		// log.Printf("delete: %v", err)
		// slog.Debug("Debug log:", "listing_id: ", listingID)
		// slog.Warn("Warnings log:", "listing_id: ", listingID)
		lh.logger.Error("Delete failed", "listings_id:", listingID, "err", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	slog.Info("Record deleted", "listing_id: ", listingID)

	w.WriteHeader(http.StatusNoContent)
}
