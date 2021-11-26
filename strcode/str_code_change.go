package strcode

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
)

//GBK转utf8
func GbkToUtf8(gbkStr string) (string, error) {
	reader := transform.NewReader(bytes.NewReader([]byte(gbkStr)), simplifiedchinese.GBK.NewDecoder())
	utf8Str, err := ioutil.ReadAll(reader)
	if err != nil {
		return gbkStr, err
	}
	return string(utf8Str), nil
}

// UTF-8 转 GBK
func Utf8ToGbk(utf8Str string) (string, error) {
	reader := transform.NewReader(bytes.NewReader([]byte(utf8Str)), simplifiedchinese.GBK.NewEncoder())
	gbkStr, err := ioutil.ReadAll(reader)
	if err != nil {
		return utf8Str, err
	}
	return string(gbkStr), nil
}
