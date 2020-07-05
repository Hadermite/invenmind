package util

import "strings"

// IsAnyStringEmpty - Checks if any of the given strings is empty
func IsAnyStringEmpty(fields []string) bool {
	for _, field := range fields {
		if len(strings.TrimSpace(field)) == 0 {
			return true
		}
	}
	return false
}
