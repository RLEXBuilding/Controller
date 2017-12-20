package util

import (
	"unicode"
	"strings"
)

type IllegalEscapeSequence struct {
	er string
}

func (er IllegalEscapeSequence) Error() string {
	return "Illegal Escape Sequence: " + er.er
}

const (
	none     = 0
	inQuotes = 1
	inString = 2
)

var (
	escape = map[rune]rune{
		'n': '\n',
		'r': '\r',
		'a': '\a',
		'b': '\b',
		'f': '\f',
	}
)

func ParseQuotes(txt string) (result []string, finished bool, error error) {
	arr := make([]string, 0)
	mode := none
	runes := []rune(strings.TrimSpace(txt))
	var buf []rune = nil
	for i := 0; i < len(runes); i++ {
		ch := runes[i]
		if mode == none {
			if unicode.IsSpace(ch) {
				continue
			} else if ch == '"' {
				mode = inQuotes
				buf = make([]rune, 0)
			} else {
				mode = inString
				buf = make([]rune, 0)
				i--
			}
		} else if ch == '\\' {
			if i == len(runes)-1 {
				return nil, false, nil // Unfinished escape sequence
			}
			i++
			nxtch := runes[i]
			if nxtch == 'x' { // Hex literal in the form of \x0A
				// TODO
			}
			instead, ok := escape[nxtch]
			if !ok {
				return nil, false, IllegalEscapeSequence{string(nxtch)}
			}
			buf = append(buf, instead)
		} else if mode == inQuotes && ch == '"' {
			mode = none
			arr = append(arr, string(buf))
			buf = nil
		} else if mode == inString && unicode.IsSpace(ch) {
			mode = none
			arr = append(arr, string(buf))
			buf = nil
		} else {
			buf = append(buf, ch)
		}
	}
	if mode == inQuotes {
		return nil, false, nil // Unfinished quotes. Continue parsing with more inputs
	}
	if mode == inString && buf != nil {
		arr = append(arr, string(buf))
	}
	return arr, true, nil
}

const (
	// String rune is the rune of the string
	StringRune = '"'
)

func DetectString(args []string) (startIndex int, endIndex int) {
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
