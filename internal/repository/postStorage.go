package repository

import (
	"database/sql"
	"fmt"
	"forumv2/internal/models"

	"github.com/gofrs/uuid"
)

type PostStorage struct {
	db *sql.DB
}

func NewPostSQLite(db *sql.DB) *PostStorage {
	return &PostStorage{
		db: db,
	}
}

// Создать пост
func (p *PostStorage) CreatePost(post models.Post) (int64, error) {
	query, err := p.db.Prepare(`INSERT INTO post(uuid,title,content,author,createdat,categories) VALUES ($1,$2,$3,$4,$5,$6)`)
	if err != nil {
		return 0, fmt.Errorf("[PostStorage]:Error with CreatePost method in repository: %w", err)
	}

	res, err := query.Exec(post.Uuid, post.Title, post.Content, post.Author, post.CreatedAt, post.Categories)
	if err != nil {
		return 0, fmt.Errorf("[PostStorage]:Error with CreatePost method in repository: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	fmt.Println("Post created successfully!")

	return id, nil
}

// Запрос на все посты
func (p *PostStorage) GetAllPost() ([]models.Post, error) {
	row, err := p.db.Query("SELECT id,uuid,title,content,author,createdAt,categories FROM post")
	if err != nil {
		return nil, fmt.Errorf("[PostStorage]:Error with GetAllPost method in repository: %w", err)
	}

	temp := models.Post{}
	allPost := []models.Post{}

	for row.Next() {
		err := row.Scan(&temp.ID, &temp.Uuid, &temp.Title, &temp.Content, &temp.Author, &temp.CreatedAt, &temp.Categories)
		if err != nil {
			return nil, fmt.Errorf("[PostStorage]:Error with GetAllPost method in repository: %w", err)
		}
		allPost = append(allPost, temp)
	}

	return allPost, nil
}

// Найти пост по ID
func (p *PostStorage) GetPostByID(id int64) (models.Post, error) {
	row := p.db.QueryRow("SELECT id,uuid,title,content,author,createdAt,categories, like, dislike FROM post WHERE id=$1", id)

	temp := models.Post{}
	err := row.Scan(&temp.ID, &temp.Uuid, &temp.Title, &temp.Content, &temp.Author, &temp.CreatedAt, &temp.Categories, &temp.Like, &temp.Dislike)
	if err != nil {
		return temp, fmt.Errorf("[PostStorage]:Error with GetPostByID method in repository: %w", err)
	}
	return temp, nil
}

// Запрос на посты который создал юзер
func (p *PostStorage) GetUsersPost(uuid uuid.UUID) ([]models.Post, error) {
	row, err := p.db.Query("SELECT id,uuid,title,content,author,createdAt,categories,like,dislike FROM post WHERE uuid=$1", uuid)
	if err != nil {
		return nil, fmt.Errorf("[PostStorage]:Error with GetUsersPost method in repository: %w", err)
	}

	temp := models.Post{}
	usersPost := []models.Post{}

	for row.Next() {
		err := row.Scan(&temp.ID, &temp.Uuid, &temp.Title, &temp.Content, &temp.Author, &temp.CreatedAt, &temp.Categories, &temp.Like, &temp.Dislike)
		if err != nil {
			return nil, fmt.Errorf("[PostStorage]:Error with GetUsersPost method in repository: %w", err)
		}
		usersPost = append(usersPost, temp)
	}
	return usersPost, nil
}

// Запрос на ID поста юзера
func (p *PostStorage) GetPostIdWithUUID(uuid uuid.UUID) ([]int64, error) {
	row, err := p.db.Query("SELECT postID FROM likePost WHERE userID==$1 AND status==$2", uuid, 1)
	if err != nil {
		return nil, fmt.Errorf("[PostStorage]:Error with GetPostIdWithUUID method in repository: %w", err)
	}

	temp := models.LikePost{}
	result := []int64{}

	for row.Next() {
		err = row.Scan(&temp.PostID)
		if err != nil {
			return nil, fmt.Errorf("[PostStorage]:Error with GetPostIdWithUUID method in repository: %w", err)
		}
		result = append(result, temp.PostID)
	}

	return result, nil
}

func (p *PostStorage) GetUsersLikePosts(postIdArray []int64) ([]models.Post, error) {
	result := []models.Post{}

	for j := 0; j < len(postIdArray); j++ {
		temp := models.Post{}
		row := p.db.QueryRow("SELECT id,uuid,title,content,author,createdAt,categories,like,dislike FROM post WHERE id=$1", postIdArray[j])

		err := row.Scan(&temp.ID, &temp.Uuid, &temp.Title, &temp.Content, &temp.Author, &temp.CreatedAt, &temp.Categories, &temp.Like, &temp.Dislike)
		if err != nil {
			return nil, fmt.Errorf("[ReactionStorage]:Error with GetUsersLikePosts method in repository: %w", err)
		}
		result = append(result, temp)

	}
	return result, nil
}

func (p *PostStorage) UpdatePost(post models.Post) error {
	stmt := `UPDATE post SET id=$1,uuid=$2,title=$3,content=$4,author=$5,createdat=$6,categories=$7,like=$8,dislike=$9 WHERE id == $1`
	query, err := p.db.Prepare(stmt)
	if err != nil {
		return err
	}
	_, err = query.Exec(&post.ID, &post.Uuid, &post.Title, &post.Content, &post.Author, &post.CreatedAt, &post.Categories, &post.Like, &post.Dislike)
	if err != nil {
		return err
	}
	return nil
}

func (c *PostStorage) GetPostsByCategory(category models.Category) ([]models.Post, error) {
	stmt := `SELECT id, uuid, title, content, author, createdat, categories, like, dislike FROM post WHERE categories&$1 != 0`
	query, err := c.db.Prepare(stmt)
	if err != nil {
		return nil, err
	}
	var res []models.Post
	values, err := query.Query(category)
	if err != nil {
		return nil, err
	}
	for values.Next() {
		var post models.Post
		if err := values.Scan(&post.ID, &post.Uuid, &post.Title, &post.Content, &post.Author, &post.CreatedAt, &post.Categories, &post.Like, &post.Dislike); err != nil {
			return nil, err
		}
		res = append(res, post)
	}
	return res, nil
}
