package homework

func reverse(input []int64) (result []int64) {
	for j := len(input) - 1; j >= 0; j-- {
		result = append(result, input[j])
	}
	return
}
