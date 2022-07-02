package core_test

import (
	. "go-ddd-practice/cmd/snackmachine/core"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Money", func() {
	It("Should add correctly", func() {
		m1 := NewMoney(1, 2, 3, 4, 5, 6)
		m2 := NewMoney(1, 2, 3, 4, 5, 6)

		result := m1.Add(m2)
		Expect(result.OneCentCount).Should(Equal(2))
		Expect(result.TenCentCount).Should(Equal(4))
		Expect(result.QuarterCount).Should(Equal(6))
		Expect(result.OneDollarCount).Should(Equal(8))
		Expect(result.FiveDollarCount).Should(Equal(10))
		Expect(result.TwentyDollarCount).Should(Equal(12))
	})

	It("Should subtract correctly", func() {
		m1 := NewMoney(10, 10, 10, 10, 10, 10)
		m2 := NewMoney(1, 2, 3, 4, 5, 6)

		result := m1.Subtract(m2)
		Expect(result.OneCentCount).Should(Equal(9))
		Expect(result.TenCentCount).Should(Equal(8))
		Expect(result.QuarterCount).Should(Equal(7))
		Expect(result.OneDollarCount).Should(Equal(6))
		Expect(result.FiveDollarCount).Should(Equal(5))
		Expect(result.TwentyDollarCount).Should(Equal(4))
	})

	It("Should equal if contain the same amount", func() {
		m1 := NewMoney(1, 2, 3, 4, 5, 6)
		m2 := NewMoney(1, 2, 3, 4, 5, 6)

		Expect(m1.Equal(m2)).Should(BeTrue())
	})

	It("Should not equal if contain the different amount", func() {
		m1 := NewMoney(1, 2, 3, 4, 5, 6)
		m2 := NewMoney(1, 0, 3, 4, 5, 6)

		Expect(m1.Equal(m2)).Should(BeFalse())
	})

	DescribeTable("should return correct amount",
		func(money Money, expected float64) {
			Expect(money.Amount()).Should(Equal(expected))
		},
		Entry("0", NewMoney(0, 0, 0, 0, 0, 0), float64(0)),
		Entry(".01", NewMoney(1, 0, 0, 0, 0, 0), float64(0.01)),
		Entry(".1", NewMoney(0, 1, 0, 0, 0, 0), float64(.1)),
		Entry(".25", NewMoney(0, 0, 1, 0, 0, 0), float64(.25)),
		Entry("1", NewMoney(0, 0, 0, 1, 0, 0), float64(1)),
		Entry("5", NewMoney(0, 0, 0, 0, 1, 0), float64(5)),
		Entry("20", NewMoney(0, 0, 0, 0, 0, 1), float64(20)),
	)
})
