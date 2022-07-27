package utils

// Function to convert int to string
func IntsToString[num int | int32 | int64](number num) (res string) {
	// estimate the string length
	strLen := 0
	temp := number
	for temp > 0 {
		strLen++
		temp /= 10
	}
	// Create the result(in bytes)
	resByte := make([]byte, strLen)

	// Convert to string
	i := (strLen - 1)
	for number > 0 {
		resByte[i] = (48 + byte(number%10))
		i--
		number /= 10
	}
	res = string(resByte)

	return
}
