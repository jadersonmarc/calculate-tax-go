package main

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestProcessOperations(t *testing.T) {
	ops := []Operation{
		{
			Type:     Buy,
			UnitCost: decimal.NewFromInt(10),
			Quantity: 100,
		},
		{
			Type:     Sell,
			UnitCost: decimal.NewFromInt(15),
			Quantity: 100,
		},
	}

	result := ProcessOperations(ops)

	if len(result) != 2 {
		t.Fatalf("expected 2 results, got %d", len(result))
	}

	// primeira operação: compra → imposto 0
	if !result[0].Tax.Equal(zero) {
		t.Error("expected zero tax for buy")
	}

	// segunda operação: venda abaixo de 20k → isento
	if !result[1].Tax.Equal(zero) {
		t.Error("expected zero tax under exemption")
	}
}

func TestProcessOperationsWithTax(t *testing.T) {
	ops := []Operation{
		{Type: Buy, UnitCost: decimal.NewFromInt(10), Quantity: 2000},
		{Type: Sell, UnitCost: decimal.NewFromInt(20), Quantity: 2000},
	}

	result := ProcessOperations(ops)

	expected := decimal.NewFromInt(1000)

	if !result[1].Tax.Equal(expected) {
		t.Errorf("expected %s, got %s", expected, result[1].Tax)
	}
}
