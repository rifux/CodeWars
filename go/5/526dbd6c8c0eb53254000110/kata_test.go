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

var _ = Describe("Alphanumeric Checker", func() {
	It("should validate alphanumeric strings", func() {
		// Existing tests
		Expect(alphanumeric(".*?")).To(BeFalse())
		Expect(alphanumeric("a")).To(BeTrue())
		Expect(alphanumeric("Mazinkaiser")).To(BeTrue())
		Expect(alphanumeric("hello world_")).To(BeFalse())
		Expect(alphanumeric("PassW0rd")).To(BeTrue())
		Expect(alphanumeric("     ")).To(BeFalse())
		Expect(alphanumeric("")).To(BeFalse())
		Expect(alphanumeric("\n\t\n")).To(BeFalse())
		Expect(alphanumeric("ciao\n$$_")).To(BeFalse())
		Expect(alphanumeric("__ * __")).To(BeFalse())
		Expect(alphanumeric("&)))(((")).To(BeFalse())
		Expect(alphanumeric("43534h56jmTHHF3k")).To(BeTrue())

		// Add new test cases here
		Expect(alphanumeric("123456")).To(BeTrue())    // Numeric string
		Expect(alphanumeric("abc123!")).To(BeFalse())  // Contains special char
		Expect(alphanumeric("Unicode©")).To(BeFalse()) // Unicode character
	})
})
