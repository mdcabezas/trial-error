package exercises

import "strconv"

func StringToInt(s string) (int, string) {
	var rI int
	var rS string
	i, _ := strconv.Atoi(s)

	if i > 100 {
		rI = i
		rS = "It is greater than 100"
	} else {
		rS = "It is less than 100"
	}
	return rI, rS
}
