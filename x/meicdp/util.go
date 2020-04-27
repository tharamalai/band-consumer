package meicdp

import (
	"math/big"
)

func calculateCollateralRatio(discountedCollateralValue *big.Int, totalDebtAmount *big.Int) *big.Float {
	// // TODO: calculate ratio
	// discountedCollateralValueFloat := new(big.Float).SetInt(discountedCollateralValue)
	// totalDebtAmountFloat := new(big.Float).SetInt(totalDebtAmount)
	// toPercent := new(big.Float).SetFloat64(100)

	// collateralRatio := new(big.Float)
	// collateralRatio.Quo(discountedCollateralValueFloat, totalDebtAmountFloat)
	// collateralRatio.Mul(collateralRatio, toPercent)

	//Mock
	collateralRatio := new(big.Float).SetFloat64(151)
	return collateralRatio
}
