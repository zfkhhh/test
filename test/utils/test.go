package utils

import "strings"

func TestStr(str string) bool {
	index := strings.LastIndex(str, "/")

	if index != -1 {
		return false
	}else {
		return true
	}
}