package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/yebology/travique-go/model"
)

func GenerateJwt(user model.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.Id,
		"email": user.Email,
		"name": user.Name,
	})

	signed, err := token.SignedString([]byte("secret-key"))
	if err != nil {
		return "", err
	}

	return signed, nil
	
}