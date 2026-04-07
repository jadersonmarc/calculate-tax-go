package main

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestNewCalculator(t *testing.T) {
	c := Calculator{}

	op := Operation{
		Type:     Buy,
		Quantity: 100,
	}

	c.Buy(op)

	if c.totalShares != 100 {
		t.Errorf("expected 100 shares got %d", c.totalShares)
	}
}

func TestWeightedAveragePrice(t *testing.T) {
	c := Calculator{}

	op := Operation{
		Type:     Buy,
		UnitCost: decimal.NewFromFloat(10),
		Quantity: 100,
	}

	op2 := Operation{
		Type:     Buy,
		UnitCost: decimal.NewFromFloat(20),
		Quantity: 50,
	}

	c.Buy(op)
	c.Buy(op2)

	expected := decimal.NewFromInt(2000).Div(decimal.NewFromInt(150))

	if !c.avgPrice.Equal(expected) {
		t.Errorf("expected avg price%s got %s", expected, c.avgPrice)
	}
}

func TestWeightedAverageWithDifferentQuantity(t *testing.T) {
	c := Calculator{}

	c.Buy(Operation{
		Type:     Buy,
		UnitCost: decimal.NewFromFloat(10),
		Quantity: 100,
	})

	c.Buy(Operation{
		Type:     Buy,
		UnitCost: decimal.NewFromFloat(20),
		Quantity: 50,
	})

	expected := decimal.NewFromInt(1333).Div(decimal.NewFromInt(100))
	if !c.avgPrice.Round(2).Equal(expected.Round(2)) {
		t.Errorf("expected avg price %s got %s", expected, c.avgPrice)
	}
}

func TestSellWithLoss(t *testing.T) {
	c := Calculator{}

	c.Buy(Operation{
		Type:     Buy,
		UnitCost: decimal.NewFromInt(10),
		Quantity: 100,
	})

	tax := c.Sell(Operation{
		Type:     Sell,
		UnitCost: decimal.NewFromInt(5),
		Quantity: 100,
	})

	expectedLoss := decimal.NewFromInt(500)

	if !c.accumulatedLoss.Equal(expectedLoss) {
		t.Errorf("expected loss %s got %s", expectedLoss, c.accumulatedLoss)
	}

	if !tax.Tax.Equal(zero) {
		t.Errorf("expected zero tax on loss")

	}
}

func TestSellProfitWithLossCompensationFully(t *testing.T) {
	c := Calculator{}

	// gera prejuízo
	c.Buy(Operation{
		Type:     Buy,
		UnitCost: decimal.NewFromInt(10),
		Quantity: 100,
	})

	c.Sell(Operation{
		Type:     Sell,
		UnitCost: decimal.NewFromInt(5),
		Quantity: 100,
	})

	// novo cenário com lucro menor que prejuízo
	c.Buy(Operation{
		Type:     Buy,
		UnitCost: decimal.NewFromInt(10),
		Quantity: 100,
	})

	tax := c.Sell(Operation{
		Type:     Sell,
		UnitCost: decimal.NewFromInt(15),
		Quantity: 100,
	})

	// prejuízo deve zerar
	expectedLoss := decimal.NewFromInt(0)

	if !c.accumulatedLoss.Equal(expectedLoss) {
		t.Errorf("expected remaining loss %s, got %s", expectedLoss, c.accumulatedLoss)
	}

	if !tax.Tax.Equal(zero) {
		t.Error("expected zero tax after compensation")
	}
}

func TestSellProfitWithLossCompensationPartial(t *testing.T) {
	c := Calculator{}

	// prejuízo de 500
	c.Buy(Operation{Type: Buy, UnitCost: decimal.NewFromInt(10), Quantity: 100})
	c.Sell(Operation{Type: Sell, UnitCost: decimal.NewFromInt(5), Quantity: 100})

	// novo lucro maior
	c.Buy(Operation{Type: Buy, UnitCost: decimal.NewFromInt(10), Quantity: 100})

	tax := c.Sell(Operation{
		Type:     Sell,
		UnitCost: decimal.NewFromInt(20),
		Quantity: 100,
	})

	// prejuízo deve zerar
	if !c.accumulatedLoss.Equal(zero) {
		t.Error("expected zero accumulated loss")
	}

	// imposto = (1000 - 500) * 0.05 = 25
	expectedTax := decimal.NewFromInt(25)

	if !tax.Tax.Equal(expectedTax) {
		t.Errorf("expected tax %s, got %s", expectedTax, tax.Tax)
	}
}

func TestSellWithProfitUnderExemption(t *testing.T) {
	c := Calculator{}

	c.Buy(Operation{
		Type:     Buy,
		UnitCost: decimal.NewFromInt(10),
		Quantity: 100,
	})

	tax := c.Sell(Operation{
		Type:     Sell,
		UnitCost: decimal.NewFromInt(15),
		Quantity: 100, // total = 1500
	})

	if !tax.Tax.Equal(zero) {
		t.Error("expected zero tax under exemption limit")
	}
}

func TestSellWithTax(t *testing.T) {
	c := Calculator{}

	c.Buy(Operation{
		Type:     Buy,
		UnitCost: decimal.NewFromInt(10),
		Quantity: 2000,
	})

	tax := c.Sell(Operation{
		Type:     Sell,
		UnitCost: decimal.NewFromInt(20),
		Quantity: 2000, // total = 40000
	})

	// lucro = (20-10)*2000 = 20000
	// imposto = 5% = 1000
	expected := decimal.NewFromInt(1000)

	if !tax.Tax.Equal(expected) {
		t.Errorf("expected %s, got %s", expected, tax.Tax)
	}
}
