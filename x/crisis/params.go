package crisis

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
)

// Default parameter namespace
const (
	DefaultParamspace = ModuleName
)

var (
	// key for constant fee parameter
	ParamStoreKeyConstantFee = []byte("ConstantFee")
)

// type declaration for parameters
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable(
		ParamStoreKeyConstantFee, sdk.Coin{},
	)
}

// GetConstantFee get's the constant fee from the paramSpace
func (k Keeper) GetConstantFee(ctx sdk.Context) (constantFee sdk.Coin) {
	if err := k.paramSpace.Get(ctx, ParamStoreKeyConstantFee, &constantFee); err != nil {
		// TODO: return error - needs rewrite interfaces
		// and handle error on the caller side
		// check PR #3782
	}
	return
}

// GetConstantFee set's the constant fee in the paramSpace
func (k Keeper) SetConstantFee(ctx sdk.Context, constantFee sdk.Coin) {
	if err := k.paramSpace.Set(ctx, ParamStoreKeyConstantFee, constantFee); err != nil {
		// TODO: return error - needs rewrite interfaces
		// and handle error on the caller side
		// check PR #3782
	}
}
