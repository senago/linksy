package helpers

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	. "github.com/onsi/gomega"
)

var (
	client = &http.Client{}
)

const (
	apiEntry = "http://127.0.0.1:8080/api"
)

func Request(method, url string, request interface{}, adapters ...RequestAdapter) *http.Response {
	body, err := json.Marshal(request)
	Expect(err).NotTo(HaveOccurred())

	httpRequest, err := http.NewRequest(method, apiEntry+url, bytes.NewReader(body))
	Expect(err).NotTo(HaveOccurred())

	if method == http.MethodPost {
		httpRequest.Header.Add("Content-Type", "application/json")
	}

	AdaptRequest(httpRequest, adapters...)

	response, err := client.Do(httpRequest)
	Expect(err).NotTo(HaveOccurred())

	return response
}

type RequestAdapter func(*http.Request) *http.Request

func AdaptRequest(req *http.Request, adapters ...RequestAdapter) *http.Request {
	for _, adapter := range adapters {
		req = adapter(req)
	}
	return req
}

func WithQuery(params map[string]string) RequestAdapter {
	return func(req *http.Request) *http.Request {
		if params != nil {
			query := req.URL.Query()
			for key, value := range params {
				query.Add(key, value)
			}
			req.URL.RawQuery = query.Encode()
		}
		return req
	}
}

func ReadResponseBody(body io.ReadCloser) string {
	bytes, err := ioutil.ReadAll(body)
	Expect(err).NotTo(HaveOccurred())
	return string(bytes)
}

func WriteResponse(resp *http.Response, to interface{}) {
	err := json.NewDecoder(resp.Body).Decode(to)
	Expect(err).NotTo(HaveOccurred())
	Expect(resp.StatusCode).To(Equal(http.StatusOK))
}
