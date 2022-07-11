package dto

type ShortenRequest struct {
	URL string `json:"url" validate:"required"`
}

type ShortenResponse struct {
	Hash string `json:"hash"`
}

type RetrieveRequest struct {
	Hash string `query:"hash" validate:"required"`
}

type RetrieveResponse struct {
	URL string `json:"url"`
}
