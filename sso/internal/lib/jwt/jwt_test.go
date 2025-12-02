package jwt

import (
	"sso/internal/domain/models"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
)

func TestNewToken(t *testing.T) {
	user := models.User{
		ID:    123,
		Email: "test@example.com",
	}

	app := models.App{
		ID:     10,
		Secret: "super-secret-key",
	}

	duration := time.Hour

	tokenString, err := NewToken(user, app, duration)
	require.NoError(t, err)
	require.NotEmpty(t, tokenString)

	// Парсим токен для проверки
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Проверяем алгоритм подписи
		require.Equal(t, jwt.SigningMethodHS256, token.Method)
		return []byte(app.Secret), nil
	})

	require.NoError(t, err)
	require.True(t, parsedToken.Valid)

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	require.True(t, ok)

	// Проверяем поля
	require.Equal(t, float64(user.ID), claims["uid"])
	require.Equal(t, user.Email, claims["email"])
	require.Equal(t, float64(app.ID), claims["app_id"])

	// exp должен быть > текущего времени
	expFloat, ok := claims["exp"].(float64)
	require.True(t, ok)
	require.Greater(t, int64(expFloat), time.Now().Unix())
}
