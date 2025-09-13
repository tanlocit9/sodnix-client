package auth

import "github.com/google/uuid"

// MessageResponse represents a generic message response.
type MessageResponse struct {
	Message string `json:"message" example:"Operation successful"`
}

// ProfileResponse represents the user profile data.
type ProfileResponse struct {
	Message string    `json:"message" example:"Welcome"`
	UserID  uuid.UUID `json:"userId" example:"550e8400-e29b-41d4-a716-446655440000"`
}

type LoginRequest struct {
	Email    string `json:"email" example:"admin@example.com"`
	Password string `json:"password" example:"admin"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	Message      string `json:"message"`
}
