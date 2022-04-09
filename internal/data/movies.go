package data

import (
	"http-go-movie/internal/validator"
	"time"
)

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"` // Use the - directive
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Runtime   Runtime   `json:"runtime,omitempty"` //in minutes
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"` // inc each time the movie info is updates
}

func ValidateMovie(v *validator.Validator, movie *Movie) {
	// Check Title
	v.Check(movie.Title != "", "title", "must be provided")
	v.Check(len(movie.Title) <= 500, "title", "must not be more than 500 bytes long")

	// Check Year
	v.Check(movie.Year != 0, "year", "must be provided")
	v.Check(movie.Year >= 1888, "year", "must be grater than 1888")
	v.Check(movie.Year <= int32(time.Now().Year()), "year", "must not be in future")

	// Check Genres
	v.Check(movie.Genres != nil, "genres", "must be provided")
	v.Check(len(movie.Genres) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(movie.Genres) <= 5, "genres", "must not contain more than 5 genres")
	v.Check(validator.Unique(movie.Genres), "genres", "must not contain duplicate values")
}
