package leap

const testVersion = 3 

func IsLeapYear(leap int) bool {
	return leap%4 == 0 && (leap%100 != 0 || leap%400 == 0)
}

