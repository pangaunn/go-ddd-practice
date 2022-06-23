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

func (m *Money) Subtract(mo Money) Money {
	oneCentCount := m.OneCentCount - mo.OneCentCount
	tenCentCount := m.TenCentCount - mo.TenCentCount
	quarterCount := m.QuarterCount - mo.QuarterCount
	oneDollarCount := m.OneDollarCount - mo.OneDollarCount
	fiveDollarCount := m.FiveDollarCount - mo.FiveDollarCount
	twentyDollarCount := m.TwentyDollarCount - mo.TwentyDollarCount
	return NewMoney(oneCentCount, tenCentCount, quarterCount, oneDollarCount, fiveDollarCount, twentyDollarCount)
}

func (m *Money) Clear() Money {
	return Money{}
}

func (m *Money) Amount() float64 {
	return float64(m.OneCentCount)*0.01 +
		float64(m.TenCentCount)*0.1 +
		float64(m.QuarterCount)*0.25 +
		float64(m.OneDollarCount) +
		float64(m.FiveDollarCount)*5 +
		float64(m.TwentyDollarCount)*20
}

func (m *Money) GetHashCodeCore() int {
	hashCode := m.OneCentCount
	hashCode = (hashCode * 397) ^ m.TenCentCount
	hashCode = (hashCode * 397) ^ m.QuarterCount
	hashCode = (hashCode * 397) ^ m.OneDollarCount
	hashCode = (hashCode * 397) ^ m.FiveDollarCount
	hashCode = (hashCode * 397) ^ m.TwentyDollarCount
	return hashCode
}

func (m *Money) EqualsCore(money Money) bool {
	return m.OneCentCount == money.OneCentCount && m.TenCentCount == money.TenCentCount && m.QuarterCount == money.QuarterCount && m.OneDollarCount == money.OneDollarCount && m.FiveDollarCount == money.FiveDollarCount && m.TwentyDollarCount == money.TwentyDollarCount
}
