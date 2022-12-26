package data

type Actor struct {
	ID           int64    `json:"id"`
	FirstName    string   `json:"firstName"`
	LastName     string   `json:"lastName"`
	DateOfBirth  int32    `json:"dateOfBirth,omitempty"`
	MoviesCasted []string `json:"moviesCasted,omitempty"`
	// time the movie information is updated
}
