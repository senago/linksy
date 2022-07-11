package dto

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	Hash string `json:"hash"`
}

type RetrieveRequest struct {
	Hash string `query:"hash"`
}

type RetrieveResponse struct {
	URL string `json:"url"`
}
