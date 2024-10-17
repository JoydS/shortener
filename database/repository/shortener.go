package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"shortener/database/model"
	"shortener/database/postgres"
)

func UpsertShortener(shortener *model.Shortener) error {
	db := postgres.GetGormDB()

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(shortener).Error; err != nil {
			return fmt.Errorf("failed to save shortener: %w", err)
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to upsert shortener: %w", err)
	}

	return nil
}

func IncrementShortenerHitCount(shortener *model.Shortener) error {
	db := postgres.GetGormDB()

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(shortener).Update("hit_count", gorm.Expr("hit_count + ?", 1)).Error; err != nil {
			return fmt.Errorf("failed to update hit count: %w", err)
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to increment hit count: %w", err)
	}

	return nil
}

func FindShortenerFromOriginalURL(originalURL string) (*model.Shortener, error) {
	db := postgres.GetGormDB()

	shortener := &model.Shortener{}

	if err := db.Where("original_url = ?", originalURL).First(shortener).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find shortener: %w", err)
	}

	return shortener, nil
}

func FindShortenerFromSlug(slug string) (*model.Shortener, error) {
	db := postgres.GetGormDB()

	shortener := &model.Shortener{}

	if err := db.Where("slug = ?", slug).First(shortener).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find shortener: %w", err)
	}

	return shortener, nil
}
