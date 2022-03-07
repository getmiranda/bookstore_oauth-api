package http

import (
	"net/http"

	"github.com/getmiranda/bookstore_oauth-api/src/services/access_token"

	"github.com/getmiranda/bookstore_oauth-api/src/utils/errors"

	atDomain "github.com/getmiranda/bookstore_oauth-api/src/domain/access_token"
	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(c *gin.Context)
}

type accessTokenHandler struct {
	atService access_token.AccessTokenService
}

func NewAccessTokenHandler(service access_token.AccessTokenService) AccessTokenHandler {
	return &accessTokenHandler{
		atService: service,
	}
}

func (h *accessTokenHandler) GetById(c *gin.Context) {
	accesToken, err := h.atService.GetById(c.Param("access_token_id"))
	if err != nil {
		err, _ := err.(errors.APIError)
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, accesToken)
}

// Create a new access token for a user
func (h *accessTokenHandler) Create(c *gin.Context) {
	var request atDomain.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	accessToken, err := h.atService.Create(&request)
	if err != nil {
		err, _ := err.(errors.APIError)
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, accessToken)
}
