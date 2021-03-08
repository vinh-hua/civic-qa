package parse

import "strconv"

func ParseUintOrDefault(uintString string) uint {
	u, err := strconv.ParseUint(uintString, 10, 64)
	if err != nil {
		return 0
	}

	return uint(u)
}

func ParseBoolOrDefault(boolString string) bool {
	b, err := strconv.ParseBool(boolString)
	if err != nil {
		return false
	}

	return b
}
