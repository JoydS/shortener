package decoder

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
	"net/url"
)

type ShortenRequest struct {
	URL string `json:"original_url"`
}

func DecodeShortenRequest(r *http.Request) (*url.URL, error) {
	shortenRequest := new(ShortenRequest)
	err := json.NewDecoder(r.Body).Decode(shortenRequest)
	if err != nil {
		slog.Error("failed to decode shorten request", "Error", err.Error())

		return nil, err
	}

	if shortenRequest.URL == "" {
		slog.Error("url is empty")

		return nil, err
	}

	originURL, err := url.ParseRequestURI(shortenRequest.URL)
	if err != nil {
		slog.Error("failed to parse url", "Error", err.Error())

		return nil, errors.New("invalid url")
	}

	return originURL, nil
}

func DecodeShortURLRequest(r *http.Request) string {
	slug := chi.URLParam(r, "slug")

	return slug
}
