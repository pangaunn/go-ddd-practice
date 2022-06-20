package model_test

import (
	"go-ddd-practice/cmd/snackmachine/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Money", func() {
	It("Add should product correct result", func() {
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
