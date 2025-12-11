package db

import (
	"log/slog"
	"time"

	"github.com/infinity/infinity-service/internal/entity"
	"gorm.io/gorm"
)

type SessionRepository struct {
	Repository[entity.UserSession]
	Logger *slog.Logger
}

func NewSessionRepository(logger *slog.Logger) *SessionRepository {
	return &SessionRepository{
		Logger: logger,
	}
}

// CreateSession creates a new user session
func (r *SessionRepository) CreateSession(db *gorm.DB, session *entity.UserSession) error {
	return db.Create(session).Error
}

// FindByToken finds a valid (non-expired) session by token
func (r *SessionRepository) FindByToken(db *gorm.DB, session *entity.UserSession, token string) error {
	return db.Where("token = ? AND expires_at > ?", token, time.Now()).First(session).Error
}

// FindBySessionID finds a session by session ID
func (r *SessionRepository) FindBySessionID(db *gorm.DB, session *entity.UserSession, sessionID string) error {
	return db.Where("session_id = ?", sessionID).First(session).Error
}

// DeleteSession deletes a session (logout)
func (r *SessionRepository) DeleteSession(db *gorm.DB, token string) error {
	return db.Where("token = ?", token).Delete(&entity.UserSession{}).Error
}

// DeleteUserSessions deletes all sessions for a user
func (r *SessionRepository) DeleteUserSessions(db *gorm.DB, userID string) error {
	return db.Where("user_id = ?", userID).Delete(&entity.UserSession{}).Error
}

// DeleteExpiredSessions cleans up expired sessions
func (r *SessionRepository) DeleteExpiredSessions(db *gorm.DB) error {
	return db.Where("expires_at < ?", time.Now()).Delete(&entity.UserSession{}).Error
}
