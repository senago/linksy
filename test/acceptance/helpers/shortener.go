package helpers

import (
	"net/http"

	"github.com/senago/linksy/internal/model/dto"
)

func ShortenURL(url string) *http.Response {
	httpResponse := Request(
		http.MethodPost,
		"/shorten",
		dto.ShortenRequest{
			URL: url,
		},
	)
	return httpResponse
}

func RetrieveURL(hash string) *http.Response {
	httpResponse := Request(
		http.MethodGet,
		"/retrieve",
		dto.ShortenRequest{},
		WithQuery(map[string]string{"hash": hash}),
	)
	return httpResponse
}
