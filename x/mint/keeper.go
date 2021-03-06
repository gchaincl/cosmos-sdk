package mint

import (
	"errors"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
)

const (
	// ModuleName is the name of the module
	ModuleName = "minting"

	// default paramspace for params keeper
	DefaultParamspace = "mint"

	// StoreKey is the default store key for mint
	StoreKey = "mint"

	// QuerierRoute is the querier route for the minting store.
	QuerierRoute = StoreKey
)

// keeper of the staking store
type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        *codec.Codec
	paramSpace params.Subspace
	sk         StakingKeeper
	fck        FeeCollectionKeeper
}

func NewKeeper(cdc *codec.Codec, key sdk.StoreKey,
	paramSpace params.Subspace, sk StakingKeeper, fck FeeCollectionKeeper) Keeper {

	keeper := Keeper{
		storeKey:   key,
		cdc:        cdc,
		paramSpace: paramSpace.WithKeyTable(ParamKeyTable()),
		sk:         sk,
		fck:        fck,
	}
	return keeper
}

//____________________________________________________________________
// Keys

var (
	minterKey = []byte{0x00} // the one key to use for the keeper store

	// params store for inflation params
	ParamStoreKeyParams = []byte("params")
)

// ParamTable for staking module
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable(
		ParamStoreKeyParams, Params{},
	)
}

//______________________________________________________________________

// get the minter
func (k Keeper) GetMinter(ctx sdk.Context) (minter Minter, err error) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(minterKey)
	if b == nil {
		err = errors.New("Stored fee pool should not have been nil")
		return
	}
	k.cdc.MustUnmarshalBinaryLengthPrefixed(b, &minter)
	return
}

// set the minter
func (k Keeper) SetMinter(ctx sdk.Context, minter Minter) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshalBinaryLengthPrefixed(minter)
	store.Set(minterKey, b)
}

//______________________________________________________________________

// get inflation params from the global param store
func (k Keeper) GetParams(ctx sdk.Context) (Params, error) {
	var params Params
	err := k.paramSpace.Get(ctx, ParamStoreKeyParams, &params)
	return params, err
}

// set inflation params from the global param store
func (k Keeper) SetParams(ctx sdk.Context, params Params) error {
	return k.paramSpace.Set(ctx, ParamStoreKeyParams, &params)
}
