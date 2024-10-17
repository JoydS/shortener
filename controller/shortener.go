package controller

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"shortener/database/manager"
	"shortener/database/model"
	"shortener/database/repository"
	"shortener/decoder"
	"shortener/encoder"
	"shortener/shortener"
)

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	originalURL, err := decoder.DecodeShortenRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	short, err := repository.FindShortenerFromOriginalURL(originalURL.String())
	if err != nil {
		slog.Error("failed to find shortener", "Error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
	if short == nil {
		slug, err := shortener.SlugURL(originalURL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		short = &model.Shortener{
			Slug:        slug,
			OriginalURL: originalURL.String(),
		}
		err = manager.CreateShortener(short)
		if err != nil {
			slog.Error("failed to create shortener", "Error", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}
	}

	shortenResponse := encoder.EncodeShortenResponse(os.Getenv("BASE_URL") + "/" + short.Slug)
	res, err := json.Marshal(shortenResponse)
	if err != nil {
		http.Error(w, "Unable to encode request", http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		slog.Error("failed to write response", "Error", err.Error())
	}
}

func ShortURLHandler(w http.ResponseWriter, r *http.Request) {
	slug := decoder.DecodeShortURLRequest(r)
	if slug == "" {
		http.Error(w, "this is not a valid short url", http.StatusBadRequest)

		return
	}

	short, err := repository.FindShortenerFromSlug(slug)
	if err != nil {
		slog.Error("failed to find shortener", "Error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
	if short == nil {
		http.Error(w, "this is not a valid short url", http.StatusNotFound)

		return
	}

	go func(short *model.Shortener) {
		err = repository.IncrementShortenerHitCount(short)
		if err != nil {
			slog.Error("failed to increment hit count", "Error", err.Error())
		}
	}(short)

	http.Redirect(w, r, short.OriginalURL, http.StatusMovedPermanently)
}
