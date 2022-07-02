package core

type SnackMachine struct {
	OneCentCount      int
	TenCentCount      int
	QuarterCentCount  int
	OneDollarCount    int
	FiveDollarCount   int
	TwentyDollarCount int

	OneCentCountInTransaction      int
	TenCentCountInTransaction      int
	QuarterCentCountInTransaction  int
	OneDollarCountInTransaction    int
	FiveDollarCountInTransaction   int
	TwentyDollarCountInTransaction int
}

func (s *SnackMachine) InsertMoney(
	oneCent, tenCent, quarterCent,
	oneDollar, fiveDollar, twentyDollar int) {
	s.OneCentCountInTransaction += oneCent
	s.TenCentCountInTransaction += tenCent
	s.QuarterCentCountInTransaction += quarterCent
	s.OneDollarCountInTransaction += oneDollar
	s.FiveDollarCountInTransaction += fiveDollar
	s.TwentyDollarCountInTransaction += twentyDollar
}
