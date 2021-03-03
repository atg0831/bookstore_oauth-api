package db

import (
	"github.com/atg0831/msabookstore/bookstore_oauth-api/src/domain/access_token"
	"github.com/atg0831/msabookstore/bookstore_oauth-api/src/utils/errors"
)

type DbReposriotry interface {
	GetByID(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func NewRepository() DbReposriotry {
	return &dbRepository{}
}

func (r *dbRepository) GetByID(id string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, errors.NewInternalServerError("database connection not implemented yet")

}
