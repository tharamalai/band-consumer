package meicdp

func calculateCollateralRatio(discountedCollateralValue float64, totalDebtAmount float64) float64 {
	return (discountedCollateralValue / totalDebtAmount) * 100
}
