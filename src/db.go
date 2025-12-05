package src

import (
	"github.com/snghnaveen/workshop/config"
	"github.com/snghnaveen/workshop/models"
)

func GetFirstURL(url string) (string, bool) {
	var firstURL models.Urls
	err := config.DB.Where("long_url = ?", url).First(&firstURL).Error
	if err != nil || firstURL.LongURL != "" {
		return "", false
	}
	return firstURL.ShortURL, true
}

func CreateAndGetShortURL(url string) (string, error) {
	shortCode, err := CodeGen.Next()

	if err != nil {
		return "", err
	}
	var newFirstURL models.Urls
	newFirstURL.LongURL = url
	newFirstURL.ShortURL = shortCode

	if err := config.DB.Create(&newFirstURL).Error; err != nil {
		return "", err
	}
	return newFirstURL.ShortURL, nil
}

func ShortUrlToLongUrl(shortCode string) (string, error) {
	var firstURL models.Urls
	err := config.DB.Where("short_url = ?", shortCode).First(&firstURL).Error
	if err != nil {
		return "", err
	}
	return firstURL.LongURL, nil
}
