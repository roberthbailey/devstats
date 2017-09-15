package gha2db

import (
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// StripUnicode strip non-unicode and control characters from a string
// From: https://rosettacode.org/wiki/Strip_control_codes_and_extended_characters_from_a_string#Go
func StripUnicode(str string) string {
	isOk := func(r rune) bool {
		return r < 32 || r >= 127
	}
	t := transform.Chain(norm.NFKD, transform.RemoveFunc(isOk))
	str, _, _ = transform.String(t, str)
	return str
}
