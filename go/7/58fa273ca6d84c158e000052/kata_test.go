package kata

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestKata(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Kata Suite")
}

func dotest(n uint64, expected int) {
	Expect(Digits(n)).To(Equal(expected), "With n = %d", n)
}

var _ = Describe("Tests", func() {
	It("Sample tests", func() {
		dotest(5, 1)
		dotest(12345, 5)
		dotest(9876543210, 10)
	})
})
