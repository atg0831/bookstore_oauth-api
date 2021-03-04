package access_token

import (
	"fmt"
	"strings"

	"github.com/atg0831/msabookstore/bookstore_oauth-api/src/utils/errors"
)

type Service interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpriationTime(AccessToken) *errors.RestErr
}

type Repository interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpriationTime(AccessToken) *errors.RestErr
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
		fmt.Println("service.go에서 db error")
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(at AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}

	return s.repository.Create(at)
}

func (s *service) UpdateExpriationTime(at AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}

	return s.repository.UpdateExpriationTime(at)
}
