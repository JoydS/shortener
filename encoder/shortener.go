package encoder

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

func EncodeShortenResponse(shortURL string) *ShortenResponse {
	return &ShortenResponse{
		ShortURL: shortURL,
	}
}
