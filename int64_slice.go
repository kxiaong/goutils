package slices

// Int64InSlice 检查slice是否包含target
func Int64InSlice(target int64, s []int64) bool {
	for _, e := range s {
		if target == e {
			return true
		}
	}
	return false
}

// Int64SliceTo2D 一维数组转二维
// @param s 原数组
// @param arrLen 二维数组每一个元素的长度
// @return [][]int64 二维数组 arrLen * (len(s)/arrLen + 1)
func Int64SliceTo2D(s []int64, arrLen int) [][]int64 {
	res := make([][]int64, 0, arrLen)
	tmp := make([]int64, 0, arrLen)
	for i, e := range s {
		tmp = append(tmp, e)
		if (i+1)%arrLen == 0 {
			res = append(res, tmp)
			tmp = make([]int64, 0, arrLen)
		}
	}
	if len(tmp) > 0 {
		res = append(res, tmp)
	}
	return res
}
