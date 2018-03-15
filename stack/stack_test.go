package stack_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "learn/stack"
)

var _ = Describe("Stack", func() {
	s := new(Stack)

	BeforeEach(func() {
		s.Push(3)
	})

	Describe("with push some int", func() {
		Context("test the pop func", func() {
			It("shoud pop 2", func() {
				p := s.Pop()
				Expect(p).To(Equal(3))
			})

		})
	})
})
