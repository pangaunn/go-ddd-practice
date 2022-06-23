package model_test

import (
	"go-ddd-practice/cmd/snackmachine/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Money", func() {
	DescribeTable("should return correct amount",
		func(money model.Money, expected float64) {
			Expect(money.Amount()).Should(Equal(expected))
		},
		Entry(model.NewMoney(0, 0, 0, 0, 0, 0), 0),
		Entry(model.NewMoney(1, 0, 0, 0, 0, 0), 10),
		Entry(model.NewMoney(0, 1, 0, 0, 0, 0), 10),
		Entry(model.NewMoney(0, 0, 1, 0, 0, 0), 10),
		Entry(model.NewMoney(0, 0, 0, 1, 0, 0), 10),
		Entry(model.NewMoney(0, 0, 0, 0, 1, 0), 10),
		Entry(model.NewMoney(0, 0, 0, 0, 0, 1), 10),
	)

	When("Add", func() {
		It("produce correct result", func() {
			m1 := model.NewMoney(1, 2, 3, 4, 5, 6)
			m2 := model.NewMoney(1, 2, 3, 4, 5, 6)
			m3 := m1.Add(m2)

			Expect(m3.OneCentCount).Should(Equal(2))
			Expect(m3.TenCentCount).Should(Equal(4))
			Expect(m3.QuarterCount).Should(Equal(6))
			Expect(m3.OneDollarCount).Should(Equal(8))
			Expect(m3.FiveDollarCount).Should(Equal(10))
			Expect(m3.TwentyDollarCount).Should(Equal(12))
		})
	})
	When("calling value object function", func() {
		It("should call GetHashCodeCore then produce correct result", func() {
			m1 := model.NewMoney(1, 2, 3, 4, 5, 6)
			m2 := model.NewMoney(1, 2, 3, 4, 5, 6)
			Expect(m1.GetHashCodeCore()).Should((Equal(m2.GetHashCodeCore())))

			m1 = model.NewMoney(0, 2, 3, 4, 5, 6)
			Expect(m1.GetHashCodeCore()).ShouldNot((Equal(m2.GetHashCodeCore())))
		})
		It("should call EqualsCore then produce correct result", func() {
			m1 := model.NewMoney(1, 2, 3, 4, 5, 6)
			m2 := model.NewMoney(1, 2, 3, 4, 5, 6)

			result := m1.EqualsCore(m2)
			Expect(result).Should(Equal(true))

			m1 = model.NewMoney(0, 2, 3, 4, 5, 6)
			result = m1.EqualsCore(m2)
			Expect(result).ShouldNot(Equal(true))
		})
	})
})
