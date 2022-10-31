package piglatin

import "strings"

func Sentence(sentence string) string {

	pieces := strings.Split(sentence, " ")
	for i, word := range pieces {
		pieces[i] = magic(word)
	}

	return strings.Join(pieces, " ")
}

func magic(sentence string) string {
	first := sentence[0]

	switch first {
	case 'a':
		fallthrough
	case 'e':
		fallthrough
	case 'i':
		fallthrough
	case 'o':
		fallthrough
	case 'u':
		return sentence + "ay"
	}

	if sentence[0] == 'y' && (sentence[1] == 'a' || sentence[1] == 'e' || sentence[1] == 'i' || sentence[1] == 'o' || sentence[1] == 'u') {
		return sentence[1:] + "yay"
	}

	if sentence[0] == 'y' {
		return sentence + "ay"
	}

	if sentence[1] == 'y' && len(sentence) == 2 {
		return string(sentence[1]) + string(first) + "ay"
	}

	if sentence[0:2] == "xr" {
		return sentence + "ay"
	}

	if sentence[0:2] == "qu" {
		return sentence[2:] + sentence[0:2] + "ay"
	}

	if sentence[1:3] == "qu" {
		return sentence[3:] + sentence[0:3] + "ay"
	}

	for i, c := range sentence {
		if c != 'a' && c != 'e' && c != 'i' && c != 'o' && c != 'u' && c != 'y' {
			continue
		}
		return sentence[i:] + sentence[:i] + "ay"
	}

	return sentence[1:] + string(first) + "ay"
}
