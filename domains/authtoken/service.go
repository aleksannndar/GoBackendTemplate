package authtoken

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type IJWTService interface {
	GenerateJWT(userId string) (string, error)
	ValidateJWT(token string) (*Claims, error)
}

type JWTService struct {
	secretKey   []byte
	tokenIssuer string
}

func NewJWTService() IJWTService {
	return &JWTService{
		secretKey:   []byte("exampleSecret"),
		tokenIssuer: "TemplateIssue",
	}
}

func (s *JWTService) GenerateJWT(userId string) (string, error) {
	claims := NewClaims(userId, s.tokenIssuer)
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	signedToken, err := token.SignedString(s.secretKey)
	if err != nil {
		return "", fmt.Errorf("failed to sgin the token: %v", err)
	}

	return signedToken, nil
}

func (s *JWTService) ValidateJWT(signedToken string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(signedToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return s.secretKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
