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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxStorageProviders\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxAllocation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minTimePeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxTimePeriod\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"CollateralAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"LiquidStakingPoolRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"provider\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"name\":\"StorageProviderAccruedRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"provider\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allocationLimit\",\"type\":\"uint256\"}],\"name\":\"StorageProviderAllocationLimitUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"provider\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usedAllocation\",\"type\":\"uint256\"}],\"name\":\"StorageProviderAllocationUsed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_provider\",\"type\":\"bytes\"}],\"name\":\"StorageProviderBeneficiaryAddressAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_beneficiaryAddress\",\"type\":\"address\"}],\"name\":\"StorageProviderBeneficiaryAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"provider\",\"type\":\"bytes\"}],\"name\":\"StorageProviderDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"provider\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"name\":\"StorageProviderLockedRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"provider\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"StorageProviderMaxRedeemablePeriodUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"provider\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"miner\",\"type\":\"bytes\"}],\"name\":\"StorageProviderMinerAddressUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"provider\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"miner\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"targetPool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allocationLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxPeriod\",\"type\":\"uint256\"}],\"name\":\"StorageProviderRegistered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_provider\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_beneficiaryAddress\",\"type\":\"address\"}],\"name\":\"acceptBeneficiaryAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_beneficiaryAddress\",\"type\":\"address\"}],\"name\":\"changeBeneficiaryAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateral\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_provider\",\"type\":\"bytes\"}],\"name\":\"deactivateStorageProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_provider\",\"type\":\"bytes\"}],\"name\":\"getStorageProvider\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalActiveStorageProviders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalStorageProviders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_provider\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_accuredRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lockedRewards\",\"type\":\"uint256\"}],\"name\":\"increaseRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_provider\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_allocated\",\"type\":\"uint256\"}],\"name\":\"increaseUsedAllocation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_provider\",\"type\":\"bytes\"}],\"name\":\"isActiveProvider\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxAllocation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxStorageProviders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTimePeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minTimePeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pools\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_miner\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_targetPool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_allocationLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"registerPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_provider\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_allocationLimit\",\"type\":\"uint256\"}],\"name\":\"setAllocationLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_collateral\",\"type\":\"address\"}],\"name\":\"setCollateralAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_provider\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"}],\"name\":\"setMaxRedeemablePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_provider\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_miner\",\"type\":\"bytes\"}],\"name\":\"setMinerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"storageProviders\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"targetPool\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"miner\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"allocationLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usedAllocation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accruedRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockedRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRedeemablePeriod\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalInactiveStorageProviders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStorageProviders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(address)
func (_StorageProviderRegistry *StorageProviderRegistryCaller) Collateral(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StorageProviderRegistry.contract.Call(opts, &out, "collateral")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(address)
func (_StorageProviderRegistry *StorageProviderRegistrySession) Collateral() (common.Address, error) {
	return _StorageProviderRegistry.Contract.Collateral(&_StorageProviderRegistry.CallOpts)
}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(address)
func (_StorageProviderRegistry *StorageProviderRegistryCallerSession) Collateral() (common.Address, error) {
	return _StorageProviderRegistry.Contract.Collateral(&_StorageProviderRegistry.CallOpts)
}

// GetStorageProvider is a free data retrieval call binding the contract method 0x50390774.
//
// Solidity: function getStorageProvider(bytes _provider) view returns(bool, address, bytes, uint256, uint256, uint256, uint256, uint256)
func (_StorageProviderRegistry *StorageProviderRegistryCaller) GetStorageProvider(opts *bind.CallOpts, _provider []byte) (bool, common.Address, []byte, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _StorageProviderRegistry.contract.Call(opts, &out, "getStorageProvider", _provider)

	if err != nil {
		return *new(bool), *new(common.Address), *new([]byte), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	out6 := *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	out7 := *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, out6, out7, err

}

// GetStorageProvider is a free data retrieval call binding the contract method 0x50390774.
//
// Solidity: function getStorageProvider(bytes _provider) view returns(bool, address, bytes, uint256, uint256, uint256, uint256, uint256)
func (_StorageProviderRegistry *StorageProviderRegistrySession) GetStorageProvider(_provider []byte) (bool, common.Address, []byte, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _StorageProviderRegistry.Contract.GetStorageProvider(&_StorageProviderRegistry.CallOpts, _provider)
}

// GetStorageProvider is a free data retrieval call binding the contract method 0x50390774.
//
// Solidity: function getStorageProvider(bytes _provider) view returns(bool, address, bytes, uint256, uint256, uint256, uint256, uint256)
func (_StorageProviderRegistry *StorageProviderRegistryCallerSession) GetStorageProvider(_provider []byte) (bool, common.Address, []byte, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _StorageProviderRegistry.Contract.GetStorageProvider(&_StorageProviderRegistry.CallOpts, _provider)
}

// GetTotalActiveStorageProviders is a free data retrieval call binding the contract method 0x0ecce720.
//
// Solidity: function getTotalActiveStorageProviders() view returns(uint256)
func (_StorageProviderRegistry *StorageProviderRegistryCaller) GetTotalActiveStorageProviders(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StorageProviderRegistry.contract.Call(opts, &out, "getTotalActiveStorageProviders")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalActiveStorageProviders is a free data retrieval call binding the contract method 0x0ecce720.
//
// Solidity: function getTotalActiveStorageProviders() view returns(uint256)
func (_StorageProviderRegistry *StorageProviderRegistrySession) GetTotalActiveStorageProviders() (*big.Int, error) {
	return _StorageProviderRegistry.Contract.GetTotalActiveStorageProviders(&_StorageProviderRegistry.CallOpts)
}

// GetTotalActiveStorageProviders is a free data retrieval call binding the contract method 0x0ecce720.
//
// Solidity: function getTotalActiveStorageProviders() view returns(uint256)
func (_StorageProviderRegistry *StorageProviderRegistryCallerSession) GetTotalActiveStorageProviders() (*big.Int, error) {
	return _StorageProviderRegistry.Contract.GetTotalActiveStorageProviders(&_StorageProviderRegistry.CallOpts)
}

// GetTotalStorageProviders is a free data retrieval call binding the contract method 0x7b5ae0cd.
//
// Solidity: function getTotalStorageProviders() view returns(uint256)
func (_StorageProviderRegistry *StorageProviderRegistryCaller) GetTotalStorageProviders(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StorageProviderRegistry.contract.Call(opts, &out, "getTotalStorageProviders")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalStorageProviders is a free data retrieval call binding the contract method 0x7b5ae0cd.
//
// Solidity: function getTotalStorageProviders() view returns(uint256)
func (_StorageProviderRegistry *StorageProviderRegistrySession) GetTotalStorageProviders() (*big.Int, error) {
	return _StorageProviderRegistry.Contract.GetTotalStorageProviders(&_StorageProviderRegistry.CallOpts)
}

// GetTotalStorageProviders is a free data retrieval call binding the contract method 0x7b5ae0cd.
//
// Solidity: function getTotalStorageProviders() view returns(uint256)
func (_StorageProviderRegistry *StorageProviderRegistryCallerSession) GetTotalStorageProviders() (*big.Int, error) {
	return _StorageProviderRegistry.Contract.GetTotalStorageProviders(&_StorageProviderRegistry.CallOpts)
}

// IsActiveProvider is a free data retrieval call binding the contract method 0x2550c3b4.
//
// Solidity: function isActiveProvider(bytes _provider) view returns(bool status)
func (_StorageProviderRegistry *StorageProviderRegistryCaller) IsActiveProvider(opts *bind.CallOpts, _provider []byte) (bool, error) {
	var out []interface{}
	err := _StorageProviderRegistry.contract.Call(opts, &out, "isActiveProvider", _provider)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActiveProvider is a free data retrieval call binding the contract method 0x2550c3b4.
//
// Solidity: function isActiveProvider(bytes _provider) view returns(bool status)
func (_StorageProviderRegistry *StorageProviderRegistrySession) IsActiveProvider(_provider []byte) (bool, error) {
	return _StorageProviderRegistry.Contract.IsActiveProvider(&_StorageProviderRegistry.CallOpts, _provider)
}

// IsActiveProvider is a free data retrieval call binding the contract method 0x2550c3b4.
//
// Solidity: function isActiveProvider(bytes _provider) view returns(bool status)
func (_StorageProviderRegistry *StorageProviderRegistryCallerSession) IsActiveProvider(_provider []byte) (bool, error) {
	return _StorageProviderRegistry.Contract.IsActiveProvider(&_StorageProviderRegistry.CallOpts, _provider)
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

// MaxStorageProviders is a free data retrieval call binding the contract method 0x4978fcdc.
//
// Solidity: function maxStorageProviders() view returns(uint256)
func (_StorageProviderRegistry *StorageProviderRegistryCaller) MaxStorageProviders(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StorageProviderRegistry.contract.Call(opts, &out, "maxStorageProviders")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxStorageProviders is a free data retrieval call binding the contract method 0x4978fcdc.
//
// Solidity: function maxStorageProviders() view returns(uint256)
func (_StorageProviderRegistry *StorageProviderRegistrySession) MaxStorageProviders() (*big.Int, error) {
	return _StorageProviderRegistry.Contract.MaxStorageProviders(&_StorageProviderRegistry.CallOpts)
}

// MaxStorageProviders is a free data retrieval call binding the contract method 0x4978fcdc.
//
// Solidity: function maxStorageProviders() view returns(uint256)
func (_StorageProviderRegistry *StorageProviderRegistryCallerSession) MaxStorageProviders() (*big.Int, error) {
	return _StorageProviderRegistry.Contract.MaxStorageProviders(&_StorageProviderRegistry.CallOpts)
}

// MaxTimePeriod is a free data retrieval call binding the contract method 0xd1a87f64.
//
// Solidity: function maxTimePeriod() view returns(uint256)
func (_StorageProviderRegistry *StorageProviderRegistryCaller) MaxTimePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StorageProviderRegistry.contract.Call(opts, &out, "maxTimePeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTimePeriod is a free data retrieval call binding the contract method 0xd1a87f64.
//
// Solidity: function maxTimePeriod() view returns(uint256)
func (_StorageProviderRegistry *StorageProviderRegistrySession) MaxTimePeriod() (*big.Int, error) {
	return _StorageProviderRegistry.Contract.MaxTimePeriod(&_StorageProviderRegistry.CallOpts)
}

// MaxTimePeriod is a free data retrieval call binding the contract method 0xd1a87f64.
//
// Solidity: function maxTimePeriod() view returns(uint256)
func (_StorageProviderRegistry *StorageProviderRegistryCallerSession) MaxTimePeriod() (*big.Int, error) {
	return _StorageProviderRegistry.Contract.MaxTimePeriod(&_StorageProviderRegistry.CallOpts)
}

// MinTimePeriod is a free data retrieval call binding the contract method 0x7b8d8144.
//
// Solidity: function minTimePeriod() view returns(uint256)
func (_StorageProviderRegistry *StorageProviderRegistryCaller) MinTimePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StorageProviderRegistry.contract.Call(opts, &out, "minTimePeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinTimePeriod is a free data retrieval call binding the contract method 0x7b8d8144.
//
// Solidity: function minTimePeriod() view returns(uint256)
func (_StorageProviderRegistry *StorageProviderRegistrySession) MinTimePeriod() (*big.Int, error) {
	return _StorageProviderRegistry.Contract.MinTimePeriod(&_StorageProviderRegistry.CallOpts)
}

// MinTimePeriod is a free data retrieval call binding the contract method 0x7b8d8144.
//
// Solidity: function minTimePeriod() view returns(uint256)
func (_StorageProviderRegistry *StorageProviderRegistryCallerSession) MinTimePeriod() (*big.Int, error) {
	return _StorageProviderRegistry.Contract.MinTimePeriod(&_StorageProviderRegistry.CallOpts)
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

// StorageProviders is a free data retrieval call binding the contract method 0xd903adad.
//
// Solidity: function storageProviders(bytes ) view returns(bool active, address targetPool, bytes miner, uint256 allocationLimit, uint256 usedAllocation, uint256 accruedRewards, uint256 lockedRewards, uint256 maxRedeemablePeriod)
func (_StorageProviderRegistry *StorageProviderRegistryCaller) StorageProviders(opts *bind.CallOpts, arg0 []byte) (struct {
	Active              bool
	TargetPool          common.Address
	Miner               []byte
	AllocationLimit     *big.Int
	UsedAllocation      *big.Int
	AccruedRewards      *big.Int
	LockedRewards       *big.Int
	MaxRedeemablePeriod *big.Int
}, error) {
	var out []interface{}
	err := _StorageProviderRegistry.contract.Call(opts, &out, "storageProviders", arg0)

	outstruct := new(struct {
		Active              bool
		TargetPool          common.Address
		Miner               []byte
		AllocationLimit     *big.Int
		UsedAllocation      *big.Int
		AccruedRewards      *big.Int
		LockedRewards       *big.Int
		MaxRedeemablePeriod *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Active = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.TargetPool = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Miner = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.AllocationLimit = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.UsedAllocation = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.AccruedRewards = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.LockedRewards = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.MaxRedeemablePeriod = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// StorageProviders is a free data retrieval call binding the contract method 0xd903adad.
//
// Solidity: function storageProviders(bytes ) view returns(bool active, address targetPool, bytes miner, uint256 allocationLimit, uint256 usedAllocation, uint256 accruedRewards, uint256 lockedRewards, uint256 maxRedeemablePeriod)
func (_StorageProviderRegistry *StorageProviderRegistrySession) StorageProviders(arg0 []byte) (struct {
	Active              bool
	TargetPool          common.Address
	Miner               []byte
	AllocationLimit     *big.Int
	UsedAllocation      *big.Int
	AccruedRewards      *big.Int
	LockedRewards       *big.Int
	MaxRedeemablePeriod *big.Int
}, error) {
	return _StorageProviderRegistry.Contract.StorageProviders(&_StorageProviderRegistry.CallOpts, arg0)
}

// StorageProviders is a free data retrieval call binding the contract method 0xd903adad.
//
// Solidity: function storageProviders(bytes ) view returns(bool active, address targetPool, bytes miner, uint256 allocationLimit, uint256 usedAllocation, uint256 accruedRewards, uint256 lockedRewards, uint256 maxRedeemablePeriod)
func (_StorageProviderRegistry *StorageProviderRegistryCallerSession) StorageProviders(arg0 []byte) (struct {
	Active              bool
	TargetPool          common.Address
	Miner               []byte
	AllocationLimit     *big.Int
	UsedAllocation      *big.Int
	AccruedRewards      *big.Int
	LockedRewards       *big.Int
	MaxRedeemablePeriod *big.Int
}, error) {
	return _StorageProviderRegistry.Contract.StorageProviders(&_StorageProviderRegistry.CallOpts, arg0)
}

// TotalInactiveStorageProviders is a free data retrieval call binding the contract method 0xb2f7be76.
//
// Solidity: function totalInactiveStorageProviders() view returns(uint256 _value)
func (_StorageProviderRegistry *StorageProviderRegistryCaller) TotalInactiveStorageProviders(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StorageProviderRegistry.contract.Call(opts, &out, "totalInactiveStorageProviders")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalInactiveStorageProviders is a free data retrieval call binding the contract method 0xb2f7be76.
//
// Solidity: function totalInactiveStorageProviders() view returns(uint256 _value)
func (_StorageProviderRegistry *StorageProviderRegistrySession) TotalInactiveStorageProviders() (*big.Int, error) {
	return _StorageProviderRegistry.Contract.TotalInactiveStorageProviders(&_StorageProviderRegistry.CallOpts)
}

// TotalInactiveStorageProviders is a free data retrieval call binding the contract method 0xb2f7be76.
//
// Solidity: function totalInactiveStorageProviders() view returns(uint256 _value)
func (_StorageProviderRegistry *StorageProviderRegistryCallerSession) TotalInactiveStorageProviders() (*big.Int, error) {
	return _StorageProviderRegistry.Contract.TotalInactiveStorageProviders(&_StorageProviderRegistry.CallOpts)
}

// TotalStorageProviders is a free data retrieval call binding the contract method 0xa1e25bd0.
//
// Solidity: function totalStorageProviders() view returns(uint256 _value)
func (_StorageProviderRegistry *StorageProviderRegistryCaller) TotalStorageProviders(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StorageProviderRegistry.contract.Call(opts, &out, "totalStorageProviders")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStorageProviders is a free data retrieval call binding the contract method 0xa1e25bd0.
//
// Solidity: function totalStorageProviders() view returns(uint256 _value)
func (_StorageProviderRegistry *StorageProviderRegistrySession) TotalStorageProviders() (*big.Int, error) {
	return _StorageProviderRegistry.Contract.TotalStorageProviders(&_StorageProviderRegistry.CallOpts)
}

// TotalStorageProviders is a free data retrieval call binding the contract method 0xa1e25bd0.
//
// Solidity: function totalStorageProviders() view returns(uint256 _value)
func (_StorageProviderRegistry *StorageProviderRegistryCallerSession) TotalStorageProviders() (*big.Int, error) {
	return _StorageProviderRegistry.Contract.TotalStorageProviders(&_StorageProviderRegistry.CallOpts)
}

// AcceptBeneficiaryAddress is a paid mutator transaction binding the contract method 0x8588b598.
//
// Solidity: function acceptBeneficiaryAddress(bytes _provider, address _beneficiaryAddress) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactor) AcceptBeneficiaryAddress(opts *bind.TransactOpts, _provider []byte, _beneficiaryAddress common.Address) (*types.Transaction, error) {
	return _StorageProviderRegistry.contract.Transact(opts, "acceptBeneficiaryAddress", _provider, _beneficiaryAddress)
}

// AcceptBeneficiaryAddress is a paid mutator transaction binding the contract method 0x8588b598.
//
// Solidity: function acceptBeneficiaryAddress(bytes _provider, address _beneficiaryAddress) returns()
func (_StorageProviderRegistry *StorageProviderRegistrySession) AcceptBeneficiaryAddress(_provider []byte, _beneficiaryAddress common.Address) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.AcceptBeneficiaryAddress(&_StorageProviderRegistry.TransactOpts, _provider, _beneficiaryAddress)
}

// AcceptBeneficiaryAddress is a paid mutator transaction binding the contract method 0x8588b598.
//
// Solidity: function acceptBeneficiaryAddress(bytes _provider, address _beneficiaryAddress) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactorSession) AcceptBeneficiaryAddress(_provider []byte, _beneficiaryAddress common.Address) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.AcceptBeneficiaryAddress(&_StorageProviderRegistry.TransactOpts, _provider, _beneficiaryAddress)
}

// ChangeBeneficiaryAddress is a paid mutator transaction binding the contract method 0x12d0e65a.
//
// Solidity: function changeBeneficiaryAddress(address _beneficiaryAddress) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactor) ChangeBeneficiaryAddress(opts *bind.TransactOpts, _beneficiaryAddress common.Address) (*types.Transaction, error) {
	return _StorageProviderRegistry.contract.Transact(opts, "changeBeneficiaryAddress", _beneficiaryAddress)
}

// ChangeBeneficiaryAddress is a paid mutator transaction binding the contract method 0x12d0e65a.
//
// Solidity: function changeBeneficiaryAddress(address _beneficiaryAddress) returns()
func (_StorageProviderRegistry *StorageProviderRegistrySession) ChangeBeneficiaryAddress(_beneficiaryAddress common.Address) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.ChangeBeneficiaryAddress(&_StorageProviderRegistry.TransactOpts, _beneficiaryAddress)
}

// ChangeBeneficiaryAddress is a paid mutator transaction binding the contract method 0x12d0e65a.
//
// Solidity: function changeBeneficiaryAddress(address _beneficiaryAddress) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactorSession) ChangeBeneficiaryAddress(_beneficiaryAddress common.Address) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.ChangeBeneficiaryAddress(&_StorageProviderRegistry.TransactOpts, _beneficiaryAddress)
}

// DeactivateStorageProvider is a paid mutator transaction binding the contract method 0xabb75c9a.
//
// Solidity: function deactivateStorageProvider(bytes _provider) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactor) DeactivateStorageProvider(opts *bind.TransactOpts, _provider []byte) (*types.Transaction, error) {
	return _StorageProviderRegistry.contract.Transact(opts, "deactivateStorageProvider", _provider)
}

// DeactivateStorageProvider is a paid mutator transaction binding the contract method 0xabb75c9a.
//
// Solidity: function deactivateStorageProvider(bytes _provider) returns()
func (_StorageProviderRegistry *StorageProviderRegistrySession) DeactivateStorageProvider(_provider []byte) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.DeactivateStorageProvider(&_StorageProviderRegistry.TransactOpts, _provider)
}

// DeactivateStorageProvider is a paid mutator transaction binding the contract method 0xabb75c9a.
//
// Solidity: function deactivateStorageProvider(bytes _provider) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactorSession) DeactivateStorageProvider(_provider []byte) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.DeactivateStorageProvider(&_StorageProviderRegistry.TransactOpts, _provider)
}

// IncreaseRewards is a paid mutator transaction binding the contract method 0x1f545e31.
//
// Solidity: function increaseRewards(bytes _provider, uint256 _accuredRewards, uint256 _lockedRewards) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactor) IncreaseRewards(opts *bind.TransactOpts, _provider []byte, _accuredRewards *big.Int, _lockedRewards *big.Int) (*types.Transaction, error) {
	return _StorageProviderRegistry.contract.Transact(opts, "increaseRewards", _provider, _accuredRewards, _lockedRewards)
}

// IncreaseRewards is a paid mutator transaction binding the contract method 0x1f545e31.
//
// Solidity: function increaseRewards(bytes _provider, uint256 _accuredRewards, uint256 _lockedRewards) returns()
func (_StorageProviderRegistry *StorageProviderRegistrySession) IncreaseRewards(_provider []byte, _accuredRewards *big.Int, _lockedRewards *big.Int) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.IncreaseRewards(&_StorageProviderRegistry.TransactOpts, _provider, _accuredRewards, _lockedRewards)
}

// IncreaseRewards is a paid mutator transaction binding the contract method 0x1f545e31.
//
// Solidity: function increaseRewards(bytes _provider, uint256 _accuredRewards, uint256 _lockedRewards) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactorSession) IncreaseRewards(_provider []byte, _accuredRewards *big.Int, _lockedRewards *big.Int) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.IncreaseRewards(&_StorageProviderRegistry.TransactOpts, _provider, _accuredRewards, _lockedRewards)
}

// IncreaseUsedAllocation is a paid mutator transaction binding the contract method 0xdb389409.
//
// Solidity: function increaseUsedAllocation(bytes _provider, uint256 _allocated) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactor) IncreaseUsedAllocation(opts *bind.TransactOpts, _provider []byte, _allocated *big.Int) (*types.Transaction, error) {
	return _StorageProviderRegistry.contract.Transact(opts, "increaseUsedAllocation", _provider, _allocated)
}

// IncreaseUsedAllocation is a paid mutator transaction binding the contract method 0xdb389409.
//
// Solidity: function increaseUsedAllocation(bytes _provider, uint256 _allocated) returns()
func (_StorageProviderRegistry *StorageProviderRegistrySession) IncreaseUsedAllocation(_provider []byte, _allocated *big.Int) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.IncreaseUsedAllocation(&_StorageProviderRegistry.TransactOpts, _provider, _allocated)
}

// IncreaseUsedAllocation is a paid mutator transaction binding the contract method 0xdb389409.
//
// Solidity: function increaseUsedAllocation(bytes _provider, uint256 _allocated) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactorSession) IncreaseUsedAllocation(_provider []byte, _allocated *big.Int) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.IncreaseUsedAllocation(&_StorageProviderRegistry.TransactOpts, _provider, _allocated)
}

