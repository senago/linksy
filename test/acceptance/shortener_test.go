package acceptance

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/senago/linksy/internal/model/dto"
	"github.com/senago/linksy/test/acceptance/helpers"
)

var _ = Describe("Testing url shortening", func() {
	defer GinkgoRecover()
	When("sending a url", func() {
		httpResponse := helpers.ShortenURL("google.com")
		shortenResponse := &dto.ShortenResponse{}
		helpers.WriteResponse(httpResponse, shortenResponse)

		It("should return a string of length 10", func() {
			Expect(shortenResponse.Hash).To(HaveLen(10))
		})
	})

	When("sending same url twice", func() {
		httpResponse1 := helpers.ShortenURL("google.com")
		shortenResponse1 := &dto.ShortenResponse{}
		helpers.WriteResponse(httpResponse1, shortenResponse1)

		httpResponse2 := helpers.ShortenURL("google.com")
		shortenResponse2 := &dto.ShortenResponse{}
		helpers.WriteResponse(httpResponse2, shortenResponse2)

		It("should return different values", func() {
			Expect(shortenResponse1.Hash).NotTo(Equal(shortenResponse2.Hash))
		})
	})
})

var _ = Describe("Testing url retrieving", func() {
	defer GinkgoRecover()
	When("trying an invalid value", func() {
		httpResponse := helpers.RetrieveURL("42")
		It("should return 404", func() {
			Expect(httpResponse.StatusCode).To(Equal(http.StatusNotFound))
		})
	})

	When("trying a valid value", func() {
		url := "google.com"

		httpResponse := helpers.ShortenURL(url)
		shortenResponse := &dto.ShortenResponse{}
		helpers.WriteResponse(httpResponse, shortenResponse)

		httpResponse = helpers.RetrieveURL(shortenResponse.Hash)
		retrieveResponse := &dto.RetrieveResponse{}
		helpers.WriteResponse(httpResponse, retrieveResponse)

		It("should return the original url", func() {
			Expect(retrieveResponse.URL).To(Equal(url))
		})
	})
})
