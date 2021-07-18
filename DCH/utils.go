package DCH

import "strings"

func SplitToArguments(content string) []string {
	inQuotes := false
	var prevChar = '0'
	split := strings.FieldsFunc(content, func(r rune) bool {
		if r == '"' || r == '\'' {
			inQuotes = !inQuotes
		}
		prevChar = r
		return !inQuotes && r == ' '
	})

	for i, elem := range split {
		split[i] = strings.ReplaceAll(elem, "\"", "")
		split[i] = strings.ReplaceAll(elem, "'", "")
	}

	return split
}
