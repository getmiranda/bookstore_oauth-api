package users_api_http

import (
	"net/http"
	"time"

	"github.com/getmiranda/bookstore_oauth-api/src/config"

	"github.com/getmiranda/go-httpclient/gohttp"
	"github.com/getmiranda/go-httpclient/gomime"
)

var (
	Client = getHttpClient()
)

func getHttpClient() gohttp.Client {
	headers := make(http.Header)
	headers.Set(gomime.HeaderContentType, gomime.ContentTypeJson)

	client := gohttp.NewBuilder().
		SetHeaders(headers).
		SetBaseUrl(config.GetUsersApiHostRepository()).
		SetConnectionTimeout(2 * time.Second).
		SetResponseTimeout(3 * time.Second).
		SetUserAgent("oauth-api").
		Build()
	return client
}
