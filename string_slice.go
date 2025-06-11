package slices

import (
	"math/rand"
	"time"
)

// StringInSlice 检查slice是否包含target
func StringInSlice(target string, slice []string) bool {
	for _, s := range slice {
		if target == s {
			return true
		}
	}
	return false
}

// StringSliceTo2D 一维数组转二维
// @param s 原数组
// @param arrLen 二维数组每一个元素的长度
// @return [][]int64 二维数组 arrLen * (len(s)/arrLen + 1)
func StringSliceTo2D(s []string, arrLen int) [][]string {
	if arrLen <= 0 {
		return nil
	}
	if len(s) <= arrLen {
		return [][]string{s}
	}
	res := make([][]string, 0, arrLen)
	tmp := make([]string, 0, arrLen)
	for i, str := range s {
		tmp = append(tmp, str)
		if (i+1)%arrLen == 0 {
			res = append(res, tmp)
			tmp = make([]string, 0, arrLen)
		}
	}
	if len(tmp) > 0 {
		res = append(res, tmp)
	}
	return res
}

// RandomSelectFromStringSlice 从source中随机选择cap个元素，返回一个新的slice
func RandomSelectFromStringSlice(source []string, cap int) []string {
	if cap > len(source) {
		return source
	}
	rdr := rand.New(rand.NewSource(time.Now().UnixNano()))

	selectedMap := make(map[string]struct{}, cap)
	selected := make([]string, 0, cap)
	for len(selected) < cap {
		idx := rdr.Intn(len(source))
		str := source[idx]
		if _, ok := selectedMap[str]; !ok {
			selectedMap[str] = struct{}{}
			selected = append(selected, str)
		}
	}
	return selected
}
