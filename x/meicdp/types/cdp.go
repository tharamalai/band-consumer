package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RouterKey is they name of the bank module
const RouterKey = ModuleName

//CDP is a Collateralized Debt Position of an user account
type CDP struct {
	CollateralAmount sdk.Coins      `json:"collateralAmount"`
	DebtAmount       sdk.Coins      `json:"debtAmount"`
	Owner            sdk.AccAddress `json:"owner"`
}

//NewCDP creates a new CDP instance.
func NewCDP(
	collateralAmount sdk.Coins,
	debtAmount sdk.Coins,
	owner sdk.AccAddress,
) CDP {
	return CDP{
		CollateralAmount: collateralAmount,
		DebtAmount:       debtAmount,
		Owner:            owner,
	}
}
