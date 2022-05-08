package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	indices := []int{}
	for key := range input {
		indices = append(indices, key)
	}
	sort.Ints(indices)
	for index := range indices {
		result = append(result, input[index])
	}
	return
}
