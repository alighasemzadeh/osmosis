package pool_yield

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/gov/simulation"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/c-osmosis/osmosis/x/pool-yield/keeper"
	"github.com/c-osmosis/osmosis/x/pool-yield/types"
)

var (
	_ module.AppModule           = AppModule{}
	_ module.AppModuleBasic      = AppModuleBasic{}
	_ module.AppModuleSimulation = AppModule{}
)

type AppModuleBasic struct {
	cdc codec.Marshaler
}

// Name returns the pool-yield module's name.
func (AppModuleBasic) Name() string { return types.ModuleName }

// RegisterLegacyAminoCodec registers the pool-yield module's types on the LegacyAmino codec.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	// noop
}

// DefaultGenesis returns default genesis state as raw bytes for the pool-yield
// module.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONMarshaler) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesisState())
}

// ValidateGenesis performs genesis state validation for the pool-yield module.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONMarshaler, config client.TxEncodingConfig, bz json.RawMessage) error {
	var data types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &data); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	}

	return types.ValidateGenesis(&data)
}

//---------------------------------------
// Interfaces
func (b AppModuleBasic) RegisterRESTRoutes(ctx client.Context, r *mux.Router) {
	// noop
}

func (b AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx))
}

func (b AppModuleBasic) GetTxCmd() *cobra.Command {
	return nil
}

func (b AppModuleBasic) GetQueryCmd() *cobra.Command {
	// TODO
	return nil
}

// RegisterInterfaces registers interfaces and implementations of the pool-yield module.
func (AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	// noop
}

type AppModule struct {
	AppModuleBasic

	keeper keeper.Keeper
}

func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterQueryServer(cfg.QueryServer(), am.keeper)
}

func NewAppModule(cdc codec.Marshaler, keeper keeper.Keeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{cdc: cdc},
		keeper:         keeper,
	}
}

// RegisterInvariants registers the pool-yield module invariants.
func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {
	return
}

// Route returns the message routing key for the pool-yield module.
func (am AppModule) Route() sdk.Route {
	return sdk.NewRoute(types.RouterKey, nil)
}

// QuerierRoute returns the pool-yield module's querier route name.
func (AppModule) QuerierRoute() string { return types.RouterKey }

// LegacyQuerierHandler returns the pool-yield module sdk.Querier.
func (am AppModule) LegacyQuerierHandler(legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return nil
}

// InitGenesis performs genesis initialization for the pool-yield module.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONMarshaler, data json.RawMessage) []abci.ValidatorUpdate {
	var genesisState types.GenesisState

	cdc.MustUnmarshalJSON(data, &genesisState)

	am.keeper.InitGenesis(ctx, &genesisState)
	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns the exported genesis state as raw bytes for the pool-yield
// module.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONMarshaler) json.RawMessage {
	return nil
}

// BeginBlock performs a no-op.
func (AppModule) BeginBlock(_ sdk.Context, _ abci.RequestBeginBlock) {}

// EndBlock returns the end blocker for the pool-yield module. It returns no validator
// updates.
func (AppModule) EndBlock(_ sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

//____________________________________________________________________________

// AppModuleSimulation functions

// GenerateGenesisState creates a randomized GenState of the pool-yield module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	simulation.RandomizedGenState(simState)
}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(simState module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized pool-yield param changes for the simulator.
func (AppModule) RandomizedParams(r *rand.Rand) []simtypes.ParamChange {
	return simulation.ParamChanges(r)
}

// RegisterStoreDecoder registers a decoder for supply module's types
func (am AppModule) RegisterStoreDecoder(sdr sdk.StoreDecoderRegistry) {
	return
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	return nil // TODO
}