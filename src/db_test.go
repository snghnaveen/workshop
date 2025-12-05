package src

import (
	"os"
	"testing"

	"github.com/snghnaveen/workshop/config"
	"github.com/snghnaveen/workshop/models"
	"github.com/stretchr/testify/assert"
)

func TestShortUrlToLongUrl(t *testing.T) {
	os.Setenv("DB_NAME", "test_workshop")

	config.ConnectDatabase()

	// Create test record
	var newFirstURL models.Urls
	newFirstURL.LongURL = "http://google.com"
	newFirstURL.ShortURL = "1234"
	assert.NoError(t, config.DB.Create(&newFirstURL).Error)

	// test case
}
