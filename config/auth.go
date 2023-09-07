package config

import (
	"errors"
	"os"
	"path"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

// ada dua fungsi yang harus dibikin
// 1. generate token
// 2. validasi token

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

func GetSecret() []byte {
	godotenv.Load(path.Join(os.Getenv("HOME"), "../.env"))
	var JWT_SECRET_KEY = []byte(os.Getenv("JWT_SECRET_KEY"))
	return JWT_SECRET_KEY

}

// memanggil dalam bentuk global

func (s *jwtService) GenerateToken(userID int) (string, error) {

	claim := jwt.MapClaims{}
	claim["user_id"] = userID
	claim["exp"] = time.Now().Add(time.Minute * 60).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(GetSecret())
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {

	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return GetSecret(), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
