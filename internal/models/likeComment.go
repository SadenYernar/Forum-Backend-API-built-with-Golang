package models

import "github.com/gofrs/uuid"

type LikeComment struct {
	UserID     uuid.UUID  `json:"userid"`
	CommentsID int        `json:"commentsid"`
	Status     LikeStatus `json:"status"`
}
