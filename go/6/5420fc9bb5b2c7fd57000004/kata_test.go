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

var _ = Describe("Tests", func() {
	Describe("Sample tests", func() {
		It("Sample test 1: 12, 10, 8, 12, 7, 6, 4, 10, 12", func() {
			Expect(HighestRank([]int{12, 10, 8, 12, 7, 6, 4, 10, 12})).To(Equal(12))
		})
	})
})
