// TODO: replace with your own tests (TDD). An example to get you started is included below.
// Ginkgo BDD Testing Framework <http://onsi.github.io/ginkgo/>
// Gomega Matcher Library <http://onsi.github.io/gomega/>

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
	It("Sample tests", func() {
		Expect(Order("is2 Thi1s T4est 3a")).To(Equal("Thi1s is2 3a T4est"))
		Expect(Order("4of Fo1r pe6ople g3ood th5e the2")).To(Equal("Fo1r the2 g3ood 4of th5e pe6ople"))
		Expect(Order("")).To(Equal(""))
		Expect(Order("3 6 4 2 8 7 5 1 9")).To(Equal("1 2 3 4 5 6 7 8 9"))
	})
})
