package cities

import (
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"strings"
	"unicode"
)
import _ "golang.org/x/text/unicode/norm"

func NormalizeString(str string) string {
	transformer := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	normalized, _, _ := transform.String(transformer, strings.ToLower(str))
	return normalized
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}
