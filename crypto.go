package lib

import (
	"crypto/sha256"
	"fmt"
)

func GetSha256Hash(str string) string {
	h := sha256.New()
	bs := h.Sum([]byte(str))
	fmt.Printf("origin: %s, sha256 hash: %x\n", str, bs)
	return string(bs)
}