// Register is a paid mutator transaction binding the contract method 0xb4699eb7.
//
// Solidity: function register(bytes _miner, address _targetPool, uint256 _allocationLimit, uint256 _period) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactor) Register(opts *bind.TransactOpts, _miner []byte, _targetPool common.Address, _allocationLimit *big.Int, _period *big.Int) (*types.Transaction, error) {
	return _StorageProviderRegistry.contract.Transact(opts, "register", _miner, _targetPool, _allocationLimit, _period)
}

// Register is a paid mutator transaction binding the contract method 0xb4699eb7.
//
// Solidity: function register(bytes _miner, address _targetPool, uint256 _allocationLimit, uint256 _period) returns()
func (_StorageProviderRegistry *StorageProviderRegistrySession) Register(_miner []byte, _targetPool common.Address, _allocationLimit *big.Int, _period *big.Int) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.Register(&_StorageProviderRegistry.TransactOpts, _miner, _targetPool, _allocationLimit, _period)
}

// Register is a paid mutator transaction binding the contract method 0xb4699eb7.
//
// Solidity: function register(bytes _miner, address _targetPool, uint256 _allocationLimit, uint256 _period) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactorSession) Register(_miner []byte, _targetPool common.Address, _allocationLimit *big.Int, _period *big.Int) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.Register(&_StorageProviderRegistry.TransactOpts, _miner, _targetPool, _allocationLimit, _period)
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

