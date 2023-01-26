package service

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"forumv2/internal/models"
	"forumv2/internal/repository"
	"log"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

const TOKEN_SECRET = 15

type UserService struct {
	repo repository.User
	gen  uuid.Generator
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{
		gen:  uuid.NewGen(),
		repo: repo,
	}
}

// Добавление нового юзера в базу,запрос на репу
func (u *UserService) CreateUserService(user models.User) (int, error) {
	var err error

	if !userValidation(user) {
		return http.StatusBadRequest, fmt.Errorf("Create user in service: %w", err)
	}
	user.Uuid, err = u.gen.NewV4()
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Create user in service: %w", err)
	}
	user.Password, err = generateHashPassword(user.Password)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Create user in service: %w", err)
	}
	return u.repo.CreateUser(user)
}

// Проверка на авторизацию
func (u *UserService) AuthorizationUserService(user models.User) (string, error) {
	var err error
	checkUser, err := u.repo.GetUserInfo(user)
	if err != nil {
		return "User is exist", err
	}
	if checkUser.Email != user.Email {
		return "Not correct email", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(checkUser.Password), []byte(user.Password))
	if err != nil {
		return "Not correct password", err
	}

	value, err := u.CreateSessionService(checkUser)
	if err != nil {
		return "Session not created", err
	}
	return value, err
}

// Получения данных юзера из БД
func (u *UserService) GetUserInfoService(user models.User) (models.User, error) {
	userInfo, err := u.repo.GetUserInfo(user) // Получает информацию с помощью почты
	if err != nil {
		return models.User{}, err
	}
	return userInfo, nil
}

// Получения данных юзера из БД c помощью токена
func (u *UserService) GetUsersInfoByUUIDService(id uuid.UUID) (models.User, error) {
	userInfo, err := u.repo.GetUsersInfoByUUID(id)
	if err != nil {
		return models.User{}, err
	}
	return userInfo, nil
}

// Cоздает токен и время токена и отправляет в БД
func (u *UserService) CreateSessionService(user models.User) (string, error) {
	token := CreateToken()
	expireTime := time.Now()
	return token, u.repo.SetSession(user, token, expireTime)
}

func CreateToken() string {
	b := make([]byte, TOKEN_SECRET)
	if _, err := rand.Read(b); err != nil {
		log.Print("Token for user not created")
	}
	return hex.EncodeToString(b)
}

func (u *UserService) CheckUserEmail(email string) (bool, error) {
	return u.repo.CheckUserEmail(email)
}

func (u *UserService) CheckUserUsername(username string) (bool, error) {
	return u.repo.CheckUserUsername(username)
}
