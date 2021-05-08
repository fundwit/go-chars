package chars_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGoChars(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Chars Suite")
}