// SetAllocationLimit is a paid mutator transaction binding the contract method 0x82f2efc7.
//
// Solidity: function setAllocationLimit(bytes _provider, uint256 _allocationLimit) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactor) SetAllocationLimit(opts *bind.TransactOpts, _provider []byte, _allocationLimit *big.Int) (*types.Transaction, error) {
	return _StorageProviderRegistry.contract.Transact(opts, "setAllocationLimit", _provider, _allocationLimit)
}

// SetAllocationLimit is a paid mutator transaction binding the contract method 0x82f2efc7.
//
// Solidity: function setAllocationLimit(bytes _provider, uint256 _allocationLimit) returns()
func (_StorageProviderRegistry *StorageProviderRegistrySession) SetAllocationLimit(_provider []byte, _allocationLimit *big.Int) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.SetAllocationLimit(&_StorageProviderRegistry.TransactOpts, _provider, _allocationLimit)
}

// SetAllocationLimit is a paid mutator transaction binding the contract method 0x82f2efc7.
//
// Solidity: function setAllocationLimit(bytes _provider, uint256 _allocationLimit) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactorSession) SetAllocationLimit(_provider []byte, _allocationLimit *big.Int) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.SetAllocationLimit(&_StorageProviderRegistry.TransactOpts, _provider, _allocationLimit)
}

