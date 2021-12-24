package utils

type AnyType interface{}

func Map[T, S](items []T, action func(t T) S) []S {
	size := len(items)
	var array = make([]S, size)
	for i, item := range items {
		array[i] = action(item)
	}
	return array
}

func Reduce[T](items []T, actions func(T, T) T) T {
	pre := items[0]
	for i := 1; i < len(items); i++ {
		pre = actions(pre, items[i])
	}
	return pre
}

func Filter[T](items []T, action func(T) bool) []T {
	var array = make([]T, len(items))
	index := 0
	for _, item := range items {
		if action(item) {
			array[index] = item
			index++
		}
	}
	return array[:index]
}

func Any[T](items []T, action func(T) bool) bool {
	for _, item := range items {
		if action(item) {
			return true
		}
	}
	return false
}

func ForEach[T](items []T, action func(T)) {
	for _, item := range items {
		action(item)
	}
}
