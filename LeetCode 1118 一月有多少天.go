package main

func numberOfDays(year int, month int) int {
	flag := false
	if year%400 == 0 || (year%100 != 0 && year%4 == 0) {
		flag = true
	}
	if month == 2 {
		if flag == false {
			return 28
		} else {
			return 29
		}
	} else if month == 1 || month == 3 || month == 5 || month == 7 || month == 8 || month == 10 || month == 12 {
		return 31
	} else {
		return 30
	}
}
