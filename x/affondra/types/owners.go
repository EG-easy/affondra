package types

import (
	"fmt"
	"sort"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ItemCollection struct {
	IDs SortedStringArray `json:"ids" yaml:"ids"`
}

type SortedStringArray []string

type Owner struct {
	Address         sdk.AccAddress   `json:"address" yaml:"address"`
	ItemCollections []ItemCollection `json:"itemCollections" yaml:"itemCollections"`
}

func NewItemCollection(ids []string) ItemCollection {
	return ItemCollection{
		IDs: SortedStringArray(ids).Sort(),
	}
}

func (owner Owner) Supply() int {
	total := 0
	for _, itemCollection := range owner.ItemCollections {
		total += itemCollection.Supply()
	}
	return total
}

func NewOwner(owner sdk.AccAddress, itemCollections ...ItemCollection) Owner {
	return Owner{
		Address:         owner,
		ItemCollections: itemCollections,
	}
}

// ========== SortedStringArray implementation ==========
func (sa SortedStringArray) String() string {
	return strings.Join(sa[:], ",")
}

//func (itemCollection ItemCollection) Exsits(id string) (exsits bool) {
//	index := itemCollection.IDs.find(id)
//	return index != -1
//}

//func (itemCollection ItemCollection) AddID(id string) ItemCollection {
//	itemCollection.IDs = append(itemCollection.IDs, id).Sort()
//	return itemCollection
//}

// ========== itemCollection implementation ==========
func (itemCollection ItemCollection) Supply() int {
	return len(itemCollection.IDs)
}

func (itemCollection ItemCollection) String() string {
	return fmt.Sprintf(`IDs:%s`, strings.Join(itemCollection.IDs, ","))
}

// ========== Sort implementation ==========
func (sa SortedStringArray) Len() int {
	return len(sa)
}

func (sa SortedStringArray) Less(i, j int) bool {
	return strings.Compare(sa[i], sa[j]) == -1
}

func (sa SortedStringArray) Swap(i, j int) {
	sa[i], sa[j] = sa[j], sa[i]
}

func (sa SortedStringArray) Sort() SortedStringArray {
	sort.Sort(sa)
	return sa
}
