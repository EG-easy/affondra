package types

import (
	"fmt"
	"sort"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type SortedStringArray []string

type Owner struct {
	Address sdk.AccAddress    `json:"address" yaml:"address"`
	IDs     SortedStringArray `json:"ids" yaml:"ids"`
}

func NewOwner(owner sdk.AccAddress, ids []string) Owner {
	return Owner{
		Address: owner,
		IDs:     SortedStringArray(ids).Sort(),
	}
}

func (owner Owner) Exists(id string) (exsits bool) {
	index := owner.IDs.find(id)
	return index != -1
}

func (owner Owner) AddID(id string) Owner {
	owner.IDs = append(owner.IDs, id).Sort()
	return owner
}

func (owner Owner) DeleteID(id string) (Owner, error) {
	index := owner.IDs.find(id)
	if index == -1 {
		return owner, sdkerrors.Wrap(ErrUnknownID, fmt.Sprintf("Item #%s doesn't exsit in this Owner %s", id, owner.Address))
	}

	owner.IDs = append(owner.IDs[:index], owner.IDs[index+1:]...)

	return owner, nil
}

func (owner Owner) Supply() int {
	return len(owner.IDs)
}

func (owner Owner) String() string {
	return fmt.Sprintf(`Address:%s
		IDs:%s`,
		owner.Address,
		strings.Join(owner.IDs, ","),
	)
}

// ========== SortedStringArray implementation ==========
func (sa SortedStringArray) String() string {
	return strings.Join(sa[:], ",")
}

func (sa SortedStringArray) ElAtIndex(index int) string {
	return sa[index]
}

func (sa SortedStringArray) find(el string) (idx int) {
	return FindUtil(sa, el)
}

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
