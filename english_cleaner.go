package cleaner

import (
	"strings"
)

type EnglishCleaner struct {
}

func (c *EnglishCleaner) removeSpace(input string) string {
	// re := regexp.MustCompile(`([A-Za-z0-9\-])\s+([A-Za-z0-9\-])`)
	// normalizedInput := toHalfWidth(input)
	// return re.ReplaceAllString(normalizedInput, `$1$2`)
	utf8Bytes := convertToUTF8(input)
	return string(utf8Bytes)
}

func toHalfWidth(input string) string {
	var builder strings.Builder
	for _, char := range input {
		if char >= 'Ａ' && char <= 'Ｚ' {
			builder.WriteRune(char - 'Ａ' + 'A')
		} else if char >= 'ａ' && char <= 'ｚ' {
			builder.WriteRune(char - 'ａ' + 'a')
		} else if char >= '０' && char <= '９' {
			builder.WriteRune(char - '０' + '0')
		} else if char == '－' {
			builder.WriteRune('-')
		} else {
			builder.WriteRune(char)
		}
	}
	return builder.String()
}

func convertToUTF8(input string) []byte {
	halfWidth := toHalfWidth(input)
	return []byte(halfWidth)
}
