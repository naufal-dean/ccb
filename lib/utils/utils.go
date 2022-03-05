package utils

import (
	"crypto/sha256"
	"fmt"
)

func GetMapKeys(m map[string]bool) []string {
	res := make([]string, len(m))
	i := 0
	for k := range m {
		res[i] = k
		i++
	}
	return res
}

func Sha256(o interface{}) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", o)))
	return fmt.Sprintf("%x", h.Sum(nil))
}
