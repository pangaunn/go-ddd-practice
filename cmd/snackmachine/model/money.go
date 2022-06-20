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

func (m *Money) Add(mo Money) {
	m.OneCentCount += mo.OneCentCount
	m.TenCentCount += mo.TenCentCount
	m.QuarterCount += mo.QuarterCount
	m.OneCentCount += mo.OneDollarCount
	m.FiveDollarCount += mo.FiveDollarCount
	m.TwentyDollarCount += mo.TwentyDollarCount
}

func (m *Money) Clear() {
	m.OneCentCount = 0
	m.TenCentCount = 0
	m.QuarterCount = 0
	m.OneCentCount = 0
	m.FiveDollarCount = 0
	m.TwentyDollarCount = 0
}
