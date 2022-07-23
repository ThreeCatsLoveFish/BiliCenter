package manager

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

func CryptoSign(data string, key string, rules []int) string {
	h0 := hmac.New(md5.New, []byte(key))
	h1 := hmac.New(sha1.New, []byte(key))
	h2 := hmac.New(sha256.New, []byte(key))
	h3 := hmac.New(sha256.New224, []byte(key))
	h4 := hmac.New(sha512.New, []byte(key))
	h5 := hmac.New(sha512.New384, []byte(key))

	for _, rule := range rules {
		switch rule {
		case 0:
			h0.Write([]byte(data))
			data = hex.EncodeToString(h0.Sum(nil))
		case 1:
			h1.Write([]byte(data))
			data = hex.EncodeToString(h1.Sum(nil))
		case 2:
			h2.Write([]byte(data))
			data = hex.EncodeToString(h2.Sum(nil))
		case 3:
			h3.Write([]byte(data))
			data = hex.EncodeToString(h3.Sum(nil))
		case 4:
			h4.Write([]byte(data))
			data = hex.EncodeToString(h4.Sum(nil))
		case 5:
			h5.Write([]byte(data))
			data = hex.EncodeToString(h5.Sum(nil))
		}
	}
	return data
}
