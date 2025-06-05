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

var _ = Describe("FindUniq", func() {
	It("should woaaaaaaark for some basic cases", func() {
		Expect(FindUniq([]float32{1.0, 1.0, 1.0, 2.0, 1.0, 1.0})).To(Equal(float32(2)))
		Expect(FindUniq([]float32{0, 0, 0.55, 0, 0})).To(Equal(float32(0.55)))
	})
})
