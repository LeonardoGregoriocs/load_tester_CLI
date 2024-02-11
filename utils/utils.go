package utils

import (
	"strconv"
)

func CheckFlags(flags ...string) bool {
	for _, flag := range flags {
		if flag == "" {
			return false
		}
	}
	return true
}

func ConvertNumber(flagString string) int {
	newFlag, _ := strconv.Atoi(flagString)
	return newFlag
}
