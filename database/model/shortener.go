package model

type Shortener struct {
	ID          int64  `json:"id" gorm:"primary_key"`
	Slug        string `json:"slug" gorm:"unique"`
	OriginalURL string `json:"original_url" gorm:"unique"`
	HitCount    int64  `json:"hit_count" gorm:"default:0"`
}

func (Shortener) TableName() string {
	return "shortener"
}
