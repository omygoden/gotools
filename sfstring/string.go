package sfstring

import (
	"fmt"
	"github.com/omygoden/gotools/sfconst"
	"github.com/omygoden/gotools/sfrand"
	"math"
	"strings"
	"sync"
	"time"
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

//生成订单号
//支持万级并发不重复
var UniqueNoMap = make(map[string]int)
var UniqueNoMapMute sync.Mutex

func GenerateUniqueNo(prefixs ...string) string {
	UniqueNoMapMute.Lock()
	defer UniqueNoMapMute.Unlock()
	var prefix string
	if len(prefixs) > 0 {
		prefix = prefixs[0]
	}
	m := time.Now().UnixMicro() - time.Now().Unix()*1000000
	s := prefix + time.Now().Format(sfconst.GO_TIME_WITH_FULL) + fmt.Sprintf("%06d", m) + fmt.Sprintf("%d", sfrand.RandRange(1000, 9999))
	if _, ok := UniqueNoMap[s]; ok {
		return GenerateUniqueNo(prefixs...)
	}
	return s
}
