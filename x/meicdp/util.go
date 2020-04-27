package meicdp

import (
	"math/big"
)

func calculateCollateralRatio(discountedCollateralValue *big.Float, totalDebtAmount *big.Float) *big.Float {
	// // TODO: calculate ratio
	toPercent := new(big.Float).SetFloat64(100)

	collateralRatio := new(big.Float)
	collateralRatio.Quo(discountedCollateralValue, totalDebtAmountFloat)
	collateralRatio.Mul(collateralRatio, toPercent)

	//Mock
	// collateralRatio := new(big.Float).SetFloat64(151)
	return collateralRatio
}
