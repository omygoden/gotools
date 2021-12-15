package encrypts

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
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

//生成hash密码
func HashPwdGenerate(pwd string) string {
	hashPwd, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)

	return string(hashPwd)
}

//验证码hash密码
func HashPwdVerify(hashPwd string, pwd string) bool {
	res := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(pwd))
	if res != nil {
		return false
	}
	return true
}
