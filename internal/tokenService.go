package internal

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
	"todo-list/protogen/token"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

type TokenServer struct {
	token.UnimplementedTokenServer
}

func NewTokenServer() TokenServer {
	return TokenServer{}
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading env file %s", err.Error())
	}
}

func (t *TokenServer) CreateToken(
	ctx context.Context,
	userId *token.UserId,
) (*token.TokenString, error) {
	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId.GetUserId(),
		"nbf":    time.Now().Unix(),
		"exp":    time.Now().Add(time.Minute * 10).Unix(),
		"iat":    time.Now().Unix(),
	})

	tokenString, err := tkn.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return nil, err
	}

	return &token.TokenString{Token: tokenString}, nil
}

func (t *TokenServer) VerifyToken(
	ctx context.Context,
	tokenString *token.TokenString,
) (*token.UserId, error) {
	tkn, err := jwt.Parse(tokenString.GetToken(), func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, errors.New("invalid token")
	}

	claims, ok := tkn.Claims.(jwt.MapClaims)

	if !ok {
		return nil, errors.New("invalid claims")
	}

	userId := claims["userId"].(float64)

	return &token.UserId{UserId: int64(userId)}, nil
}
