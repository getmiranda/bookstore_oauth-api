package access_token

import (
	"strings"

	"github.com/getmiranda/bookstore_oauth-api/src/domain/access_token"
	"github.com/getmiranda/bookstore_oauth-api/src/repository/http"

	"github.com/getmiranda/bookstore_oauth-api/src/repository/db"
	"github.com/getmiranda/bookstore_oauth-api/src/utils/errors"
)

type AccessTokenService interface {
	GetById(string) (*access_token.AccessToken, error)
	Create(*access_token.AccessTokenRequest) (*access_token.AccessToken, error)
	UpdateExpirationTime(*access_token.AccessToken) error
}

type accessTokenService struct {
	dbRepo        db.DBRepository
	restUsersRepo http.HttpUsersRepository
}

func (s *accessTokenService) GetById(accessTokenId string) (*access_token.AccessToken, error) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if accessTokenId == "" {
		return nil, errors.NewBadRequestError("access token id cannot be empty")
	}
	return s.dbRepo.GetById(accessTokenId)
}

func (s *accessTokenService) Create(request *access_token.AccessTokenRequest) (*access_token.AccessToken, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	//TODO: Support both grant types: client_credentials and password

	// Authenticate the user against the Users API:
	user, err := s.restUsersRepo.LoginUser(request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	// Generate a new access token:
	at := access_token.GetNewAccessToken(user.Id)
	at.Generate()

	// Save the new access token in Cassandra:
	if err := s.dbRepo.Create(&at); err != nil {
		return nil, err
	}
	return &at, nil
}

func (s *accessTokenService) UpdateExpirationTime(at *access_token.AccessToken) error {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.dbRepo.UpdateExpirationTime(at)
}

func NewAccessTokenService(usersRepo http.HttpUsersRepository, dbRepo db.DBRepository) AccessTokenService {
	return &accessTokenService{
		restUsersRepo: usersRepo,
		dbRepo:        dbRepo,
	}
}
