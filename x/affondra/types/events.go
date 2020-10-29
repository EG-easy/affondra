package types

// affondra module event types
const (
	// TODO: Create your event types
	// EventType<Action>    		= "action"
	EventTypeCreateItem         = "create_item"
	EventTypeBuyItem            = "buy_item"
	EventTypeChangeInSaleStatus = "change_insale_status"
	EventTypeSetItem            = "set_item"
	EventTypeDeleteItem         = "delete_item"

	// TODO: Create keys fo your events, the values will be derivided from the msg
	// AttributeKeyAddress  		= "address"
	AttributeKeyDenom    = "denom"
	AttributeKeyNftId    = "nft_id"
	AttributeKeyReceiver = "receiver"
	AttributeKeyID       = "ID"

	// TODO: Some events may not have values for that reason you want to emit that something happened.
	// AttributeValueDoubleSign = "double_sign"

	AttributeValueCategory = ModuleName
)
