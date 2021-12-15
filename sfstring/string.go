package sfstring

import (
	"fmt"
	"math"
	"strings"
)

//判断字符串是否包含切片中的其中一个
func StringsMulContains(s string, substr []string) bool {
	for _, v := range substr {
		if strings.Contains(s, v) {
			return true
		}
	}
	return false
}

func Letters(num int) []string {
	letter := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	if num <= 26 {
		return letter[0:num]
	} else {
		var res = make([]string, 26)
		page := int(math.Ceil(float64(num) / 26))
		allPage := letter[:page]
		copy(res, letter)
		surplus := num - 26
		for k, v := range allPage {
			ii := surplus - (k * 26)
			if ii > 26 {
				ii = 26
			}
			for i := 0; i < ii; i++ {
				res = append(res, fmt.Sprintf("%s%s", v, letter[i]))
			}
		}
		return res
	}
}
