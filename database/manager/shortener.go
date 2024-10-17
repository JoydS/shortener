package manager

import (
	"shortener/database/model"
	"shortener/database/repository"
)

func CreateShortener(s *model.Shortener) error {
	return UpdateShortener(s)
}

func UpdateShortener(s *model.Shortener) error {
	return repository.UpsertShortener(s)
}
