package mock

import (
	"context"
	"errors"
	"token-management-service/protogen/token"

	"google.golang.org/grpc"
)

type MockTokenClient struct {
	Err ErrMock
}

const (
	_ ErrMock = iota
	ErrInTokenCreation
	ErrInTokenVerification
	StatusOkInTokenService
)

func NewMockTokenClient() MockTokenClient {
	return MockTokenClient{}
}

func (t *MockTokenClient) CreateToken(
	ctx context.Context,
	userId *token.UserId,
	opts ...grpc.CallOption,
) (*token.TokenString, error) {
	if t.Err == ErrInTokenCreation {
		return nil, errors.New("error in token creation")
	}

	return &token.TokenString{Token: "Token"}, nil
}

func (t *MockTokenClient) VerifyToken(
	ctx context.Context,
	tokenString *token.TokenString,
	opts ...grpc.CallOption,
) (*token.UserId, error) {
	if t.Err == ErrInTokenVerification {
		return nil, errors.New("error in token verification")
	}
	return &token.UserId{UserId: 1}, nil
}
