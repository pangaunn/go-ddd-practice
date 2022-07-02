package core

type Money struct {
	OneCentCount      int
	TenCentCount      int
	QuarterCount      int
	OneDollarCount    int
	FiveDollarCount   int
	TwentyDollarCount int
}

func NewMoney(oneCent, tenCent, quarter, oneDollar, fiveDollar, twentyDollar int) Money {
	return Money{
		OneCentCount:      oneCent,
		TenCentCount:      tenCent,
		QuarterCount:      quarter,
		OneDollarCount:    oneDollar,
		FiveDollarCount:   fiveDollar,
		TwentyDollarCount: twentyDollar,
	}
}

func (m *Money) Equal(money Money) bool {
	return m.OneCentCount == money.OneCentCount &&
		m.TenCentCount == money.TenCentCount &&
		m.QuarterCount == money.QuarterCount &&
		m.OneDollarCount == money.OneDollarCount &&
		m.FiveDollarCount == money.FiveDollarCount &&
		m.TwentyDollarCount == money.TwentyDollarCount
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
