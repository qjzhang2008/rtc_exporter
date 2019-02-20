package utils

import (
	"strings"
)

func SliceWords(line []byte) []string {

	str := string(line)
	words := strings.Split(str, ",")
	return words
}
