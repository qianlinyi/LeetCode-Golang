package main

func replaceDigits(s string) string {
	res := []byte(s)
	for i := 1; i < len(res); i += 2 {
		res[i] = res[i-1] + res[i] - '0'
	}
	return string(res)
}
