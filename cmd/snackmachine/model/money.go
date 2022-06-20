package model

type Money struct {
	OneCentCount      int
	TenCentCount      int
	QuarterCount      int
	OneDollarCount    int
	FiveDollarCount   int
	TwentyDollarCount int
}

func NewMoney(oneCentCount, tenCentCount, quarterCount, oneDollarCount, fiveDollarCount, twentyDollarCount int) Money {
	return Money{
		OneCentCount:      oneCentCount,
		TenCentCount:      tenCentCount,
		QuarterCount:      quarterCount,
		OneDollarCount:    oneDollarCount,
		FiveDollarCount:   fiveDollarCount,
		TwentyDollarCount: twentyDollarCount,
	}
}

func (m *Money) Add(mo Money) Money {
	oneCentCount := m.OneCentCount + mo.OneCentCount
	tenCentCount := m.TenCentCount + mo.TenCentCount
	quarterCount := m.QuarterCount + mo.QuarterCount
	oneDollarCount := m.OneDollarCount + mo.OneDollarCount
	fiveDollarCount := m.FiveDollarCount + mo.FiveDollarCount
	twentyDollarCount := m.TwentyDollarCount + mo.TwentyDollarCount
	return NewMoney(oneCentCount, tenCentCount, quarterCount, oneDollarCount, fiveDollarCount, twentyDollarCount)
}

func (m *Money) Clear() Money {
	return Money{}
}
