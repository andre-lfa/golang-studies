// Package bob should respond with the correct string based on a input.
package bob

import (
	"strings"
	"unicode"
)

// Hey should respond with the correct string based on a input.
func Hey(remark string) string {

	remark = strings.TrimSpace(remark)

	if isAskingAndIsUpper(remark) {
		return "Calm down, I know what I'm doing!"
	} else if isAQuestion(remark) {
		return "Sure."
	} else if isUpper(remark) {
		return "Whoa, chill out!"
	} else if remark == "" {
		return "Fine. Be that way!"
	} else {
		return "Whatever."
	}
}

func isAQuestion(srt string) bool {
	return strings.HasSuffix(srt, "?")
}

func isUpper(str string) bool {
	isLetters := strings.IndexFunc(str, unicode.IsLetter) >= 0
	isUpper := strings.ToUpper(str) == str

	return isLetters && isUpper
}

func isAskingAndIsUpper(str string) bool {
	return isAQuestion(str) && isUpper(str)
}
