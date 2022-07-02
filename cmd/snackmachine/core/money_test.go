package core_test

import (
	"go-ddd-practice/cmd/snackmachine/core"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Money", func() {
	It("Should add correct result", func() {
		m1 := core.NewMoney(1, 2, 3, 4, 5, 6)
		m2 := core.NewMoney(1, 2, 3, 4, 5, 6)

		sum := m1.Add(m2)
		Expect(sum.OneCentCount).Should(Equal(2))
		Expect(sum.TenCentCount).Should(Equal(4))
		Expect(sum.QuarterCount).Should(Equal(6))
		Expect(sum.OneDollarCount).Should(Equal(8))
		Expect(sum.FiveDollarCount).Should(Equal(10))
		Expect(sum.TwentyDollarCount).Should(Equal(12))
	})
})