// SetCollateralAddress is a paid mutator transaction binding the contract method 0x029d262e.
//
// Solidity: function setCollateralAddress(address _collateral) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactor) SetCollateralAddress(opts *bind.TransactOpts, _collateral common.Address) (*types.Transaction, error) {
	return _StorageProviderRegistry.contract.Transact(opts, "setCollateralAddress", _collateral)
}

// SetCollateralAddress is a paid mutator transaction binding the contract method 0x029d262e.
//
// Solidity: function setCollateralAddress(address _collateral) returns()
func (_StorageProviderRegistry *StorageProviderRegistrySession) SetCollateralAddress(_collateral common.Address) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.SetCollateralAddress(&_StorageProviderRegistry.TransactOpts, _collateral)
}

// SetCollateralAddress is a paid mutator transaction binding the contract method 0x029d262e.
//
// Solidity: function setCollateralAddress(address _collateral) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactorSession) SetCollateralAddress(_collateral common.Address) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.SetCollateralAddress(&_StorageProviderRegistry.TransactOpts, _collateral)
}

// SetMaxRedeemablePeriod is a paid mutator transaction binding the contract method 0xc30f4a27.
//
// Solidity: function setMaxRedeemablePeriod(bytes _provider, uint256 _period) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactor) SetMaxRedeemablePeriod(opts *bind.TransactOpts, _provider []byte, _period *big.Int) (*types.Transaction, error) {
	return _StorageProviderRegistry.contract.Transact(opts, "setMaxRedeemablePeriod", _provider, _period)
}

