package util

import "strconv"

func ToInt(data string) int {
	val, _ := strconv.Atoi(data)
	return val
}
