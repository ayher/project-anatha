package fee

import (
	"github.com/ayher/project-anatha/x/fee/internal/keeper"
	"github.com/ayher/project-anatha/x/fee/internal/types"
)

const (
	ModuleName               = types.ModuleName
	RouterKey                = types.RouterKey
	StoreKey                 = types.StoreKey
	DefaultParamspace        = types.DefaultParamspace
	QuerierRoute             = types.QuerierRoute
)

var (
	// functions aliases
	NewKeeper                          = keeper.NewKeeper
	NewQuerier                         = keeper.NewQuerier
	RegisterCodec                      = types.RegisterCodec
	NewGenesisState                    = types.NewGenesisState
	DefaultGenesisState                = types.DefaultGenesisState
	ValidateGenesis                    = types.ValidateGenesis

	NewParams                          = types.NewParams

	// variable aliases
	ModuleCdc     = types.ModuleCdc
)

type (
	Keeper       = keeper.Keeper
	GenesisState = types.GenesisState
	Params       = types.Params
)