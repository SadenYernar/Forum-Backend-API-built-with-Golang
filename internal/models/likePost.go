package models

import "github.com/gofrs/uuid"

type LikeStatus int

const (
	Like    LikeStatus = 1
	DisLike LikeStatus = -1
	NoLike  LikeStatus = 0
)

type LikePost struct {
	UserID uuid.UUID  `json:"userid"`
	PostID int64      `json:"postid"`
	Status LikeStatus `json:"status"`
}
