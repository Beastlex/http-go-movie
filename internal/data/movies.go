package data

import "time"

type Movie struct {
	ID        int64
	CreatedAt time.Time
	Title     string
	Year      int32
	Runtime   int32 //in minutes
	Genres    []string
	Version   int32 // inc each time the movie info is updates
}
