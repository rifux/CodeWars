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

var _ = Describe("Example Tests", func() {
	It("TestBob27Male", func() {
		dm := NewDinglemouse().SetName("Bob").SetAge(27).SetSex('M')
		expected := "Hello. My name is Bob. I am 27. I am male."
		Expect(dm.Hello()).To(Equal(expected))
	})

	It("Test27MaleBob", func() {
		dm := NewDinglemouse().SetAge(27).SetSex('M').SetName("Bob")
		expected := "Hello. I am 27. I am male. My name is Bob."
		Expect(dm.Hello()).To(Equal(expected))
	})

	It("TestAliceFemale", func() {
		dm := NewDinglemouse().SetName("Alice").SetSex('F')
		expected := "Hello. My name is Alice. I am female."
		Expect(dm.Hello()).To(Equal(expected))
	})

	It("TestBatman", func() {
		dm := NewDinglemouse().SetName("Batman")
		expected := "Hello. My name is Batman."
		Expect(dm.Hello()).To(Equal(expected))
	})
})
