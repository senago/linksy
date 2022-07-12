package util

import (
	"crypto/md5"

	"github.com/google/uuid"
	"github.com/mr-tron/base58"
)

const (
	HASH_LENGTH = 10
)

func hashOf(value string) []byte {
	hasher := md5.New()
	hasher.Write([]byte(value))
	return hasher.Sum(nil)
}

func Shorten(url string) string {
	u := uuid.New()
	hash := hashOf(url + u.String())
	encoded := base58.Encode(hash)
	return encoded[:HASH_LENGTH]
}
