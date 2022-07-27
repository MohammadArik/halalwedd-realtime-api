package utils

func BubbleSort(slice []int32) {
	var temp int32

	for i := len(slice); i > 1; i-- {
		for j := 1; j < i; j++ {
			if slice[j] < slice[j-1] {
				temp = slice[j]
				slice[j] = slice[(j - 1)]
				slice[j-1] = temp
			}
		}

	}
}
