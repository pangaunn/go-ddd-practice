package core_test

import (
	. "go-ddd-practice/cmd/snackmachine/core"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SnackMachine", func() {
	It("Should return empty money in transaction", func() {
		sm := SnackMachine{}
		sm.InsertMoney(NewMoney(0, 0, 0, 1, 0, 0))
		sm.ReturnMoney()

		Expect(sm.MoneyIntransaction.Amount()).Should(Equal(0.0))
	})

	It("Should have correct amount when insertMoney", func() {
		sm := SnackMachine{}
		sm.InsertMoney(NewMoney(1, 0, 0, 0, 0, 0))
		sm.InsertMoney(NewMoney(0, 0, 0, 1, 0, 0))

		Expect(sm.MoneyIntransaction.Amount()).Should(Equal(1.01))
	})

	It("Money in transaction should goes inside after purchase", func() {
		sm := SnackMachine{}
		sm.InsertMoney(NewMoney(0, 0, 0, 1, 0, 0))
		sm.InsertMoney(NewMoney(0, 0, 0, 1, 0, 0))

		sm.BuySnack()

		Expect(sm.MoneyIntransaction.Amount()).Should(Equal(0.0))
		Expect(sm.MoneyInside.Amount()).Should(Equal(2.0))
	})
})
