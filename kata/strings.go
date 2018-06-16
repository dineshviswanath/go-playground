package kata

import (
	"strings"
)

func ConvertCamelCase(s string) string {
	var sep string
	if strings.ContainsAny(s, "-") {
		sep = "-"
	} else {
		sep = "_"
	}
	splited := strings.Split(s, sep)
	var out string
	for i, part := range splited {
		if i != 0 {
			out += strings.Title(part)
		} else {
			out += part
		}
	}
	return out
}