// SetMaxRedeemablePeriod is a paid mutator transaction binding the contract method 0xc30f4a27.
//
// Solidity: function setMaxRedeemablePeriod(bytes _provider, uint256 _period) returns()
func (_StorageProviderRegistry *StorageProviderRegistrySession) SetMaxRedeemablePeriod(_provider []byte, _period *big.Int) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.SetMaxRedeemablePeriod(&_StorageProviderRegistry.TransactOpts, _provider, _period)
}

// SetMaxRedeemablePeriod is a paid mutator transaction binding the contract method 0xc30f4a27.
//
// Solidity: function setMaxRedeemablePeriod(bytes _provider, uint256 _period) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactorSession) SetMaxRedeemablePeriod(_provider []byte, _period *big.Int) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.SetMaxRedeemablePeriod(&_StorageProviderRegistry.TransactOpts, _provider, _period)
}

// SetMinerAddress is a paid mutator transaction binding the contract method 0x9703d285.
//
// Solidity: function setMinerAddress(bytes _provider, bytes _miner) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactor) SetMinerAddress(opts *bind.TransactOpts, _provider []byte, _miner []byte) (*types.Transaction, error) {
	return _StorageProviderRegistry.contract.Transact(opts, "setMinerAddress", _provider, _miner)
}

// SetMinerAddress is a paid mutator transaction binding the contract method 0x9703d285.
//
// Solidity: function setMinerAddress(bytes _provider, bytes _miner) returns()
func (_StorageProviderRegistry *StorageProviderRegistrySession) SetMinerAddress(_provider []byte, _miner []byte) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.SetMinerAddress(&_StorageProviderRegistry.TransactOpts, _provider, _miner)
}

// SetMinerAddress is a paid mutator transaction binding the contract method 0x9703d285.
//
// Solidity: function setMinerAddress(bytes _provider, bytes _miner) returns()
func (_StorageProviderRegistry *StorageProviderRegistryTransactorSession) SetMinerAddress(_provider []byte, _miner []byte) (*types.Transaction, error) {
	return _StorageProviderRegistry.Contract.SetMinerAddress(&_StorageProviderRegistry.TransactOpts, _provider, _miner)
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
	Provider []byte
	Rewards  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStorageProviderAccruedRewards is a free log retrieval operation binding the contract event 0xf913241f9eaabb10e21e29dee673dac2b0dcc086bed112cefcbd8ff5d6a91f35.
//
// Solidity: event StorageProviderAccruedRewards(bytes provider, uint256 rewards)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) FilterStorageProviderAccruedRewards(opts *bind.FilterOpts) (*StorageProviderRegistryStorageProviderAccruedRewardsIterator, error) {

	logs, sub, err := _StorageProviderRegistry.contract.FilterLogs(opts, "StorageProviderAccruedRewards")
	if err != nil {
		return nil, err
	}
	return &StorageProviderRegistryStorageProviderAccruedRewardsIterator{contract: _StorageProviderRegistry.contract, event: "StorageProviderAccruedRewards", logs: logs, sub: sub}, nil
}

// WatchStorageProviderAccruedRewards is a free log subscription operation binding the contract event 0xf913241f9eaabb10e21e29dee673dac2b0dcc086bed112cefcbd8ff5d6a91f35.
//
// Solidity: event StorageProviderAccruedRewards(bytes provider, uint256 rewards)
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

