package homework

import "sort"

func SortMapValues(input map[int]string) (result []string) {
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
