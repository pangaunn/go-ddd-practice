package core

type SnackMachine struct {
	OneCentCount      int
	TenCentCount      int
	QuarterCentCount  int
	OneDollarCount    int
	FiveDollarCount   int
	TwentyDollarCount int
}

func (s *SnackMachine) InsertMoney(
	oneCent, tenCent, quarterCent,
	oneDollar, fiveDollar, twentyDollar int) {
	s.OneCentCount += oneCent
	s.TenCentCount += tenCent
	s.QuarterCentCount += quarterCent
	s.OneDollarCount += oneDollar
	s.FiveDollarCount += fiveDollar
	s.TwentyDollarCount += twentyDollar
}
