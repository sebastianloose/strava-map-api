package auth

import (
	"errors"
	"os"
	"strings"

	"github.com/sebastianloose/strava-map-api/model"

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

func GenerateToken(user model.User) (string, error) {
	claims := jwt.MapClaims{}

	claims["userId"] = (user.UserId)
	claims["expiresAt"] = user.ExpiresAt
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString(jwtKey)
}

func ExtractUserID(tokenString string) (int64, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return -1, errors.New("unexpected signing method")
		}
		return jwtKey, nil
	})

	if err != nil {
		return -1, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return -1, errors.New("jwt invalid")
	}

	userId := int64(claims["userId"].(float64))
	return userId, nil
}
