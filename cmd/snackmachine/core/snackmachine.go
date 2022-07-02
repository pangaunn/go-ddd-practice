package core

type SnackMachine struct {
	ID                 int
	MoneyInside        Money
	MoneyIntransaction Money
}

func (s *SnackMachine) Equal(sm SnackMachine) bool {
	return s.ID == sm.ID
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
