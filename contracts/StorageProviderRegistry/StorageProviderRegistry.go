// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package StorageProviderRegistry

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// StorageProviderRegistryMetaData contains all meta data concerning the StorageProviderRegistry contract.
var StorageProviderRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ActivePool\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"errorCode\",\"type\":\"int256\"}],\"name\":\"ActorError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ActorNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllocationOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailToCallActor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailToCallActor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InactiveActor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InactivePool\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InactiveSP\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAccess\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"CommonTypes.FilActorId\",\"name\":\"actorId\",\"type\":\"uint64\"}],\"name\":\"InvalidActorID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAllocation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"InvalidCodec\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDailyAllocation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRepayment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidResponseLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"NotEnoughBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerProposed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegisteredSP\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"CollateralAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"LiquidStakingPoolRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"ownerId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"name\":\"StorageProviderAccruedRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"ownerId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allocationLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dailyAllocation\",\"type\":\"uint256\"}],\"name\":\"StorageProviderAllocationLimitRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"ownerId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allocationLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dailyAllocation\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayment\",\"type\":\"uint256\"}],\"name\":\"StorageProviderAllocationLimitUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"ownerId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usedAllocation\",\"type\":\"uint256\"}],\"name\":\"StorageProviderAllocationUsed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"ownerId\",\"type\":\"uint64\"}],\"name\":\"StorageProviderBeneficiaryAddressAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"}],\"name\":\"StorageProviderBeneficiaryAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"ownerId\",\"type\":\"uint64\"}],\"name\":\"StorageProviderDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"ownerId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"int64\",\"name\":\"lastEpoch\",\"type\":\"int64\"}],\"name\":\"StorageProviderLastEpochUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"ownerId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"miner\",\"type\":\"uint64\"}],\"name\":\"StorageProviderMinerAddressUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"ownerId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"restakingRatio\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"restakingAddress\",\"type\":\"address\"}],\"name\":\"StorageProviderMinerRestakingRatioUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"ownerId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"minerId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allocationLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dailyAllocation\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayment\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int64\",\"name\":\"lastEpoch\",\"type\":\"int64\"}],\"name\":\"StorageProviderOnboarded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"owner\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"ownerId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"minerId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"targetPool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allocationLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dailyAllocation\",\"type\":\"uint256\"}],\"name\":\"StorageProviderRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"ownerId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pledge\",\"type\":\"uint256\"}],\"name\":\"StorageProviderRepaidPledge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxAllocation\",\"type\":\"uint256\"}],\"name\":\"UpdateMaxAllocation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ownerId\",\"type\":\"uint64\"}],\"name\":\"acceptBeneficiaryAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"allocationRequests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allocationLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dailyAllocation\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"allocations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allocationLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"repayment\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usedAllocation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dailyAllocation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accruedRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"repaidPledge\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"beneficiaryStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"dailyUsages\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ownerId\",\"type\":\"uint64\"}],\"name\":\"deactivateStorageProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ownerId\",\"type\":\"uint64\"}],\"name\":\"getRepayment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ownerId\",\"type\":\"uint64\"}],\"name\":\"getStorageProvider\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"int64\",\"name\":\"\",\"type\":\"int64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ownerId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_repaidPledge\",\"type\":\"uint256\"}],\"name\":\"increasePledgeRepayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ownerId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_accuredRewards\",\"type\":\"uint256\"}],\"name\":\"increaseRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ownerId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_allocated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"increaseUsedAllocation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxAllocation\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_resolver\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"isActivePool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ownerId\",\"type\":\"uint64\"}],\"name\":\"isActiveProvider\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxAllocation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_minerId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_allocationLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dailyAllocation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_repayment\",\"type\":\"uint256\"},{\"internalType\":\"int64\",\"name\":\"_lastEpoch\",\"type\":\"int64\"}],\"name\":\"onboardStorageProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pools\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_minerId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_allocationLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dailyAllocation\",\"type\":\"uint256\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"registerPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_allocationLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dailyAllocation\",\"type\":\"uint256\"}],\"name\":\"requestAllocationLimitUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"restakings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"restakingRatio\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"restakingAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"sectorSizes\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ownerId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_minerId\",\"type\":\"uint64\"}],\"name\":\"setMinerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_restakingRatio\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_restakingAddress\",\"type\":\"address\"}],\"name\":\"setRestaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"storageProviders\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onboarded\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"targetPool\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"minerId\",\"type\":\"uint64\"},{\"internalType\":\"int64\",\"name\":\"lastEpoch\",\"type\":\"int64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ownerId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_allocationLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dailyAllocation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_repaymentAmount\",\"type\":\"uint256\"}],\"name\":\"updateAllocationLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"allocation\",\"type\":\"uint256\"}],\"name\":\"updateMaxAllocation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// StorageProviderRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageProviderRegistryMetaData.ABI instead.
var StorageProviderRegistryABI = StorageProviderRegistryMetaData.ABI

// StorageProviderRegistry is an auto generated Go binding around an Ethereum contract.
type StorageProviderRegistry struct {
	StorageProviderRegistryCaller     // Read-only binding to the contract
	StorageProviderRegistryTransactor // Write-only binding to the contract
	StorageProviderRegistryFilterer   // Log filterer for contract events
}

// StorageProviderRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageProviderRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageProviderRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageProviderRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageProviderRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageProviderRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageProviderRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageProviderRegistrySession struct {
	Contract     *StorageProviderRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StorageProviderRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageProviderRegistryCallerSession struct {
	Contract *StorageProviderRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// StorageProviderRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageProviderRegistryTransactorSession struct {
	Contract     *StorageProviderRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// StorageProviderRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageProviderRegistryRaw struct {
	Contract *StorageProviderRegistry // Generic contract binding to access the raw methods on
}

// StorageProviderRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageProviderRegistryCallerRaw struct {
	Contract *StorageProviderRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// StorageProviderRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageProviderRegistryTransactorRaw struct {
	Contract *StorageProviderRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorageProviderRegistry creates a new instance of StorageProviderRegistry, bound to a specific deployed contract.
func NewStorageProviderRegistry(address common.Address, backend bind.ContractBackend) (*StorageProviderRegistry, error) {
	contract, err := bindStorageProviderRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StorageProviderRegistry{StorageProviderRegistryCaller: StorageProviderRegistryCaller{contract: contract}, StorageProviderRegistryTransactor: StorageProviderRegistryTransactor{contract: contract}, StorageProviderRegistryFilterer: StorageProviderRegistryFilterer{contract: contract}}, nil
}

// NewStorageProviderRegistryCaller creates a new read-only instance of StorageProviderRegistry, bound to a specific deployed contract.
func NewStorageProviderRegistryCaller(address common.Address, caller bind.ContractCaller) (*StorageProviderRegistryCaller, error) {
	contract, err := bindStorageProviderRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageProviderRegistryCaller{contract: contract}, nil
}

// NewStorageProviderRegistryTransactor creates a new write-only instance of StorageProviderRegistry, bound to a specific deployed contract.
func NewStorageProviderRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageProviderRegistryTransactor, error) {
	contract, err := bindStorageProviderRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageProviderRegistryTransactor{contract: contract}, nil
}

// NewStorageProviderRegistryFilterer creates a new log filterer instance of StorageProviderRegistry, bound to a specific deployed contract.
func NewStorageProviderRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageProviderRegistryFilterer, error) {
	contract, err := bindStorageProviderRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageProviderRegistryFilterer{contract: contract}, nil
}

// bindStorageProviderRegistry binds a generic wrapper to an already deployed contract.
func bindStorageProviderRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StorageProviderRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StorageProviderRegistry *StorageProviderRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StorageProviderRegistry.Contract.StorageProviderRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StorageProviderRegistry *StorageProviderRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.StorageProviderRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StorageProviderRegistry *StorageProviderRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.StorageProviderRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StorageProviderRegistry *StorageProviderRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StorageProviderRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StorageProviderRegistry *StorageProviderRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StorageProviderRegistry *StorageProviderRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StorageProviderRegistry *StorageProviderRegistryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StorageProviderRegistry.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StorageProviderRegistry *StorageProviderRegistrySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StorageProviderRegistry.Contract.DEFAULTADMINROLE(&_StorageProviderRegistry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StorageProviderRegistry *StorageProviderRegistryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StorageProviderRegistry.Contract.DEFAULTADMINROLE(&_StorageProviderRegistry.CallOpts)
}

