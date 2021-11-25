package types

import sdk "github.com/ayher/anatha/types"

type NameHooks interface {
	AfterFirstNameCreated(ctx sdk.Context, address sdk.AccAddress) error
	AfterLastNameRemoved(ctx sdk.Context, address sdk.AccAddress) error
}
