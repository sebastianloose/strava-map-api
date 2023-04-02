package auth

import (
	"errors"
	"os"
	"strings"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

func GetTokenFromRequest(c *gin.Context) string {
	token := c.Request.Header.Get("Authorization")
	tokenParts := strings.Split(token, " ")
	if len(tokenParts) == 2 {
		return tokenParts[1]
	}
	return ""
}

func GenerateToken(userId uuid.UUID, expiresAt int) (string, error) {
	claims := jwt.MapClaims{}

	claims["userId"] = userId
	claims["expiresAt"] = expiresAt
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString(jwtKey)
}

func ExtractUserID(tokenString string) (uuid.UUID, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return uuid.UUID{}, errors.New("unexpected signing method")
		}
		return jwtKey, nil
	})

	if err != nil {
		return uuid.UUID{}, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return uuid.UUID{}, errors.New("jwt invalid")
	}

	userId, _ := uuid.Parse(claims["userId"].(string))
	return userId, nil
}
