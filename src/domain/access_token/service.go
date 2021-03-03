package access_token

import (
	"strings"

	"github.com/atg0831/msabookstore/bookstore_oauth-api/src/utils/errors"
)

type Service interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
}

type Repository interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
}
type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetByID(accessTokenID string) (*AccessToken, *errors.RestErr) {
	accessTokenID = strings.TrimSpace(accessTokenID)
	if len(accessTokenID) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}

	//s.repository =>type은 DbReposriotry(from db의 NewRepository()에서)
	accessToken, err := s.repository.GetByID(accessTokenID)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}
