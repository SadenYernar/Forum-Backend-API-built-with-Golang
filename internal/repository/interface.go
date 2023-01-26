package repository

import (
	"database/sql"
	"forumv2/internal/models"
	"time"

	"github.com/gofrs/uuid"
)

type Repository struct {
	User
	Post
	Session
	Comments
	Reactions
}

type User interface {
	SetSession(user models.User, token string, time time.Time) error
	CreateUser(user models.User) (int, error)
	GetUserInfo(user models.User) (models.User, error)
	GetUsersEmail(user models.User) (models.User, error)
	GetUsersInfoByUUID(id uuid.UUID) (models.User, error) //++
	CheckUserEmail(email string) (bool, error)
	CheckUserUsername(username string) (bool, error)
}

type Post interface {
	CreatePost(post models.Post) (int64, error)

	GetAllPost() ([]models.Post, error)
	GetPostByID(id int64) (models.Post, error)
	GetUsersPost(uuid uuid.UUID) ([]models.Post, error)
	GetPostIdWithUUID(uuid uuid.UUID) ([]int64, error)
	GetUsersLikePosts(i []int64) ([]models.Post, error)
	GetPostsByCategory(category models.Category) ([]models.Post, error)

	UpdatePost(models.Post) error
}

type Session interface {
	GetSessionFromDB(token string) (uuid.UUID, error)
	DeleteSessionFromDB(uuid.UUID) error
}

type Comments interface {
	CreateComments(models.Comment) error

	GetAllComments() ([]models.Comment, error)
	GetCommentsByID(postID int64) ([]models.Comment, error)
	GetCommentByCommentID(commentID int) (models.Comment, error)

	UpdateComment(models.Comment) error
}

type Reactions interface {
	CreateLikeForPost(like models.LikePost) (models.LikePost, error)
	CreateLikeForComment(like models.LikeComment) (models.LikeComment, error)

	GetUserIDfromLikePost(like models.LikePost) (int64, error)
	GetLikeStatusByPostAndUserID(like models.LikePost) (models.LikeStatus, error)
	GetLikeStatusByCommentAndUserID(like models.LikeComment) (models.LikeStatus, error)

	UpdatePostLikeStatus(like models.LikePost) error
	UpdateCommentLikeStatus(like models.LikeComment) error

	DeletePostLike(models.LikePost) error
	DeleteCommentLike(models.LikeComment) error
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		User:      NewUserSQLite(db),
		Post:      NewPostSQLite(db),
		Session:   NewSessionSQLite(db),
		Comments:  NewCommentsSQLite(db),
		Reactions: NewReactionsSQLite(db),
	}
}
