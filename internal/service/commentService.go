package service

import (
	"errors"
	"forumv2/internal/models"
	"forumv2/internal/repository"
	"strings"
)

type CommentService struct {
	repo repository.Comments
}

func NewCommentsService(repo repository.Comments) *CommentService {
	return &CommentService{
		repo: repo,
	}
}

func (c *CommentService) CheckCommentInput(comment models.Comment) error {
	if comment := strings.Trim(comment.Content, "\r\n "); len(comment) == 0 {
		return errors.New("empty title")
	}
	if len(comment.Content) == 0 {
		return errors.New("empty comment")
	}
	if len(comment.Content) > 500 {
		return errors.New("comment too long")
	}
	return nil
}

func (c *CommentService) GetAllCommentsInService() ([]models.Comment, error) {
	return c.repo.GetAllComments()
}

func (c *CommentService) GetCommentsByIDinService(postID int64) ([]models.Comment, error) {
	return c.repo.GetCommentsByID(postID)
}

func (c *CommentService) CreateCommentsInService(com models.Comment) error {
	return c.repo.CreateComments(com)
}
