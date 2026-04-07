package main

import "github.com/shopspring/decimal"

type Tax struct {
	Tax decimal.Decimal
}

type TaxOutput struct {
	Tax string `json:"tax"`
}

func FormatTaxes(taxes []Tax) []TaxOutput {
	results := make([]TaxOutput, 0, len(taxes))

	for _, tax := range taxes {
		results = append(results, TaxOutput{Tax: tax.Tax.StringFixed(2)})
	}
	return results
}
