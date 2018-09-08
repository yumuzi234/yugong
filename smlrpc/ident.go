package smlrpc

import (
	"strings"
	"unicode"
)

func lowerIdent(in string) (string, error) {
	if in == "" {
		return in, nil
	}
	runes := []rune(in)
	start := runes[0]
	if !unicode.IsUpper(start) && !unicode.IsLower(start) {
		return in, nil
	}
	var b strings.Builder
	// change the first letter to lower case
	cur := runes[0]
	_, err := b.WriteRune(unicode.ToLower(cur))
	if err != nil {
		return "", err
	}
	i := 1
	foundLower := unicode.IsLower(cur)
	for ; i < len(runes)-1 && !foundLower; i++ {
		cur := runes[i]
		next := runes[i+1]
		if unicode.IsLower(next) {
			foundLower = true
		} else {
			cur = unicode.ToLower(cur)
		}
		_, err := b.WriteRune(cur)
		if err != nil {
			return "", err
		}
	}

	for ; i < len(in); i++ {
		cur := runes[i]
		toWrite := cur
		if !foundLower {
			toWrite = unicode.ToLower(cur)
		}
		foundLower = unicode.IsLower(cur)
		_, err := b.WriteRune(toWrite)
		if err != nil {
			return "", err
		}
	}
	return b.String(), nil
}
