package src

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateURL(t *testing.T) {
	validUrl1 := "http://google.com"
	assert.NoError(t, ValidateURL(validUrl1))

	invalidUrl1 := "fake"
	assert.Error(t, ValidateURL(invalidUrl1))
}
