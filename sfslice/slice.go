package sfslice


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