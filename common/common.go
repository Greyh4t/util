package common

func Min(num ...int) int {
	min := num[0]
	for _, n := range num {
		if min > n {
			min = n
		}
	}
	return min
}

func Max(num ...int) int {
	max := num[0]
	for _, n := range num {
		if max < n {
			max = n
		}
	}
	return max
}
