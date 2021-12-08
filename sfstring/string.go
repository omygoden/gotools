package sfstring

import "strings"

//判断字符串是否包含切片中的其中一个
func StringsMulContains(s string, substr []string) bool {
	for _, v := range substr {
		if strings.Contains(s, v) {
			return true
		}
	}
	return false
}
