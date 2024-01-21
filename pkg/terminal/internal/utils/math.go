package utils

const MaxInt = int(^uint(0) >> 1)

type Number interface {
	int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | float32 | float64
}

func Min[T Number](arr ...T) T {
	if len(arr) == 0 {
		var zero T

		return zero
	}

	min := arr[0]

	for _, v := range arr {
		if v < min {
			min = v
		}
	}

	return min
}
