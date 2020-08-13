package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/ryukazari/twinstgo/models"
)

// GenerateJWT generate a token for the user session
func GenerateJWT(user models.User) (string, error) {
	privateKey := []byte("ryukazari_twinstgo")
	payload := jwt.MapClaims{
		"email":     user.Email,
		"name":      user.Name,
		"lastName":  user.LastName,
		"birthday":  user.Birthday,
		"ubication": user.Ubication,
		"webSite":   user.WebSite,
		"_id":       user.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 20).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(privateKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
