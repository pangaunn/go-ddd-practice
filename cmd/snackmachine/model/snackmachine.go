package model

type SnackMachine interface {
	InsertMoney(oneCentCount, tenCentCount, quarterCount, oneDollarCount, fiveDollarCount, twentyDollarCount int)
}

type snackMachine struct {
	OneCentCount      int
	TenCentCount      int
	QuarterCount      int
	OneDollarCount    int
	FiveDollarCount   int
	TwentyDollarCount int
}

func NewSnackMachine() SnackMachine {
	return &snackMachine{}
}

func (s *snackMachine) InsertMoney(oneCentCount, tenCentCount, quarterCount, oneDollarCount, fiveDollarCount, twentyDollarCount int) {
	s.OneCentCount += oneCentCount
	s.TenCentCount += tenCentCount
	s.QuarterCount += quarterCount
	s.OneDollarCount += oneDollarCount
	s.FiveDollarCount += fiveDollarCount
	s.TwentyDollarCount += twentyDollarCount
}
