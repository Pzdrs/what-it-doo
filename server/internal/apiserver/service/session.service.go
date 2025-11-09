package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"pycrs.cz/what-it-doo/internal/apiserver/model"
	"pycrs.cz/what-it-doo/internal/apiserver/repository"
	"pycrs.cz/what-it-doo/internal/config"
)

type SessionService interface {
	// Create creates a new user session for the given user ID, device type, and device OS.
	Create(ctx context.Context, userID uuid.UUID, deviceType, deviceOs string) (model.UserSession, error)
	// GetByToken retrieves a user session by its token.
	GetByToken(ctx context.Context, token string) (model.UserSession, bool)
}

type sessionService struct {
	repository        repository.UserRepository
	sessionRepository repository.SessionRepository
	config            config.Configuration
}

func NewSessionService(repo repository.UserRepository, sessionRepo repository.SessionRepository, config config.Configuration) SessionService {
	return &sessionService{
		repository:        repo,
		sessionRepository: sessionRepo,
		config:            config,
	}
}

func (s *sessionService) Create(ctx context.Context, userID uuid.UUID, deviceType, deviceOs string) (model.UserSession, error) {
	session := model.UserSession{
		UserID: userID,
		// TODO: Use a more secure token generation method
		Token:      uuid.NewString(),
		DeviceType: deviceType,
		DeviceOs:   deviceOs,
		ExpiresAt:  time.Now().Add(7 * 24 * time.Hour),
	}

	if err := s.sessionRepository.Create(ctx, &session); err != nil {
		return model.UserSession{}, err
	}

	return session, nil
}

func (s *sessionService) GetByToken(ctx context.Context, token string) (model.UserSession, bool) {
	session, err := s.sessionRepository.GetByToken(ctx, token)
	if err != nil {
		return model.UserSession{}, false
	}
	if session.ExpiresAt.Before(time.Now()) || session.RevokedAt != nil {
		return *session, false
	}
	return *session, true
}

var _ SessionService = (*sessionService)(nil)
