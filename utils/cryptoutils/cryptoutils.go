package cryptoutils

import (
	"crypto/md5"
	"encoding/hex"
)

// GetMd5 - returns an md5 string hash of the passed in input
func GetMd5(input string) string {
	hash := md5.New()
	defer hash.Reset()

	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}