// ParseStorageProviderAccruedRewards is a log parse operation binding the contract event 0xf913241f9eaabb10e21e29dee673dac2b0dcc086bed112cefcbd8ff5d6a91f35.
//
// Solidity: event StorageProviderAccruedRewards(bytes provider, uint256 rewards)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) ParseStorageProviderAccruedRewards(log types.Log) (*StorageProviderRegistryStorageProviderAccruedRewards, error) {
	event := new(StorageProviderRegistryStorageProviderAccruedRewards)
	if err := _StorageProviderRegistry.contract.UnpackLog(event, "StorageProviderAccruedRewards", log); err != nil {
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
	Provider        []byte
	AllocationLimit *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStorageProviderAllocationLimitUpdate is a free log retrieval operation binding the contract event 0x5829760c5f5d66e198f7802533e19bf232b6219fd8bce4d85f1fd741f6c0fd39.
//
// Solidity: event StorageProviderAllocationLimitUpdate(bytes provider, uint256 allocationLimit)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) FilterStorageProviderAllocationLimitUpdate(opts *bind.FilterOpts) (*StorageProviderRegistryStorageProviderAllocationLimitUpdateIterator, error) {

	logs, sub, err := _StorageProviderRegistry.contract.FilterLogs(opts, "StorageProviderAllocationLimitUpdate")
	if err != nil {
		return nil, err
	}
	return &StorageProviderRegistryStorageProviderAllocationLimitUpdateIterator{contract: _StorageProviderRegistry.contract, event: "StorageProviderAllocationLimitUpdate", logs: logs, sub: sub}, nil
}

// WatchStorageProviderAllocationLimitUpdate is a free log subscription operation binding the contract event 0x5829760c5f5d66e198f7802533e19bf232b6219fd8bce4d85f1fd741f6c0fd39.
//
// Solidity: event StorageProviderAllocationLimitUpdate(bytes provider, uint256 allocationLimit)
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

// ParseStorageProviderAllocationLimitUpdate is a log parse operation binding the contract event 0x5829760c5f5d66e198f7802533e19bf232b6219fd8bce4d85f1fd741f6c0fd39.
//
// Solidity: event StorageProviderAllocationLimitUpdate(bytes provider, uint256 allocationLimit)
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
	Provider       []byte
	UsedAllocation *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterStorageProviderAllocationUsed is a free log retrieval operation binding the contract event 0x3649f4e411a63b43c9ae90b162ef7b14d19a0d6ace211b0f66a07615b05b45ab.
//
// Solidity: event StorageProviderAllocationUsed(bytes provider, uint256 usedAllocation)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) FilterStorageProviderAllocationUsed(opts *bind.FilterOpts) (*StorageProviderRegistryStorageProviderAllocationUsedIterator, error) {

	logs, sub, err := _StorageProviderRegistry.contract.FilterLogs(opts, "StorageProviderAllocationUsed")
	if err != nil {
		return nil, err
	}
	return &StorageProviderRegistryStorageProviderAllocationUsedIterator{contract: _StorageProviderRegistry.contract, event: "StorageProviderAllocationUsed", logs: logs, sub: sub}, nil
}

// WatchStorageProviderAllocationUsed is a free log subscription operation binding the contract event 0x3649f4e411a63b43c9ae90b162ef7b14d19a0d6ace211b0f66a07615b05b45ab.
//
// Solidity: event StorageProviderAllocationUsed(bytes provider, uint256 usedAllocation)
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

// ParseStorageProviderAllocationUsed is a log parse operation binding the contract event 0x3649f4e411a63b43c9ae90b162ef7b14d19a0d6ace211b0f66a07615b05b45ab.
//
// Solidity: event StorageProviderAllocationUsed(bytes provider, uint256 usedAllocation)
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
	Provider []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStorageProviderBeneficiaryAddressAccepted is a free log retrieval operation binding the contract event 0x8b32ebd74b8f29efad1c38a9c13977ffe07b909ddf14da665c327681ad0172a2.
//
// Solidity: event StorageProviderBeneficiaryAddressAccepted(bytes _provider)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) FilterStorageProviderBeneficiaryAddressAccepted(opts *bind.FilterOpts) (*StorageProviderRegistryStorageProviderBeneficiaryAddressAcceptedIterator, error) {

	logs, sub, err := _StorageProviderRegistry.contract.FilterLogs(opts, "StorageProviderBeneficiaryAddressAccepted")
	if err != nil {
		return nil, err
	}
	return &StorageProviderRegistryStorageProviderBeneficiaryAddressAcceptedIterator{contract: _StorageProviderRegistry.contract, event: "StorageProviderBeneficiaryAddressAccepted", logs: logs, sub: sub}, nil
}

// WatchStorageProviderBeneficiaryAddressAccepted is a free log subscription operation binding the contract event 0x8b32ebd74b8f29efad1c38a9c13977ffe07b909ddf14da665c327681ad0172a2.
//
// Solidity: event StorageProviderBeneficiaryAddressAccepted(bytes _provider)
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

// ParseStorageProviderBeneficiaryAddressAccepted is a log parse operation binding the contract event 0x8b32ebd74b8f29efad1c38a9c13977ffe07b909ddf14da665c327681ad0172a2.
//
// Solidity: event StorageProviderBeneficiaryAddressAccepted(bytes _provider)
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
// Solidity: event StorageProviderBeneficiaryAddressUpdated(address _beneficiaryAddress)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) FilterStorageProviderBeneficiaryAddressUpdated(opts *bind.FilterOpts) (*StorageProviderRegistryStorageProviderBeneficiaryAddressUpdatedIterator, error) {

	logs, sub, err := _StorageProviderRegistry.contract.FilterLogs(opts, "StorageProviderBeneficiaryAddressUpdated")
	if err != nil {
		return nil, err
	}
	return &StorageProviderRegistryStorageProviderBeneficiaryAddressUpdatedIterator{contract: _StorageProviderRegistry.contract, event: "StorageProviderBeneficiaryAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchStorageProviderBeneficiaryAddressUpdated is a free log subscription operation binding the contract event 0x906cdab773befe815c0a81a3837dc4df80e433c0c26e4ee402007515bb777c1c.
//
// Solidity: event StorageProviderBeneficiaryAddressUpdated(address _beneficiaryAddress)
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
// Solidity: event StorageProviderBeneficiaryAddressUpdated(address _beneficiaryAddress)
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
	Provider []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStorageProviderDeactivated is a free log retrieval operation binding the contract event 0x4c02ca55b73202eaa7b064870cac49386835e51269059394aa7ceb9a69db3493.
//
// Solidity: event StorageProviderDeactivated(bytes provider)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) FilterStorageProviderDeactivated(opts *bind.FilterOpts) (*StorageProviderRegistryStorageProviderDeactivatedIterator, error) {

	logs, sub, err := _StorageProviderRegistry.contract.FilterLogs(opts, "StorageProviderDeactivated")
	if err != nil {
		return nil, err
	}
	return &StorageProviderRegistryStorageProviderDeactivatedIterator{contract: _StorageProviderRegistry.contract, event: "StorageProviderDeactivated", logs: logs, sub: sub}, nil
}

