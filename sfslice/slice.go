package sfslice

import (
	"strings"
)

//判断字符串是否包含切片中的其中一个
func StringsMulContains(s string,substr []string) bool {
	for _,v := range substr {
		if strings.Contains(s,v) {
			return true
		}
	}
	return false
}

//map切片根据map里的某个字段排序，并返回字符串类型的切片
func SliceMapSortByField(smap []map[string]interface{},sortField string,valueField string) []string {
	var res = make([]string,len(smap))
	for i:=0;i<len(smap);i++ {
		for j:=0;j<(len(smap)-i-1);j++ {
			if smap[j][sortField].(int) > smap[j+1][sortField].(int) {
				smap[j],smap[j+1] = smap[j+1],smap[j]
			}
		}
	}
	for k,v := range smap {
		res[k] = v[valueField].(string)
	}
	return res
}


//判断s切片是否包含i变量
func SliceContainInt(iList []int,i int) bool {
	for _,v := range iList {
		if v == i {
			return true
		}
	}
	return false
}

//判断s切片是否包含i变量
func SliceContainString(strList []string,s string) bool {
	for _,v := range strList {
		if v == s {
			return true
		}
	}
	return false
}