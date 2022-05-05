package util

import (
	"crypto/sha1"
	"fmt"
)

func GetHash(obj interface{}) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(fmt.Sprintf("%v", obj))))
}
