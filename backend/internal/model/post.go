package model

import "time"

type Post struct {
	ID uint `json:"id"`

	Title string `json:"title"`

	Content string `json:"content"`

	AuthorID uint `json:"author_id"`

	Status string `json:"status"`

	CreatedAt time.Time `json:"created_at"`

	UpdatedAt time.Time `json:"updated_at"`
}
