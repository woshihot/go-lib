package str

func Contains(v []string, s string) bool {
	result := false
	for _, value := range v {
		if value == s {
			result = true
		}
	}
	return result
}

func AppendSet(arr []string, appends ...string) (newArr []string) {

	return HashSet(append(arr, appends...))
}

func HashSet(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}
