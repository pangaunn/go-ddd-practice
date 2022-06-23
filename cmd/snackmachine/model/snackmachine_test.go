package model_test

import (
	"go-ddd-practice/cmd/snackmachine/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Snackmachine", func() {
	When("Add", func() {
		It("should empty money in trasaction", func() {
			sm := model.NewSnackMachine()

			sm.InsertMoney(model.NewMoney(1, 1, 1, 1, 1, 1))
			sm.ReturnMoney()

			Expect(sm.MoneyInTransaction.Amount()).Should(Equal(float64(0)))
		})
	})
})
