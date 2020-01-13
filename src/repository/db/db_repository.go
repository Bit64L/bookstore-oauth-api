package db

import (
	"bookstore-oauth-api/src/domain/access_token"
	"bookstore-oauth-api/src/utils/errors"
)

func New() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct{}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, errors.NewInternalServerError("db connection not be implemented!")
}
