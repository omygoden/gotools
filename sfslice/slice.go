package sfslice

import (
	"sort"
	"strconv"
)

//map切片根据map里的某个字段排序，并返回字符串类型的切片
func SliceMapSortByField(smap []map[string]interface{}, sortField string, valueField string) []string {
	var res = make([]string, len(smap))
	for i := 0; i < len(smap); i++ {
		for j := 0; j < (len(smap) - i - 1); j++ {
			if smap[j][sortField].(int) > smap[j+1][sortField].(int) {
				smap[j], smap[j+1] = smap[j+1], smap[j]
			}
		}
	}
	for k, v := range smap {
		res[k] = v[valueField].(string)
	}
	return res
}

//判断s切片是否包含i变量
func SliceContainInt(iList []int, i int) bool {
	for _, v := range iList {
		if v == i {
			return true
		}
	}
	return false
}

//判断s切片是否包含i变量
func SliceContainString(strList []string, s string) bool {
	for _, v := range strList {
		if v == s {
			return true
		}
	}
	return false
}

//字符串切片转in64切片
func SliceStrToSliceInt64(s []string) []int64 {
	i := make([]int64, len(s))
	for k, v := range s {
		ii, _ := strconv.Atoi(v)
		i[k] = int64(ii)
	}
	return i
}

//切片正序
func SliceSort(slice interface{}) {
	switch slice.(type) {
	case []string:
		sort.Strings(slice.([]string))
	case []int:
		sort.Ints(slice.([]int))
	case []float64:
		sort.Float64s(slice.([]float64))
	}
}

//切片倒序
func SliceSortReverse(slice interface{}) {
	switch slice.(type) {
	case []string:
		sort.Sort(sort.Reverse(sort.StringSlice(slice.([]string))))
	case []int:
		sort.Sort(sort.Reverse(sort.IntSlice(slice.([]int))))
	case []float64:
		sort.Sort(sort.Reverse(sort.Float64Slice(slice.([]float64))))
	}
}
