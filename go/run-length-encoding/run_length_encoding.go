package encode

import (
	"strconv"
	"strings"
)

//func RunLengthEncode(input string) string {
//	if len(input) == 0 {
//		return ""
//	}
//	var result strings.Builder
//
//	var token struct {
//		char   byte
//		length int
//	}
//	for i := range input {
//		curr := input[i]
//		if i == 0 {
//			token.char = curr
//			token.length = 1
//			continue
//		}
//
//		if curr == token.char {
//			token.length++
//			continue
//		}
//
//		if token.length == 1 {
//			result.WriteByte(token.char)
//		} else {
//			result.WriteString(fmt.Sprint(token.length, string(token.char)))
//		}
//
//		token.char = curr
//		token.length = 1
//	}
//
//	if token.length == 1 {
//		result.WriteByte(token.char)
//	} else {
//		result.WriteString(fmt.Sprint(token.length, string(token.char)))
//	}
//
//	return result.String()
//}

func RunLengthEncode(s string) string {
	var encoded strings.Builder

	for len(s) > 0 {

		letter := s[0]
		slen := len(s)
		s = strings.TrimLeft(s, string(letter))

		if n := slen - len(s); n > 1 {
			encoded.WriteString(strconv.Itoa(n))
		}
		encoded.WriteByte(letter)
	}

	return encoded.String()

}

func RunLengthDecode(input string) string {
	if len(input) == 0 {
		return ""
	}

	var result strings.Builder
	var currLength strings.Builder
	for i, rune := range input {
		curr := input[i]
		if rune >= '0' && rune <= '9' {
			currLength.WriteByte(curr)
			continue
		}

		if currLength.Len() == 0 {
			result.WriteByte(curr)
			continue
		}

		length, _ := strconv.Atoi(currLength.String())
		result.WriteString(strings.Repeat(string(curr), length))

		currLength.Reset()
	}

	return result.String()
}
