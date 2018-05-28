package gitpls

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	wordSep         = " "
	truncatedPrefix = "..."
	truncatedSuffix = "..."
)

var (
	normalizer = regexp.MustCompile(`\s`)
)

type MessageTruncator struct {
	MaxLength int
}

func (m *MessageTruncator) doTruncate(src string, keyword string) string {
	idx := strings.Index(strings.ToLower(src), strings.ToLower(keyword))
	if idx == -1 {
		panic(fmt.Errorf("'%s' not found in source string", keyword))
	}

	left := strings.Split(strings.TrimSpace(src[:idx]), " ")
	right := strings.Split(strings.TrimSpace(src[idx+len(keyword):len(src)]), " ")

	result := keyword

	pickLeft := true

	for {
		// Keywords left?
		if len(left) == 0 && len(right) == 0 {
			break
		}

		if len(left) > 0 && (len(right) == 0 || pickLeft) {
			var prefix string
			prefix, left = left[len(left)-1], left[:len(left)-1]

			if len(truncatedPrefix)+len(prefix)+len(wordSep)+len(result)+len(truncatedSuffix) > m.MaxLength {
				break
			}

			result = prefix + wordSep + result
		} else {
			var suffix string
			suffix, right = right[0], right[1:]

			if len(truncatedPrefix)+len(result)+len(wordSep)+len(suffix)+len(truncatedSuffix) > m.MaxLength {
				break
			}

			result += wordSep + suffix
		}

		pickLeft = !pickLeft
	}

	if len(left) > 0 {
		result = truncatedPrefix + result
	}

	if len(right) > 0 {
		result += truncatedSuffix
	}

	return result
}

func (m *MessageTruncator) Truncate(message *string, match string) string {
	s := normalizer.ReplaceAllString(strings.TrimSpace(*message), " ")

	// Do we even need to truncate?
	if len(s) < m.MaxLength {
		return s
	} else {
		return m.doTruncate(s, match)
	}
}