// AllocationRequests is a free data retrieval call binding the contract method 0x853d941d.
//
// Solidity: function allocationRequests(uint64 ) view returns(uint256 allocationLimit, uint256 dailyAllocation)
func (_StorageProviderRegistry *StorageProviderRegistryCaller) AllocationRequests(opts *bind.CallOpts, arg0 uint64) (struct {
	AllocationLimit *big.Int
	DailyAllocation *big.Int
}, error) {
	var out []interface{}
	err := _StorageProviderRegistry.contract.Call(opts, &out, "allocationRequests", arg0)

	outstruct := new(struct {
		AllocationLimit *big.Int
		DailyAllocation *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AllocationLimit = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.DailyAllocation = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// AllocationRequests is a free data retrieval call binding the contract method 0x853d941d.
//
// Solidity: function allocationRequests(uint64 ) view returns(uint256 allocationLimit, uint256 dailyAllocation)
func (_StorageProviderRegistry *StorageProviderRegistrySession) AllocationRequests(arg0 uint64) (struct {
	AllocationLimit *big.Int
	DailyAllocation *big.Int
}, error) {
	return _StorageProviderRegistry.Contract.AllocationRequests(&_StorageProviderRegistry.CallOpts, arg0)
}

// AllocationRequests is a free data retrieval call binding the contract method 0x853d941d.
//
// Solidity: function allocationRequests(uint64 ) view returns(uint256 allocationLimit, uint256 dailyAllocation)
func (_StorageProviderRegistry *StorageProviderRegistryCallerSession) AllocationRequests(arg0 uint64) (struct {
	AllocationLimit *big.Int
	DailyAllocation *big.Int
}, error) {
	return _StorageProviderRegistry.Contract.AllocationRequests(&_StorageProviderRegistry.CallOpts, arg0)
}

// Allocations is a free data retrieval call binding the contract method 0x2e54941e.
//
// Solidity: function allocations(uint64 ) view returns(uint256 allocationLimit, uint256 repayment, uint256 usedAllocation, uint256 dailyAllocation, uint256 accruedRewards, uint256 repaidPledge)
func (_StorageProviderRegistry *StorageProviderRegistryCaller) Allocations(opts *bind.CallOpts, arg0 uint64) (struct {
	AllocationLimit *big.Int
	Repayment       *big.Int
	UsedAllocation  *big.Int
	DailyAllocation *big.Int
	AccruedRewards  *big.Int
	RepaidPledge    *big.Int
}, error) {
	var out []interface{}
	err := _StorageProviderRegistry.contract.Call(opts, &out, "allocations", arg0)

	outstruct := new(struct {
		AllocationLimit *big.Int
		Repayment       *big.Int
		UsedAllocation  *big.Int
		DailyAllocation *big.Int
		AccruedRewards  *big.Int
		RepaidPledge    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AllocationLimit = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Repayment = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.UsedAllocation = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.DailyAllocation = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AccruedRewards = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.RepaidPledge = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Allocations is a free data retrieval call binding the contract method 0x2e54941e.
//
// Solidity: function allocations(uint64 ) view returns(uint256 allocationLimit, uint256 repayment, uint256 usedAllocation, uint256 dailyAllocation, uint256 accruedRewards, uint256 repaidPledge)
func (_StorageProviderRegistry *StorageProviderRegistrySession) Allocations(arg0 uint64) (struct {
	AllocationLimit *big.Int
	Repayment       *big.Int
	UsedAllocation  *big.Int
	DailyAllocation *big.Int
	AccruedRewards  *big.Int
	RepaidPledge    *big.Int
}, error) {
	return _StorageProviderRegistry.Contract.Allocations(&_StorageProviderRegistry.CallOpts, arg0)
}

// Allocations is a free data retrieval call binding the contract method 0x2e54941e.
//
// Solidity: function allocations(uint64 ) view returns(uint256 allocationLimit, uint256 repayment, uint256 usedAllocation, uint256 dailyAllocation, uint256 accruedRewards, uint256 repaidPledge)
func (_StorageProviderRegistry *StorageProviderRegistryCallerSession) Allocations(arg0 uint64) (struct {
	AllocationLimit *big.Int
	Repayment       *big.Int
	UsedAllocation  *big.Int
	DailyAllocation *big.Int
	AccruedRewards  *big.Int
	RepaidPledge    *big.Int
}, error) {
	return _StorageProviderRegistry.Contract.Allocations(&_StorageProviderRegistry.CallOpts, arg0)
}

// BeneficiaryStatus is a free data retrieval call binding the contract method 0x0610f3ad.
//
// Solidity: function beneficiaryStatus(uint64 ) view returns(bool)
func (_StorageProviderRegistry *StorageProviderRegistryCaller) BeneficiaryStatus(opts *bind.CallOpts, arg0 uint64) (bool, error) {
	var out []interface{}
	err := _StorageProviderRegistry.contract.Call(opts, &out, "beneficiaryStatus", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BeneficiaryStatus is a free data retrieval call binding the contract method 0x0610f3ad.
//
// Solidity: function beneficiaryStatus(uint64 ) view returns(bool)
func (_StorageProviderRegistry *StorageProviderRegistrySession) BeneficiaryStatus(arg0 uint64) (bool, error) {
	return _StorageProviderRegistry.Contract.BeneficiaryStatus(&_StorageProviderRegistry.CallOpts, arg0)
}

// BeneficiaryStatus is a free data retrieval call binding the contract method 0x0610f3ad.
//
// Solidity: function beneficiaryStatus(uint64 ) view returns(bool)
func (_StorageProviderRegistry *StorageProviderRegistryCallerSession) BeneficiaryStatus(arg0 uint64) (bool, error) {
	return _StorageProviderRegistry.Contract.BeneficiaryStatus(&_StorageProviderRegistry.CallOpts, arg0)
}

// DailyUsages is a free data retrieval call binding the contract method 0xd2d3164b.
//
// Solidity: function dailyUsages(bytes32 ) view returns(uint256)
func (_StorageProviderRegistry *StorageProviderRegistryCaller) DailyUsages(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _StorageProviderRegistry.contract.Call(opts, &out, "dailyUsages", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DailyUsages is a free data retrieval call binding the contract method 0xd2d3164b.
//
// Solidity: function dailyUsages(bytes32 ) view returns(uint256)
func (_StorageProviderRegistry *StorageProviderRegistrySession) DailyUsages(arg0 [32]byte) (*big.Int, error) {
	return _StorageProviderRegistry.Contract.DailyUsages(&_StorageProviderRegistry.CallOpts, arg0)
}

// DailyUsages is a free data retrieval call binding the contract method 0xd2d3164b.
//
// Solidity: function dailyUsages(bytes32 ) view returns(uint256)
func (_StorageProviderRegistry *StorageProviderRegistryCallerSession) DailyUsages(arg0 [32]byte) (*big.Int, error) {
	return _StorageProviderRegistry.Contract.DailyUsages(&_StorageProviderRegistry.CallOpts, arg0)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_StorageProviderRegistry *StorageProviderRegistryCaller) GetImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StorageProviderRegistry.contract.Call(opts, &out, "getImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_StorageProviderRegistry *StorageProviderRegistrySession) GetImplementation() (common.Address, error) {
	return _StorageProviderRegistry.Contract.GetImplementation(&_StorageProviderRegistry.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_StorageProviderRegistry *StorageProviderRegistryCallerSession) GetImplementation() (common.Address, error) {
	return _StorageProviderRegistry.Contract.GetImplementation(&_StorageProviderRegistry.CallOpts)
}

// GetRepayment is a free data retrieval call binding the contract method 0x67f47a94.
//
// Solidity: function getRepayment(uint64 _ownerId) view returns(uint256)
func (_StorageProviderRegistry *StorageProviderRegistryCaller) GetRepayment(opts *bind.CallOpts, _ownerId uint64) (*big.Int, error) {
	var out []interface{}
	err := _StorageProviderRegistry.contract.Call(opts, &out, "getRepayment", _ownerId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRepayment is a free data retrieval call binding the contract method 0x67f47a94.
//
// Solidity: function getRepayment(uint64 _ownerId) view returns(uint256)
func (_StorageProviderRegistry *StorageProviderRegistrySession) GetRepayment(_ownerId uint64) (*big.Int, error) {
	return _StorageProviderRegistry.Contract.GetRepayment(&_StorageProviderRegistry.CallOpts, _ownerId)
}

// GetRepayment is a free data retrieval call binding the contract method 0x67f47a94.
//
// Solidity: function getRepayment(uint64 _ownerId) view returns(uint256)
func (_StorageProviderRegistry *StorageProviderRegistryCallerSession) GetRepayment(_ownerId uint64) (*big.Int, error) {
	return _StorageProviderRegistry.Contract.GetRepayment(&_StorageProviderRegistry.CallOpts, _ownerId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StorageProviderRegistry *StorageProviderRegistryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _StorageProviderRegistry.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StorageProviderRegistry *StorageProviderRegistrySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StorageProviderRegistry.Contract.GetRoleAdmin(&_StorageProviderRegistry.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StorageProviderRegistry *StorageProviderRegistryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StorageProviderRegistry.Contract.GetRoleAdmin(&_StorageProviderRegistry.CallOpts, role)
}

// GetStorageProvider is a free data retrieval call binding the contract method 0xb0298c31.
//
// Solidity: function getStorageProvider(uint64 _ownerId) view returns(bool, address, uint64, int64)
func (_StorageProviderRegistry *StorageProviderRegistryCaller) GetStorageProvider(opts *bind.CallOpts, _ownerId uint64) (bool, common.Address, uint64, int64, error) {
	var out []interface{}
	err := _StorageProviderRegistry.contract.Call(opts, &out, "getStorageProvider", _ownerId)

	if err != nil {
		return *new(bool), *new(common.Address), *new(uint64), *new(int64), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(uint64)).(*uint64)
	out3 := *abi.ConvertType(out[3], new(int64)).(*int64)

	return out0, out1, out2, out3, err

}

// GetStorageProvider is a free data retrieval call binding the contract method 0xb0298c31.
//
// Solidity: function getStorageProvider(uint64 _ownerId) view returns(bool, address, uint64, int64)
func (_StorageProviderRegistry *StorageProviderRegistrySession) GetStorageProvider(_ownerId uint64) (bool, common.Address, uint64, int64, error) {
	return _StorageProviderRegistry.Contract.GetStorageProvider(&_StorageProviderRegistry.CallOpts, _ownerId)
}

// GetStorageProvider is a free data retrieval call binding the contract method 0xb0298c31.
//
// Solidity: function getStorageProvider(uint64 _ownerId) view returns(bool, address, uint64, int64)
func (_StorageProviderRegistry *StorageProviderRegistryCallerSession) GetStorageProvider(_ownerId uint64) (bool, common.Address, uint64, int64, error) {
	return _StorageProviderRegistry.Contract.GetStorageProvider(&_StorageProviderRegistry.CallOpts, _ownerId)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StorageProviderRegistry *StorageProviderRegistryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _StorageProviderRegistry.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StorageProviderRegistry *StorageProviderRegistrySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StorageProviderRegistry.Contract.HasRole(&_StorageProviderRegistry.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StorageProviderRegistry *StorageProviderRegistryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StorageProviderRegistry.Contract.HasRole(&_StorageProviderRegistry.CallOpts, role, account)
}

// IsActivePool is a free data retrieval call binding the contract method 0x8097354f.
//
// Solidity: function isActivePool(address _pool) view returns(bool)
func (_StorageProviderRegistry *StorageProviderRegistryCaller) IsActivePool(opts *bind.CallOpts, _pool common.Address) (bool, error) {
	var out []interface{}
	err := _StorageProviderRegistry.contract.Call(opts, &out, "isActivePool", _pool)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActivePool is a free data retrieval call binding the contract method 0x8097354f.
//
// Solidity: function isActivePool(address _pool) view returns(bool)
func (_StorageProviderRegistry *StorageProviderRegistrySession) IsActivePool(_pool common.Address) (bool, error) {
	return _StorageProviderRegistry.Contract.IsActivePool(&_StorageProviderRegistry.CallOpts, _pool)
}

// IsActivePool is a free data retrieval call binding the contract method 0x8097354f.
//
// Solidity: function isActivePool(address _pool) view returns(bool)
func (_StorageProviderRegistry *StorageProviderRegistryCallerSession) IsActivePool(_pool common.Address) (bool, error) {
	return _StorageProviderRegistry.Contract.IsActivePool(&_StorageProviderRegistry.CallOpts, _pool)
}

// IsActiveProvider is a free data retrieval call binding the contract method 0x154caabb.
//
// Solidity: function isActiveProvider(uint64 _ownerId) view returns(bool status)
func (_StorageProviderRegistry *StorageProviderRegistryCaller) IsActiveProvider(opts *bind.CallOpts, _ownerId uint64) (bool, error) {
	var out []interface{}
	err := _StorageProviderRegistry.contract.Call(opts, &out, "isActiveProvider", _ownerId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActiveProvider is a free data retrieval call binding the contract method 0x154caabb.
//
// Solidity: function isActiveProvider(uint64 _ownerId) view returns(bool status)
func (_StorageProviderRegistry *StorageProviderRegistrySession) IsActiveProvider(_ownerId uint64) (bool, error) {
	return _StorageProviderRegistry.Contract.IsActiveProvider(&_StorageProviderRegistry.CallOpts, _ownerId)
}

// IsActiveProvider is a free data retrieval call binding the contract method 0x154caabb.
//
// Solidity: function isActiveProvider(uint64 _ownerId) view returns(bool status)
func (_StorageProviderRegistry *StorageProviderRegistryCallerSession) IsActiveProvider(_ownerId uint64) (bool, error) {
	return _StorageProviderRegistry.Contract.IsActiveProvider(&_StorageProviderRegistry.CallOpts, _ownerId)
}

// MaxAllocation is a free data retrieval call binding the contract method 0x9b3ba79f.
//
// Solidity: function maxAllocation() view returns(uint256)
func (_StorageProviderRegistry *StorageProviderRegistryCaller) MaxAllocation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StorageProviderRegistry.contract.Call(opts, &out, "maxAllocation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAllocation is a free data retrieval call binding the contract method 0x9b3ba79f.
//
// Solidity: function maxAllocation() view returns(uint256)
func (_StorageProviderRegistry *StorageProviderRegistrySession) MaxAllocation() (*big.Int, error) {
	return _StorageProviderRegistry.Contract.MaxAllocation(&_StorageProviderRegistry.CallOpts)
}

// MaxAllocation is a free data retrieval call binding the contract method 0x9b3ba79f.
//
// Solidity: function maxAllocation() view returns(uint256)
func (_StorageProviderRegistry *StorageProviderRegistryCallerSession) MaxAllocation() (*big.Int, error) {
	return _StorageProviderRegistry.Contract.MaxAllocation(&_StorageProviderRegistry.CallOpts)
}

// Pools is a free data retrieval call binding the contract method 0xa4063dbc.
//
// Solidity: function pools(address ) view returns(bool)
func (_StorageProviderRegistry *StorageProviderRegistryCaller) Pools(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _StorageProviderRegistry.contract.Call(opts, &out, "pools", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Pools is a free data retrieval call binding the contract method 0xa4063dbc.
//
// Solidity: function pools(address ) view returns(bool)
func (_StorageProviderRegistry *StorageProviderRegistrySession) Pools(arg0 common.Address) (bool, error) {
	return _StorageProviderRegistry.Contract.Pools(&_StorageProviderRegistry.CallOpts, arg0)
}

// Pools is a free data retrieval call binding the contract method 0xa4063dbc.
//
// Solidity: function pools(address ) view returns(bool)
func (_StorageProviderRegistry *StorageProviderRegistryCallerSession) Pools(arg0 common.Address) (bool, error) {
	return _StorageProviderRegistry.Contract.Pools(&_StorageProviderRegistry.CallOpts, arg0)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_StorageProviderRegistry *StorageProviderRegistryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StorageProviderRegistry.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_StorageProviderRegistry *StorageProviderRegistrySession) ProxiableUUID() ([32]byte, error) {
	return _StorageProviderRegistry.Contract.ProxiableUUID(&_StorageProviderRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_StorageProviderRegistry *StorageProviderRegistryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _StorageProviderRegistry.Contract.ProxiableUUID(&_StorageProviderRegistry.CallOpts)
}

// Restakings is a free data retrieval call binding the contract method 0x9e5c1399.
//
// Solidity: function restakings(uint64 ) view returns(uint256 restakingRatio, address restakingAddress)
func (_StorageProviderRegistry *StorageProviderRegistryCaller) Restakings(opts *bind.CallOpts, arg0 uint64) (struct {
	RestakingRatio   *big.Int
	RestakingAddress common.Address
}, error) {
	var out []interface{}
	err := _StorageProviderRegistry.contract.Call(opts, &out, "restakings", arg0)

	outstruct := new(struct {
		RestakingRatio   *big.Int
		RestakingAddress common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RestakingRatio = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RestakingAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Restakings is a free data retrieval call binding the contract method 0x9e5c1399.
//
// Solidity: function restakings(uint64 ) view returns(uint256 restakingRatio, address restakingAddress)
func (_StorageProviderRegistry *StorageProviderRegistrySession) Restakings(arg0 uint64) (struct {
	RestakingRatio   *big.Int
	RestakingAddress common.Address
}, error) {
	return _StorageProviderRegistry.Contract.Restakings(&_StorageProviderRegistry.CallOpts, arg0)
}

// Restakings is a free data retrieval call binding the contract method 0x9e5c1399.
//
// Solidity: function restakings(uint64 ) view returns(uint256 restakingRatio, address restakingAddress)
func (_StorageProviderRegistry *StorageProviderRegistryCallerSession) Restakings(arg0 uint64) (struct {
	RestakingRatio   *big.Int
	RestakingAddress common.Address
}, error) {
	return _StorageProviderRegistry.Contract.Restakings(&_StorageProviderRegistry.CallOpts, arg0)
}

// SectorSizes is a free data retrieval call binding the contract method 0x76246e33.
//
// Solidity: function sectorSizes(uint64 ) view returns(uint64)
func (_StorageProviderRegistry *StorageProviderRegistryCaller) SectorSizes(opts *bind.CallOpts, arg0 uint64) (uint64, error) {
	var out []interface{}
	err := _StorageProviderRegistry.contract.Call(opts, &out, "sectorSizes", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// SectorSizes is a free data retrieval call binding the contract method 0x76246e33.
//
// Solidity: function sectorSizes(uint64 ) view returns(uint64)
func (_StorageProviderRegistry *StorageProviderRegistrySession) SectorSizes(arg0 uint64) (uint64, error) {
	return _StorageProviderRegistry.Contract.SectorSizes(&_StorageProviderRegistry.CallOpts, arg0)
}

// SectorSizes is a free data retrieval call binding the contract method 0x76246e33.
//
// Solidity: function sectorSizes(uint64 ) view returns(uint64)
func (_StorageProviderRegistry *StorageProviderRegistryCallerSession) SectorSizes(arg0 uint64) (uint64, error) {
	return _StorageProviderRegistry.Contract.SectorSizes(&_StorageProviderRegistry.CallOpts, arg0)
}

// StorageProviders is a free data retrieval call binding the contract method 0xdd6460ce.
//
// Solidity: function storageProviders(uint64 ) view returns(bool active, bool onboarded, address targetPool, uint64 minerId, int64 lastEpoch)
func (_StorageProviderRegistry *StorageProviderRegistryCaller) StorageProviders(opts *bind.CallOpts, arg0 uint64) (struct {
	Active     bool
	Onboarded  bool
	TargetPool common.Address
	MinerId    uint64
	LastEpoch  int64
}, error) {
	var out []interface{}
	err := _StorageProviderRegistry.contract.Call(opts, &out, "storageProviders", arg0)

	outstruct := new(struct {
		Active     bool
		Onboarded  bool
		TargetPool common.Address
		MinerId    uint64
		LastEpoch  int64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Active = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Onboarded = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.TargetPool = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.MinerId = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.LastEpoch = *abi.ConvertType(out[4], new(int64)).(*int64)

	return *outstruct, err

}

// StorageProviders is a free data retrieval call binding the contract method 0xdd6460ce.
//
// Solidity: function storageProviders(uint64 ) view returns(bool active, bool onboarded, address targetPool, uint64 minerId, int64 lastEpoch)
func (_StorageProviderRegistry *StorageProviderRegistrySession) StorageProviders(arg0 uint64) (struct {
	Active     bool
	Onboarded  bool
	TargetPool common.Address
	MinerId    uint64
	LastEpoch  int64
}, error) {
	return _StorageProviderRegistry.Contract.StorageProviders(&_StorageProviderRegistry.CallOpts, arg0)
}

// StorageProviders is a free data retrieval call binding the contract method 0xdd6460ce.
//
// Solidity: function storageProviders(uint64 ) view returns(bool active, bool onboarded, address targetPool, uint64 minerId, int64 lastEpoch)
func (_StorageProviderRegistry *StorageProviderRegistryCallerSession) StorageProviders(arg0 uint64) (struct {
	Active     bool
	Onboarded  bool
	TargetPool common.Address
	MinerId    uint64
	LastEpoch  int64
}, error) {
	return _StorageProviderRegistry.Contract.StorageProviders(&_StorageProviderRegistry.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StorageProviderRegistry *StorageProviderRegistryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _StorageProviderRegistry.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StorageProviderRegistry *StorageProviderRegistrySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StorageProviderRegistry.Contract.SupportsInterface(&_StorageProviderRegistry.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StorageProviderRegistry *StorageProviderRegistryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StorageProviderRegistry.Contract.SupportsInterface(&_StorageProviderRegistry.CallOpts, interfaceId)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_StorageProviderRegistry *StorageProviderRegistryCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StorageProviderRegistry.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_StorageProviderRegistry *StorageProviderRegistrySession) Version() (string, error) {
	return _StorageProviderRegistry.Contract.Version(&_StorageProviderRegistry.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_StorageProviderRegistry *StorageProviderRegistryCallerSession) Version() (string, error) {
	return _StorageProviderRegistry.Contract.Version(&_StorageProviderRegistry.CallOpts)
}

// AcceptBeneficiaryAddress is a paid mutator transaction binding the contract method 0xfa2755c3.
//
// Solidity: function acceptBeneficiaryAddress(uint64 _ownerId) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactor) AcceptBeneficiaryAddress(opts *bind.TransactOpts, _ownerId uint64) (*types.Transaction, error) {
	return _StorageProviderRegistry.contract.Transact(opts, "acceptBeneficiaryAddress", _ownerId)
}

// AcceptBeneficiaryAddress is a paid mutator transaction binding the contract method 0xfa2755c3.
//
// Solidity: function acceptBeneficiaryAddress(uint64 _ownerId) returns()
func (_StorageProviderRegistry *StorageProviderRegistrySession) AcceptBeneficiaryAddress(_ownerId uint64) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.AcceptBeneficiaryAddress(&_StorageProviderRegistry.TransactOpts, _ownerId)
}

// AcceptBeneficiaryAddress is a paid mutator transaction binding the contract method 0xfa2755c3.
//
// Solidity: function acceptBeneficiaryAddress(uint64 _ownerId) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactorSession) AcceptBeneficiaryAddress(_ownerId uint64) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.AcceptBeneficiaryAddress(&_StorageProviderRegistry.TransactOpts, _ownerId)
}

// DeactivateStorageProvider is a paid mutator transaction binding the contract method 0x4f85bf32.
//
// Solidity: function deactivateStorageProvider(uint64 _ownerId) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactor) DeactivateStorageProvider(opts *bind.TransactOpts, _ownerId uint64) (*types.Transaction, error) {
	return _StorageProviderRegistry.contract.Transact(opts, "deactivateStorageProvider", _ownerId)
}

// DeactivateStorageProvider is a paid mutator transaction binding the contract method 0x4f85bf32.
//
// Solidity: function deactivateStorageProvider(uint64 _ownerId) returns()
func (_StorageProviderRegistry *StorageProviderRegistrySession) DeactivateStorageProvider(_ownerId uint64) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.DeactivateStorageProvider(&_StorageProviderRegistry.TransactOpts, _ownerId)
}

// DeactivateStorageProvider is a paid mutator transaction binding the contract method 0x4f85bf32.
//
// Solidity: function deactivateStorageProvider(uint64 _ownerId) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactorSession) DeactivateStorageProvider(_ownerId uint64) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.DeactivateStorageProvider(&_StorageProviderRegistry.TransactOpts, _ownerId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StorageProviderRegistry.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StorageProviderRegistry *StorageProviderRegistrySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.GrantRole(&_StorageProviderRegistry.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.GrantRole(&_StorageProviderRegistry.TransactOpts, role, account)
}

// IncreasePledgeRepayment is a paid mutator transaction binding the contract method 0x261d8315.
//
// Solidity: function increasePledgeRepayment(uint64 _ownerId, uint256 _repaidPledge) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactor) IncreasePledgeRepayment(opts *bind.TransactOpts, _ownerId uint64, _repaidPledge *big.Int) (*types.Transaction, error) {
	return _StorageProviderRegistry.contract.Transact(opts, "increasePledgeRepayment", _ownerId, _repaidPledge)
}

// IncreasePledgeRepayment is a paid mutator transaction binding the contract method 0x261d8315.
//
// Solidity: function increasePledgeRepayment(uint64 _ownerId, uint256 _repaidPledge) returns()
func (_StorageProviderRegistry *StorageProviderRegistrySession) IncreasePledgeRepayment(_ownerId uint64, _repaidPledge *big.Int) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.IncreasePledgeRepayment(&_StorageProviderRegistry.TransactOpts, _ownerId, _repaidPledge)
}

// IncreasePledgeRepayment is a paid mutator transaction binding the contract method 0x261d8315.
//
// Solidity: function increasePledgeRepayment(uint64 _ownerId, uint256 _repaidPledge) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactorSession) IncreasePledgeRepayment(_ownerId uint64, _repaidPledge *big.Int) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.IncreasePledgeRepayment(&_StorageProviderRegistry.TransactOpts, _ownerId, _repaidPledge)
}

// IncreaseRewards is a paid mutator transaction binding the contract method 0x7f2d1a60.
//
// Solidity: function increaseRewards(uint64 _ownerId, uint256 _accuredRewards) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactor) IncreaseRewards(opts *bind.TransactOpts, _ownerId uint64, _accuredRewards *big.Int) (*types.Transaction, error) {
	return _StorageProviderRegistry.contract.Transact(opts, "increaseRewards", _ownerId, _accuredRewards)
}

// IncreaseRewards is a paid mutator transaction binding the contract method 0x7f2d1a60.
//
// Solidity: function increaseRewards(uint64 _ownerId, uint256 _accuredRewards) returns()
func (_StorageProviderRegistry *StorageProviderRegistrySession) IncreaseRewards(_ownerId uint64, _accuredRewards *big.Int) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.IncreaseRewards(&_StorageProviderRegistry.TransactOpts, _ownerId, _accuredRewards)
}

// IncreaseRewards is a paid mutator transaction binding the contract method 0x7f2d1a60.
//
// Solidity: function increaseRewards(uint64 _ownerId, uint256 _accuredRewards) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactorSession) IncreaseRewards(_ownerId uint64, _accuredRewards *big.Int) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.IncreaseRewards(&_StorageProviderRegistry.TransactOpts, _ownerId, _accuredRewards)
}

// IncreaseUsedAllocation is a paid mutator transaction binding the contract method 0x4978e4b7.
//
// Solidity: function increaseUsedAllocation(uint64 _ownerId, uint256 _allocated, uint256 _timestamp) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactor) IncreaseUsedAllocation(opts *bind.TransactOpts, _ownerId uint64, _allocated *big.Int, _timestamp *big.Int) (*types.Transaction, error) {
	return _StorageProviderRegistry.contract.Transact(opts, "increaseUsedAllocation", _ownerId, _allocated, _timestamp)
}

// IncreaseUsedAllocation is a paid mutator transaction binding the contract method 0x4978e4b7.
//
// Solidity: function increaseUsedAllocation(uint64 _ownerId, uint256 _allocated, uint256 _timestamp) returns()
func (_StorageProviderRegistry *StorageProviderRegistrySession) IncreaseUsedAllocation(_ownerId uint64, _allocated *big.Int, _timestamp *big.Int) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.IncreaseUsedAllocation(&_StorageProviderRegistry.TransactOpts, _ownerId, _allocated, _timestamp)
}

// IncreaseUsedAllocation is a paid mutator transaction binding the contract method 0x4978e4b7.
//
// Solidity: function increaseUsedAllocation(uint64 _ownerId, uint256 _allocated, uint256 _timestamp) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactorSession) IncreaseUsedAllocation(_ownerId uint64, _allocated *big.Int, _timestamp *big.Int) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.IncreaseUsedAllocation(&_StorageProviderRegistry.TransactOpts, _ownerId, _allocated, _timestamp)
}

// Initialize is a paid mutator transaction binding the contract method 0xda35a26f.
//
// Solidity: function initialize(uint256 _maxAllocation, address _resolver) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactor) Initialize(opts *bind.TransactOpts, _maxAllocation *big.Int, _resolver common.Address) (*types.Transaction, error) {
	return _StorageProviderRegistry.contract.Transact(opts, "initialize", _maxAllocation, _resolver)
}

// Initialize is a paid mutator transaction binding the contract method 0xda35a26f.
//
// Solidity: function initialize(uint256 _maxAllocation, address _resolver) returns()
func (_StorageProviderRegistry *StorageProviderRegistrySession) Initialize(_maxAllocation *big.Int, _resolver common.Address) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.Initialize(&_StorageProviderRegistry.TransactOpts, _maxAllocation, _resolver)
}

// Initialize is a paid mutator transaction binding the contract method 0xda35a26f.
//
// Solidity: function initialize(uint256 _maxAllocation, address _resolver) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactorSession) Initialize(_maxAllocation *big.Int, _resolver common.Address) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.Initialize(&_StorageProviderRegistry.TransactOpts, _maxAllocation, _resolver)
}

// OnboardStorageProvider is a paid mutator transaction binding the contract method 0x0e2fc9e1.
//
// Solidity: function onboardStorageProvider(uint64 _minerId, uint256 _allocationLimit, uint256 _dailyAllocation, uint256 _repayment, int64 _lastEpoch) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactor) OnboardStorageProvider(opts *bind.TransactOpts, _minerId uint64, _allocationLimit *big.Int, _dailyAllocation *big.Int, _repayment *big.Int, _lastEpoch int64) (*types.Transaction, error) {
	return _StorageProviderRegistry.contract.Transact(opts, "onboardStorageProvider", _minerId, _allocationLimit, _dailyAllocation, _repayment, _lastEpoch)
}

// OnboardStorageProvider is a paid mutator transaction binding the contract method 0x0e2fc9e1.
//
// Solidity: function onboardStorageProvider(uint64 _minerId, uint256 _allocationLimit, uint256 _dailyAllocation, uint256 _repayment, int64 _lastEpoch) returns()
func (_StorageProviderRegistry *StorageProviderRegistrySession) OnboardStorageProvider(_minerId uint64, _allocationLimit *big.Int, _dailyAllocation *big.Int, _repayment *big.Int, _lastEpoch int64) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.OnboardStorageProvider(&_StorageProviderRegistry.TransactOpts, _minerId, _allocationLimit, _dailyAllocation, _repayment, _lastEpoch)
}

// OnboardStorageProvider is a paid mutator transaction binding the contract method 0x0e2fc9e1.
//
// Solidity: function onboardStorageProvider(uint64 _minerId, uint256 _allocationLimit, uint256 _dailyAllocation, uint256 _repayment, int64 _lastEpoch) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactorSession) OnboardStorageProvider(_minerId uint64, _allocationLimit *big.Int, _dailyAllocation *big.Int, _repayment *big.Int, _lastEpoch int64) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.OnboardStorageProvider(&_StorageProviderRegistry.TransactOpts, _minerId, _allocationLimit, _dailyAllocation, _repayment, _lastEpoch)
}

// Register is a paid mutator transaction binding the contract method 0x31ea602b.
//
// Solidity: function register(uint64 _minerId, uint256 _allocationLimit, uint256 _dailyAllocation) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactor) Register(opts *bind.TransactOpts, _minerId uint64, _allocationLimit *big.Int, _dailyAllocation *big.Int) (*types.Transaction, error) {
	return _StorageProviderRegistry.contract.Transact(opts, "register", _minerId, _allocationLimit, _dailyAllocation)
}

// Register is a paid mutator transaction binding the contract method 0x31ea602b.
//
// Solidity: function register(uint64 _minerId, uint256 _allocationLimit, uint256 _dailyAllocation) returns()
func (_StorageProviderRegistry *StorageProviderRegistrySession) Register(_minerId uint64, _allocationLimit *big.Int, _dailyAllocation *big.Int) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.Register(&_StorageProviderRegistry.TransactOpts, _minerId, _allocationLimit, _dailyAllocation)
}

// Register is a paid mutator transaction binding the contract method 0x31ea602b.
//
// Solidity: function register(uint64 _minerId, uint256 _allocationLimit, uint256 _dailyAllocation) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactorSession) Register(_minerId uint64, _allocationLimit *big.Int, _dailyAllocation *big.Int) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.Register(&_StorageProviderRegistry.TransactOpts, _minerId, _allocationLimit, _dailyAllocation)
}

// RegisterPool is a paid mutator transaction binding the contract method 0xabd90846.
//
// Solidity: function registerPool(address _pool) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactor) RegisterPool(opts *bind.TransactOpts, _pool common.Address) (*types.Transaction, error) {
	return _StorageProviderRegistry.contract.Transact(opts, "registerPool", _pool)
}

// RegisterPool is a paid mutator transaction binding the contract method 0xabd90846.
//
// Solidity: function registerPool(address _pool) returns()
func (_StorageProviderRegistry *StorageProviderRegistrySession) RegisterPool(_pool common.Address) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.RegisterPool(&_StorageProviderRegistry.TransactOpts, _pool)
}

// RegisterPool is a paid mutator transaction binding the contract method 0xabd90846.
//
// Solidity: function registerPool(address _pool) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactorSession) RegisterPool(_pool common.Address) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.RegisterPool(&_StorageProviderRegistry.TransactOpts, _pool)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StorageProviderRegistry.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StorageProviderRegistry *StorageProviderRegistrySession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.RenounceRole(&_StorageProviderRegistry.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.RenounceRole(&_StorageProviderRegistry.TransactOpts, role, account)
}

// RequestAllocationLimitUpdate is a paid mutator transaction binding the contract method 0xa93fbf86.
//
// Solidity: function requestAllocationLimitUpdate(uint256 _allocationLimit, uint256 _dailyAllocation) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactor) RequestAllocationLimitUpdate(opts *bind.TransactOpts, _allocationLimit *big.Int, _dailyAllocation *big.Int) (*types.Transaction, error) {
	return _StorageProviderRegistry.contract.Transact(opts, "requestAllocationLimitUpdate", _allocationLimit, _dailyAllocation)
}

// RequestAllocationLimitUpdate is a paid mutator transaction binding the contract method 0xa93fbf86.
//
// Solidity: function requestAllocationLimitUpdate(uint256 _allocationLimit, uint256 _dailyAllocation) returns()
func (_StorageProviderRegistry *StorageProviderRegistrySession) RequestAllocationLimitUpdate(_allocationLimit *big.Int, _dailyAllocation *big.Int) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.RequestAllocationLimitUpdate(&_StorageProviderRegistry.TransactOpts, _allocationLimit, _dailyAllocation)
}

// RequestAllocationLimitUpdate is a paid mutator transaction binding the contract method 0xa93fbf86.
//
// Solidity: function requestAllocationLimitUpdate(uint256 _allocationLimit, uint256 _dailyAllocation) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactorSession) RequestAllocationLimitUpdate(_allocationLimit *big.Int, _dailyAllocation *big.Int) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.RequestAllocationLimitUpdate(&_StorageProviderRegistry.TransactOpts, _allocationLimit, _dailyAllocation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StorageProviderRegistry.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StorageProviderRegistry *StorageProviderRegistrySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.RevokeRole(&_StorageProviderRegistry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.RevokeRole(&_StorageProviderRegistry.TransactOpts, role, account)
}

// SetMinerAddress is a paid mutator transaction binding the contract method 0x2d0c628a.
//
// Solidity: function setMinerAddress(uint64 _ownerId, uint64 _minerId) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactor) SetMinerAddress(opts *bind.TransactOpts, _ownerId uint64, _minerId uint64) (*types.Transaction, error) {
	return _StorageProviderRegistry.contract.Transact(opts, "setMinerAddress", _ownerId, _minerId)
}

// SetMinerAddress is a paid mutator transaction binding the contract method 0x2d0c628a.
//
// Solidity: function setMinerAddress(uint64 _ownerId, uint64 _minerId) returns()
func (_StorageProviderRegistry *StorageProviderRegistrySession) SetMinerAddress(_ownerId uint64, _minerId uint64) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.SetMinerAddress(&_StorageProviderRegistry.TransactOpts, _ownerId, _minerId)
}

// SetMinerAddress is a paid mutator transaction binding the contract method 0x2d0c628a.
//
// Solidity: function setMinerAddress(uint64 _ownerId, uint64 _minerId) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactorSession) SetMinerAddress(_ownerId uint64, _minerId uint64) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.SetMinerAddress(&_StorageProviderRegistry.TransactOpts, _ownerId, _minerId)
}

