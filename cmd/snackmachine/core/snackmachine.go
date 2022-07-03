package core

type SnackMachine struct {
	ID                 int
	MoneyInside        Money
	MoneyIntransaction Money
	Slots              []Slot
}

func NewSnackMachine() SnackMachine {
	s := SnackMachine{}
	s.Slots = []Slot{
		{SnackMachine: s, Position: 1},
		{SnackMachine: s, Position: 2},
		{SnackMachine: s, Position: 3},
	}
	return s
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

func (s *SnackMachine) BuySnack(position int) {
	var slotIndex int
	for index, s := range s.Slots {
		if s.Position == position {
			slotIndex = index
			break
		}
	}
	s.Slots[slotIndex].Quantity--
	s.MoneyInside = s.MoneyInside.Add(s.MoneyIntransaction)
	s.MoneyIntransaction = Money{}
}

func (s *SnackMachine) LoadSnack(position int, snack Snack, quantity int, price float64) {
	var slotIndex int
	for index, s := range s.Slots {
		if s.Position == position {
			slotIndex = index
			break
		}
	}

	s.Slots[slotIndex].Snack = snack
	s.Slots[slotIndex].Quantity = quantity
	s.Slots[slotIndex].Price = price
}
