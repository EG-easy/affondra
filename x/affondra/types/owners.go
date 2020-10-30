package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// type SortedStringArray []string

// Owner defines key(address) -> value(items) for query
type Owner struct {
	Address sdk.AccAddress `json:"address" yaml:"address"`
	Items   Items          `json:"items" yaml:"items"`
}

// NewOwner is a constructor function for Owner
func NewOwner(owner sdk.AccAddress, items Items) Owner {
	return Owner{
		Address: owner,
		Items:   items,
	}
}

// GetItem gets a Item from the owner
func (owner Owner) GetItem(id string) (item Item, err error) {
	item, found := owner.Items.Find(id)
	if found {
		return item, nil
	}
	return Item{}, sdkerrors.Wrap(ErrUnknownItem,
		fmt.Sprintf("Item #%s doesn't exist in owner %s", id, owner.Address),
	)
}

// ContainsItem returns whether or not a owner contains an Item
func (owner Owner) ContainsItem(id string) bool {
	_, err := owner.GetItem(id)
	return err == nil
}

// AddItem adds an Item to the owner
func (owner Owner) AddItem(item Item) (Owner, error) {
	id := item.GetID()
	exists := owner.ContainsItem(id)
	if exists {
		return owner, sdkerrors.Wrap(ErrItemAlreadyExists,
			fmt.Sprintf("Item #%s already exists in owner %s", id, owner.Address),
		)
	}
	owner.Items = owner.Items.Append(item)
	return owner, nil
}

// UpdateItem updates an Item from a owner
func (owner Owner) UpdateItem(item Item) (Owner, error) {
	items, ok := owner.Items.Update(item.GetID(), item)

	if !ok {
		return owner, sdkerrors.Wrap(ErrUnknownItem,
			fmt.Sprintf("Item #%s doesn't exist on owner %s", item.GetID(), owner.Address),
		)
	}
	owner.Items = items
	return owner, nil
}

// DeleteItem deletes an Item from a owner
func (owner Owner) DeleteItem(item Item) (Owner, error) {
	items, ok := owner.Items.Remove(item.GetID())
	if !ok {
		return owner, sdkerrors.Wrap(ErrUnknownItem,
			fmt.Sprintf("Item #%s doesn't exist on owner %s", item.GetID(), owner.Address),
		)
	}
	owner.Items = items
	return owner, nil
}

// Supply return total item
func (owner Owner) Supply() int {
	return len(owner.Items)
}

func (owner Owner) String() string {
	return fmt.Sprintf(`Address:%s
		Items:%s`,
		owner.Address,
		owner.Items,
	)
}

//// ========== SortedStringArray implementation ==========
//func (sa SortedStringArray) String() string {
//	return strings.Join(sa[:], ",")
//}
//
//func (sa SortedStringArray) ElAtIndex(index int) string {
//	return sa[index]
//}
//
//func (sa SortedStringArray) find(el string) (idx int) {
//	return FindUtil(sa, el)
//}
//
//func (sa SortedStringArray) Len() int {
//	return len(sa)
//}
//
//func (sa SortedStringArray) Less(i, j int) bool {
//	return strings.Compare(sa[i], sa[j]) == -1
//}
//
//func (sa SortedStringArray) Swap(i, j int) {
//	sa[i], sa[j] = sa[j], sa[i]
//}
//
//func (sa SortedStringArray) Sort() SortedStringArray {
//	sort.Sort(sa)
//	return sa
//}
