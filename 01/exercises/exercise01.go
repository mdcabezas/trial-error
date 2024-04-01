package exercises

import "strconv"

func StringToInt(s string) (int, string) {

	i, err := strconv.Atoi(s)

	if err != nil {
		return 0, "Error" + err.Error()
	}

	if i > 100 {
		return i, "It is greater than 100"
	} else {
		return i, "It is less than 100"
	}

	/*
		// First attempt

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
	*/
}
