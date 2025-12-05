package src

import (
	"errors"
	"net/url"
	"strings"
)

func ValidateURL(input string) error {
	if input == "" {
		return errors.New("url is empty")
	}

	parsed, err := url.ParseRequestURI(input)
	if err != nil {
		return errors.New("invalid url format")
	}

	// Must have a scheme: http or https
	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return errors.New("url must start with http or https")
	}

	// Must have a domain/host
	if parsed.Host == "" {
		return errors.New("url missing host")
	}

	// Basic safety check
	if strings.Contains(parsed.Host, " ") {
		return errors.New("url contains spaces")
	}

	return nil
}
