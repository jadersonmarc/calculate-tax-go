package main

func ProcessOperations(ops []Operation) []Tax {
	calc := Calculator{}
	results := make([]Tax, 0, len(ops))

	for _, op := range ops {
		var tax Tax

		if op.Type == Buy {
			calc.Buy(op)
			tax = Tax{Tax: zero}
		} else {
			tax = calc.Sell(op)
		}
		results = append(results, tax)
	}

	return results
}
