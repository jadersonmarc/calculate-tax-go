package main

import "github.com/shopspring/decimal"

type Calculator struct {
	avgPrice        decimal.Decimal
	totalShares     int
	accumulatedLoss decimal.Decimal
}

var (
	zero           = decimal.NewFromInt(0)
	taxRate        = decimal.NewFromFloat(0.20)
	exemptionLimit = decimal.NewFromInt(20000)
)

func (c *Calculator) Buy(op Operation) {
	qty := decimal.NewFromInt(int64(op.Quantity))

	// custo total atual = preço médio * quantidade total
	currentTotal := c.avgPrice.Mul(decimal.NewFromInt(int64(c.totalShares)))

	// custo total da nova operação = preço unitário * quantidade
	newTotal := op.UnitCost.Mul(qty)

	// custo total acumulado = custo total atual + custo total da nova operação
	totalCost := currentTotal.Add(newTotal)

	//atualiza quantidade total
	c.totalShares += op.Quantity

	// evita divisão por zero
	if c.totalShares > 0 {
		c.avgPrice = totalCost.Div(decimal.NewFromInt(int64(c.totalShares)))
	}
}

func (c *Calculator) Sell(op Operation) Tax {
	qty := decimal.NewFromInt(int64(op.Quantity))
	totalValue := op.UnitCost.Mul(qty)
	profit := op.UnitCost.Sub(c.avgPrice).Mul(qty)

	// atualiza posição
	c.totalShares -= op.Quantity

	// verifica prejuízo primeiro
	if profit.LessThan(zero) {
		c.accumulatedLoss = c.accumulatedLoss.Add(profit.Abs())
		return Tax{Tax: zero}
	}

	// compensação de prejuízo
	compensated := false
	if c.accumulatedLoss.GreaterThan(zero) {
		if profit.LessThanOrEqual(c.accumulatedLoss) {
			c.accumulatedLoss = c.accumulatedLoss.Sub(profit)
			return Tax{Tax: zero}
		}
		profit = profit.Sub(c.accumulatedLoss)
		c.accumulatedLoss = zero
		compensated = true
	}

	// regra da isenção
	if !compensated && totalValue.LessThanOrEqual(exemptionLimit) {
		return Tax{Tax: zero}
	}

	tax := profit.Mul(taxRate)
	return Tax{Tax: tax}
}
