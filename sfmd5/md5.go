package sfmd5

import (
	"crypto/md5"
	"encoding/hex"
)

//获取md5的值
func Md5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}
