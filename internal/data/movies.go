package data

import "time"

// By default, the keys in the JSON object are equal to the field names in the struct ( ID,
// CreatedAt, Title and so on).
type Movie struct {
	ID        int64     `json:"id"`                       // Unique integer ID for the movie
	CreatedAt time.Time `json:"-"`                        // Timestamp for when the movie is added to our database, "-" directive, hidden in response
	Title     string    `json:"title"`                    // Movie title
	Year      int32     `json:"year,omitempty"`           // Movie release year, "omitempty" - hide from response if empty
	Runtime   int32     `json:"runtime,omitempty,string"` // Movie runtime (in minutes), "string" - convert int to string
	//		  Runtime										I'm just curious
	Genres  []string `json:"genres,omitempty"` // Slice of genres for the movie (romance, comedy, etc.)
	Version int32    `json:"version"`          // The version number starts at 1 and will be incremented each
	// time the movie information is updated
}