// SetRestaking is a paid mutator transaction binding the contract method 0xbe7aebcd.
//
// Solidity: function setRestaking(uint256 _restakingRatio, address _restakingAddress) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactor) SetRestaking(opts *bind.TransactOpts, _restakingRatio *big.Int, _restakingAddress common.Address) (*types.Transaction, error) {
	return _StorageProviderRegistry.contract.Transact(opts, "setRestaking", _restakingRatio, _restakingAddress)
}

// SetRestaking is a paid mutator transaction binding the contract method 0xbe7aebcd.
//
// Solidity: function setRestaking(uint256 _restakingRatio, address _restakingAddress) returns()
func (_StorageProviderRegistry *StorageProviderRegistrySession) SetRestaking(_restakingRatio *big.Int, _restakingAddress common.Address) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.SetRestaking(&_StorageProviderRegistry.TransactOpts, _restakingRatio, _restakingAddress)
}

// SetRestaking is a paid mutator transaction binding the contract method 0xbe7aebcd.
//
// Solidity: function setRestaking(uint256 _restakingRatio, address _restakingAddress) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactorSession) SetRestaking(_restakingRatio *big.Int, _restakingAddress common.Address) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.SetRestaking(&_StorageProviderRegistry.TransactOpts, _restakingRatio, _restakingAddress)
}

