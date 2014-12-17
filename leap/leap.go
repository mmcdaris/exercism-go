package leap

// leap years meet the following requirements:
//   Divisible by 4 AND NOT by 100...
//   UNLESS also Divisible by 400(century leap year)
func IsLeapYear(year int) (leap bool) {
	if year%4 == 0 {
		leap = (year%100 != 0) || (year%400 == 0)
	}
	return leap
}
