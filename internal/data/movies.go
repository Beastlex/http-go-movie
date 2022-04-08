package data

import "time"

type Movie struct {
	ID        int64     `json: "id"`
	CreatedAt time.Time `json: "created_at"`
	Title     string    `json: "title"`
	Year      int32     `json: "year"`
	Runtime   int32     `json: "runtime"` //in minutes
	Genres    []string  `json: "genres"`
	Version   int32     `json: "version"` // inc each time the movie info is updates
}
