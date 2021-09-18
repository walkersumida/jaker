package upper

import "strings"

func First(str string) string {
	top := str[:1]
	others := str[1:]
	return strings.ToUpper(top) + others
}
