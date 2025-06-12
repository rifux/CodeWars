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

func dotest(a, b, margin float64, expected int) {
	Expect(CloseCompare(a, b, margin)).To(Equal(expected), "With a = %f, b = %f, margin = %f", a, b, margin)
}

var _ = Describe("Tests", func() {
	It("Sample tests", func() {
		dotest(4.0, 5.0, 0.0, -1)
		dotest(5.0, 5.0, 0.0, 0)
		dotest(6.0, 5.0, 0.0, 1)
		dotest(2.0, 5.0, 3.0, 0)
		dotest(5.0, 5.0, 3.0, 0)
		dotest(8.0, 5.0, 3.0, 0)
		dotest(8.1, 5.0, 3.0, 1)
		dotest(1.99, 5.0, 3.0, -1)
	})
})
