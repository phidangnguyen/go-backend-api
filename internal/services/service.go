package services

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Skip jwt authentication for these below api
var SkipJWTAuth = []string{
	"user/login",
}

type BaseService struct {
	Claim    *JWTCustomClaim
	seretKey string
}

type JWTCustomClaim struct {
	Email string `json:"email"`
	Id    int64  `json:"id"`
	jwt.StandardClaims
}

func (s *BaseService) GenerateToken(id int64, email string) (string, error) {

	claims := &JWTCustomClaim{
		email,
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encode token
	t, err := token.SignedString([]byte(s.seretKey))

	if err != nil {
		return "", err
	}

	return t, nil
}

func (s *BaseService) SetSecretKey(key string) {
	s.seretKey = key
}
