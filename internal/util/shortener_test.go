package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashLength(t *testing.T) {
	url := "google.com"
	hash := Shorten(url)
	assert.Equal(t, len(hash), 10)
}

func TestHashNondeterministic(t *testing.T) {
	url := "google.com"
	hash1 := Shorten(url)
	hash2 := Shorten(url)
	assert.NotEqual(t, hash1, hash2)
}
