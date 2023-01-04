package data

import (
	"context"
	"database/sql"
	"fmt"
	"time"
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

func (t TrailerModel) GetAll(trailer_name string, filters Filters) ([]*Trailer, error) {
	query := fmt.Sprintf(`
SELECT id, trailer_name, duration, premier_date, version
FROM trailers
WHERE (to_tsvector('english', trailer_name) @@ plainto_tsquery('english', $1) OR $1 = '')
ORDER BY %s %s, id ASC`, filters.sortColumn(), filters.sortDirection())

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []interface{}{trailer_name}

	rows, err := t.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	// Importantly, defer a call to rows.Close() to ensure that the resultset is closed
	// before GetAll() returns.
	defer rows.Close()

	trailers := []*Trailer{}
	for rows.Next() {
		// Initialize an empty Movie struct to hold the data for an individual movie.
		var trailer Trailer
		// Scan the values from the row into the Movie struct. Again, note that we're
		// using the pq.Array() adapter on the genres field here.
		err := rows.Scan(
			&trailer.ID,
			&trailer.TrailerName,
			&trailer.Duration,
			&trailer.PremierDate,
			&trailer.Version,
		)
		if err != nil {
			return nil, err
		}
		// Add the Movie struct to the slice.
		trailers = append(trailers, &trailer)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return trailers, nil
}
