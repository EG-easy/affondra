package types

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Item defines item struct
type Item struct {
	Creator     sdk.AccAddress `json:"creator" yaml:"creator"`
	ID          string         `json:"id" yaml:"id"`
	Denom       string         `json:"denom" yaml:"denom"`
	NftId       string         `json:"nftId" yaml:"nftId"`
	TokenURI    string         `json:"token_uri" yaml:"token_uri"`
	Price       sdk.Coin       `json:"price" yaml:"price"`
	Affiliate   sdk.Coin       `json:"affiliate" yaml:"affiliate"`
	Receiver    sdk.AccAddress `json:"receiver" yaml:"receiver"`
	Description string         `json:"description" yaml:"description"`
	InSale      bool           `json:"inSale" yaml:"inSale"`
}

// NewItem is a constructor function for Item
func NewItem(id string, owner sdk.AccAddress, denom string, nftId string, tokenURI string, price sdk.Coin, affiliate sdk.Coin, description string, inSale bool) Item {
	return Item{
		Creator:     owner,
		ID:          id,
		Denom:       denom,
		NftId:       nftId,
		TokenURI:    tokenURI,
		Price:       price,
		Affiliate:   affiliate,
		Receiver:    nil, //Receiver will be decided after buying
		Description: description,
		InSale:      inSale,
	}
}

// GetID returns item id
func (item Item) GetID() string {
	return item.ID
}

// GetOwner returns creator address
func (item Item) GetOwner() sdk.AccAddress {
	return item.Creator
}

// GetDenom returns item denom
func (item Item) GetDenom() string {
	return item.Denom
}

// GetNftid returns item nft id
func (item Item) GetNftid() string {
	return item.NftId
}

// GetTokenURI returns item token uri
func (item Item) GetTokenURI() string {
	return item.TokenURI
}

// GetPrice returns item price
func (item Item) GetPrice() sdk.Coin {
	return item.Price
}

// GetAffiliate returns affiliate reward
func (item Item) GetAffiliate() sdk.Coin {
	return item.Affiliate
}

// GetReceiver returns receiver address
func (item Item) GetReceiver() sdk.AccAddress {
	return item.Receiver
}

// GetDescription returns item description
func (item Item) GetDescription() string {
	return item.Description
}

// GetInSale returns item is in sale or not
func (item Item) GetInSale() bool {
	return item.InSale
}

// SetOwner set item owner address
func (item *Item) SetOwner(addr sdk.AccAddress) {
	item.Creator = addr
}

// SetTokenURI set item URI
func (item *Item) SetTokenURI(tokenURI string) {
	item.TokenURI = tokenURI
}

// SetReceiver set item receiver
func (item *Item) SetReceiver(addr sdk.AccAddress) {
	item.Receiver = addr
}

// SetDescription set item description
func (item *Item) SetDescription(desc string) {
	item.Description = desc
}

// ChangeInSaleStatus set item sale status
func (item *Item) ChangeInSaleStatus() {
	item.InSale = !item.InSale
}

func (item Item) String() string {
	return fmt.Sprintf(`ID:%s
Creator:%s
Denom:%s
NftId:%s
TokenURI:%s
Price:%s
Affiliate:%s
Receiver:%s
Description:%s
InSale:%v`,
		item.ID,
		item.Creator,
		item.Denom,
		item.NftId,
		item.TokenURI,
		item.Price,
		item.Affiliate,
		item.Receiver,
		item.Description,
		item.InSale,
	)
}

// Items define a list of Item
type Items []Item

// NewItems is a constructor function for Items
func NewItems(items ...Item) Items {
	if len(items) == 0 {
		return Items{}
	}
	return Items(items).Sort()
}

// Append appends two set of Items
func (items Items) Append(itemsB ...Item) Items {
	return append(items, itemsB...).Sort()
}

// Find returns searched item from the set
func (items Items) Find(id string) (item Item, found bool) {
	index := items.find(id)
	if index == -1 {
		return item, false
	}
	return items[index], true
}

// Update removes and replaces an item from the set
func (items Items) Update(id string, item Item) (Items, bool) {
	index := items.find(id)
	if index == -1 {
		return items, false
	}
	return append(append(items[:index], item), items[index+1:]...), true
}

// Remove removes an Item from the set of Items
func (items Items) Remove(id string) (Items, bool) {
	index := items.find(id)
	if index == -1 {
		return items, false
	}

	return append(items[:index], items[index+1:]...), true
}

// String follows stringer interface
func (items Items) String() string {
	if len(items) == 0 {
		return ""
	}

	out := ""
	for _, item := range items {
		out += fmt.Sprintf("%v\n", item.String())
	}
	return out[:len(out)-1]
}

// Empty returns true if there are no items and false otherwise.
func (items Items) Empty() bool {
	return len(items) == 0
}

func (items Items) find(id string) int {
	return FindUtil(items, id)
}

// ----------------------------------------------------------------------------
// Encoding

// ItemJSON is the exported Item format for clients
type ItemJSON map[string]Item

// MarshalJSON for Items
func (items Items) MarshalJSON() ([]byte, error) {
	itemJSON := make(ItemJSON)
	for _, item := range items {
		id := item.GetID()
		item := NewItem(id, item.GetOwner(), item.GetDenom(), item.GetNftid(), item.GetTokenURI(), item.GetPrice(), item.GetAffiliate(), item.GetDescription(), item.GetInSale())
		itemJSON[id] = item
	}
	return json.Marshal(itemJSON)
}

// UnmarshalJSON for Items
func (items *Items) UnmarshalJSON(b []byte) error {
	itemJSON := make(ItemJSON)
	if err := json.Unmarshal(b, &itemJSON); err != nil {
		return err
	}

	for id, item := range itemJSON {
		baseitem := NewItem(id, item.GetOwner(), item.GetDenom(), item.GetNftid(), item.GetTokenURI(), item.GetPrice(), item.GetAffiliate(), item.GetDescription(), item.GetInSale())
		*items = append(*items, baseitem)
	}
	return nil
}

// ElAtIndex is Findable and Sort interface
func (items Items) ElAtIndex(index int) string { return items[index].GetID() }
func (items Items) Len() int                   { return len(items) }
func (items Items) Less(i, j int) bool {
	return strings.Compare(items[i].GetID(), items[j].GetID()) == -1
}
func (items Items) Swap(i, j int) { items[i], items[j] = items[j], items[i] }

var _ sort.Interface = Items{}

// Sort is a helper function to sort the set of coins in place
func (items Items) Sort() Items {
	sort.Sort(items)
	return items
}
