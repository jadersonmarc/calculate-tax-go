package money

import (
	"errors"
	"fmt"
)

type Money struct {
	cents    int64
	currency string
}

// New cria a partir de centavos
func New(cents int64, currency string) Money {
	if currency == "" {
		currency = "BRL"
	}

	return Money{
		cents:    cents,
		currency: currency,
	}
}

// NewFromFloat (uso controlado)
func NewFromFloat(amount float64, currency string) Money {
	return New(int64(amount*100), currency)
}

// Add soma valores
func (m Money) Add(other Money) (Money, error) {
	if m.currency != other.currency {
		return Money{}, errors.New("currency mismatch")
	}

	return Money{
		cents:    m.cents + other.cents,
		currency: m.currency,
	}, nil
}

// Sub subtrai valores
func (m Money) Sub(other Money) (Money, error) {
	if m.currency != other.currency {
		return Money{}, errors.New("currency mismatch")
	}

	return Money{
		cents:    m.cents - other.cents,
		currency: m.currency,
	}, nil
}

// Mul multiplica (ex: quantidade)
func (m Money) Mul(multiplier int64) Money {
	return Money{
		cents:    m.cents * multiplier,
		currency: m.currency,
	}
}

// MulFloat (usar com cuidado)
func (m Money) MulFloat(multiplier float64) Money {
	return Money{
		cents:    int64(float64(m.cents) * multiplier),
		currency: m.currency,
	}
}

// IsZero verifica se é zero
func (m Money) IsZero() bool {
	return m.cents == 0
}

// IsNegative verifica se é negativo
func (m Money) IsNegative() bool {
	return m.cents < 0
}

// Equals compara
func (m Money) Equals(other Money) bool {
	return m.cents == other.cents && m.currency == other.currency
}

// Float64 (compatibilidade)
func (m Money) Float64() float64 {
	return float64(m.cents) / 100
}

// String formata (2 casas)
func (m Money) String() string {
	return fmt.Sprintf("%.2f", m.Float64())
}

// Cents retorna valor bruto
func (m Money) Cents() int64 {
	return m.cents
}

// Currency retorna moeda
func (m Money) Currency() string {
	return m.currency
}
