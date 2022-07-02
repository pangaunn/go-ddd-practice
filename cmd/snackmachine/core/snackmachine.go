package core

type SnackMachine struct {
	MoneyInside        Money
	MoneyIntransaction Money
}

func (s *SnackMachine) InsertMoney(money Money) {
	s.MoneyIntransaction = s.MoneyIntransaction.Add(money)
}

func (s *SnackMachine) ReturnMoney() {
	s.MoneyIntransaction = Money{}
}

func (s *SnackMachine) BuySnack() {
	s.MoneyInside = s.MoneyInside.Add(s.MoneyIntransaction)
	s.MoneyIntransaction = Money{}
}
