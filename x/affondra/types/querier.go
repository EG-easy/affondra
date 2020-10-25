package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const QueryListPoll = "list-poll"
const QueryGetPoll = "get-poll"

const QueryListVote = "list-vote"
const QueryGetVote = "get-vote"

const QueryListItem = "list-item"
const QueryGetItem = "get-item"

const QuerySupply = "supply"
const QueryOwner = "owner"

type QueryOwnerParams struct {
	Owner sdk.AccAddress
}
