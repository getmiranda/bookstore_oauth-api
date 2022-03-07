package db

import (
	"log"

	"github.com/getmiranda/bookstore_oauth-api/src/clients/cassandra"
	"github.com/gocql/gocql"

	"github.com/getmiranda/bookstore_oauth-api/src/domain/access_token"
	"github.com/getmiranda/bookstore_oauth-api/src/utils/errors"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?);"
	queryUpdateExpires     = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)

type DBRepository interface {
	GetById(string) (*access_token.AccessToken, error)
	Create(*access_token.AccessToken) error
	UpdateExpirationTime(*access_token.AccessToken) error
}

type dbRepository struct{}

func NewDBRepository() DBRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, error) {
	var result access_token.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.ClientId,
		&result.Expires,
	); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError("no access token found with given id")
		}
		log.Println(err)
		return nil, errors.NewInternalServerError("error while fetching access token")
	}

	return &result, nil
}

func (r *dbRepository) Create(at *access_token.AccessToken) error {
	if err := cassandra.GetSession().Query(queryCreateAccessToken,
		at.AccessToken,
		at.UserId,
		at.ClientId,
		at.Expires,
	).Exec(); err != nil {
		return errors.NewInternalServerError("error while creating access token")
	}
	return nil
}

func (r *dbRepository) UpdateExpirationTime(at *access_token.AccessToken) error {
	if err := cassandra.GetSession().Query(queryUpdateExpires,
		at.Expires,
		at.AccessToken,
	).Exec(); err != nil {
		return errors.NewInternalServerError("error when trying to update current resource")
	}
	return nil
}