// WatchStorageProviderDeactivated is a free log subscription operation binding the contract event 0x4c02ca55b73202eaa7b064870cac49386835e51269059394aa7ceb9a69db3493.
//
// Solidity: event StorageProviderDeactivated(bytes provider)
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

// ParseStorageProviderDeactivated is a log parse operation binding the contract event 0x4c02ca55b73202eaa7b064870cac49386835e51269059394aa7ceb9a69db3493.
//
// Solidity: event StorageProviderDeactivated(bytes provider)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) ParseStorageProviderDeactivated(log types.Log) (*StorageProviderRegistryStorageProviderDeactivated, error) {
	event := new(StorageProviderRegistryStorageProviderDeactivated)
	if err := _StorageProviderRegistry.contract.UnpackLog(event, "StorageProviderDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageProviderRegistryStorageProviderLockedRewardsIterator is returned from FilterStorageProviderLockedRewards and is used to iterate over the raw logs and unpacked data for StorageProviderLockedRewards events raised by the StorageProviderRegistry contract.
type StorageProviderRegistryStorageProviderLockedRewardsIterator struct {
	Event *StorageProviderRegistryStorageProviderLockedRewards // Event containing the contract specifics and raw log

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
func (it *StorageProviderRegistryStorageProviderLockedRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProviderRegistryStorageProviderLockedRewards)
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
		it.Event = new(StorageProviderRegistryStorageProviderLockedRewards)
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
func (it *StorageProviderRegistryStorageProviderLockedRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProviderRegistryStorageProviderLockedRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProviderRegistryStorageProviderLockedRewards represents a StorageProviderLockedRewards event raised by the StorageProviderRegistry contract.
type StorageProviderRegistryStorageProviderLockedRewards struct {
	Provider []byte
	Rewards  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStorageProviderLockedRewards is a free log retrieval operation binding the contract event 0x922d39de9bd52bd6cb20993efb950a62b2918aacc8098f262b9fa02d5031abb5.
//
// Solidity: event StorageProviderLockedRewards(bytes provider, uint256 rewards)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) FilterStorageProviderLockedRewards(opts *bind.FilterOpts) (*StorageProviderRegistryStorageProviderLockedRewardsIterator, error) {

	logs, sub, err := _StorageProviderRegistry.contract.FilterLogs(opts, "StorageProviderLockedRewards")
	if err != nil {
		return nil, err
	}
	return &StorageProviderRegistryStorageProviderLockedRewardsIterator{contract: _StorageProviderRegistry.contract, event: "StorageProviderLockedRewards", logs: logs, sub: sub}, nil
}

// WatchStorageProviderLockedRewards is a free log subscription operation binding the contract event 0x922d39de9bd52bd6cb20993efb950a62b2918aacc8098f262b9fa02d5031abb5.
//
// Solidity: event StorageProviderLockedRewards(bytes provider, uint256 rewards)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) WatchStorageProviderLockedRewards(opts *bind.WatchOpts, sink chan<- *StorageProviderRegistryStorageProviderLockedRewards) (event.Subscription, error) {

	logs, sub, err := _StorageProviderRegistry.contract.WatchLogs(opts, "StorageProviderLockedRewards")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProviderRegistryStorageProviderLockedRewards)
				if err := _StorageProviderRegistry.contract.UnpackLog(event, "StorageProviderLockedRewards", log); err != nil {
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

// ParseStorageProviderLockedRewards is a log parse operation binding the contract event 0x922d39de9bd52bd6cb20993efb950a62b2918aacc8098f262b9fa02d5031abb5.
//
// Solidity: event StorageProviderLockedRewards(bytes provider, uint256 rewards)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) ParseStorageProviderLockedRewards(log types.Log) (*StorageProviderRegistryStorageProviderLockedRewards, error) {
	event := new(StorageProviderRegistryStorageProviderLockedRewards)
	if err := _StorageProviderRegistry.contract.UnpackLog(event, "StorageProviderLockedRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageProviderRegistryStorageProviderMaxRedeemablePeriodUpdateIterator is returned from FilterStorageProviderMaxRedeemablePeriodUpdate and is used to iterate over the raw logs and unpacked data for StorageProviderMaxRedeemablePeriodUpdate events raised by the StorageProviderRegistry contract.
type StorageProviderRegistryStorageProviderMaxRedeemablePeriodUpdateIterator struct {
	Event *StorageProviderRegistryStorageProviderMaxRedeemablePeriodUpdate // Event containing the contract specifics and raw log

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
func (it *StorageProviderRegistryStorageProviderMaxRedeemablePeriodUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProviderRegistryStorageProviderMaxRedeemablePeriodUpdate)
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
		it.Event = new(StorageProviderRegistryStorageProviderMaxRedeemablePeriodUpdate)
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
func (it *StorageProviderRegistryStorageProviderMaxRedeemablePeriodUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProviderRegistryStorageProviderMaxRedeemablePeriodUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProviderRegistryStorageProviderMaxRedeemablePeriodUpdate represents a StorageProviderMaxRedeemablePeriodUpdate event raised by the StorageProviderRegistry contract.
type StorageProviderRegistryStorageProviderMaxRedeemablePeriodUpdate struct {
	Provider []byte
	Period   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStorageProviderMaxRedeemablePeriodUpdate is a free log retrieval operation binding the contract event 0x33eb7e48fc0b27df04308755aa8646ccff8e3be14acf92bfde7d8282852833da.
//
// Solidity: event StorageProviderMaxRedeemablePeriodUpdate(bytes provider, uint256 period)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) FilterStorageProviderMaxRedeemablePeriodUpdate(opts *bind.FilterOpts) (*StorageProviderRegistryStorageProviderMaxRedeemablePeriodUpdateIterator, error) {

	logs, sub, err := _StorageProviderRegistry.contract.FilterLogs(opts, "StorageProviderMaxRedeemablePeriodUpdate")
	if err != nil {
		return nil, err
	}
	return &StorageProviderRegistryStorageProviderMaxRedeemablePeriodUpdateIterator{contract: _StorageProviderRegistry.contract, event: "StorageProviderMaxRedeemablePeriodUpdate", logs: logs, sub: sub}, nil
}

// WatchStorageProviderMaxRedeemablePeriodUpdate is a free log subscription operation binding the contract event 0x33eb7e48fc0b27df04308755aa8646ccff8e3be14acf92bfde7d8282852833da.
//
// Solidity: event StorageProviderMaxRedeemablePeriodUpdate(bytes provider, uint256 period)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) WatchStorageProviderMaxRedeemablePeriodUpdate(opts *bind.WatchOpts, sink chan<- *StorageProviderRegistryStorageProviderMaxRedeemablePeriodUpdate) (event.Subscription, error) {

	logs, sub, err := _StorageProviderRegistry.contract.WatchLogs(opts, "StorageProviderMaxRedeemablePeriodUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProviderRegistryStorageProviderMaxRedeemablePeriodUpdate)
				if err := _StorageProviderRegistry.contract.UnpackLog(event, "StorageProviderMaxRedeemablePeriodUpdate", log); err != nil {
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

// ParseStorageProviderMaxRedeemablePeriodUpdate is a log parse operation binding the contract event 0x33eb7e48fc0b27df04308755aa8646ccff8e3be14acf92bfde7d8282852833da.
//
// Solidity: event StorageProviderMaxRedeemablePeriodUpdate(bytes provider, uint256 period)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) ParseStorageProviderMaxRedeemablePeriodUpdate(log types.Log) (*StorageProviderRegistryStorageProviderMaxRedeemablePeriodUpdate, error) {
	event := new(StorageProviderRegistryStorageProviderMaxRedeemablePeriodUpdate)
	if err := _StorageProviderRegistry.contract.UnpackLog(event, "StorageProviderMaxRedeemablePeriodUpdate", log); err != nil {
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
	Provider []byte
	Miner    []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStorageProviderMinerAddressUpdate is a free log retrieval operation binding the contract event 0x33dffc993955f301c60d13e36d9afc2fc2f67a204a47d6e1c897369b41beeabf.
//
// Solidity: event StorageProviderMinerAddressUpdate(bytes provider, bytes miner)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) FilterStorageProviderMinerAddressUpdate(opts *bind.FilterOpts) (*StorageProviderRegistryStorageProviderMinerAddressUpdateIterator, error) {

	logs, sub, err := _StorageProviderRegistry.contract.FilterLogs(opts, "StorageProviderMinerAddressUpdate")
	if err != nil {
		return nil, err
	}
	return &StorageProviderRegistryStorageProviderMinerAddressUpdateIterator{contract: _StorageProviderRegistry.contract, event: "StorageProviderMinerAddressUpdate", logs: logs, sub: sub}, nil
}

// WatchStorageProviderMinerAddressUpdate is a free log subscription operation binding the contract event 0x33dffc993955f301c60d13e36d9afc2fc2f67a204a47d6e1c897369b41beeabf.
//
// Solidity: event StorageProviderMinerAddressUpdate(bytes provider, bytes miner)
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

// ParseStorageProviderMinerAddressUpdate is a log parse operation binding the contract event 0x33dffc993955f301c60d13e36d9afc2fc2f67a204a47d6e1c897369b41beeabf.
//
// Solidity: event StorageProviderMinerAddressUpdate(bytes provider, bytes miner)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) ParseStorageProviderMinerAddressUpdate(log types.Log) (*StorageProviderRegistryStorageProviderMinerAddressUpdate, error) {
	event := new(StorageProviderRegistryStorageProviderMinerAddressUpdate)
	if err := _StorageProviderRegistry.contract.UnpackLog(event, "StorageProviderMinerAddressUpdate", log); err != nil {
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
	Provider        []byte
	Miner           []byte
	TargetPool      common.Address
	AllocationLimit *big.Int
	MaxPeriod       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStorageProviderRegistered is a free log retrieval operation binding the contract event 0x192c847b11e7489f68fb7e25fbf6a795432e4e9c9c9f174a85b77cc86211d9d4.
//
// Solidity: event StorageProviderRegistered(bytes provider, bytes miner, address targetPool, uint256 allocationLimit, uint256 maxPeriod)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) FilterStorageProviderRegistered(opts *bind.FilterOpts) (*StorageProviderRegistryStorageProviderRegisteredIterator, error) {

	logs, sub, err := _StorageProviderRegistry.contract.FilterLogs(opts, "StorageProviderRegistered")
	if err != nil {
		return nil, err
	}
	return &StorageProviderRegistryStorageProviderRegisteredIterator{contract: _StorageProviderRegistry.contract, event: "StorageProviderRegistered", logs: logs, sub: sub}, nil
}

// WatchStorageProviderRegistered is a free log subscription operation binding the contract event 0x192c847b11e7489f68fb7e25fbf6a795432e4e9c9c9f174a85b77cc86211d9d4.
//
// Solidity: event StorageProviderRegistered(bytes provider, bytes miner, address targetPool, uint256 allocationLimit, uint256 maxPeriod)
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

// ParseStorageProviderRegistered is a log parse operation binding the contract event 0x192c847b11e7489f68fb7e25fbf6a795432e4e9c9c9f174a85b77cc86211d9d4.
//
// Solidity: event StorageProviderRegistered(bytes provider, bytes miner, address targetPool, uint256 allocationLimit, uint256 maxPeriod)
func (_StorageProviderRegistry *StorageProviderRegistryFilterer) ParseStorageProviderRegistered(log types.Log) (*StorageProviderRegistryStorageProviderRegistered, error) {
	event := new(StorageProviderRegistryStorageProviderRegistered)
	if err := _StorageProviderRegistry.contract.UnpackLog(event, "StorageProviderRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