// UpdateAllocationLimit is a paid mutator transaction binding the contract method 0x7aca42fe.
//
// Solidity: function updateAllocationLimit(uint64 _ownerId, uint256 _allocationLimit, uint256 _dailyAllocation, uint256 _repaymentAmount) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactor) UpdateAllocationLimit(opts *bind.TransactOpts, _ownerId uint64, _allocationLimit *big.Int, _dailyAllocation *big.Int, _repaymentAmount *big.Int) (*types.Transaction, error) {
	return _StorageProviderRegistry.contract.Transact(opts, "updateAllocationLimit", _ownerId, _allocationLimit, _dailyAllocation, _repaymentAmount)
}

// UpdateAllocationLimit is a paid mutator transaction binding the contract method 0x7aca42fe.
//
// Solidity: function updateAllocationLimit(uint64 _ownerId, uint256 _allocationLimit, uint256 _dailyAllocation, uint256 _repaymentAmount) returns()
func (_StorageProviderRegistry *StorageProviderRegistrySession) UpdateAllocationLimit(_ownerId uint64, _allocationLimit *big.Int, _dailyAllocation *big.Int, _repaymentAmount *big.Int) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.UpdateAllocationLimit(&_StorageProviderRegistry.TransactOpts, _ownerId, _allocationLimit, _dailyAllocation, _repaymentAmount)
}

// UpdateAllocationLimit is a paid mutator transaction binding the contract method 0x7aca42fe.
//
// Solidity: function updateAllocationLimit(uint64 _ownerId, uint256 _allocationLimit, uint256 _dailyAllocation, uint256 _repaymentAmount) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactorSession) UpdateAllocationLimit(_ownerId uint64, _allocationLimit *big.Int, _dailyAllocation *big.Int, _repaymentAmount *big.Int) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.UpdateAllocationLimit(&_StorageProviderRegistry.TransactOpts, _ownerId, _allocationLimit, _dailyAllocation, _repaymentAmount)
}

// UpdateMaxAllocation is a paid mutator transaction binding the contract method 0x24da48a3.
//
// Solidity: function updateMaxAllocation(uint256 allocation) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactor) UpdateMaxAllocation(opts *bind.TransactOpts, allocation *big.Int) (*types.Transaction, error) {
	return _StorageProviderRegistry.contract.Transact(opts, "updateMaxAllocation", allocation)
}

// UpdateMaxAllocation is a paid mutator transaction binding the contract method 0x24da48a3.
//
// Solidity: function updateMaxAllocation(uint256 allocation) returns()
func (_StorageProviderRegistry *StorageProviderRegistrySession) UpdateMaxAllocation(allocation *big.Int) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.UpdateMaxAllocation(&_StorageProviderRegistry.TransactOpts, allocation)
}

// UpdateMaxAllocation is a paid mutator transaction binding the contract method 0x24da48a3.
//
// Solidity: function updateMaxAllocation(uint256 allocation) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactorSession) UpdateMaxAllocation(allocation *big.Int) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.UpdateMaxAllocation(&_StorageProviderRegistry.TransactOpts, allocation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _StorageProviderRegistry.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_StorageProviderRegistry *StorageProviderRegistrySession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.UpgradeTo(&_StorageProviderRegistry.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.UpgradeTo(&_StorageProviderRegistry.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _StorageProviderRegistry.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_StorageProviderRegistry *StorageProviderRegistrySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.UpgradeToAndCall(&_StorageProviderRegistry.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.UpgradeToAndCall(&_StorageProviderRegistry.TransactOpts, newImplementation, data)
}

// StorageProviderRegistryAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the StorageProviderRegistry contract.
type StorageProviderRegistryAdminChangedIterator struct {
	Event *StorageProviderRegistryAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageProviderRegistryAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProviderRegistryAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageProviderRegistryAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageProviderRegistryAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProviderRegistryAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProviderRegistryAdminChanged represents a AdminChanged event raised by the StorageProviderRegistry contract.
type StorageProviderRegistryAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*StorageProviderRegistryAdminChangedIterator, error) {

	logs, sub, err := _StorageProviderRegistry.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &StorageProviderRegistryAdminChangedIterator{contract: _StorageProviderRegistry.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *StorageProviderRegistryAdminChanged) (event.Subscription, error) {

	logs, sub, err := _StorageProviderRegistry.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProviderRegistryAdminChanged)
				if err := _StorageProviderRegistry.contract.UnpackLog(event, "AdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) ParseAdminChanged(log types.Log) (*StorageProviderRegistryAdminChanged, error) {
	event := new(StorageProviderRegistryAdminChanged)
	if err := _StorageProviderRegistry.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageProviderRegistryBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the StorageProviderRegistry contract.
type StorageProviderRegistryBeaconUpgradedIterator struct {
	Event *StorageProviderRegistryBeaconUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageProviderRegistryBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProviderRegistryBeaconUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageProviderRegistryBeaconUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageProviderRegistryBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProviderRegistryBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProviderRegistryBeaconUpgraded represents a BeaconUpgraded event raised by the StorageProviderRegistry contract.
type StorageProviderRegistryBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*StorageProviderRegistryBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _StorageProviderRegistry.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &StorageProviderRegistryBeaconUpgradedIterator{contract: _StorageProviderRegistry.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *StorageProviderRegistryBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _StorageProviderRegistry.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProviderRegistryBeaconUpgraded)
				if err := _StorageProviderRegistry.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) ParseBeaconUpgraded(log types.Log) (*StorageProviderRegistryBeaconUpgraded, error) {
	event := new(StorageProviderRegistryBeaconUpgraded)
	if err := _StorageProviderRegistry.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageProviderRegistryCollateralAddressUpdatedIterator is returned from FilterCollateralAddressUpdated and is used to iterate over the raw logs and unpacked data for CollateralAddressUpdated events raised by the StorageProviderRegistry contract.
type StorageProviderRegistryCollateralAddressUpdatedIterator struct {
	Event *StorageProviderRegistryCollateralAddressUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageProviderRegistryCollateralAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProviderRegistryCollateralAddressUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageProviderRegistryCollateralAddressUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageProviderRegistryCollateralAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProviderRegistryCollateralAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProviderRegistryCollateralAddressUpdated represents a CollateralAddressUpdated event raised by the StorageProviderRegistry contract.
type StorageProviderRegistryCollateralAddressUpdated struct {
	Collateral common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCollateralAddressUpdated is a free log retrieval operation binding the contract event 0x823c58e2969bcdf1301b932fcbced62d5d4617880949dd434df13d001b3a1162.
//
// Solidity: event CollateralAddressUpdated(address collateral)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) FilterCollateralAddressUpdated(opts *bind.FilterOpts) (*StorageProviderRegistryCollateralAddressUpdatedIterator, error) {

	logs, sub, err := _StorageProviderRegistry.contract.FilterLogs(opts, "CollateralAddressUpdated")
	if err != nil {
		return nil, err
	}
	return &StorageProviderRegistryCollateralAddressUpdatedIterator{contract: _StorageProviderRegistry.contract, event: "CollateralAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchCollateralAddressUpdated is a free log subscription operation binding the contract event 0x823c58e2969bcdf1301b932fcbced62d5d4617880949dd434df13d001b3a1162.
//
// Solidity: event CollateralAddressUpdated(address collateral)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) WatchCollateralAddressUpdated(opts *bind.WatchOpts, sink chan<- *StorageProviderRegistryCollateralAddressUpdated) (event.Subscription, error) {

	logs, sub, err := _StorageProviderRegistry.contract.WatchLogs(opts, "CollateralAddressUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProviderRegistryCollateralAddressUpdated)
				if err := _StorageProviderRegistry.contract.UnpackLog(event, "CollateralAddressUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCollateralAddressUpdated is a log parse operation binding the contract event 0x823c58e2969bcdf1301b932fcbced62d5d4617880949dd434df13d001b3a1162.
//
// Solidity: event CollateralAddressUpdated(address collateral)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) ParseCollateralAddressUpdated(log types.Log) (*StorageProviderRegistryCollateralAddressUpdated, error) {
	event := new(StorageProviderRegistryCollateralAddressUpdated)
	if err := _StorageProviderRegistry.contract.UnpackLog(event, "CollateralAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageProviderRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StorageProviderRegistry contract.
type StorageProviderRegistryInitializedIterator struct {
	Event *StorageProviderRegistryInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageProviderRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProviderRegistryInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageProviderRegistryInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageProviderRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProviderRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProviderRegistryInitialized represents a Initialized event raised by the StorageProviderRegistry contract.
type StorageProviderRegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*StorageProviderRegistryInitializedIterator, error) {

	logs, sub, err := _StorageProviderRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StorageProviderRegistryInitializedIterator{contract: _StorageProviderRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StorageProviderRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _StorageProviderRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProviderRegistryInitialized)
				if err := _StorageProviderRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) ParseInitialized(log types.Log) (*StorageProviderRegistryInitialized, error) {
	event := new(StorageProviderRegistryInitialized)
	if err := _StorageProviderRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageProviderRegistryLiquidStakingPoolRegisteredIterator is returned from FilterLiquidStakingPoolRegistered and is used to iterate over the raw logs and unpacked data for LiquidStakingPoolRegistered events raised by the StorageProviderRegistry contract.
type StorageProviderRegistryLiquidStakingPoolRegisteredIterator struct {
	Event *StorageProviderRegistryLiquidStakingPoolRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageProviderRegistryLiquidStakingPoolRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProviderRegistryLiquidStakingPoolRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageProviderRegistryLiquidStakingPoolRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageProviderRegistryLiquidStakingPoolRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProviderRegistryLiquidStakingPoolRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProviderRegistryLiquidStakingPoolRegistered represents a LiquidStakingPoolRegistered event raised by the StorageProviderRegistry contract.
type StorageProviderRegistryLiquidStakingPoolRegistered struct {
	Pool common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLiquidStakingPoolRegistered is a free log retrieval operation binding the contract event 0xbba7dd87057c9346bdb652d4c49948e76139a24eb1975262e98556f0455466ac.
//
// Solidity: event LiquidStakingPoolRegistered(address pool)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) FilterLiquidStakingPoolRegistered(opts *bind.FilterOpts) (*StorageProviderRegistryLiquidStakingPoolRegisteredIterator, error) {

	logs, sub, err := _StorageProviderRegistry.contract.FilterLogs(opts, "LiquidStakingPoolRegistered")
	if err != nil {
		return nil, err
	}
	return &StorageProviderRegistryLiquidStakingPoolRegisteredIterator{contract: _StorageProviderRegistry.contract, event: "LiquidStakingPoolRegistered", logs: logs, sub: sub}, nil
}

// WatchLiquidStakingPoolRegistered is a free log subscription operation binding the contract event 0xbba7dd87057c9346bdb652d4c49948e76139a24eb1975262e98556f0455466ac.
//
// Solidity: event LiquidStakingPoolRegistered(address pool)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) WatchLiquidStakingPoolRegistered(opts *bind.WatchOpts, sink chan<- *StorageProviderRegistryLiquidStakingPoolRegistered) (event.Subscription, error) {

	logs, sub, err := _StorageProviderRegistry.contract.WatchLogs(opts, "LiquidStakingPoolRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProviderRegistryLiquidStakingPoolRegistered)
				if err := _StorageProviderRegistry.contract.UnpackLog(event, "LiquidStakingPoolRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLiquidStakingPoolRegistered is a log parse operation binding the contract event 0xbba7dd87057c9346bdb652d4c49948e76139a24eb1975262e98556f0455466ac.
//
// Solidity: event LiquidStakingPoolRegistered(address pool)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) ParseLiquidStakingPoolRegistered(log types.Log) (*StorageProviderRegistryLiquidStakingPoolRegistered, error) {
	event := new(StorageProviderRegistryLiquidStakingPoolRegistered)
	if err := _StorageProviderRegistry.contract.UnpackLog(event, "LiquidStakingPoolRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageProviderRegistryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the StorageProviderRegistry contract.
type StorageProviderRegistryRoleAdminChangedIterator struct {
	Event *StorageProviderRegistryRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageProviderRegistryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProviderRegistryRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageProviderRegistryRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageProviderRegistryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProviderRegistryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProviderRegistryRoleAdminChanged represents a RoleAdminChanged event raised by the StorageProviderRegistry contract.
type StorageProviderRegistryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*StorageProviderRegistryRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _StorageProviderRegistry.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &StorageProviderRegistryRoleAdminChangedIterator{contract: _StorageProviderRegistry.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *StorageProviderRegistryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _StorageProviderRegistry.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProviderRegistryRoleAdminChanged)
				if err := _StorageProviderRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) ParseRoleAdminChanged(log types.Log) (*StorageProviderRegistryRoleAdminChanged, error) {
	event := new(StorageProviderRegistryRoleAdminChanged)
	if err := _StorageProviderRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageProviderRegistryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the StorageProviderRegistry contract.
type StorageProviderRegistryRoleGrantedIterator struct {
	Event *StorageProviderRegistryRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageProviderRegistryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProviderRegistryRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageProviderRegistryRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageProviderRegistryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProviderRegistryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProviderRegistryRoleGranted represents a RoleGranted event raised by the StorageProviderRegistry contract.
type StorageProviderRegistryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StorageProviderRegistryRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _StorageProviderRegistry.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StorageProviderRegistryRoleGrantedIterator{contract: _StorageProviderRegistry.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *StorageProviderRegistryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _StorageProviderRegistry.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProviderRegistryRoleGranted)
				if err := _StorageProviderRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) ParseRoleGranted(log types.Log) (*StorageProviderRegistryRoleGranted, error) {
	event := new(StorageProviderRegistryRoleGranted)
	if err := _StorageProviderRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageProviderRegistryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the StorageProviderRegistry contract.
type StorageProviderRegistryRoleRevokedIterator struct {
	Event *StorageProviderRegistryRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageProviderRegistryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProviderRegistryRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageProviderRegistryRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageProviderRegistryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProviderRegistryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProviderRegistryRoleRevoked represents a RoleRevoked event raised by the StorageProviderRegistry contract.
type StorageProviderRegistryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StorageProviderRegistryRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _StorageProviderRegistry.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StorageProviderRegistryRoleRevokedIterator{contract: _StorageProviderRegistry.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *StorageProviderRegistryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _StorageProviderRegistry.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProviderRegistryRoleRevoked)
				if err := _StorageProviderRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) ParseRoleRevoked(log types.Log) (*StorageProviderRegistryRoleRevoked, error) {
	event := new(StorageProviderRegistryRoleRevoked)
	if err := _StorageProviderRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageProviderRegistryStorageProviderAccruedRewardsIterator is returned from FilterStorageProviderAccruedRewards and is used to iterate over the raw logs and unpacked data for StorageProviderAccruedRewards events raised by the StorageProviderRegistry contract.
type StorageProviderRegistryStorageProviderAccruedRewardsIterator struct {
	Event *StorageProviderRegistryStorageProviderAccruedRewards // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageProviderRegistryStorageProviderAccruedRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProviderRegistryStorageProviderAccruedRewards)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageProviderRegistryStorageProviderAccruedRewards)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageProviderRegistryStorageProviderAccruedRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProviderRegistryStorageProviderAccruedRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProviderRegistryStorageProviderAccruedRewards represents a StorageProviderAccruedRewards event raised by the StorageProviderRegistry contract.
type StorageProviderRegistryStorageProviderAccruedRewards struct {
	OwnerId uint64
	Rewards *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStorageProviderAccruedRewards is a free log retrieval operation binding the contract event 0xd3adca66167eaebe0bf908f97055d75ea826041f45a2ef406a2f03e07a65f758.
//
// Solidity: event StorageProviderAccruedRewards(uint64 ownerId, uint256 rewards)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) FilterStorageProviderAccruedRewards(opts *bind.FilterOpts) (*StorageProviderRegistryStorageProviderAccruedRewardsIterator, error) {

	logs, sub, err := _StorageProviderRegistry.contract.FilterLogs(opts, "StorageProviderAccruedRewards")
	if err != nil {
		return nil, err
	}
	return &StorageProviderRegistryStorageProviderAccruedRewardsIterator{contract: _StorageProviderRegistry.contract, event: "StorageProviderAccruedRewards", logs: logs, sub: sub}, nil
}

// WatchStorageProviderAccruedRewards is a free log subscription operation binding the contract event 0xd3adca66167eaebe0bf908f97055d75ea826041f45a2ef406a2f03e07a65f758.
//
// Solidity: event StorageProviderAccruedRewards(uint64 ownerId, uint256 rewards)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) WatchStorageProviderAccruedRewards(opts *bind.WatchOpts, sink chan<- *StorageProviderRegistryStorageProviderAccruedRewards) (event.Subscription, error) {

	logs, sub, err := _StorageProviderRegistry.contract.WatchLogs(opts, "StorageProviderAccruedRewards")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProviderRegistryStorageProviderAccruedRewards)
				if err := _StorageProviderRegistry.contract.UnpackLog(event, "StorageProviderAccruedRewards", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStorageProviderAccruedRewards is a log parse operation binding the contract event 0xd3adca66167eaebe0bf908f97055d75ea826041f45a2ef406a2f03e07a65f758.
//
// Solidity: event StorageProviderAccruedRewards(uint64 ownerId, uint256 rewards)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) ParseStorageProviderAccruedRewards(log types.Log) (*StorageProviderRegistryStorageProviderAccruedRewards, error) {
	event := new(StorageProviderRegistryStorageProviderAccruedRewards)
	if err := _StorageProviderRegistry.contract.UnpackLog(event, "StorageProviderAccruedRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageProviderRegistryStorageProviderAllocationLimitRequestIterator is returned from FilterStorageProviderAllocationLimitRequest and is used to iterate over the raw logs and unpacked data for StorageProviderAllocationLimitRequest events raised by the StorageProviderRegistry contract.
type StorageProviderRegistryStorageProviderAllocationLimitRequestIterator struct {
	Event *StorageProviderRegistryStorageProviderAllocationLimitRequest // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageProviderRegistryStorageProviderAllocationLimitRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProviderRegistryStorageProviderAllocationLimitRequest)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageProviderRegistryStorageProviderAllocationLimitRequest)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageProviderRegistryStorageProviderAllocationLimitRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProviderRegistryStorageProviderAllocationLimitRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProviderRegistryStorageProviderAllocationLimitRequest represents a StorageProviderAllocationLimitRequest event raised by the StorageProviderRegistry contract.
type StorageProviderRegistryStorageProviderAllocationLimitRequest struct {
	OwnerId         uint64
	AllocationLimit *big.Int
	DailyAllocation *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStorageProviderAllocationLimitRequest is a free log retrieval operation binding the contract event 0xd0198e2535194b2efbec973cb73349e6d1edd434a271491d44c2f620a997a3c3.
//
// Solidity: event StorageProviderAllocationLimitRequest(uint64 ownerId, uint256 allocationLimit, uint256 dailyAllocation)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) FilterStorageProviderAllocationLimitRequest(opts *bind.FilterOpts) (*StorageProviderRegistryStorageProviderAllocationLimitRequestIterator, error) {

	logs, sub, err := _StorageProviderRegistry.contract.FilterLogs(opts, "StorageProviderAllocationLimitRequest")
	if err != nil {
		return nil, err
	}
	return &StorageProviderRegistryStorageProviderAllocationLimitRequestIterator{contract: _StorageProviderRegistry.contract, event: "StorageProviderAllocationLimitRequest", logs: logs, sub: sub}, nil
}

// WatchStorageProviderAllocationLimitRequest is a free log subscription operation binding the contract event 0xd0198e2535194b2efbec973cb73349e6d1edd434a271491d44c2f620a997a3c3.
//
// Solidity: event StorageProviderAllocationLimitRequest(uint64 ownerId, uint256 allocationLimit, uint256 dailyAllocation)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) WatchStorageProviderAllocationLimitRequest(opts *bind.WatchOpts, sink chan<- *StorageProviderRegistryStorageProviderAllocationLimitRequest) (event.Subscription, error) {

	logs, sub, err := _StorageProviderRegistry.contract.WatchLogs(opts, "StorageProviderAllocationLimitRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProviderRegistryStorageProviderAllocationLimitRequest)
				if err := _StorageProviderRegistry.contract.UnpackLog(event, "StorageProviderAllocationLimitRequest", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStorageProviderAllocationLimitRequest is a log parse operation binding the contract event 0xd0198e2535194b2efbec973cb73349e6d1edd434a271491d44c2f620a997a3c3.
//
// Solidity: event StorageProviderAllocationLimitRequest(uint64 ownerId, uint256 allocationLimit, uint256 dailyAllocation)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) ParseStorageProviderAllocationLimitRequest(log types.Log) (*StorageProviderRegistryStorageProviderAllocationLimitRequest, error) {
	event := new(StorageProviderRegistryStorageProviderAllocationLimitRequest)
	if err := _StorageProviderRegistry.contract.UnpackLog(event, "StorageProviderAllocationLimitRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageProviderRegistryStorageProviderAllocationLimitUpdateIterator is returned from FilterStorageProviderAllocationLimitUpdate and is used to iterate over the raw logs and unpacked data for StorageProviderAllocationLimitUpdate events raised by the StorageProviderRegistry contract.
type StorageProviderRegistryStorageProviderAllocationLimitUpdateIterator struct {
	Event *StorageProviderRegistryStorageProviderAllocationLimitUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageProviderRegistryStorageProviderAllocationLimitUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProviderRegistryStorageProviderAllocationLimitUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageProviderRegistryStorageProviderAllocationLimitUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageProviderRegistryStorageProviderAllocationLimitUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProviderRegistryStorageProviderAllocationLimitUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProviderRegistryStorageProviderAllocationLimitUpdate represents a StorageProviderAllocationLimitUpdate event raised by the StorageProviderRegistry contract.
type StorageProviderRegistryStorageProviderAllocationLimitUpdate struct {
	OwnerId         uint64
	AllocationLimit *big.Int
	DailyAllocation *big.Int
	Repayment       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStorageProviderAllocationLimitUpdate is a free log retrieval operation binding the contract event 0x6eaa57f14bb4acd5ffc19285fec33d57c58d9fdf748cc9502219d2bded72f611.
//
// Solidity: event StorageProviderAllocationLimitUpdate(uint64 ownerId, uint256 allocationLimit, uint256 dailyAllocation, uint256 repayment)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) FilterStorageProviderAllocationLimitUpdate(opts *bind.FilterOpts) (*StorageProviderRegistryStorageProviderAllocationLimitUpdateIterator, error) {

	logs, sub, err := _StorageProviderRegistry.contract.FilterLogs(opts, "StorageProviderAllocationLimitUpdate")
	if err != nil {
		return nil, err
	}
	return &StorageProviderRegistryStorageProviderAllocationLimitUpdateIterator{contract: _StorageProviderRegistry.contract, event: "StorageProviderAllocationLimitUpdate", logs: logs, sub: sub}, nil
}

// WatchStorageProviderAllocationLimitUpdate is a free log subscription operation binding the contract event 0x6eaa57f14bb4acd5ffc19285fec33d57c58d9fdf748cc9502219d2bded72f611.
//
// Solidity: event StorageProviderAllocationLimitUpdate(uint64 ownerId, uint256 allocationLimit, uint256 dailyAllocation, uint256 repayment)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) WatchStorageProviderAllocationLimitUpdate(opts *bind.WatchOpts, sink chan<- *StorageProviderRegistryStorageProviderAllocationLimitUpdate) (event.Subscription, error) {

	logs, sub, err := _StorageProviderRegistry.contract.WatchLogs(opts, "StorageProviderAllocationLimitUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProviderRegistryStorageProviderAllocationLimitUpdate)
				if err := _StorageProviderRegistry.contract.UnpackLog(event, "StorageProviderAllocationLimitUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStorageProviderAllocationLimitUpdate is a log parse operation binding the contract event 0x6eaa57f14bb4acd5ffc19285fec33d57c58d9fdf748cc9502219d2bded72f611.
//
// Solidity: event StorageProviderAllocationLimitUpdate(uint64 ownerId, uint256 allocationLimit, uint256 dailyAllocation, uint256 repayment)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) ParseStorageProviderAllocationLimitUpdate(log types.Log) (*StorageProviderRegistryStorageProviderAllocationLimitUpdate, error) {
	event := new(StorageProviderRegistryStorageProviderAllocationLimitUpdate)
	if err := _StorageProviderRegistry.contract.UnpackLog(event, "StorageProviderAllocationLimitUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageProviderRegistryStorageProviderAllocationUsedIterator is returned from FilterStorageProviderAllocationUsed and is used to iterate over the raw logs and unpacked data for StorageProviderAllocationUsed events raised by the StorageProviderRegistry contract.
type StorageProviderRegistryStorageProviderAllocationUsedIterator struct {
	Event *StorageProviderRegistryStorageProviderAllocationUsed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageProviderRegistryStorageProviderAllocationUsedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProviderRegistryStorageProviderAllocationUsed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageProviderRegistryStorageProviderAllocationUsed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageProviderRegistryStorageProviderAllocationUsedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProviderRegistryStorageProviderAllocationUsedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProviderRegistryStorageProviderAllocationUsed represents a StorageProviderAllocationUsed event raised by the StorageProviderRegistry contract.
type StorageProviderRegistryStorageProviderAllocationUsed struct {
	OwnerId        uint64
	UsedAllocation *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterStorageProviderAllocationUsed is a free log retrieval operation binding the contract event 0x5fc112892cdeb5d0ea544d634fb55a2eb4ba066d9ad730b67f34791462ebf51b.
//
// Solidity: event StorageProviderAllocationUsed(uint64 ownerId, uint256 usedAllocation)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) FilterStorageProviderAllocationUsed(opts *bind.FilterOpts) (*StorageProviderRegistryStorageProviderAllocationUsedIterator, error) {

	logs, sub, err := _StorageProviderRegistry.contract.FilterLogs(opts, "StorageProviderAllocationUsed")
	if err != nil {
		return nil, err
	}
	return &StorageProviderRegistryStorageProviderAllocationUsedIterator{contract: _StorageProviderRegistry.contract, event: "StorageProviderAllocationUsed", logs: logs, sub: sub}, nil
}

// WatchStorageProviderAllocationUsed is a free log subscription operation binding the contract event 0x5fc112892cdeb5d0ea544d634fb55a2eb4ba066d9ad730b67f34791462ebf51b.
//
// Solidity: event StorageProviderAllocationUsed(uint64 ownerId, uint256 usedAllocation)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) WatchStorageProviderAllocationUsed(opts *bind.WatchOpts, sink chan<- *StorageProviderRegistryStorageProviderAllocationUsed) (event.Subscription, error) {

	logs, sub, err := _StorageProviderRegistry.contract.WatchLogs(opts, "StorageProviderAllocationUsed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProviderRegistryStorageProviderAllocationUsed)
				if err := _StorageProviderRegistry.contract.UnpackLog(event, "StorageProviderAllocationUsed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStorageProviderAllocationUsed is a log parse operation binding the contract event 0x5fc112892cdeb5d0ea544d634fb55a2eb4ba066d9ad730b67f34791462ebf51b.
//
// Solidity: event StorageProviderAllocationUsed(uint64 ownerId, uint256 usedAllocation)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) ParseStorageProviderAllocationUsed(log types.Log) (*StorageProviderRegistryStorageProviderAllocationUsed, error) {
	event := new(StorageProviderRegistryStorageProviderAllocationUsed)
	if err := _StorageProviderRegistry.contract.UnpackLog(event, "StorageProviderAllocationUsed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageProviderRegistryStorageProviderBeneficiaryAddressAcceptedIterator is returned from FilterStorageProviderBeneficiaryAddressAccepted and is used to iterate over the raw logs and unpacked data for StorageProviderBeneficiaryAddressAccepted events raised by the StorageProviderRegistry contract.
type StorageProviderRegistryStorageProviderBeneficiaryAddressAcceptedIterator struct {
	Event *StorageProviderRegistryStorageProviderBeneficiaryAddressAccepted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageProviderRegistryStorageProviderBeneficiaryAddressAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProviderRegistryStorageProviderBeneficiaryAddressAccepted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageProviderRegistryStorageProviderBeneficiaryAddressAccepted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageProviderRegistryStorageProviderBeneficiaryAddressAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProviderRegistryStorageProviderBeneficiaryAddressAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProviderRegistryStorageProviderBeneficiaryAddressAccepted represents a StorageProviderBeneficiaryAddressAccepted event raised by the StorageProviderRegistry contract.
type StorageProviderRegistryStorageProviderBeneficiaryAddressAccepted struct {
	OwnerId uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStorageProviderBeneficiaryAddressAccepted is a free log retrieval operation binding the contract event 0x5087366327b1aef4b5e946461225164c036e5dda2b25c458ea1bb444586b6a50.
//
// Solidity: event StorageProviderBeneficiaryAddressAccepted(uint64 ownerId)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) FilterStorageProviderBeneficiaryAddressAccepted(opts *bind.FilterOpts) (*StorageProviderRegistryStorageProviderBeneficiaryAddressAcceptedIterator, error) {

	logs, sub, err := _StorageProviderRegistry.contract.FilterLogs(opts, "StorageProviderBeneficiaryAddressAccepted")
	if err != nil {
		return nil, err
	}
	return &StorageProviderRegistryStorageProviderBeneficiaryAddressAcceptedIterator{contract: _StorageProviderRegistry.contract, event: "StorageProviderBeneficiaryAddressAccepted", logs: logs, sub: sub}, nil
}

// WatchStorageProviderBeneficiaryAddressAccepted is a free log subscription operation binding the contract event 0x5087366327b1aef4b5e946461225164c036e5dda2b25c458ea1bb444586b6a50.
//
// Solidity: event StorageProviderBeneficiaryAddressAccepted(uint64 ownerId)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) WatchStorageProviderBeneficiaryAddressAccepted(opts *bind.WatchOpts, sink chan<- *StorageProviderRegistryStorageProviderBeneficiaryAddressAccepted) (event.Subscription, error) {

	logs, sub, err := _StorageProviderRegistry.contract.WatchLogs(opts, "StorageProviderBeneficiaryAddressAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProviderRegistryStorageProviderBeneficiaryAddressAccepted)
				if err := _StorageProviderRegistry.contract.UnpackLog(event, "StorageProviderBeneficiaryAddressAccepted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStorageProviderBeneficiaryAddressAccepted is a log parse operation binding the contract event 0x5087366327b1aef4b5e946461225164c036e5dda2b25c458ea1bb444586b6a50.
//
// Solidity: event StorageProviderBeneficiaryAddressAccepted(uint64 ownerId)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) ParseStorageProviderBeneficiaryAddressAccepted(log types.Log) (*StorageProviderRegistryStorageProviderBeneficiaryAddressAccepted, error) {
	event := new(StorageProviderRegistryStorageProviderBeneficiaryAddressAccepted)
	if err := _StorageProviderRegistry.contract.UnpackLog(event, "StorageProviderBeneficiaryAddressAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageProviderRegistryStorageProviderBeneficiaryAddressUpdatedIterator is returned from FilterStorageProviderBeneficiaryAddressUpdated and is used to iterate over the raw logs and unpacked data for StorageProviderBeneficiaryAddressUpdated events raised by the StorageProviderRegistry contract.
type StorageProviderRegistryStorageProviderBeneficiaryAddressUpdatedIterator struct {
	Event *StorageProviderRegistryStorageProviderBeneficiaryAddressUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageProviderRegistryStorageProviderBeneficiaryAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProviderRegistryStorageProviderBeneficiaryAddressUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageProviderRegistryStorageProviderBeneficiaryAddressUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageProviderRegistryStorageProviderBeneficiaryAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProviderRegistryStorageProviderBeneficiaryAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProviderRegistryStorageProviderBeneficiaryAddressUpdated represents a StorageProviderBeneficiaryAddressUpdated event raised by the StorageProviderRegistry contract.
type StorageProviderRegistryStorageProviderBeneficiaryAddressUpdated struct {
	BeneficiaryAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterStorageProviderBeneficiaryAddressUpdated is a free log retrieval operation binding the contract event 0x906cdab773befe815c0a81a3837dc4df80e433c0c26e4ee402007515bb777c1c.
//
// Solidity: event StorageProviderBeneficiaryAddressUpdated(address beneficiaryAddress)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) FilterStorageProviderBeneficiaryAddressUpdated(opts *bind.FilterOpts) (*StorageProviderRegistryStorageProviderBeneficiaryAddressUpdatedIterator, error) {

	logs, sub, err := _StorageProviderRegistry.contract.FilterLogs(opts, "StorageProviderBeneficiaryAddressUpdated")
	if err != nil {
		return nil, err
	}
	return &StorageProviderRegistryStorageProviderBeneficiaryAddressUpdatedIterator{contract: _StorageProviderRegistry.contract, event: "StorageProviderBeneficiaryAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchStorageProviderBeneficiaryAddressUpdated is a free log subscription operation binding the contract event 0x906cdab773befe815c0a81a3837dc4df80e433c0c26e4ee402007515bb777c1c.
//
// Solidity: event StorageProviderBeneficiaryAddressUpdated(address beneficiaryAddress)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) WatchStorageProviderBeneficiaryAddressUpdated(opts *bind.WatchOpts, sink chan<- *StorageProviderRegistryStorageProviderBeneficiaryAddressUpdated) (event.Subscription, error) {

	logs, sub, err := _StorageProviderRegistry.contract.WatchLogs(opts, "StorageProviderBeneficiaryAddressUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProviderRegistryStorageProviderBeneficiaryAddressUpdated)
				if err := _StorageProviderRegistry.contract.UnpackLog(event, "StorageProviderBeneficiaryAddressUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStorageProviderBeneficiaryAddressUpdated is a log parse operation binding the contract event 0x906cdab773befe815c0a81a3837dc4df80e433c0c26e4ee402007515bb777c1c.
//
// Solidity: event StorageProviderBeneficiaryAddressUpdated(address beneficiaryAddress)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) ParseStorageProviderBeneficiaryAddressUpdated(log types.Log) (*StorageProviderRegistryStorageProviderBeneficiaryAddressUpdated, error) {
	event := new(StorageProviderRegistryStorageProviderBeneficiaryAddressUpdated)
	if err := _StorageProviderRegistry.contract.UnpackLog(event, "StorageProviderBeneficiaryAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageProviderRegistryStorageProviderDeactivatedIterator is returned from FilterStorageProviderDeactivated and is used to iterate over the raw logs and unpacked data for StorageProviderDeactivated events raised by the StorageProviderRegistry contract.
type StorageProviderRegistryStorageProviderDeactivatedIterator struct {
	Event *StorageProviderRegistryStorageProviderDeactivated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageProviderRegistryStorageProviderDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProviderRegistryStorageProviderDeactivated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageProviderRegistryStorageProviderDeactivated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageProviderRegistryStorageProviderDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProviderRegistryStorageProviderDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProviderRegistryStorageProviderDeactivated represents a StorageProviderDeactivated event raised by the StorageProviderRegistry contract.
type StorageProviderRegistryStorageProviderDeactivated struct {
	OwnerId uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStorageProviderDeactivated is a free log retrieval operation binding the contract event 0xec36a8776cbb637a0a427dc03df67a0283f596365d3d127696243e82d612c88f.
//
// Solidity: event StorageProviderDeactivated(uint64 ownerId)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) FilterStorageProviderDeactivated(opts *bind.FilterOpts) (*StorageProviderRegistryStorageProviderDeactivatedIterator, error) {

	logs, sub, err := _StorageProviderRegistry.contract.FilterLogs(opts, "StorageProviderDeactivated")
	if err != nil {
		return nil, err
	}
	return &StorageProviderRegistryStorageProviderDeactivatedIterator{contract: _StorageProviderRegistry.contract, event: "StorageProviderDeactivated", logs: logs, sub: sub}, nil
}

// WatchStorageProviderDeactivated is a free log subscription operation binding the contract event 0xec36a8776cbb637a0a427dc03df67a0283f596365d3d127696243e82d612c88f.
//
// Solidity: event StorageProviderDeactivated(uint64 ownerId)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) WatchStorageProviderDeactivated(opts *bind.WatchOpts, sink chan<- *StorageProviderRegistryStorageProviderDeactivated) (event.Subscription, error) {

	logs, sub, err := _StorageProviderRegistry.contract.WatchLogs(opts, "StorageProviderDeactivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProviderRegistryStorageProviderDeactivated)
				if err := _StorageProviderRegistry.contract.UnpackLog(event, "StorageProviderDeactivated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStorageProviderDeactivated is a log parse operation binding the contract event 0xec36a8776cbb637a0a427dc03df67a0283f596365d3d127696243e82d612c88f.
//
// Solidity: event StorageProviderDeactivated(uint64 ownerId)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) ParseStorageProviderDeactivated(log types.Log) (*StorageProviderRegistryStorageProviderDeactivated, error) {
	event := new(StorageProviderRegistryStorageProviderDeactivated)
	if err := _StorageProviderRegistry.contract.UnpackLog(event, "StorageProviderDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageProviderRegistryStorageProviderLastEpochUpdateIterator is returned from FilterStorageProviderLastEpochUpdate and is used to iterate over the raw logs and unpacked data for StorageProviderLastEpochUpdate events raised by the StorageProviderRegistry contract.
type StorageProviderRegistryStorageProviderLastEpochUpdateIterator struct {
	Event *StorageProviderRegistryStorageProviderLastEpochUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageProviderRegistryStorageProviderLastEpochUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProviderRegistryStorageProviderLastEpochUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageProviderRegistryStorageProviderLastEpochUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageProviderRegistryStorageProviderLastEpochUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProviderRegistryStorageProviderLastEpochUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProviderRegistryStorageProviderLastEpochUpdate represents a StorageProviderLastEpochUpdate event raised by the StorageProviderRegistry contract.
type StorageProviderRegistryStorageProviderLastEpochUpdate struct {
	OwnerId   uint64
	LastEpoch int64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStorageProviderLastEpochUpdate is a free log retrieval operation binding the contract event 0xef9bdb18018734892a7eab99e55de15e0e8b9af41fc9e8f31ca9bfed22ceee3e.
//
// Solidity: event StorageProviderLastEpochUpdate(uint64 ownerId, int64 lastEpoch)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) FilterStorageProviderLastEpochUpdate(opts *bind.FilterOpts) (*StorageProviderRegistryStorageProviderLastEpochUpdateIterator, error) {

	logs, sub, err := _StorageProviderRegistry.contract.FilterLogs(opts, "StorageProviderLastEpochUpdate")
	if err != nil {
		return nil, err
	}
	return &StorageProviderRegistryStorageProviderLastEpochUpdateIterator{contract: _StorageProviderRegistry.contract, event: "StorageProviderLastEpochUpdate", logs: logs, sub: sub}, nil
}

// WatchStorageProviderLastEpochUpdate is a free log subscription operation binding the contract event 0xef9bdb18018734892a7eab99e55de15e0e8b9af41fc9e8f31ca9bfed22ceee3e.
//
// Solidity: event StorageProviderLastEpochUpdate(uint64 ownerId, int64 lastEpoch)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) WatchStorageProviderLastEpochUpdate(opts *bind.WatchOpts, sink chan<- *StorageProviderRegistryStorageProviderLastEpochUpdate) (event.Subscription, error) {

	logs, sub, err := _StorageProviderRegistry.contract.WatchLogs(opts, "StorageProviderLastEpochUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProviderRegistryStorageProviderLastEpochUpdate)
				if err := _StorageProviderRegistry.contract.UnpackLog(event, "StorageProviderLastEpochUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStorageProviderLastEpochUpdate is a log parse operation binding the contract event 0xef9bdb18018734892a7eab99e55de15e0e8b9af41fc9e8f31ca9bfed22ceee3e.
//
// Solidity: event StorageProviderLastEpochUpdate(uint64 ownerId, int64 lastEpoch)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) ParseStorageProviderLastEpochUpdate(log types.Log) (*StorageProviderRegistryStorageProviderLastEpochUpdate, error) {
	event := new(StorageProviderRegistryStorageProviderLastEpochUpdate)
	if err := _StorageProviderRegistry.contract.UnpackLog(event, "StorageProviderLastEpochUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageProviderRegistryStorageProviderMinerAddressUpdateIterator is returned from FilterStorageProviderMinerAddressUpdate and is used to iterate over the raw logs and unpacked data for StorageProviderMinerAddressUpdate events raised by the StorageProviderRegistry contract.
type StorageProviderRegistryStorageProviderMinerAddressUpdateIterator struct {
	Event *StorageProviderRegistryStorageProviderMinerAddressUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageProviderRegistryStorageProviderMinerAddressUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProviderRegistryStorageProviderMinerAddressUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageProviderRegistryStorageProviderMinerAddressUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageProviderRegistryStorageProviderMinerAddressUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProviderRegistryStorageProviderMinerAddressUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProviderRegistryStorageProviderMinerAddressUpdate represents a StorageProviderMinerAddressUpdate event raised by the StorageProviderRegistry contract.
type StorageProviderRegistryStorageProviderMinerAddressUpdate struct {
	OwnerId uint64
	Miner   uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStorageProviderMinerAddressUpdate is a free log retrieval operation binding the contract event 0x0ebad5580422e1bba5ee173f19698e506fd96a2d04253a0901c1db86f654c849.
//
// Solidity: event StorageProviderMinerAddressUpdate(uint64 ownerId, uint64 miner)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) FilterStorageProviderMinerAddressUpdate(opts *bind.FilterOpts) (*StorageProviderRegistryStorageProviderMinerAddressUpdateIterator, error) {

	logs, sub, err := _StorageProviderRegistry.contract.FilterLogs(opts, "StorageProviderMinerAddressUpdate")
	if err != nil {
		return nil, err
	}
	return &StorageProviderRegistryStorageProviderMinerAddressUpdateIterator{contract: _StorageProviderRegistry.contract, event: "StorageProviderMinerAddressUpdate", logs: logs, sub: sub}, nil
}

// WatchStorageProviderMinerAddressUpdate is a free log subscription operation binding the contract event 0x0ebad5580422e1bba5ee173f19698e506fd96a2d04253a0901c1db86f654c849.
//
// Solidity: event StorageProviderMinerAddressUpdate(uint64 ownerId, uint64 miner)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) WatchStorageProviderMinerAddressUpdate(opts *bind.WatchOpts, sink chan<- *StorageProviderRegistryStorageProviderMinerAddressUpdate) (event.Subscription, error) {

	logs, sub, err := _StorageProviderRegistry.contract.WatchLogs(opts, "StorageProviderMinerAddressUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProviderRegistryStorageProviderMinerAddressUpdate)
				if err := _StorageProviderRegistry.contract.UnpackLog(event, "StorageProviderMinerAddressUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStorageProviderMinerAddressUpdate is a log parse operation binding the contract event 0x0ebad5580422e1bba5ee173f19698e506fd96a2d04253a0901c1db86f654c849.
//
// Solidity: event StorageProviderMinerAddressUpdate(uint64 ownerId, uint64 miner)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) ParseStorageProviderMinerAddressUpdate(log types.Log) (*StorageProviderRegistryStorageProviderMinerAddressUpdate, error) {
	event := new(StorageProviderRegistryStorageProviderMinerAddressUpdate)
	if err := _StorageProviderRegistry.contract.UnpackLog(event, "StorageProviderMinerAddressUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageProviderRegistryStorageProviderMinerRestakingRatioUpdateIterator is returned from FilterStorageProviderMinerRestakingRatioUpdate and is used to iterate over the raw logs and unpacked data for StorageProviderMinerRestakingRatioUpdate events raised by the StorageProviderRegistry contract.
type StorageProviderRegistryStorageProviderMinerRestakingRatioUpdateIterator struct {
	Event *StorageProviderRegistryStorageProviderMinerRestakingRatioUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageProviderRegistryStorageProviderMinerRestakingRatioUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProviderRegistryStorageProviderMinerRestakingRatioUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageProviderRegistryStorageProviderMinerRestakingRatioUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageProviderRegistryStorageProviderMinerRestakingRatioUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProviderRegistryStorageProviderMinerRestakingRatioUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProviderRegistryStorageProviderMinerRestakingRatioUpdate represents a StorageProviderMinerRestakingRatioUpdate event raised by the StorageProviderRegistry contract.
type StorageProviderRegistryStorageProviderMinerRestakingRatioUpdate struct {
	OwnerId          uint64
	RestakingRatio   *big.Int
	RestakingAddress common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterStorageProviderMinerRestakingRatioUpdate is a free log retrieval operation binding the contract event 0xfec0818a0fae8fa0693f4ab6209e6ace854f156067b4069ebede3c8c02622511.
//
// Solidity: event StorageProviderMinerRestakingRatioUpdate(uint64 ownerId, uint256 restakingRatio, address restakingAddress)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) FilterStorageProviderMinerRestakingRatioUpdate(opts *bind.FilterOpts) (*StorageProviderRegistryStorageProviderMinerRestakingRatioUpdateIterator, error) {

	logs, sub, err := _StorageProviderRegistry.contract.FilterLogs(opts, "StorageProviderMinerRestakingRatioUpdate")
	if err != nil {
		return nil, err
	}
	return &StorageProviderRegistryStorageProviderMinerRestakingRatioUpdateIterator{contract: _StorageProviderRegistry.contract, event: "StorageProviderMinerRestakingRatioUpdate", logs: logs, sub: sub}, nil
}

// WatchStorageProviderMinerRestakingRatioUpdate is a free log subscription operation binding the contract event 0xfec0818a0fae8fa0693f4ab6209e6ace854f156067b4069ebede3c8c02622511.
//
// Solidity: event StorageProviderMinerRestakingRatioUpdate(uint64 ownerId, uint256 restakingRatio, address restakingAddress)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) WatchStorageProviderMinerRestakingRatioUpdate(opts *bind.WatchOpts, sink chan<- *StorageProviderRegistryStorageProviderMinerRestakingRatioUpdate) (event.Subscription, error) {

	logs, sub, err := _StorageProviderRegistry.contract.WatchLogs(opts, "StorageProviderMinerRestakingRatioUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProviderRegistryStorageProviderMinerRestakingRatioUpdate)
				if err := _StorageProviderRegistry.contract.UnpackLog(event, "StorageProviderMinerRestakingRatioUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStorageProviderMinerRestakingRatioUpdate is a log parse operation binding the contract event 0xfec0818a0fae8fa0693f4ab6209e6ace854f156067b4069ebede3c8c02622511.
//
// Solidity: event StorageProviderMinerRestakingRatioUpdate(uint64 ownerId, uint256 restakingRatio, address restakingAddress)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) ParseStorageProviderMinerRestakingRatioUpdate(log types.Log) (*StorageProviderRegistryStorageProviderMinerRestakingRatioUpdate, error) {
	event := new(StorageProviderRegistryStorageProviderMinerRestakingRatioUpdate)
	if err := _StorageProviderRegistry.contract.UnpackLog(event, "StorageProviderMinerRestakingRatioUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageProviderRegistryStorageProviderOnboardedIterator is returned from FilterStorageProviderOnboarded and is used to iterate over the raw logs and unpacked data for StorageProviderOnboarded events raised by the StorageProviderRegistry contract.
type StorageProviderRegistryStorageProviderOnboardedIterator struct {
	Event *StorageProviderRegistryStorageProviderOnboarded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageProviderRegistryStorageProviderOnboardedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProviderRegistryStorageProviderOnboarded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageProviderRegistryStorageProviderOnboarded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageProviderRegistryStorageProviderOnboardedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProviderRegistryStorageProviderOnboardedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProviderRegistryStorageProviderOnboarded represents a StorageProviderOnboarded event raised by the StorageProviderRegistry contract.
type StorageProviderRegistryStorageProviderOnboarded struct {
	OwnerId         uint64
	MinerId         uint64
	AllocationLimit *big.Int
	DailyAllocation *big.Int
	Repayment       *big.Int
	LastEpoch       int64
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStorageProviderOnboarded is a free log retrieval operation binding the contract event 0xa023ff258c2cf4bbb89ff7b29a652f01c9132b5ad8bca0f51767a4e972ec22a7.
//
// Solidity: event StorageProviderOnboarded(uint64 ownerId, uint64 minerId, uint256 allocationLimit, uint256 dailyAllocation, uint256 repayment, int64 lastEpoch)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) FilterStorageProviderOnboarded(opts *bind.FilterOpts) (*StorageProviderRegistryStorageProviderOnboardedIterator, error) {

	logs, sub, err := _StorageProviderRegistry.contract.FilterLogs(opts, "StorageProviderOnboarded")
	if err != nil {
		return nil, err
	}
	return &StorageProviderRegistryStorageProviderOnboardedIterator{contract: _StorageProviderRegistry.contract, event: "StorageProviderOnboarded", logs: logs, sub: sub}, nil
}

// WatchStorageProviderOnboarded is a free log subscription operation binding the contract event 0xa023ff258c2cf4bbb89ff7b29a652f01c9132b5ad8bca0f51767a4e972ec22a7.
//
// Solidity: event StorageProviderOnboarded(uint64 ownerId, uint64 minerId, uint256 allocationLimit, uint256 dailyAllocation, uint256 repayment, int64 lastEpoch)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) WatchStorageProviderOnboarded(opts *bind.WatchOpts, sink chan<- *StorageProviderRegistryStorageProviderOnboarded) (event.Subscription, error) {

	logs, sub, err := _StorageProviderRegistry.contract.WatchLogs(opts, "StorageProviderOnboarded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProviderRegistryStorageProviderOnboarded)
				if err := _StorageProviderRegistry.contract.UnpackLog(event, "StorageProviderOnboarded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStorageProviderOnboarded is a log parse operation binding the contract event 0xa023ff258c2cf4bbb89ff7b29a652f01c9132b5ad8bca0f51767a4e972ec22a7.
//
// Solidity: event StorageProviderOnboarded(uint64 ownerId, uint64 minerId, uint256 allocationLimit, uint256 dailyAllocation, uint256 repayment, int64 lastEpoch)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) ParseStorageProviderOnboarded(log types.Log) (*StorageProviderRegistryStorageProviderOnboarded, error) {
	event := new(StorageProviderRegistryStorageProviderOnboarded)
	if err := _StorageProviderRegistry.contract.UnpackLog(event, "StorageProviderOnboarded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageProviderRegistryStorageProviderRegisteredIterator is returned from FilterStorageProviderRegistered and is used to iterate over the raw logs and unpacked data for StorageProviderRegistered events raised by the StorageProviderRegistry contract.
type StorageProviderRegistryStorageProviderRegisteredIterator struct {
	Event *StorageProviderRegistryStorageProviderRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageProviderRegistryStorageProviderRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProviderRegistryStorageProviderRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageProviderRegistryStorageProviderRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageProviderRegistryStorageProviderRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProviderRegistryStorageProviderRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProviderRegistryStorageProviderRegistered represents a StorageProviderRegistered event raised by the StorageProviderRegistry contract.
type StorageProviderRegistryStorageProviderRegistered struct {
	Owner           []byte
	OwnerId         uint64
	MinerId         uint64
	TargetPool      common.Address
	AllocationLimit *big.Int
	DailyAllocation *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStorageProviderRegistered is a free log retrieval operation binding the contract event 0x9346111d414b53e51a9c5d64a17479d23c5c981f0ae8f5ef3d211a91011ef4ae.
//
// Solidity: event StorageProviderRegistered(bytes owner, uint64 ownerId, uint64 minerId, address targetPool, uint256 allocationLimit, uint256 dailyAllocation)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) FilterStorageProviderRegistered(opts *bind.FilterOpts) (*StorageProviderRegistryStorageProviderRegisteredIterator, error) {

	logs, sub, err := _StorageProviderRegistry.contract.FilterLogs(opts, "StorageProviderRegistered")
	if err != nil {
		return nil, err
	}
	return &StorageProviderRegistryStorageProviderRegisteredIterator{contract: _StorageProviderRegistry.contract, event: "StorageProviderRegistered", logs: logs, sub: sub}, nil
}

// WatchStorageProviderRegistered is a free log subscription operation binding the contract event 0x9346111d414b53e51a9c5d64a17479d23c5c981f0ae8f5ef3d211a91011ef4ae.
//
// Solidity: event StorageProviderRegistered(bytes owner, uint64 ownerId, uint64 minerId, address targetPool, uint256 allocationLimit, uint256 dailyAllocation)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) WatchStorageProviderRegistered(opts *bind.WatchOpts, sink chan<- *StorageProviderRegistryStorageProviderRegistered) (event.Subscription, error) {

	logs, sub, err := _StorageProviderRegistry.contract.WatchLogs(opts, "StorageProviderRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProviderRegistryStorageProviderRegistered)
				if err := _StorageProviderRegistry.contract.UnpackLog(event, "StorageProviderRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStorageProviderRegistered is a log parse operation binding the contract event 0x9346111d414b53e51a9c5d64a17479d23c5c981f0ae8f5ef3d211a91011ef4ae.
//
// Solidity: event StorageProviderRegistered(bytes owner, uint64 ownerId, uint64 minerId, address targetPool, uint256 allocationLimit, uint256 dailyAllocation)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) ParseStorageProviderRegistered(log types.Log) (*StorageProviderRegistryStorageProviderRegistered, error) {
	event := new(StorageProviderRegistryStorageProviderRegistered)
	if err := _StorageProviderRegistry.contract.UnpackLog(event, "StorageProviderRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageProviderRegistryStorageProviderRepaidPledgeIterator is returned from FilterStorageProviderRepaidPledge and is used to iterate over the raw logs and unpacked data for StorageProviderRepaidPledge events raised by the StorageProviderRegistry contract.
type StorageProviderRegistryStorageProviderRepaidPledgeIterator struct {
	Event *StorageProviderRegistryStorageProviderRepaidPledge // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageProviderRegistryStorageProviderRepaidPledgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProviderRegistryStorageProviderRepaidPledge)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageProviderRegistryStorageProviderRepaidPledge)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageProviderRegistryStorageProviderRepaidPledgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProviderRegistryStorageProviderRepaidPledgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProviderRegistryStorageProviderRepaidPledge represents a StorageProviderRepaidPledge event raised by the StorageProviderRegistry contract.
type StorageProviderRegistryStorageProviderRepaidPledge struct {
	OwnerId uint64
	Pledge  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStorageProviderRepaidPledge is a free log retrieval operation binding the contract event 0x1124506509059ae4a33548514242566cb78aeb3d1a720b18dbec366182bae588.
//
// Solidity: event StorageProviderRepaidPledge(uint64 ownerId, uint256 pledge)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) FilterStorageProviderRepaidPledge(opts *bind.FilterOpts) (*StorageProviderRegistryStorageProviderRepaidPledgeIterator, error) {

	logs, sub, err := _StorageProviderRegistry.contract.FilterLogs(opts, "StorageProviderRepaidPledge")
	if err != nil {
		return nil, err
	}
	return &StorageProviderRegistryStorageProviderRepaidPledgeIterator{contract: _StorageProviderRegistry.contract, event: "StorageProviderRepaidPledge", logs: logs, sub: sub}, nil
}

// WatchStorageProviderRepaidPledge is a free log subscription operation binding the contract event 0x1124506509059ae4a33548514242566cb78aeb3d1a720b18dbec366182bae588.
//
// Solidity: event StorageProviderRepaidPledge(uint64 ownerId, uint256 pledge)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) WatchStorageProviderRepaidPledge(opts *bind.WatchOpts, sink chan<- *StorageProviderRegistryStorageProviderRepaidPledge) (event.Subscription, error) {

	logs, sub, err := _StorageProviderRegistry.contract.WatchLogs(opts, "StorageProviderRepaidPledge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProviderRegistryStorageProviderRepaidPledge)
				if err := _StorageProviderRegistry.contract.UnpackLog(event, "StorageProviderRepaidPledge", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStorageProviderRepaidPledge is a log parse operation binding the contract event 0x1124506509059ae4a33548514242566cb78aeb3d1a720b18dbec366182bae588.
//
// Solidity: event StorageProviderRepaidPledge(uint64 ownerId, uint256 pledge)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) ParseStorageProviderRepaidPledge(log types.Log) (*StorageProviderRegistryStorageProviderRepaidPledge, error) {
	event := new(StorageProviderRegistryStorageProviderRepaidPledge)
	if err := _StorageProviderRegistry.contract.UnpackLog(event, "StorageProviderRepaidPledge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageProviderRegistryUpdateMaxAllocationIterator is returned from FilterUpdateMaxAllocation and is used to iterate over the raw logs and unpacked data for UpdateMaxAllocation events raised by the StorageProviderRegistry contract.
type StorageProviderRegistryUpdateMaxAllocationIterator struct {
	Event *StorageProviderRegistryUpdateMaxAllocation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageProviderRegistryUpdateMaxAllocationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProviderRegistryUpdateMaxAllocation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageProviderRegistryUpdateMaxAllocation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageProviderRegistryUpdateMaxAllocationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProviderRegistryUpdateMaxAllocationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProviderRegistryUpdateMaxAllocation represents a UpdateMaxAllocation event raised by the StorageProviderRegistry contract.
type StorageProviderRegistryUpdateMaxAllocation struct {
	MaxAllocation *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdateMaxAllocation is a free log retrieval operation binding the contract event 0x39db4c142f7363a2d6a0c4a06b5f53631194e27e1c9e1ee7996b7e0acc561ce3.
//
// Solidity: event UpdateMaxAllocation(uint256 maxAllocation)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) FilterUpdateMaxAllocation(opts *bind.FilterOpts) (*StorageProviderRegistryUpdateMaxAllocationIterator, error) {

	logs, sub, err := _StorageProviderRegistry.contract.FilterLogs(opts, "UpdateMaxAllocation")
	if err != nil {
		return nil, err
	}
	return &StorageProviderRegistryUpdateMaxAllocationIterator{contract: _StorageProviderRegistry.contract, event: "UpdateMaxAllocation", logs: logs, sub: sub}, nil
}

// WatchUpdateMaxAllocation is a free log subscription operation binding the contract event 0x39db4c142f7363a2d6a0c4a06b5f53631194e27e1c9e1ee7996b7e0acc561ce3.
//
// Solidity: event UpdateMaxAllocation(uint256 maxAllocation)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) WatchUpdateMaxAllocation(opts *bind.WatchOpts, sink chan<- *StorageProviderRegistryUpdateMaxAllocation) (event.Subscription, error) {

	logs, sub, err := _StorageProviderRegistry.contract.WatchLogs(opts, "UpdateMaxAllocation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProviderRegistryUpdateMaxAllocation)
				if err := _StorageProviderRegistry.contract.UnpackLog(event, "UpdateMaxAllocation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateMaxAllocation is a log parse operation binding the contract event 0x39db4c142f7363a2d6a0c4a06b5f53631194e27e1c9e1ee7996b7e0acc561ce3.
//
// Solidity: event UpdateMaxAllocation(uint256 maxAllocation)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) ParseUpdateMaxAllocation(log types.Log) (*StorageProviderRegistryUpdateMaxAllocation, error) {
	event := new(StorageProviderRegistryUpdateMaxAllocation)
	if err := _StorageProviderRegistry.contract.UnpackLog(event, "UpdateMaxAllocation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageProviderRegistryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the StorageProviderRegistry contract.
type StorageProviderRegistryUpgradedIterator struct {
	Event *StorageProviderRegistryUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageProviderRegistryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProviderRegistryUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageProviderRegistryUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageProviderRegistryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProviderRegistryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProviderRegistryUpgraded represents a Upgraded event raised by the StorageProviderRegistry contract.
type StorageProviderRegistryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*StorageProviderRegistryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _StorageProviderRegistry.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &StorageProviderRegistryUpgradedIterator{contract: _StorageProviderRegistry.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *StorageProviderRegistryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _StorageProviderRegistry.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProviderRegistryUpgraded)
				if err := _StorageProviderRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) ParseUpgraded(log types.Log) (*StorageProviderRegistryUpgraded, error) {
	event := new(StorageProviderRegistryUpgraded)
	if err := _StorageProviderRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
