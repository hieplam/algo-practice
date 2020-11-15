package general

//ProductSum Tip: Each element of `array` can either be cast
// to `SpecialArray` or `int`.
func ProductSum(array []interface{}) int {
	return CalProductSum(array, 1)
}

//CalProductSum Calculate product sum of the array
func CalProductSum(arr []interface{}, depth int) int {
	sum := 0
	for _, v := range arr {
		if cast, ok := v.(int); ok {
			sum += cast
		} else if cast, ok := v.([]interface{}); ok {
			sum += CalProductSum(cast, depth+1)
		}
	}

	return sum * depth
}
