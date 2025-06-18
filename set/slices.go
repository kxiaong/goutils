package set

// SliceSubtract 求slice的差集， a-b
func SliceSubtract[T comparable](a, b []T) []T {
	mapA := make(map[T]struct{}, len(a))
	mapB := make(map[T]struct{}, len(b))
	for _, e := range a {
		mapA[e] = struct{}{}
	}
	for _, e := range b {
		mapB[e] = struct{}{}
	}
	return MapSubtract[T, struct{}](mapA, mapB)
}

// SliceUnion 求slice的并集， a+b
func SliceUnion[T comparable](a, b []T) []T {
	mapA := make(map[T]struct{}, len(a))
	mapB := make(map[T]struct{}, len(b))
	for _, e := range a {
		mapA[e] = struct{}{}
	}
	for _, e := range b {
		mapB[e] = struct{}{}
	}
	return MapUnion[T, struct{}](mapA, mapB)
}

// SliceIntersect 求slice的交集， a&b
func SliceIntersect[T comparable](a, b []T) []T {
	mapA := make(map[T]struct{}, len(a))
	mapB := make(map[T]struct{}, len(b))
	for _, e := range a {
		mapA[e] = struct{}{}
	}
	for _, e := range b {
		mapB[e] = struct{}{}
	}
	return MapIntersect[T, struct{}](mapA, mapB)
}
