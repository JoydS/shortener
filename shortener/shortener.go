package shortener

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"net/url"
)

func SlugURL(u *url.URL) (string, error) {
	if u == nil {
		return "", errors.New("URL is nil")
	}
	h := sha512.New()
	h.Write([]byte(u.String()))

	hashBytes := h.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)

	// Return the first 10 characters of the hash
	return hashString[:10], nil
}
