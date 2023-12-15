package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hashicorp/golang-lru/v2/expirable"
	"github.com/stretchr/testify/assert"

	"github.com/okanaganrusty/flr-ftso-price-index-demo/internal/flare"
)

const (
	// https://docs.flare.network/dev/getting-started/contract-addresses/
	FLARE_CONTRACT_REGISTRY           = "0xaD67FE66660Fb8dFE9d6b1b4240d8650e30F6019"
	FLARE_FTSO_REGISTRY_CONTRACT_NAME = "FtsoRegistry"
	FLARE_FTSO_RPC_ADDRESS            = "https://flare-api.flare.network/ext/C/rpc"
)

var cache *expirable.LRU[string, float64]

func main() {
	ctx := context.Background()

	// Make cache with 10ms TTL and 5 max keys
	cache = expirable.NewLRU[string, float64](50, nil, time.Second*10)

	bindCallerOptions := &bind.CallOpts{
		Context: ctx,
	}

	client, err := ethclient.Dial(FLARE_FTSO_RPC_ADDRESS)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	t := new(testing.T)

	if cache.Len() == 0 {
		networkId, err := client.NetworkID(ctx)
		if err != nil {
			log.Fatal(err)
		}

		blockNumber, err := client.BlockNumber(ctx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Network ID: %d\n", networkId)
		fmt.Printf("Current Block number: %d\n", blockNumber)

		registry, err := flare.NewFlareContractRegistry(common.HexToAddress(FLARE_CONTRACT_REGISTRY), client)
		if err != nil {
			log.Fatal("NewFlareContractRegistry", err)
		}

		ftsoRegistryContactAddress, err := registry.GetContractAddressByName(&bind.CallOpts{}, FLARE_FTSO_REGISTRY_CONTRACT_NAME)
		if err != nil {
			log.Fatal("GetContractAddressByName", err)
		}

		ftsoRegistryCaller, err := flare.NewFtsoRegistry(ftsoRegistryContactAddress, client)

		if err != nil {
			log.Fatal("NewFtsoRegistry", err)
		}

		supportedIndiciesAndSymbols, err := ftsoRegistryCaller.GetSupportedIndicesAndSymbols(bindCallerOptions)
		if err != nil {
			log.Fatal("GetSupportedIndicesAndSymbols", err)
		}

		assert.Equal(t,
			len(supportedIndiciesAndSymbols.SupportedSymbols),
			len(supportedIndiciesAndSymbols.SupportedIndices),
		)

		symbolMapping := make(map[int64]string, 0)

		for i := 0; i < len(supportedIndiciesAndSymbols.SupportedIndices); i++ {
			symbolMapping[supportedIndiciesAndSymbols.SupportedIndices[i].Int64()] = supportedIndiciesAndSymbols.SupportedSymbols[i]
		}

		symbolPrices, err := ftsoRegistryCaller.GetCurrentPricesBySymbols(bindCallerOptions, supportedIndiciesAndSymbols.SupportedSymbols)

		if err != nil {
			log.Fatal("GetCurrentPricesBySymbols", err)
		}

		for _, symbolPrice := range symbolPrices {
			currentPrice := float64(symbolPrice.Price.Int64()) / math.Pow(float64(10), float64(symbolPrice.Decimals.Int64()))

			ftsoIndex := symbolPrice.FtsoIndex.Int64()

			cache.Add(symbolMapping[ftsoIndex], currentPrice)
		}

		for _, key := range cache.Keys() {
			value, _ := cache.Get(key)

			fmt.Printf("Coin: %s, price: %f\n",
				key,
				value,
			)
		}
	}
}
