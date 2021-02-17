package iteration

import "strings"

// Repeat accepts a character and repeats it
func Repeat(character string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}

// RepeatWithBuiltIn accepts a character and repeats it
func RepeatWithBuiltIn(character string, repeatCount int) string {
	return strings.Repeat(character, repeatCount)
}
