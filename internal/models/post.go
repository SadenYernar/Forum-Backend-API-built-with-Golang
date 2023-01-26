package models

import "github.com/gofrs/uuid"

type Post struct {
	Uuid            uuid.UUID `json:"uuid"`
	ID              int64     `json:"id"`
	Title           string    `json:"title"`
	Content         string    `json:"content"`
	Author          string    `json:"author"`
	CreatedAt       string    `json:"createdat"`
	Categories      Category  `json:"categories"`
	CategoriesArray []string
	Like            int `json:"like"`
	Dislike         int `json:"dislike"`
}
