package chars

import (
	"github.com/mozillazg/go-pinyin"
	"strings"
	"unicode"
)

func Abbreviate(input string, width int) string {
	// word split
	works := [][]rune{}
	work := []rune{}
	lowerCase := false
	for _, r := range input {
		if unicode.IsSpace(r) && len(work) > 0 {
			works = append(works, work)
			work = []rune{}
		} else {
			if IsUpperAlphabetLetter(r) && lowerCase {
				if len(work) > 0 {
					works = append(works, work)
					work = []rune{}
				}
				lowerCase = false
			}
			if !IsUpperAlphabetLetter(r) && !lowerCase {
				lowerCase = true
			}
			work = append(work, r)
		}
	}
	if len(work) > 0 {
		works = append(works, work)
	}

	// characters normalize
	normalizedWords := [][]rune{}
	normalizedWord := []rune{}
	arg := pinyin.NewArgs()
	for _, w := range works {
		for _, r := range w {
			if r >= '0' && r <= '9' || r >= 'A' && r <= 'Z' || r >= 'a' && r <= 'z' {
				normalizedWord = append(normalizedWord, r)
			} else {
				py := pinyin.SinglePinyin(r, arg)
				if len(py) > 0 {
					if len(normalizedWord) > 0 {
						normalizedWords = append(normalizedWords, normalizedWord)
					}
					normalizedWord = []rune{}
					for _, p := range py[0] {
						normalizedWord = append(normalizedWord, p)
					}
					normalizedWords = append(normalizedWords, normalizedWord)
					normalizedWord = []rune{}
				}
			}
		}
		if len(normalizedWord) > 0 {
			normalizedWords = append(normalizedWords, normalizedWord)
			normalizedWord = []rune{}
		}
	}

	// Abbreviate
	abbreviation := []rune{}
	gap := width - len(normalizedWords)
	for _, result := range normalizedWords {
		num := len(result)
		abbreviation = append(abbreviation, result[0])
		if gap > 0 {
			if num-1 >= gap {
				for i := 1; i <= gap; i++ {
					abbreviation = append(abbreviation, result[i])
				}
				gap = 0
			} else {
				for i := 1; i < num; i++ {
					abbreviation = append(abbreviation, result[i])
				}
				gap = gap - num + 1
			}
		}
		if len(abbreviation) == width {
			break
		}
	}
	return strings.ToUpper(string(abbreviation))
}

func IsUpperAlphabetLetter(r rune) bool {
	return r >= 'A' && r <= 'Z'
}
func IsLowerAlphabetLetter(r rune) bool {
	return r >= 'a' && r <= 'z'
}
func IsAlphabetLetter(r rune) bool {
	return IsLowerAlphabetLetter(r) || IsUpperAlphabetLetter(r)
}
