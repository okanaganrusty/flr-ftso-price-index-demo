// Not all of these ABIs are needed, but just generating them so I can use in other projects
//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --abi internal/contract/FtsoManager.abi --pkg flare --type FtsoManager --out internal/flare/ftso_manager.go
//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --abi internal/contract/FtsoRegistry.abi --pkg flare --type FtsoRegistry --out internal/flare/ftso_registry.go
//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --abi internal/contract/FtsoRegistryProxy.abi --pkg flare --type FtsoRegistryProxy --out internal/flare/ftso_registry_proxy.go
//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --abi internal/contract/FtsoRewardManager.abi --pkg flare --type FtsoRewardManager --out internal/flare/ftso_reward_manager.go
//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --abi internal/contract/FlareAssetRegistry.abi --pkg flare --type FlareAssetRegistry --out internal/flare/flare_asset_registry.go
//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --abi internal/contract/FlareContractRegistry.abi --pkg flare --type FlareContractRegistry --out internal/flare/flare_contract_registry.go

package bindings
