package core

type Money struct {
	OneCentCount      int
	TenCentCount      int
	QuarterCount      int
	OneDollarCount    int
	FiveDollarCount   int
	TwentyDollarCount int
}

func (m *Money) Add(money Money) Money {
	sum := Money{}
	sum.OneCentCount = m.OneCentCount + money.OneCentCount
	sum.TenCentCount = m.TenCentCount + money.TenCentCount
	sum.QuarterCount = m.QuarterCount + money.QuarterCount
	sum.OneDollarCount = m.OneDollarCount + money.OneDollarCount
	sum.FiveDollarCount = m.FiveDollarCount + money.FiveDollarCount
	sum.TwentyDollarCount = m.TwentyDollarCount + money.TwentyDollarCount
	return sum
}
