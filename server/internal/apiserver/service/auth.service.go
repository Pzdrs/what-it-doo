package service

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
	"pycrs.cz/what-it-do/internal/apiserver/model"
	"pycrs.cz/what-it-do/internal/apiserver/repository"
	"pycrs.cz/what-it-do/internal/database"
)

type AuthService struct {
	repository        *repository.UserRepository
	sessionRepository *repository.SessionRepository
}

func NewAuthService(repo *repository.UserRepository, sessionRepo *repository.SessionRepository) *AuthService {
	return &AuthService{
		repository:        repo,
		sessionRepository: sessionRepo,
	}
}

func mapSessionToModel(session database.Session) model.UserSession {
	return model.UserSession{
		ID:         session.ID,
		UserID:     session.UserID,
		Token:      session.Token,
		DeviceType: session.DeviceType.String,
		DeviceOs:   session.DeviceOs.String,
		CreatedAt:  session.CreatedAt.Time,
		ExpiresAt:  session.ExpiresAt.Time,
		RevokedAt:  session.RevokedAt.Time,
	}
}

func (s *AuthService) RegisterUser(user model.User, password string) (model.User, error) {
	if s.repository.UserExists(user.Email) {
		return model.User{}, fmt.Errorf("user already exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	u, err := s.repository.SaveUser(database.User{
		FirstName:      pgtype.Text{String: user.FirstName, Valid: true},
		LastName:       pgtype.Text{String: user.LastName, Valid: true},
		Email:          user.Email,
		HashedPassword: pgtype.Text{String: string(hashedPassword), Valid: true},
		CreatedAt:      pgtype.Timestamptz{Time: time.Now()},
		UpdatedAt:      pgtype.Timestamptz{Time: time.Now()},
	})
	if err != nil {
		return model.User{}, err
	}
	return mapUserToModel(u), nil
}

func (s *AuthService) GetUserByEmail(email string) (model.User, error) {
	return func() (model.User, error) {
		user, err := s.repository.GetUserByEmail(email)
		if err != nil {
			return model.User{}, err
		}
		return mapUserToModel(user), nil
	}()
}

func (s *AuthService) AuthenticateUser(email, password string) bool {
	user, err := s.repository.GetUserByEmail(email)
	if err != nil {
		return false
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword.String), []byte(password))
	return err == nil
}

func (s *AuthService) CreateSession(userID uuid.UUID, deviceType, deviceOs string) (model.UserSession, error) {
	return func() (model.UserSession, error) {
		session, err := s.sessionRepository.CreateSession(database.CreateSessionParams{
			UserID: userID,
			// TODO: Use a more secure token generation method
			Token:      uuid.NewString(),
			DeviceType: pgtype.Text{String: deviceType, Valid: true},
			DeviceOs:   pgtype.Text{String: deviceOs, Valid: true},
			ExpiresAt:  pgtype.Timestamptz{Time: time.Now().Add(7 * 24 * time.Hour), Valid: true},
		})
		if err != nil {
			return model.UserSession{}, err
		}
		return mapSessionToModel(session), nil
	}()
}

func (s *AuthService) FindSession(token string) (model.UserSession, bool) {
	session, err := s.sessionRepository.GetSessionByToken(token)
	if err != nil {
		return model.UserSession{}, false
	}
	if session.ExpiresAt.Time.Before(time.Now()) || session.RevokedAt.Valid {
		return mapSessionToModel(session), false
	}
	return mapSessionToModel(session), true
}

func (s *AuthService) Logout(session model.UserSession) error {
	return s.sessionRepository.DeleteSessionByID(session.ID)
}
