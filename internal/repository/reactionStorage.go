package repository

import (
	"database/sql"
	"fmt"
	"forumv2/internal/models"
)

type ReactionsStorage struct {
	db *sql.DB
}

func NewReactionsSQLite(db *sql.DB) *ReactionsStorage {
	return &ReactionsStorage{
		db: db,
	}
}

func (r *ReactionsStorage) CreateLikeForPost(like models.LikePost) (models.LikePost, error) {
	queryForLike, err := r.db.Prepare(`INSERT INTO likePost(userID,postID, status) VALUES ($1,$2,$3)`)
	if err != nil {
		return like, fmt.Errorf("[ReactionStorage]:Error with CreateLikeForPost method in repository: %w", err)
	}
	_, err = queryForLike.Exec(like.UserID, like.PostID, like.Status)
	if err != nil {
		return like, fmt.Errorf("[ReactionStorage]:Error with CreateLikeForPost method in repository: %v", err)
	}
	return like, nil
}

func (r *ReactionsStorage) CreateLikeForComment(like models.LikeComment) (models.LikeComment, error) {
	queryForLike, err := r.db.Prepare(`INSERT INTO likeComments(userID, commentsID, status) VALUES ($1,$2,$3)`)
	if err != nil {
		return like, fmt.Errorf("[ReactionStorage]:Error with CreateLikeForComment method in repository: %w", err)
	}
	_, err = queryForLike.Exec(like.UserID, like.CommentsID, like.Status)
	if err != nil {
		return like, fmt.Errorf("[ReactionStorage]:Error with CreateLikeForComment method in repository: %v", err)
	}
	return like, nil
}

func (r *ReactionsStorage) UpdatePostLikeStatus(like models.LikePost) error {
	records := ("UPDATE likePost SET status = $1 WHERE postID = $2")
	query, err := r.db.Prepare(records)
	if err != nil {
		return fmt.Errorf("[ReactionStorage]:Error with UpdatePostLikeStatus method in repository: %v", err)
	}
	_, err = query.Exec(like.Status, like.PostID)
	if err != nil {
		return fmt.Errorf("[ReactionStorage]:Error with UpdatePostLikeStatus method in repository: %v", err)
	}
	return nil
}

func (r *ReactionsStorage) UpdateCommentLikeStatus(like models.LikeComment) error {
	records := ("UPDATE likeComments SET status = $1 WHERE commentsID = $2")
	query, err := r.db.Prepare(records)
	if err != nil {
		return fmt.Errorf("[ReactionStorage]:Error with UpdateCommentLikeStatus method in repository: %v", err)
	}
	_, err = query.Exec(like.Status, like.CommentsID)
	if err != nil {
		return fmt.Errorf("[ReactionStorage]:Error with UpdateCommentLikeStatus method in repository: %v", err)
	}
	return nil
}

func (r *ReactionsStorage) GetUserIDfromLikePost(like models.LikePost) (int64, error) {
	row := r.db.QueryRow("SELECT postID FROM likePost WHERE userID=$1", like.UserID)
	temp := models.LikePost{}
	err := row.Scan(&temp.PostID)
	if err != nil {
		return temp.PostID, fmt.Errorf("[ReactionStorage]:Error with GetUserIDfromLikePost method in repository: %v", err)
	}
	return temp.PostID, nil
}

func (r *ReactionsStorage) GetLikeStatusByPostAndUserID(like models.LikePost) (models.LikeStatus, error) {
	stmt := `SELECT status FROM likePost WHERE userID == $1 AND postID == $2`
	query, err := r.db.Prepare(stmt)
	if err != nil {
		return models.NoLike, fmt.Errorf("[ReactionStorage]:Error with GetLikeStatusByPostAndUserID method in repository: %v", err)
	}
	res := query.QueryRow(like.UserID, like.PostID)

	var status models.LikeStatus
	err = res.Scan(&status)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.NoLike, nil
		}
		return models.NoLike, fmt.Errorf("[ReactionStorage]:Error with GetLikeStatusByPostAndUserID method in repository: %v", err)
	}
	return status, nil
}

func (r *ReactionsStorage) GetLikeStatusByCommentAndUserID(like models.LikeComment) (models.LikeStatus, error) {
	stmt := `SELECT status FROM likeComments WHERE userID == $1 AND commentsID == $2`
	query, err := r.db.Prepare(stmt)
	if err != nil {
		return models.NoLike, fmt.Errorf("[ReactionStorage]:Error with GetLikeStatusByCommentAndUserID method in repository: %v", err)
	}
	res := query.QueryRow(like.UserID, like.CommentsID)
	var status models.LikeStatus
	err = res.Scan(&status)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.NoLike, nil
		}
		return models.NoLike, fmt.Errorf("[ReactionStorage]:Error with GetLikeStatusByCommentAndUserID method in repository: %v", err)
	}
	return status, nil
}

func (r *ReactionsStorage) DeleteCommentLike(like models.LikeComment) error {
	stmt := `DELETE FROM likeComments WHERE commentsID == $1 AND userID == $2`
	query, err := r.db.Prepare(stmt)
	if err != nil {
		return err
	}
	_, err = query.Exec(like.CommentsID, like.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (r *ReactionsStorage) DeletePostLike(like models.LikePost) error {
	stmt := `DELETE FROM likePost WHERE postID == $1 AND userID == $2`
	query, err := r.db.Prepare(stmt)
	if err != nil {
		return err
	}
	_, err = query.Exec(like.PostID, like.UserID)
	if err != nil {
		return err
	}
	return nil
}
