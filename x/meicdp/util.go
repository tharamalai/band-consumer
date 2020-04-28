package meicdp

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/tharamalai/meichain/x/meicdp/types"
)

func calculateCollateralRatio(discountedCollateralValue *big.Float, totalDebtAmount *big.Float) *big.Float {
	toPercent := new(big.Float).SetFloat64(100)

	collateralRatio := new(big.Float)
	collateralRatio.Quo(discountedCollateralValue, totalDebtAmount)
	collateralRatio.Mul(collateralRatio, toPercent)

	// //Mock
	// // collateralRatio := new(big.Float).SetFloat64(151)
	return collateralRatio
}

// encodeRequestParams returns byte array of encoded request coin price
func encodeRequestParams(coinName string, multiplier uint64) string {
	encoder := types.NewEncoder()
	encoder.EncodeString(coinName)
	encoder.EncodeU64(multiplier)
	return hex.EncodeToString(encoder.GetEncodedData())
}

// calculateCollateralRatioOfCDP returns collateral ratio of the CDP
func calculateCollateralRatioOfCDP(cdp types.CDP, collateralPrice uint64, collateralPriceMultipler uint64) *big.Float {

	fmt.Println("collateralPrice", collateralPrice)

	fmt.Println("collateralPriceMultipler", collateralPriceMultipler)

	// Calculate new collateral ratio. If collateral is lower than 150 percent then returns error.
	conllateralPriceFloat64 := new(big.Float).SetUint64(collateralPrice)

	// Remove multiplier from result price (USD per ATOM)
	usdPerAtom := new(big.Float).Quo(conllateralPriceFloat64, new(big.Float).SetUint64(collateralPriceMultipler))
	fmt.Println("usdPerAtom", usdPerAtom)

	// Convert USD per ATOM to USD per uATOM
	unitPerAtom := new(big.Float).SetInt64(AtomUnitPerAtom)
	usdPerUnitAtom := new(big.Float).Quo(usdPerAtom, unitPerAtom)
	fmt.Println("usdPerUnitAtom", usdPerUnitAtom)

	collateralAmount := new(big.Float).SetUint64(cdp.CollateralAmount)
	discountCollateralValueUint64 := new(big.Float).Mul(collateralAmount, usdPerUnitAtom)
	fmt.Println("discountCollateralValueUint64", discountCollateralValueUint64)

	deptAmountFloat64 := new(big.Float).SetUint64(cdp.DebtAmount)

	// Convert uMei to Mei(USD)
	deptAmountFloat64.Quo(deptAmountFloat64, new(big.Float).SetInt64(MeiUnitPerMei))
	fmt.Println("deptAmountFloat64", deptAmountFloat64)

	collateralRatioFloat := calculateCollateralRatio(discountCollateralValueUint64, deptAmountFloat64)
	fmt.Println("collateralRatioFloat", collateralRatioFloat)
	return collateralRatioFloat
}
