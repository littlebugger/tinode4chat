package auth

import (
	"github.com/littlebugger/tinode4chat/internal/service/entity"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// Secret key used for signing JWTs (in production, this should be in environment variables)
// TODO: move to env config
var jwtSecret = []byte("your_secret_key")

var (
	// Custom errors
	ErrMissingAuthHeader     = echo.NewHTTPError(http.StatusUnauthorized, "Missing or invalid Authorization header")
	ErrInvalidTokenFormat    = echo.NewHTTPError(http.StatusUnauthorized, "Invalid token format")
	ErrInvalidOrExpiredToken = echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired token")
	ErrInvalidTokenClaims    = echo.NewHTTPError(http.StatusUnauthorized, "Invalid token claims")
)

// TODO: make mockable

// JWTMiddleware validates JWT tokens for protected routes
func JWTMiddleware(c echo.Context) error {
	// Get the Authorization header
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return ErrMissingAuthHeader
	}

	// Ensure the token is in the format `Bearer <token>`
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader { // token didn't have Bearer prefix
		return ErrInvalidTokenFormat
	}

	// Parse and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidOrExpiredToken
		}
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return ErrInvalidOrExpiredToken
	}

	// Extract claims and attach to the context
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// You can extract specific claims here and add them to the context
		c.Set("userID", claims["user_id"])
		c.Set("email", claims["email"])
	} else {
		return ErrInvalidTokenClaims
	}

	return nil
}

// GenerateJWTToken generates a signed JWT token with user information
func GenerateJWTToken(user *entity.User) (string, error) {
	// Set token claims
	claims := &entity.JWTCustomClaims{
		UserID: user.ID,
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Token expires in 24 hours
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "tinode4chat",
		},
	}

	// Create the token using the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	return token.SignedString(jwtSecret)
}
