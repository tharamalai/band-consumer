package meicdp

import (
	"math/big"
)

func calculateCollateralRatio(discountedCollateralValue *big.Int, totalDebtAmount *big.Int) *big.Float {
	// TODO: calculate ratio
	// discountedCollateralValueFloat := new(big.Float).SetInt(discountedCollateralValue)
	// totalDebtAmountFloat := new(big.Float).SetInt(totalDebtAmount)

	//Mock
	collateralRatio := new(big.Float).SetUint64(150)
	return collateralRatio
}
