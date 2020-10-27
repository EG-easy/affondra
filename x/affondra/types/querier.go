package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const QueryListItem = "list-item"
const QueryGetItem = "get-item"

const QuerySupply = "supply"
const QueryOwner = "owner"
const QueryDenom = "denom"

type QueryOwnerParams struct {
	Owner sdk.AccAddress
}

// QueryCollectionParams defines the params for queries:
type QueryCollectionParams struct {
	Denom string
}

// NewQueryCollectionParams creates a new instance of QuerySupplyParams
func NewQueryCollectionParams(denom string) QueryCollectionParams {
	return QueryCollectionParams{Denom: denom}
}

// Bytes exports the Denom as bytes
func (q QueryCollectionParams) Bytes() []byte {
	return []byte(q.Denom)
}
