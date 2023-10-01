package data

import "time"

// note: all fields are exported to expose them to encoding/json package
type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	// note the custom Runtime type
	Runtime Runtime  `json:"runtime,omitempty,string"`
	Genres  []string `json:"genres,omitempty"`
	Version int32    `json:"version"`
}
