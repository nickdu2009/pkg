package xstring

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"math/rand"
)

func RandString(n int, runes []rune) string {
	if runes == nil {
		runes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	}
	b := make([]rune, n)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}

func Md5WithSalt(s string, salt string) string {
	h := md5.Sum([]byte(s + salt))
	return hex.EncodeToString(h[:])
}

func Md5(s string) string {
	return Md5WithSalt(s, "")
}

func HmacSha1(input string, key string) string {
	h := hmac.New(sha1.New, []byte(key))
	h.Write([]byte(input))
	s := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return s
}