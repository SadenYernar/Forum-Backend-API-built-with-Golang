package service

import (
	"forumv2/internal/repository"

	"github.com/gofrs/uuid"
)

type SessionService struct {
	repo repository.Session
}

func NewSessionService(repo repository.Session) *SessionService {
	return &SessionService{
		repo: repo,
	}
}

// RQ-request
// Запрос на удаления токена и время токена
func (s *SessionService) DeleteSessionService(uuid uuid.UUID) error {
	err := s.repo.DeleteSessionFromDB(uuid)
	if err != nil {
		return err
	}
	return nil
}

// Получаем по токену username
func (s *SessionService) GetSessionService(token string) (uuid.UUID, error) {
	uuid, err := s.repo.GetSessionFromDB(token)
	if err != nil {
		return uuid, err
	}
	return uuid, nil
}
