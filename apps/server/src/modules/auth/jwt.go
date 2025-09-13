package auth

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"sodnix/apps/server/src/config"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Use env or config value in real-world apps (avoid hardcoding keys in code)
var jwtKey = []byte(config.JWT_SECRET_KEY)

type Claims struct {
	UserID      uuid.UUID `json:"user_id"`
	Email       string    `json:"email"`
	DisplayName string    `json:"display_name"`
	jwt.RegisteredClaims
}

func generateAccessToken(userID uuid.UUID, email, displayName string) (string, error) {
	accessTokenExpirationMin, err := strconv.Atoi(config.ACCESS_TOKEN_EXPIRATION)
	if err != nil {
		return "", err
	}

	now := time.Now()
	accessClaims := Claims{
		UserID:      userID,
		Email:       email,
		DisplayName: displayName,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(accessTokenExpirationMin) * time.Minute)),
		},
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, err := at.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func generateRefreshToken(userID uuid.UUID, email, displayName string) (string, error) {
	refreshTokenExpirationDay, err := strconv.Atoi(config.REFRESH_TOKEN_EXPIRATION)
	if err != nil {
		return "", err
	}

	now := time.Now()
	refreshClaims := Claims{
		UserID:      userID,
		Email:       email,
		DisplayName: displayName,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(refreshTokenExpirationDay) * 24 * time.Hour)),
		},
	}

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err := rt.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return refreshToken, nil
}

func GenerateJWT(userID uuid.UUID, email, displayName string) (accessToken string, refreshToken string, err error) {
	accessToken, err = generateAccessToken(userID, email, displayName)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = generateRefreshToken(userID, email, displayName)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func ParseJWT(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {

		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtKey, nil
	})
	fmt.Println(err)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
