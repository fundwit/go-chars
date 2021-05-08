package chars_test

import (
	"github.com/fundwit/go-chars/chars"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Chars", func() {
	var _ = Describe("Abbreviate", func() {
		It("should be able to abbreviate works correctly", func() {
			Expect(chars.Abbreviate("abc def ghi", 3)).To(Equal("ADG"))
			Expect(chars.Abbreviate("abc def ghi", 2)).To(Equal("AD"))
			Expect(chars.Abbreviate("abc def ", 3)).To(Equal("ABD"))
			Expect(chars.Abbreviate("abc def ", 4)).To(Equal("ABCD"))
			Expect(chars.Abbreviate(" \tabc def ", 5)).To(Equal("ABCDE"))
			Expect(chars.Abbreviate("abc def ", 6)).To(Equal("ABCDEF"))
			Expect(chars.Abbreviate("abc def ", 7)).To(Equal("ABCDEF"))
			Expect(chars.Abbreviate("testProject def ghi", 3)).To(Equal("TPD"))
			Expect(chars.Abbreviate("testProject def", 4)).To(Equal("TEPD"))
			Expect(chars.Abbreviate("TTest222PRO100ject def ghi", 3)).To(Equal("TPD"))
			Expect(chars.Abbreviate("abc 开心 ghi", 3)).To(Equal("AKX"))
			Expect(chars.Abbreviate("abc 开心 ghi", 4)).To(Equal("AKXG"))
			Expect(chars.Abbreviate("abc bd开心88H ghi", 5)).To(Equal("ABKX8"))
		})
	})

	var _ = Describe("IsAlphabetLetter", func() {
		It("should be able to determine a rune is an alphabet letter correctly", func() {
			Expect(chars.IsAlphabetLetter('a')).To(BeTrue())
			Expect(chars.IsAlphabetLetter('z')).To(BeTrue())
			Expect(chars.IsAlphabetLetter('A')).To(BeTrue())
			Expect(chars.IsAlphabetLetter('Z')).To(BeTrue())
			Expect(chars.IsAlphabetLetter('1')).To(BeFalse())
			Expect(chars.IsAlphabetLetter(' ')).To(BeFalse())
			Expect(chars.IsAlphabetLetter('好')).To(BeFalse())
		})
	})
})
