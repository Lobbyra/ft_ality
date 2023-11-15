package lib

func StringsToString(arr []string, _tempString string, sep string) string {
	if len(arr) > 0 {
		if _tempString == "" {
			return (StringsToString(arr[1:], arr[0], sep))
		} else {
			return (StringsToString(arr[1:], _tempString+sep+arr[0], sep))
		}
	}
	return (_tempString)
}
