package util

import (
	"strings"
)

const (
	// String rune is the rune of the string
	StringRune = '"'
)

func detectString(args []string) (startIndex int, endIndex int) {
	str := strings.Join(args, " ")
	if strings.Count(str, string(StringRune)) == 2 {
		index := strings.Index(str, string(StringRune))
		secondIndex := strings.LastIndex(str, string(StringRune))
		if index != -1 {
			startIndex = index
		}
		if secondIndex != -1 {
			endIndex = secondIndex
		}

	}
	return
}
