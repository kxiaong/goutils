package set

// MapSubtract 求map key的差集， a-b
func MapSubtract[K comparable, V1, V2 any](a map[K]V1, b map[K]V2) []K {
	res := make([]K, 0, len(a)+len(b))
	for k, _ := range a {
		if _, ok := b[k]; !ok {
			res = append(res, k)
		}
	}
	return res
}

// MapUnion 求map key的并集， a+b
func MapUnion[K comparable, V1, V2 any](a map[K]V1, b map[K]V2) []K {
	res := make([]K, 0, len(a)+len(b))
	for k, _ := range a {
		res = append(res, k)
	}
	for k, _ := range b {
		if _, ok := a[k]; ok {
			continue
		}
		res = append(res, k)
	}

	return res
}

// MapIntersect 求map key的交集， a&b
func MapIntersect[K comparable, V1, V2 any](a map[K]V1, b map[K]V2) []K {
	res := make([]K, 0, len(a)+len(b))
	for k, _ := range a {
		if _, ok := b[k]; ok {
			res = append(res, k)
		}
	}
	return res
}
