package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"forumv2/internal/models"
	"log"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
)

type UserStorage struct {
	db *sql.DB
}

func NewUserSQLite(db *sql.DB) *UserStorage {
	return &UserStorage{
		db: db,
	}
}

// Добавление нового пользователя в базу
func (u *UserStorage) CreateUser(user models.User) (int, error) {
	records := `INSERT INTO users(uuid,name,username,email,password) VALUES ($1,$2,$3,$4,$5)`

	query, err := u.db.Prepare(records)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Error in CreateUser method in repository: %w", err)
	}

	_, err = query.Exec(user.Uuid, user.Name, user.Username, user.Email, user.Password)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Error in CreateUser method in repository: %w", err)
	}

	fmt.Println("User created successfully!")
	return http.StatusOK, err
}

// Создает токен и время для токена по uuid
func (u *UserStorage) SetSession(user models.User, token string, time time.Time) error {
	records := `UPDATE users SET token=$1,expiretime=$2 WHERE uuid=$3`

	query, err := u.db.Prepare(records)
	if err != nil {
		return fmt.Errorf("Error in SetSession method in repository: %w", err)
	}

	_, err = query.Exec(token, time, user.Uuid)
	if err != nil {
		return fmt.Errorf("Error in SetSession method in repository: %w", err)
	}

	fmt.Println("Session created successfully!")
	return nil
}

// Получить полную информация о юзере с помощью почты
func (u *UserStorage) GetUserInfo(user models.User) (models.User, error) {
	row := u.db.QueryRow("SELECT uuid,name,username,email,password FROM users WHERE email=$1", user.Email)

	temp := models.User{}
	err := row.Scan(&temp.Uuid, &temp.Name, &temp.Username, &temp.Email, &temp.Password)
	if err != nil {
		log.Printf("Error with GetUserInfo in repository: %v\n", err)
		return models.User{}, err
	}
	return temp, nil
}

// Получить почту юзера по username
func (u *UserStorage) GetUsersEmail(user models.User) (models.User, error) {
	row := u.db.QueryRow("SELECT email FROM users WHERE username=$1", user.Username)

	temp := models.User{}
	err := row.Scan(&temp.Email)
	if err != nil {
		log.Printf("Error with GetUsersEmail method in repository: %v\n", err)
		return models.User{}, err
	}
	return temp, nil
}

// Получить информацию юзера по uuid
func (u *UserStorage) GetUsersInfoByUUID(id uuid.UUID) (models.User, error) {
	row := u.db.QueryRow("SELECT name,username,email,password FROM users WHERE uuid=$1", id)

	temp := models.User{}
	err := row.Scan(&temp.Name, &temp.Username, &temp.Email, &temp.Password)
	if err != nil {
		log.Printf("GetUsersInfoByUUID error: %v\n", err)
		return models.User{}, err
	}
	return temp, nil
}

// CheckUserEmail - returns true if user by this email exists
func (u *UserStorage) CheckUserEmail(email string) (UserExist bool, err error) {
	stmt := `SELECT email FROM users WHERE email == $1`
	query, err := u.db.Prepare(stmt)
	if err != nil {
		return false, err
	}
	row := query.QueryRow(email)
	var mail string
	err = row.Scan(&mail)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (u *UserStorage) CheckUserUsername(username string) (UserExist bool, err error) {
	stmt := `SELECT username FROM users WHERE username == $1`
	query, err := u.db.Prepare(stmt)
	if err != nil {
		return false, err
	}
	row := query.QueryRow(username)
	var name string
	err = row.Scan(&name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
