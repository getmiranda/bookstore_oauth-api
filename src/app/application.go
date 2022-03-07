package app

import (
	"github.com/getmiranda/bookstore_oauth-api/src/repository/db"
	httpUsers "github.com/getmiranda/bookstore_oauth-api/src/repository/http"
	"github.com/gin-gonic/gin"

	"github.com/getmiranda/bookstore_oauth-api/src/http"
	"github.com/getmiranda/bookstore_oauth-api/src/services/access_token"
)

var router = gin.Default()

func StartApplication() {
	atHandler := http.NewAccessTokenHandler(
		access_token.NewAccessTokenService(httpUsers.NewHttpUsersRepository(), db.NewDBRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8081")
}
