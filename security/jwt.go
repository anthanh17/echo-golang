package security

import (
	"goEcho/model"

	"github.com/dgrijalva/jwt-go"
)

const SECRET_KEY = "ant_echo_golang"

func GenToken(user model.User) (string, error) {
	claims := &model.JwtCustomClaims{
		UserId:         user.UserId,
		Role:           user.Role,
		StandardClaims: jwt.StandardClaims{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	result, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}
	return result, nil
}
