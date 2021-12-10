package encrypts

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

//md5加密
func Md5(p interface{}) string {
	hash := md5.New()
	switch p.(type) {
	case int:
		hash.Write([]byte(strconv.Itoa(p.(int))))
	case string:
		hash.Write([]byte(p.(string)))
	default:
		panic("该类型暂不支持")
	}

	return hex.EncodeToString(hash.Sum(nil))
}

//sha1加密
func Sha1(p interface{}) string {
	hash := sha1.New()
	switch p.(type) {
	case int:
		hash.Write([]byte(strconv.Itoa(p.(int))))
	case string:
		hash.Write([]byte(p.(string)))
	default:
		panic("该类型暂不支持")
	}

	return hex.EncodeToString(hash.Sum(nil))
}

//sha1加密
func Sha256(p interface{}) string {
	hash := sha256.New()
	switch p.(type) {
	case int:
		hash.Write([]byte(strconv.Itoa(p.(int))))
	case string:
		hash.Write([]byte(p.(string)))
	default:
		panic("该类型暂不支持")
	}

	return hex.EncodeToString(hash.Sum(nil))
}
