package types

const (
	// ModuleName is the name of the module
	ModuleName = "affondra"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey to be used for routing msgs
	RouterKey = ModuleName

	// QuerierRoute to be used for querier msgs
	QuerierRoute = ModuleName
)

const (
	PollPrefix = "poll-"
)

const (
	VotePrefix = "vote-"
)

const (
	ItemPrefix = "item-"
)
