package helpers

import (
	"fmt"
	"strings"
)

func PrintSliceByComma[T uint | string](slice []T) string {
	var buf strings.Builder
	if len(slice) == 0 {
		return "'0'"
	}
	buf.WriteString(fmt.Sprintf("'%v'", slice[0]))
	for _, el := range slice[1:] {
		buf.WriteString(fmt.Sprintf(", '%v'", el))
	}
	return buf.String()
}
