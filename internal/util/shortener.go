package util

import (
	"crypto/sha256"
	"math/big"
	"strconv"

	"github.com/google/uuid"
	"github.com/mr-tron/base58"
)

const (
	HASH_LENGTH = 10
)

func sha256Of(value string) []byte {
	hasher := sha256.New()
	hasher.Write([]byte(value))
	return hasher.Sum(nil)
}

func hash2uint64(hash []byte) uint64 {
	return new(big.Int).SetBytes(hash).Uint64()
}

func Shorten(url string) string {
	u := uuid.New()
	hash := sha256Of(url + u.String())
	encoded := base58.Encode([]byte(strconv.FormatUint(hash2uint64(hash), 10)))
	return encoded[:HASH_LENGTH]
}
