package set

func Unique[K comparable](a []K) []K {
	res := make([]K, 0, len(a))
	m := make(map[K]struct{}, len(a))
	for _, e := range a {
		if _, ok := m[e]; !ok {
			m[e] = struct{}{}
			res = append(res, e)
		}
	}
	return res
}
