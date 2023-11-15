package lib

func Foreach[T any](arr []T, f func(T)) {
	if len(arr) > 0 {
		f(arr[0])
		Foreach[T](arr[1:], f)
	}
}

func DumbPrint[T any](v T) {
	println(v)
}
