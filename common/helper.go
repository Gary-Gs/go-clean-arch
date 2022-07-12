package common

import "strings"

func ContainsIgnoreCase(arr []string, target string) bool {
	target = strings.ToUpper(target)
	for _, v := range arr {
		if strings.ToUpper(v) == target {
			return true
		}
	}
	return false
}
