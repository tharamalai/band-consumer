package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

//CDP is a Collateralized Debt Position of an user account
type CDP struct {
	CollateralAmount uint64         `json:"collateralAmount"`
	DebtAmount       uint64         `json:"debtAmount"`
	Owner            sdk.AccAddress `json:"owner"`
}

//NewCDP creates a new CDP instance.
func NewCDP(
	collateralAmount uint64,
	debtAmount uint64,
	owner sdk.AccAddress,
) CDP {
	return CDP{
		CollateralAmount: collateralAmount,
		DebtAmount:       debtAmount,
		Owner:            owner,
	}
}
