package mint

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	abci "github.com/tendermint/tendermint/abci/types"
)

func TestNewQuerier(t *testing.T) {
	input := newTestInput(t)
	querier := NewQuerier(input.mintKeeper)

	query := abci.RequestQuery{
		Path: "",
		Data: []byte{},
	}

	_, err := querier(input.ctx, []string{QueryParameters}, query)
	require.NoError(t, err)

	_, err = querier(input.ctx, []string{QueryInflation}, query)
	require.NoError(t, err)

	_, err = querier(input.ctx, []string{QueryAnnualProvisions}, query)
	require.NoError(t, err)

	_, err = querier(input.ctx, []string{"foo"}, query)
	require.Error(t, err)
}

func TestQueryParams(t *testing.T) {
	input := newTestInput(t)

	var params Params

	res, sdkErr := queryParams(input.ctx, input.mintKeeper)
	require.NoError(t, sdkErr)

	err := input.cdc.UnmarshalJSON(res, &params)
	require.NoError(t, err)

	parm, err := input.mintKeeper.GetParams(input.ctx)
	require.NoError(t, err)
	require.Equal(t, parm, params)
}

func TestQueryInflation(t *testing.T) {
	input := newTestInput(t)

	var inflation sdk.Dec

	res, sdkErr := queryInflation(input.ctx, input.mintKeeper)
	require.NoError(t, sdkErr)

	err := input.cdc.UnmarshalJSON(res, &inflation)
	require.NoError(t, err)

	parm, err := input.mintKeeper.GetMinter(input.ctx)
	require.NoError(t, err)
	require.Equal(t, parm.Inflation, inflation)
}

func TestQueryAnnualProvisions(t *testing.T) {
	input := newTestInput(t)

	var annualProvisions sdk.Dec

	res, sdkErr := queryAnnualProvisions(input.ctx, input.mintKeeper)
	require.NoError(t, sdkErr)

	err := input.cdc.UnmarshalJSON(res, &annualProvisions)
	require.NoError(t, err)

	parm, err := input.mintKeeper.GetMinter(input.ctx)
	require.NoError(t, err)
	require.Equal(t, parm.AnnualProvisions, annualProvisions)
}
