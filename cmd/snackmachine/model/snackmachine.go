package model

type SnackMachine interface {
	InsertMoney(oneCentCount, tenCentCount, quarterCount, oneDollarCount, fiveDollarCount, twentyDollarCount int)
	ReturnMoney()
	BuySnack()
}

type snackMachine struct {
	OneCentCount      int
	TenCentCount      int
	QuarterCount      int
	OneDollarCount    int
	FiveDollarCount   int
	TwentyDollarCount int

	OneCentCountInTransaction      int
	TenCentCountInTransaction      int
	QuarterCountInTransaction      int
	OneDollarCountInTransaction    int
	FiveDollarCountInTransaction   int
	TwentyDollarCountInTransaction int
}

func NewSnackMachine() SnackMachine {
	return &snackMachine{}
}

func (s *snackMachine) InsertMoney(oneCentCount, tenCentCount, quarterCount, oneDollarCount, fiveDollarCount, twentyDollarCount int) {
	s.OneCentCountInTransaction += oneCentCount
	s.TenCentCountInTransaction += tenCentCount
	s.QuarterCountInTransaction += quarterCount
	s.OneDollarCountInTransaction += oneDollarCount
	s.FiveDollarCountInTransaction += fiveDollarCount
	s.TwentyDollarCountInTransaction += twentyDollarCount
}

func (s *snackMachine) ReturnMoney() {
	s.OneCentCountInTransaction = 0
	s.TenCentCountInTransaction = 0
	s.QuarterCountInTransaction = 0
	s.OneDollarCountInTransaction = 0
	s.FiveDollarCountInTransaction = 0
	s.TwentyDollarCountInTransaction = 0
}

func (s *snackMachine) BuySnack() {
	s.OneCentCount += s.OneCentCountInTransaction
	s.TenCentCount += s.TenCentCountInTransaction
	s.QuarterCount += s.QuarterCountInTransaction
	s.OneDollarCount += s.OneDollarCountInTransaction
	s.FiveDollarCount += s.FiveDollarCountInTransaction
	s.TwentyDollarCount += s.TwentyDollarCountInTransaction

	s.OneCentCountInTransaction = 0
	s.TenCentCountInTransaction = 0
	s.QuarterCountInTransaction = 0
	s.OneDollarCountInTransaction = 0
	s.FiveDollarCountInTransaction = 0
	s.TwentyDollarCountInTransaction = 0
}
