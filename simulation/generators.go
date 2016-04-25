package simulation

import (
	_ "github.com/leesper/go_rng"
)

/*
func (b *Blockchain) GenerateTransactionArrival() int {
	// TODO make this readable
	return round(b.expGen.Exp(b.TransactionLambda))
}

func (b *Blockchain) GenerateBlockArrival() int {
	// TODO find the actual parameters
	const (
		Alpha = iota
		Beta
	)
	return round(b.gammaGen.Gamma(Alpha, Beta))
}

func round(f float64) int {
	if math.Abs(f) < 0.5 {
		return 0
	}
	return int(f + math.Copysign(0.5, f))
}
*/
