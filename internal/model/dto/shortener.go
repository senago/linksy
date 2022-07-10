package dto

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ID string `json:"id"`
}

type RetrieveRequest struct {
	ID string `query:"id"`
}

type RetrieveResponse struct {
	URL string `json:"url"`
}
