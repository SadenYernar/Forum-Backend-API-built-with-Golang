package repository

import (
	"database/sql"
	"fmt"
	"forumv2/internal/models"
)

type CommentStorage struct {
	db *sql.DB
}

func NewCommentsSQLite(db *sql.DB) *CommentStorage {
	return &CommentStorage{
		db: db,
	}
}

func (c *CommentStorage) CreateComments(com models.Comment) error {
	query, err := c.db.Prepare(`INSERT INTO comments(postID,content,author,like,dislike,createdat) VALUES ($1,$2,$3,$4,$5,$6)`)
	if err != nil {
		return fmt.Errorf("[CommentStorage]:Error with CreateComments method in repository: %w", err)
	}

	_, err = query.Exec(com.PostID, com.Content, com.Author, com.Like, com.Dislike, com.CreatedAt)
	if err != nil {
		return fmt.Errorf("Create comment in repository: %w", err)
	}

	return nil
}

func (c *CommentStorage) GetAllComments() ([]models.Comment, error) {
	stmt := `SELECT id, postID,content,author,like,dislike,createdat FROM comments`
	query, err := c.db.Prepare(stmt)
	if err != nil {
		return nil, err
	}

	row, err := query.Query()
	if err != nil {
		return nil, fmt.Errorf("[CommentStorage]:Error with GetAllComments method in repository: %w", err)
	}

	temp := models.Comment{}
	allComments := []models.Comment{}

	for row.Next() {
		err := row.Scan(&temp.ID, &temp.PostID, &temp.Content, &temp.Author, &temp.Like, &temp.Dislike, &temp.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("[CommentStorage]:Error with GetAllComments method in repository: %w", err)
		}
		allComments = append(allComments, temp)
	}
	return allComments, nil
}

func (c *CommentStorage) GetCommentsByID(postID int64) ([]models.Comment, error) {
	row, err := c.db.Query("SELECT id,postID,content,author,like,dislike,createdat FROM comments WHERE postID=$1", postID)
	if err != nil {
		return nil, fmt.Errorf("[CommentStorage]:Error with GetCommentsByID method in repository: %w", err)
	}

	temp := models.Comment{}
	allComments := []models.Comment{}

	for row.Next() {
		err := row.Scan(&temp.ID, &temp.PostID, &temp.Content, &temp.Author, &temp.Like, &temp.Dislike, &temp.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("[CommentStorage]:Error with GetCommentsByID method in repository: %w", err)
		}
		allComments = append(allComments, temp)
	}
	return allComments, nil
}

func (c *CommentStorage) UpdateComment(comment models.Comment) error {
	stmt := `UPDATE comments SET id = $1, postID = $2, content = $3, author = $4, like = $5, dislike = $6, createdat = $7 WHERE id == $1`
	query, err := c.db.Prepare(stmt)
	if err != nil {
		return fmt.Errorf("error executing statement %v:\n%v", stmt, err)
	}
	_, err = query.Exec(&comment.ID, &comment.PostID, &comment.Content, &comment.Author, &comment.Like, &comment.Dislike, &comment.CreatedAt)
	if err != nil {
		return fmt.Errorf("error executing statement %v: %v", stmt, err)
	}
	return nil
}

func (c *CommentStorage) GetCommentByCommentID(commentID int) (models.Comment, error) {
	stmt := `SELECT id, postID, content, author, like, dislike, createdat FROM comments WHERE id == $1`
	query, err := c.db.Prepare(stmt)
	if err != nil {
		return models.Comment{}, fmt.Errorf("error executing statement %v: %v", stmt, err)
	}
	var res models.Comment
	err = query.QueryRow(commentID).Scan(&res.ID, &res.PostID, &res.Content, &res.Author, &res.Like, &res.Dislike, &res.CreatedAt)
	if err != nil {
		return models.Comment{}, fmt.Errorf("error executing statement %v: %v", stmt, err)
	}
	return res, nil
}
