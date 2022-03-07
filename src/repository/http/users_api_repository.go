package http

import (
	"fmt"

	"github.com/getmiranda/bookstore_oauth-api/src/clients/http/users_api_http"
	"github.com/getmiranda/bookstore_oauth-api/src/utils/errors"

	"github.com/getmiranda/bookstore_oauth-api/src/domain/users"
)

type HttpUsersRepository interface {
	LoginUser(string, string) (*users.User, error)
}

type usersRepository struct{}

func NewHttpUsersRepository() HttpUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, error) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	response, err := users_api_http.Client.Post("/users/login", request)
	if err != nil {
		fmt.Println(err)
		return nil, errors.NewInternalServerError("invalid restclient response when trying to login user")
	}

	if response.StatusCode > 299 {
		apiErr, err := errors.NewErrorFromBytes(response.Bytes())
		if err != nil {
			return nil, errors.NewInternalServerError("invalid error interface when trying to login user")
		}
		return nil, apiErr
	}

	var user users.User
	if err := response.UnmarshalJson(&user); err != nil {
		return nil, errors.NewInternalServerError("error when trying to unmarshal users login response")
	}
	return &user, nil
}
