package strings

import "strings"

// ContainsString check if a slice contains a string
func ContainsString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

// RemoveString remove string from a sliced
func RemoveString(slice []string, s string) (result []string) {
	for _, item := range slice {
		if item == s {
			continue
		}
		result = append(result, item)
	}
	return
}

func StripNewLines(s string) string {
	return strings.TrimSuffix(s, "\n")
}

func ReturnNonEmptyString(value string) string {
	if len(value) == 0 {
		return "-"
	}
	return value
}
