package tdd

type Currency int

const (
	DollarType Currency = iota
	FrancType
)

type Money struct {
	currency Currency
	amount   int
}

func NewMoney(currency Currency, amount int) *Money {
	return &Money{
		currency: currency,
		amount:   amount,
	}
}

func NewDollar(amount int) *Money {
	return &Money{
		currency: DollarType,
		amount:   amount,
	}
}

func NewFranc(amount int) *Money {
	return &Money{
		currency: FrancType,
		amount:   amount,
	}
}

func (m *Money) Equals(o interface{}) bool {
	if o == nil {
		return false
	}

	money, ok := o.(MoneyAmount)
	if !ok {
		return false
	}

	if m.Currency() != money.Currency() {
		return false
	}

	return m.amount == money.Amount()
}

func (m *Money) Times(multiplier int) *Money {
	return NewMoney(m.currency, m.amount*multiplier)
}

func (m *Money) Amount() int {
	return m.amount
}

func (m *Money) Currency() Currency {
	return m.currency
}

type MoneyAmount interface {
	Amount() int
	Currency() Currency
}
