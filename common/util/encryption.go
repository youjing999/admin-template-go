package util

import (
	"crypto/md5"
	"encoding/hex"
)

func EncryptionMd5(s string) string {
	en := md5.New()
	en.Write([]byte(s))
	return hex.EncodeToString(en.Sum(nil))
}
