// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package sqlc

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type AuthAccessToken struct {
	ID        uuid.UUID     `json:"id"`
	SessionID uuid.NullUUID `json:"session_id"`
	Token     string        `json:"token"`
	CreatedAt sql.NullTime  `json:"created_at"`
	ExpiresAt time.Time     `json:"expires_at"`
}

type AuthRefreshToken struct {
	ID        uuid.UUID     `json:"id"`
	SessionID uuid.NullUUID `json:"session_id"`
	Token     string        `json:"token"`
	CreatedAt sql.NullTime  `json:"created_at"`
	ExpiresAt time.Time     `json:"expires_at"`
	IsRevoked sql.NullBool  `json:"is_revoked"`
}

type AuthSession struct {
	ID        uuid.UUID     `json:"id"`
	UserID    uuid.NullUUID `json:"user_id"`
	CreatedAt sql.NullTime  `json:"created_at"`
	UpdatedAt sql.NullTime  `json:"updated_at"`
}

type AuthUser struct {
	ID              uuid.UUID      `json:"id"`
	Email           string         `json:"email"`
	PasswordHash    string         `json:"password_hash"`
	FirstName       sql.NullString `json:"first_name"`
	LastName        sql.NullString `json:"last_name"`
	IsEmailVerified bool           `json:"is_email_verified"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
}
