package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ErrExpiredToken = errors.New("Token has expired")
	ErrInvalidToken = errors.New("Token is invalid")
)

// Payload contains the payload data of the token
type Payload struct {
	ID                   uuid.UUID `json:"id"`
	Username             string    `json:"username"`
	jwt.RegisteredClaims           // Embed jwt.RegisteredClaims to satisfy jwt.Claims
}

// NewPayload create a new token payload with a specific username and duration:
func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:       tokenID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}
	return payload, nil
}

// Valid checks if the token payload is valid or not:
func (payload *Payload) Valid() error {
	if payload.ExpiresAt != nil && time.Now().After(payload.ExpiresAt.Time) {
		return ErrExpiredToken
	}
	return nil
}
