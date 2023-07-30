package slice

const (
	ReduceFactor = 0.625
)

func Delete[T any](src []T, index int) ([]T, T, error) {
	if index < 0 || index >= len(src) {
		var zero T
		return nil, zero, ErrOutOfRange(len(src), index)
	}
	var res = src[index]
	for i := index + 1; i < len(src); i++ {
		src[i-1] = src[i]
	}
	//TODO: shrink here?
	src = src[:len(src)-1]
	return src, res, nil
}

func Shrink[T any](src []T) []T {
	var c, l = cap(src), len(src)
	if c <= 64 || c == l || c/l == 1 {
		return src
	}

	if c < 2048 && c/l >= 4 {
		return reduceSlice(src, c/2)
	}
	if c > 2048 && c/l >= 2 {
		return reduceSlice(src, int(float32(c)*ReduceFactor))
	}
	return src
}

func reduceSlice[T any](src []T, length int) []T {
	var res = make([]T, 0, length)
	copy(res, src)
	return res
}
