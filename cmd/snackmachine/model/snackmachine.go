package model

type SnackMachine struct {
	MoneyInside        Money
	MoneyInTransaction Money
}

func NewSnackMachine() SnackMachine {
	return SnackMachine{}
}

func (s *SnackMachine) InsertMoney(money Money) {
	s.MoneyInTransaction = s.MoneyInTransaction.Add(money)
}

func (s *SnackMachine) ReturnMoney() {
	s.MoneyInTransaction = s.MoneyInTransaction.Clear()
}

func (s *SnackMachine) BuySnack() {
	s.MoneyInTransaction = s.MoneyInside.Add(s.MoneyInTransaction)
	s.MoneyInTransaction = s.MoneyInTransaction.Clear()
}
