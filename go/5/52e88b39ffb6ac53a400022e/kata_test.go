// TODO: replace with your own tests (TDD). An example to get you started is included below.
// Ginkgo BDD Testing Framework <http://onsi.github.io/ginkgo/>
// Gomega Matcher Library <http://onsi.github.io/gomega/>

package kata

import (
	"math"
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
		Expect(Int32ToIp(2154959208)).To(Equal("128.114.17.104"))
		Expect(Int32ToIp(2149583361)).To(Equal("128.32.10.1"))
		Expect(Int32ToIp(0)).To(Equal("0.0.0.0"))
		Expect(Int32ToIp(math.MaxUint32)).To(Equal("255.255.255.255"))

	})
})
