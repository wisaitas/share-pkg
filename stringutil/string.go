package stringutil

import (
	"strings"
	"unicode"
)

func ToScreamingSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) && i > 0 {
			prev := rune(s[i-1])
			if unicode.IsLower(prev) || (i+1 < len(s) && unicode.IsLower(rune(s[i+1]))) {
				result.WriteRune('_')
			}
		}
		result.WriteRune(unicode.ToUpper(r))
	}
	return result.String()
}

func ToSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) && i > 0 {
			prev := rune(s[i-1])
			if unicode.IsLower(prev) || (i+1 < len(s) && unicode.IsLower(rune(s[i+1]))) {
				result.WriteRune('_')
			}
		}
		result.WriteRune(unicode.ToLower(r))
	}
	return result.String()
}
