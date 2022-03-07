package config

import "os"

const (
	usersApiHostRepository = "USERS_API_HOST_REPOSITORY"
)

var (
	secretUsersApiHostRepository = os.Getenv(usersApiHostRepository)
)

func GetUsersApiHostRepository() string {
	return secretUsersApiHostRepository
}
