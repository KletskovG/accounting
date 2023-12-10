package services

import (
	"net/http"
	"net/url"

	"github.com/kletskovg/packages/common"
)

// Notify sends GET request to notification service. message - text, which have to be sent
func Notify(message string) (*http.Response, error) {
	return http.Get(common.Hosts().TelegramAPIURL + "/done/" + url.PathEscape(message))
}
