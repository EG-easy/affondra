package types

import (
	"fmt"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Collection is a set of items categorized by denom
type Collection struct {
	Denom string `json:"denom,omitempty" yaml:"denom"`
	Items Items  `json:"items" yaml:"items"`
}

// NewCollection is a constructor function for Collection
func NewCollection(denom string, items Items) Collection {
	return Collection{
		Denom: strings.TrimSpace(denom),
		Items: items,
	}
}

// EmptyCollection returns an empty collection
func EmptyCollection() Collection {
	return NewCollection("", NewItems())
}

// GetItem gets a Item from the collection
func (collection Collection) GetItem(id string) (item Item, err error) {
	item, found := collection.Items.Find(id)
	if found {
		return item, nil
	}
	return Item{}, sdkerrors.Wrap(ErrUnknownItem,
		fmt.Sprintf("Item #%s doesn't exist in collection %s", id, collection.Denom),
	)
}

// ContainsItem returns whether or not a Collection contains an Item
func (collection Collection) ContainsItem(id string) bool {
	_, err := collection.GetItem(id)
	return err == nil
}

// AddItem adds an Item to the collection
func (collection Collection) AddItem(item Item) (Collection, error) {
	id := item.GetID()
	exists := collection.ContainsItem(id)
	if exists {
		return collection, sdkerrors.Wrap(ErrItemAlreadyExists,
			fmt.Sprintf("Item #%s already exists in collection %s", id, collection.Denom),
		)
	}
	collection.Items = collection.Items.Append(item)
	return collection, nil
}

// UpdateItem updates an Item from a collection
func (collection Collection) UpdateItem(item Item) (Collection, error) {
	items, ok := collection.Items.Update(item.GetID(), item)

	if !ok {
		return collection, sdkerrors.Wrap(ErrUnknownItem,
			fmt.Sprintf("Item #%s doesn't exist on collection %s", item.GetID(), collection.Denom),
		)
	}
	collection.Items = items
	return collection, nil
}

// DeleteItem deletes an Item from a collection
func (collection Collection) DeleteItem(item Item) (Collection, error) {
	items, ok := collection.Items.Remove(item.GetID())
	if !ok {
		return collection, sdkerrors.Wrap(ErrUnknownItem,
			fmt.Sprintf("Item #%s doesn't exist on collection %s", item.GetID(), collection.Denom),
		)
	}
	collection.Items = items
	return collection, nil
}

// Supply gets the total supply of Items of a collection
func (collection Collection) Supply() int {
	return len(collection.Items)
}

// String follows stringer interface
func (collection Collection) String() string {
	return fmt.Sprintf(`Denom: 				%s
Items:

%s`,
		collection.Denom,
		collection.Items.String(),
	)
}
