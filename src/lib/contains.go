package lib

func Contains(toFind string, arr []string) bool {
	if len(arr) > 0 {
		if arr[0] == toFind {
			return (true)
		}
		return (Contains(toFind, arr[1:]))
	}
	return (false)
}
