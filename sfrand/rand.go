package sfrand

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"math/big"
	"strconv"
	"time"
)

//随机生成 [start,end] 内的整数
func RandRange(start,end int64) int64 {
	//伪随机，如果时间一样，返回的值都一样
	//r := rand.New(rand.NewSource(time.Now().Unix()))
	//return r.Intn(end - start + 1) + start

	//真随机，性能比伪随机差了10倍
	res,_ := rand.Int(rand.Reader,big.NewInt(end - start + 1))
	return res.Int64() + start
}

//随机返回一个唯一的md5格式的字符串--根据纳秒生成
func RandMd5Str() string {
	dataMd5 := md5.New()
	t := time.Now().UnixNano()
	s := strconv.FormatInt(t,10)
	dataMd5.Write([]byte(s))
	return hex.EncodeToString(dataMd5.Sum(nil))
}
