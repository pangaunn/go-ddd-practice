package model

type SnackMachine interface {
	InsertMoney(money Money)
	ReturnMoney()
	BuySnack()
}

type snackMachine struct {
	MoneyInside        Money
	MoneyInTransaction Money
}

func NewSnackMachine() SnackMachine {
	return &snackMachine{}
}

func (s *snackMachine) InsertMoney(money Money) {
	s.MoneyInTransaction.Add(money)
}

func (s *snackMachine) ReturnMoney() {
	s.MoneyInTransaction.Clear()
}

func (s *snackMachine) BuySnack() {
	s.MoneyInside.Add(s.MoneyInTransaction)
	s.MoneyInTransaction.Clear()
}
