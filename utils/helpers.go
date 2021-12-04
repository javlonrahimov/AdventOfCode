package utils

import "strings"

func Trim_array(strs []string) []string {
	for i, str := range strs {
		strs[i] = strings.Trim(str, " ")
	}
	return strs
}
