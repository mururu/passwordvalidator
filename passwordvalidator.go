package passwordvalidator

import (
	"regexp"
	"strings"
)

func CommonPassword(password string) bool {
	_, ok := commonPasswords[password]
	return !ok
}

var space = regexp.MustCompile(`\W+`)

const maxSimilarity = 0.7

func Similarity(password string, attrs ...string) bool {
	lowerPassword := strings.ToLower(password)
	for _, attr := range attrs {
		parts := space.Split(strings.ToLower(attr), -1)
		parts = append(parts, attr)
		for _, part := range parts {
			if sequenceMatchRatio(lowerPassword, part) >= maxSimilarity {
				return false
			}
		}
	}
	return true
}

func sequenceMatchRatio(a string, b string) float64 {
	fullbcount := map[byte]int{}
	for _, c := range []byte(b) {
		fullbcount[c] = fullbcount[c] + 1
	}
	avail := map[byte]int{}
	matches := 0
	for _, c := range []byte(a) {
		n, ok := avail[c]
		if !ok {
			n = fullbcount[c]
		}
		avail[c] = n - 1
		if n > 0 {
			matches += 1
		}
	}
	return calculateRatio(matches, len(a)+len(b))
}

func calculateRatio(matches int, length int) float64 {
	if length > 0 {
		return 2.0 * float64(matches) / float64(length)
	}
	return 1.0
}
