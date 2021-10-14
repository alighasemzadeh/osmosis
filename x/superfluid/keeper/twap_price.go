package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/osmosis-labs/osmosis/x/superfluid/types"
)

func (k Keeper) SetEpochOsmoEquivalentTWAP(ctx sdk.Context, epoch int64, poolId uint64, price sdk.Dec) {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, types.TokenPriceTwapEpochPrefix(epoch))
	priceRecord := types.EpochOsmoEquivalentTWAP{
		Epoch:  epoch,
		PoolId: poolId,
		Price:  price,
	}
	bz, err := proto.Marshal(&priceRecord)
	if err != nil {
		panic(err)
	}
	prefixStore.Set(sdk.Uint64ToBigEndian(poolId), bz)
}

func (k Keeper) DeleteEpochOsmoEquivalentTWAP(ctx sdk.Context, epoch int64, poolId uint64) {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, types.TokenPriceTwapEpochPrefix(epoch))
	prefixStore.Delete(sdk.Uint64ToBigEndian(poolId))
}

func (k Keeper) GetEpochOsmoEquivalentTWAP(ctx sdk.Context, epoch int64, poolId uint64) sdk.Dec {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, types.TokenPriceTwapEpochPrefix(epoch))
	bz := prefixStore.Get(sdk.Uint64ToBigEndian(poolId))
	if bz == nil {
		return sdk.ZeroDec()
	}
	priceRecord := types.EpochOsmoEquivalentTWAP{}
	err := proto.Unmarshal(bz, &priceRecord)
	if err != nil {
		panic(err)
	}
	return priceRecord.Price
}

func (k Keeper) GetAllEpochOsmoEquivalentTWAPs(ctx sdk.Context, epoch int64) []types.EpochOsmoEquivalentTWAP {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, types.TokenPriceTwapEpochPrefix(epoch))
	iterator := prefixStore.Iterator(nil, nil)
	defer iterator.Close()

	priceRecords := []types.EpochOsmoEquivalentTWAP{}
	for ; iterator.Valid(); iterator.Next() {
		priceRecord := types.EpochOsmoEquivalentTWAP{}

		err := proto.Unmarshal(iterator.Value(), &priceRecord)
		if err != nil {
			panic(err)
		}

		priceRecords = append(priceRecords, priceRecord)
	}
	return priceRecords
}

func (k Keeper) GetAllOsmoEquivalentTWAPs(ctx sdk.Context) []types.EpochOsmoEquivalentTWAP {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, types.KeyPrefixTokenPriceTwap)
	iterator := prefixStore.Iterator(nil, nil)
	defer iterator.Close()

	priceRecords := []types.EpochOsmoEquivalentTWAP{}
	for ; iterator.Valid(); iterator.Next() {
		priceRecord := types.EpochOsmoEquivalentTWAP{}

		err := proto.Unmarshal(iterator.Value(), &priceRecord)
		if err != nil {
			panic(err)
		}

		priceRecords = append(priceRecords, priceRecord)
	}
	return priceRecords
}