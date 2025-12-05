package src

import (
	"errors"
	"sync"
)

var CodeGen *ShortCodeGenerator

func init() {
	CodeGen = NewShortCodeGenerator()
}

var alphabet = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var base = uint64(len(alphabet))

type ShortCodeGenerator struct {
	mu   sync.Mutex
	last uint64 // numeric counter (0,1,2,...)
}

func NewShortCodeGenerator() *ShortCodeGenerator {
	return &ShortCodeGenerator{
		last: 0,
	}
}

// Next returns the next shortcode in incremental Base62 format
// MaxLength = 7 chars → max value = 62^7 - 1
func (g *ShortCodeGenerator) Next() (string, error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// max possible value for 7 chars
	maxVal := uint64(1)
	for i := 0; i < 7; i++ {
		maxVal *= base
	}
	maxVal -= 1

	if g.last > maxVal {
		return "", errors.New("shortcode limit exceeded")
	}

	// convert integer counter → base62 string
	code := encodeBase62(g.last)
	g.last++

	return code, nil
}

// converts a number to base62 string (1–7 chars)
func encodeBase62(num uint64) string {
	if num == 0 {
		return "0" // shortest possible
	}

	buf := make([]byte, 0, 7)
	for num > 0 {
		remainder := num % base
		buf = append(buf, alphabet[remainder])
		num /= base
	}

	// reverse
	for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}

	return string(buf)
}
