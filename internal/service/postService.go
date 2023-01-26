package service

import (
	"errors"
	"forumv2/internal/models"
	"forumv2/internal/repository"
	"strings"

	"github.com/gofrs/uuid"
)

var CategoriesMap = map[int64]string{
	models.Coding:  "Coding",
	models.Art:     "Art",
	models.Sports:  "Sports",
	models.Cooking: "Cooking",
	models.Music:   "Music",
	models.Other:   "Other",
}

type PostService struct {
	repo repository.Post
}

func NewPostService(repo repository.Post) Post {
	return &PostService{
		repo: repo,
	}
}

func (p *PostService) CheckPostInput(post models.Post) error {
	if len(post.Title) == 0 {
		return errors.New("empty title")
	}
	if title := strings.Trim(post.Title, "\r\n "); len(title) == 0 {
		return errors.New("empty title")
	}
	if content := strings.Trim(post.Content, "\r\n "); len(content) == 0 {
		return errors.New("empty title")
	}
	if len(post.Title) > 50 {
		return errors.New("title too long")
	}
	if len(post.Content) == 0 {
		return errors.New("empty content")
	}
	if len(post.Content) > 1000 {
		return errors.New("content too long")
	}
	return nil
}

func (p *PostService) CreatePostService(post models.Post) (int64, error) {
	return p.repo.CreatePost(post)
}

func (p *PostService) GetAllPostService() ([]models.Post, error) {
	posts, err := p.repo.GetAllPost()
	if err != nil {
		return nil, err
	}
	for i := range posts {
		posts[i].CategoriesArray = p.AssignPostCategories(posts[i].Categories)
	}
	return posts, nil
}

func (p *PostService) GetPostByIDinService(id int64) (models.Post, error) {
	post, err := p.repo.GetPostByID(id)
	if err != nil {
		return models.Post{}, err
	}
	post.CategoriesArray = p.AssignPostCategories(post.Categories)
	return post, nil
}

func (p *PostService) AssignPostCategories(category models.Category) []string {
	var res []string
	for category != 0 {
		switch {
		case category&models.Art != 0:
			category -= models.Art
			res = append(res, "Art")

		case category&models.Music != 0:
			category -= models.Music
			res = append(res, "Music")

		case category&models.Sports != 0:
			category -= models.Sports
			res = append(res, "Sports")

		case category&models.Coding != 0:
			category -= models.Coding
			res = append(res, "Coding")

		case category&models.Cooking != 0:
			category -= models.Cooking
			res = append(res, "Cooking")

		case category&models.Other != 0:
			category -= models.Other
			res = append(res, "Other")
		case category&models.All != 0:
			category -= models.All
			res = append(res, "All")
		}
	}
	return res
}

func (p *PostService) GetUsersPostInService(uuid uuid.UUID) ([]models.Post, error) {
	posts, err := p.repo.GetUsersPost(uuid)
	if err != nil {
		return nil, err
	}
	for i := range posts {
		posts[i].CategoriesArray = p.AssignPostCategories(posts[i].Categories)
	}
	return posts, nil
}

func (p *PostService) GetUserLikePostsInService(uuid uuid.UUID) ([]models.Post, error) {
	temp, err := p.repo.GetPostIdWithUUID(uuid)
	if err != nil {
		return nil, err
	}
	posts, err := p.repo.GetUsersLikePosts(temp)
	if err != nil {
		return nil, err
	}
	for i := range posts {
		posts[i].CategoriesArray = p.AssignPostCategories(posts[i].Categories)
	}
	return posts, nil
}

func (p *PostService) FilterPostsByCategories(categories []string) ([]models.Post, error) {
	category, err := p.CreateCategory(categories)
	if err != nil {
		return nil, err
	}
	posts, err := p.repo.GetPostsByCategory(category)
	if err != nil {
		return nil, err
	}
	for i := range posts {
		posts[i].CategoriesArray = p.AssignPostCategories(posts[i].Categories)
	}
	return posts, nil
}

func (p *PostService) CreateCategory(categories []string) (models.Category, error) {
	var category models.Category
	for _, val := range categories {
		switch val {
		case "Coding":
			category = category | models.Coding
		case "Music":
			category = category | models.Music
		case "Art":
			category = category | models.Art
		case "Sports":
			category = category | models.Sports
		case "Cooking":
			category = category | models.Cooking
		case "Other":
			category = category | models.Other
		case "All":
			category = category | models.All

		default:
			return models.Other, errors.New("invalid category")
		}
	}
	return category, nil
}
