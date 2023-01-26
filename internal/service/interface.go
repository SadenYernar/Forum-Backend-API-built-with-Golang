package service

import (
	"forumv2/internal/models"
	"forumv2/internal/repository"

	"github.com/gofrs/uuid"
)

type Service struct {
	User
	Post
	Session
	Comments
	Reactions
}

type User interface {
	CreateSessionService(user models.User) (string, error)
	CreateUserService(user models.User) (int, error)
	AuthorizationUserService(models.User) (string, error)
	GetUserInfoService(user models.User) (models.User, error)
	GetUsersInfoByUUIDService(id uuid.UUID) (models.User, error)
	CheckUserEmail(email string) (bool, error)
	CheckUserUsername(username string) (bool, error)
}

type Post interface {
	CreatePostService(post models.Post) (int64, error)
	GetAllPostService() ([]models.Post, error)
	GetUsersPostInService(uuid uuid.UUID) ([]models.Post, error)
	GetUserLikePostsInService(uuid uuid.UUID) ([]models.Post, error)
	GetPostByIDinService(id int64) (models.Post, error)
	FilterPostsByCategories([]string) ([]models.Post, error)
	CreateCategory([]string) (models.Category, error)
	CheckPostInput(models.Post) error
}

type Session interface {
	DeleteSessionService(uuid uuid.UUID) error
	GetSessionService(token string) (uuid.UUID, error)
}

type Comments interface {
	GetAllCommentsInService() ([]models.Comment, error)
	GetCommentsByIDinService(postID int64) ([]models.Comment, error)
	CreateCommentsInService(com models.Comment) error
	CheckCommentInput(models.Comment) error
}

type Reactions interface {
	LikePostService(like models.LikePost) error
	LikeCommentService(like models.LikeComment) error
}

func NewService(repo repository.Repository) Service {
	return Service{
		User:      NewUserService(repo.User),
		Post:      NewPostService(repo.Post),
		Session:   NewSessionService(repo.Session),
		Comments:  NewCommentsService(repo.Comments),
		Reactions: NewReactionsService(repo.Reactions, repo.Post, repo.Comments),
	}
}
