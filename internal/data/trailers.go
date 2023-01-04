package data

import (
	"database/sql"
)

type Trailer struct {
	ID          int64  `json:"id"`
	TrailerName string `json:"trailer_name"`
	Duration    int32  `json:"duration,omitempty"`
	PremierDate string `json:"premier_date,omitempty"`
	Version     int32  `json:"version"`
}

type TrailerModel struct {
	DB *sql.DB
}

func (t TrailerModel) Insert(trailer *Trailer) error {

	query := `
INSERT INTO trailers (trailer_name, duration, premier_date)
VALUES ($1, $2, $3)
RETURNING id, version`
	// Create an args slice containing the values for the placeholder parameters from
	// the movie struct. Declaring this slice immediately next to our SQL query helps to
	// make it nice and clear *what values are being used where* in the query.
	args := []interface{}{trailer.TrailerName, trailer.Duration, trailer.PremierDate}
	// Use the QueryRow() method to execute the SQL query on our connection pool,
	// passing in the args slice as a variadic parameter and scanning the system-
	// generated id, created_at and version values into the movie struct.
	return t.DB.QueryRow(query, args...).Scan(&trailer.ID, &trailer.Version)
}
