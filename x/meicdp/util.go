package meicdp

import (
	"math/big"

	"github.com/tharamalai/meichain/x/meicdp/types"
)

func calculateCollateralRatio(discountedCollateralValue *big.Float, totalDebtAmount *big.Float) *big.Float {
	// // TODO: calculate ratio
	// toPercent := new(big.Float).SetFloat64(100)

	// collateralRatio := new(big.Float)
	// collateralRatio.Quo(discountedCollateralValue, totalDebtAmount)
	// collateralRatio.Mul(collateralRatio, toPercent)

	//Mock
	collateralRatio := new(big.Float).SetFloat64(151)
	return collateralRatio
}

// encodeRequestParams returns byte array of encoded request coin price
func encodeRequestParams(coinName string, multiplier uint64) []byte {
	encoder := types.NewEncoder()
	encoder.EncodeString(coinName)
	encoder.EncodeU64(multiplier)
	return encoder.GetEncodedData()
}
