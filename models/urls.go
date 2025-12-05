package models

import "gorm.io/gorm"

type Urls struct {
	gorm.Model
	LongURL  string
	ShortURL string
}
