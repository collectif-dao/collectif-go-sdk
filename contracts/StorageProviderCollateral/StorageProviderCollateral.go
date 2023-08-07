// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package StorageProviderCollateral

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

// StorageProviderCollateralMetaData contains all meta data concerning the StorageProviderCollateral contract.
var StorageProviderCollateralMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AllocationOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InactiveActor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InactivePool\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InactiveSP\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InactiveSlashing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientCollateral\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAccess\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidParams\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"ownerId\",\"type\":\"uint64\"}],\"name\":\"ReportRecovery\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"ownerId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashingAmount\",\"type\":\"uint256\"}],\"name\":\"ReportSlashing\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"}],\"name\":\"SetRegistryAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_ownerId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StorageProviderCollateralDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_ownerId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockedCollateral\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"availableCollateral\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isUnlock\",\"type\":\"bool\"}],\"name\":\"StorageProviderCollateralRebalance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_ownerId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashingAmt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"StorageProviderCollateralSlash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_ownerId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevRequirements\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requirements\",\"type\":\"uint256\"}],\"name\":\"StorageProviderCollateralUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_ownerId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StorageProviderCollateralWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseCollateralRequirements\",\"type\":\"uint256\"}],\"name\":\"UpdateBaseCollateralRequirements\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WFIL\",\"outputs\":[{\"internalType\":\"contractIWFIL\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"activeSlashings\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseRequirements\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"collateralRequirements\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"collaterals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"availableCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockedCollateral\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ownerId\",\"type\":\"uint64\"}],\"name\":\"fit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ownerId\",\"type\":\"uint64\"}],\"name\":\"getAvailableCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ownerId\",\"type\":\"uint64\"}],\"name\":\"getCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ownerId\",\"type\":\"uint64\"}],\"name\":\"getCollateralRequirements\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ownerId\",\"type\":\"uint64\"}],\"name\":\"getDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ownerId\",\"type\":\"uint64\"}],\"name\":\"getLockedCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIWFIL\",\"name\":\"_wFIL\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_resolver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_baseRequirements\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ownerId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_minerId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_allocated\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ownerId\",\"type\":\"uint64\"}],\"name\":\"reportRecovery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ownerId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_slashingAmt\",\"type\":\"uint256\"}],\"name\":\"reportSlashing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"slashings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requirements\",\"type\":\"uint256\"}],\"name\":\"updateBaseCollateralRequirements\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ownerId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"requirements\",\"type\":\"uint256\"}],\"name\":\"updateCollateralRequirements\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// StorageProviderCollateralABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageProviderCollateralMetaData.ABI instead.
var StorageProviderCollateralABI = StorageProviderCollateralMetaData.ABI

// StorageProviderCollateral is an auto generated Go binding around an Ethereum contract.
type StorageProviderCollateral struct {
	StorageProviderCollateralCaller     // Read-only binding to the contract
	StorageProviderCollateralTransactor // Write-only binding to the contract
	StorageProviderCollateralFilterer   // Log filterer for contract events
}

// StorageProviderCollateralCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageProviderCollateralCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageProviderCollateralTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageProviderCollateralTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageProviderCollateralFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageProviderCollateralFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageProviderCollateralSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageProviderCollateralSession struct {
	Contract     *StorageProviderCollateral // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// StorageProviderCollateralCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageProviderCollateralCallerSession struct {
	Contract *StorageProviderCollateralCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// StorageProviderCollateralTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageProviderCollateralTransactorSession struct {
	Contract     *StorageProviderCollateralTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// StorageProviderCollateralRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageProviderCollateralRaw struct {
	Contract *StorageProviderCollateral // Generic contract binding to access the raw methods on
}

// StorageProviderCollateralCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageProviderCollateralCallerRaw struct {
	Contract *StorageProviderCollateralCaller // Generic read-only contract binding to access the raw methods on
}

// StorageProviderCollateralTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageProviderCollateralTransactorRaw struct {
	Contract *StorageProviderCollateralTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorageProviderCollateral creates a new instance of StorageProviderCollateral, bound to a specific deployed contract.
func NewStorageProviderCollateral(address common.Address, backend bind.ContractBackend) (*StorageProviderCollateral, error) {
	contract, err := bindStorageProviderCollateral(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StorageProviderCollateral{StorageProviderCollateralCaller: StorageProviderCollateralCaller{contract: contract}, StorageProviderCollateralTransactor: StorageProviderCollateralTransactor{contract: contract}, StorageProviderCollateralFilterer: StorageProviderCollateralFilterer{contract: contract}}, nil
}

// NewStorageProviderCollateralCaller creates a new read-only instance of StorageProviderCollateral, bound to a specific deployed contract.
func NewStorageProviderCollateralCaller(address common.Address, caller bind.ContractCaller) (*StorageProviderCollateralCaller, error) {
	contract, err := bindStorageProviderCollateral(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageProviderCollateralCaller{contract: contract}, nil
}

// NewStorageProviderCollateralTransactor creates a new write-only instance of StorageProviderCollateral, bound to a specific deployed contract.
func NewStorageProviderCollateralTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageProviderCollateralTransactor, error) {
	contract, err := bindStorageProviderCollateral(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageProviderCollateralTransactor{contract: contract}, nil
}

// NewStorageProviderCollateralFilterer creates a new log filterer instance of StorageProviderCollateral, bound to a specific deployed contract.
func NewStorageProviderCollateralFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageProviderCollateralFilterer, error) {
	contract, err := bindStorageProviderCollateral(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageProviderCollateralFilterer{contract: contract}, nil
}

// bindStorageProviderCollateral binds a generic wrapper to an already deployed contract.
func bindStorageProviderCollateral(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StorageProviderCollateralABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StorageProviderCollateral *StorageProviderCollateralRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StorageProviderCollateral.Contract.StorageProviderCollateralCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StorageProviderCollateral *StorageProviderCollateralRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageProviderCollateral.Contract.StorageProviderCollateralTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StorageProviderCollateral *StorageProviderCollateralRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StorageProviderCollateral.Contract.StorageProviderCollateralTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StorageProviderCollateral *StorageProviderCollateralCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StorageProviderCollateral.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StorageProviderCollateral *StorageProviderCollateralTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageProviderCollateral.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StorageProviderCollateral *StorageProviderCollateralTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StorageProviderCollateral.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StorageProviderCollateral *StorageProviderCollateralCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StorageProviderCollateral.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StorageProviderCollateral *StorageProviderCollateralSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StorageProviderCollateral.Contract.DEFAULTADMINROLE(&_StorageProviderCollateral.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StorageProviderCollateral *StorageProviderCollateralCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StorageProviderCollateral.Contract.DEFAULTADMINROLE(&_StorageProviderCollateral.CallOpts)
}

// WFIL is a free data retrieval call binding the contract method 0x09448a86.
//
// Solidity: function WFIL() view returns(address)
func (_StorageProviderCollateral *StorageProviderCollateralCaller) WFIL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StorageProviderCollateral.contract.Call(opts, &out, "WFIL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WFIL is a free data retrieval call binding the contract method 0x09448a86.
//
// Solidity: function WFIL() view returns(address)
func (_StorageProviderCollateral *StorageProviderCollateralSession) WFIL() (common.Address, error) {
	return _StorageProviderCollateral.Contract.WFIL(&_StorageProviderCollateral.CallOpts)
}

// WFIL is a free data retrieval call binding the contract method 0x09448a86.
//
// Solidity: function WFIL() view returns(address)
func (_StorageProviderCollateral *StorageProviderCollateralCallerSession) WFIL() (common.Address, error) {
	return _StorageProviderCollateral.Contract.WFIL(&_StorageProviderCollateral.CallOpts)
}

// ActiveSlashings is a free data retrieval call binding the contract method 0x44324a79.
//
// Solidity: function activeSlashings(uint64 ) view returns(bool)
func (_StorageProviderCollateral *StorageProviderCollateralCaller) ActiveSlashings(opts *bind.CallOpts, arg0 uint64) (bool, error) {
	var out []interface{}
	err := _StorageProviderCollateral.contract.Call(opts, &out, "activeSlashings", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ActiveSlashings is a free data retrieval call binding the contract method 0x44324a79.
//
// Solidity: function activeSlashings(uint64 ) view returns(bool)
func (_StorageProviderCollateral *StorageProviderCollateralSession) ActiveSlashings(arg0 uint64) (bool, error) {
	return _StorageProviderCollateral.Contract.ActiveSlashings(&_StorageProviderCollateral.CallOpts, arg0)
}

// ActiveSlashings is a free data retrieval call binding the contract method 0x44324a79.
//
// Solidity: function activeSlashings(uint64 ) view returns(bool)
func (_StorageProviderCollateral *StorageProviderCollateralCallerSession) ActiveSlashings(arg0 uint64) (bool, error) {
	return _StorageProviderCollateral.Contract.ActiveSlashings(&_StorageProviderCollateral.CallOpts, arg0)
}

// BaseRequirements is a free data retrieval call binding the contract method 0x8199b668.
//
// Solidity: function baseRequirements() view returns(uint256)
func (_StorageProviderCollateral *StorageProviderCollateralCaller) BaseRequirements(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StorageProviderCollateral.contract.Call(opts, &out, "baseRequirements")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseRequirements is a free data retrieval call binding the contract method 0x8199b668.
//
// Solidity: function baseRequirements() view returns(uint256)
func (_StorageProviderCollateral *StorageProviderCollateralSession) BaseRequirements() (*big.Int, error) {
	return _StorageProviderCollateral.Contract.BaseRequirements(&_StorageProviderCollateral.CallOpts)
}

// BaseRequirements is a free data retrieval call binding the contract method 0x8199b668.
//
// Solidity: function baseRequirements() view returns(uint256)
func (_StorageProviderCollateral *StorageProviderCollateralCallerSession) BaseRequirements() (*big.Int, error) {
	return _StorageProviderCollateral.Contract.BaseRequirements(&_StorageProviderCollateral.CallOpts)
}

// CollateralRequirements is a free data retrieval call binding the contract method 0xce096110.
//
// Solidity: function collateralRequirements(uint64 ) view returns(uint256)
func (_StorageProviderCollateral *StorageProviderCollateralCaller) CollateralRequirements(opts *bind.CallOpts, arg0 uint64) (*big.Int, error) {
	var out []interface{}
	err := _StorageProviderCollateral.contract.Call(opts, &out, "collateralRequirements", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollateralRequirements is a free data retrieval call binding the contract method 0xce096110.
//
// Solidity: function collateralRequirements(uint64 ) view returns(uint256)
func (_StorageProviderCollateral *StorageProviderCollateralSession) CollateralRequirements(arg0 uint64) (*big.Int, error) {
	return _StorageProviderCollateral.Contract.CollateralRequirements(&_StorageProviderCollateral.CallOpts, arg0)
}

// CollateralRequirements is a free data retrieval call binding the contract method 0xce096110.
//
// Solidity: function collateralRequirements(uint64 ) view returns(uint256)
func (_StorageProviderCollateral *StorageProviderCollateralCallerSession) CollateralRequirements(arg0 uint64) (*big.Int, error) {
	return _StorageProviderCollateral.Contract.CollateralRequirements(&_StorageProviderCollateral.CallOpts, arg0)
}

// Collaterals is a free data retrieval call binding the contract method 0x7efedebe.
//
// Solidity: function collaterals(uint64 ) view returns(uint256 availableCollateral, uint256 lockedCollateral)
func (_StorageProviderCollateral *StorageProviderCollateralCaller) Collaterals(opts *bind.CallOpts, arg0 uint64) (struct {
	AvailableCollateral *big.Int
	LockedCollateral    *big.Int
}, error) {
	var out []interface{}
	err := _StorageProviderCollateral.contract.Call(opts, &out, "collaterals", arg0)

	outstruct := new(struct {
		AvailableCollateral *big.Int
		LockedCollateral    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AvailableCollateral = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LockedCollateral = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Collaterals is a free data retrieval call binding the contract method 0x7efedebe.
//
// Solidity: function collaterals(uint64 ) view returns(uint256 availableCollateral, uint256 lockedCollateral)
func (_StorageProviderCollateral *StorageProviderCollateralSession) Collaterals(arg0 uint64) (struct {
	AvailableCollateral *big.Int
	LockedCollateral    *big.Int
}, error) {
	return _StorageProviderCollateral.Contract.Collaterals(&_StorageProviderCollateral.CallOpts, arg0)
}

// Collaterals is a free data retrieval call binding the contract method 0x7efedebe.
//
// Solidity: function collaterals(uint64 ) view returns(uint256 availableCollateral, uint256 lockedCollateral)
func (_StorageProviderCollateral *StorageProviderCollateralCallerSession) Collaterals(arg0 uint64) (struct {
	AvailableCollateral *big.Int
	LockedCollateral    *big.Int
}, error) {
	return _StorageProviderCollateral.Contract.Collaterals(&_StorageProviderCollateral.CallOpts, arg0)
}

// GetAvailableCollateral is a free data retrieval call binding the contract method 0xe460512e.
//
// Solidity: function getAvailableCollateral(uint64 _ownerId) view returns(uint256)
func (_StorageProviderCollateral *StorageProviderCollateralCaller) GetAvailableCollateral(opts *bind.CallOpts, _ownerId uint64) (*big.Int, error) {
	var out []interface{}
	err := _StorageProviderCollateral.contract.Call(opts, &out, "getAvailableCollateral", _ownerId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAvailableCollateral is a free data retrieval call binding the contract method 0xe460512e.
//
// Solidity: function getAvailableCollateral(uint64 _ownerId) view returns(uint256)
func (_StorageProviderCollateral *StorageProviderCollateralSession) GetAvailableCollateral(_ownerId uint64) (*big.Int, error) {
	return _StorageProviderCollateral.Contract.GetAvailableCollateral(&_StorageProviderCollateral.CallOpts, _ownerId)
}

// GetAvailableCollateral is a free data retrieval call binding the contract method 0xe460512e.
//
// Solidity: function getAvailableCollateral(uint64 _ownerId) view returns(uint256)
func (_StorageProviderCollateral *StorageProviderCollateralCallerSession) GetAvailableCollateral(_ownerId uint64) (*big.Int, error) {
	return _StorageProviderCollateral.Contract.GetAvailableCollateral(&_StorageProviderCollateral.CallOpts, _ownerId)
}

// GetCollateral is a free data retrieval call binding the contract method 0x4eaf14a6.
//
// Solidity: function getCollateral(uint64 _ownerId) view returns(uint256, uint256)
func (_StorageProviderCollateral *StorageProviderCollateralCaller) GetCollateral(opts *bind.CallOpts, _ownerId uint64) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _StorageProviderCollateral.contract.Call(opts, &out, "getCollateral", _ownerId)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetCollateral is a free data retrieval call binding the contract method 0x4eaf14a6.
//
// Solidity: function getCollateral(uint64 _ownerId) view returns(uint256, uint256)
func (_StorageProviderCollateral *StorageProviderCollateralSession) GetCollateral(_ownerId uint64) (*big.Int, *big.Int, error) {
	return _StorageProviderCollateral.Contract.GetCollateral(&_StorageProviderCollateral.CallOpts, _ownerId)
}

// GetCollateral is a free data retrieval call binding the contract method 0x4eaf14a6.
//
// Solidity: function getCollateral(uint64 _ownerId) view returns(uint256, uint256)
func (_StorageProviderCollateral *StorageProviderCollateralCallerSession) GetCollateral(_ownerId uint64) (*big.Int, *big.Int, error) {
	return _StorageProviderCollateral.Contract.GetCollateral(&_StorageProviderCollateral.CallOpts, _ownerId)
}

// GetCollateralRequirements is a free data retrieval call binding the contract method 0x717e9d26.
//
// Solidity: function getCollateralRequirements(uint64 _ownerId) view returns(uint256)
func (_StorageProviderCollateral *StorageProviderCollateralCaller) GetCollateralRequirements(opts *bind.CallOpts, _ownerId uint64) (*big.Int, error) {
	var out []interface{}
	err := _StorageProviderCollateral.contract.Call(opts, &out, "getCollateralRequirements", _ownerId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCollateralRequirements is a free data retrieval call binding the contract method 0x717e9d26.
//
// Solidity: function getCollateralRequirements(uint64 _ownerId) view returns(uint256)
func (_StorageProviderCollateral *StorageProviderCollateralSession) GetCollateralRequirements(_ownerId uint64) (*big.Int, error) {
	return _StorageProviderCollateral.Contract.GetCollateralRequirements(&_StorageProviderCollateral.CallOpts, _ownerId)
}

// GetCollateralRequirements is a free data retrieval call binding the contract method 0x717e9d26.
//
// Solidity: function getCollateralRequirements(uint64 _ownerId) view returns(uint256)
func (_StorageProviderCollateral *StorageProviderCollateralCallerSession) GetCollateralRequirements(_ownerId uint64) (*big.Int, error) {
	return _StorageProviderCollateral.Contract.GetCollateralRequirements(&_StorageProviderCollateral.CallOpts, _ownerId)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_StorageProviderCollateral *StorageProviderCollateralCaller) GetImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StorageProviderCollateral.contract.Call(opts, &out, "getImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_StorageProviderCollateral *StorageProviderCollateralSession) GetImplementation() (common.Address, error) {
	return _StorageProviderCollateral.Contract.GetImplementation(&_StorageProviderCollateral.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_StorageProviderCollateral *StorageProviderCollateralCallerSession) GetImplementation() (common.Address, error) {
	return _StorageProviderCollateral.Contract.GetImplementation(&_StorageProviderCollateral.CallOpts)
}

// GetLockedCollateral is a free data retrieval call binding the contract method 0x6ef2dafc.
//
// Solidity: function getLockedCollateral(uint64 _ownerId) view returns(uint256)
func (_StorageProviderCollateral *StorageProviderCollateralCaller) GetLockedCollateral(opts *bind.CallOpts, _ownerId uint64) (*big.Int, error) {
	var out []interface{}
	err := _StorageProviderCollateral.contract.Call(opts, &out, "getLockedCollateral", _ownerId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLockedCollateral is a free data retrieval call binding the contract method 0x6ef2dafc.
//
// Solidity: function getLockedCollateral(uint64 _ownerId) view returns(uint256)
func (_StorageProviderCollateral *StorageProviderCollateralSession) GetLockedCollateral(_ownerId uint64) (*big.Int, error) {
	return _StorageProviderCollateral.Contract.GetLockedCollateral(&_StorageProviderCollateral.CallOpts, _ownerId)
}

// GetLockedCollateral is a free data retrieval call binding the contract method 0x6ef2dafc.
//
// Solidity: function getLockedCollateral(uint64 _ownerId) view returns(uint256)
func (_StorageProviderCollateral *StorageProviderCollateralCallerSession) GetLockedCollateral(_ownerId uint64) (*big.Int, error) {
	return _StorageProviderCollateral.Contract.GetLockedCollateral(&_StorageProviderCollateral.CallOpts, _ownerId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StorageProviderCollateral *StorageProviderCollateralCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _StorageProviderCollateral.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StorageProviderCollateral *StorageProviderCollateralSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StorageProviderCollateral.Contract.GetRoleAdmin(&_StorageProviderCollateral.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StorageProviderCollateral *StorageProviderCollateralCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StorageProviderCollateral.Contract.GetRoleAdmin(&_StorageProviderCollateral.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StorageProviderCollateral *StorageProviderCollateralCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _StorageProviderCollateral.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StorageProviderCollateral *StorageProviderCollateralSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StorageProviderCollateral.Contract.HasRole(&_StorageProviderCollateral.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StorageProviderCollateral *StorageProviderCollateralCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StorageProviderCollateral.Contract.HasRole(&_StorageProviderCollateral.CallOpts, role, account)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_StorageProviderCollateral *StorageProviderCollateralCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StorageProviderCollateral.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_StorageProviderCollateral *StorageProviderCollateralSession) ProxiableUUID() ([32]byte, error) {
	return _StorageProviderCollateral.Contract.ProxiableUUID(&_StorageProviderCollateral.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_StorageProviderCollateral *StorageProviderCollateralCallerSession) ProxiableUUID() ([32]byte, error) {
	return _StorageProviderCollateral.Contract.ProxiableUUID(&_StorageProviderCollateral.CallOpts)
}

// Slashings is a free data retrieval call binding the contract method 0x4184acd5.
//
// Solidity: function slashings(uint64 ) view returns(uint256)
func (_StorageProviderCollateral *StorageProviderCollateralCaller) Slashings(opts *bind.CallOpts, arg0 uint64) (*big.Int, error) {
	var out []interface{}
	err := _StorageProviderCollateral.contract.Call(opts, &out, "slashings", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Slashings is a free data retrieval call binding the contract method 0x4184acd5.
//
// Solidity: function slashings(uint64 ) view returns(uint256)
func (_StorageProviderCollateral *StorageProviderCollateralSession) Slashings(arg0 uint64) (*big.Int, error) {
	return _StorageProviderCollateral.Contract.Slashings(&_StorageProviderCollateral.CallOpts, arg0)
}

// Slashings is a free data retrieval call binding the contract method 0x4184acd5.
//
// Solidity: function slashings(uint64 ) view returns(uint256)
func (_StorageProviderCollateral *StorageProviderCollateralCallerSession) Slashings(arg0 uint64) (*big.Int, error) {
	return _StorageProviderCollateral.Contract.Slashings(&_StorageProviderCollateral.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StorageProviderCollateral *StorageProviderCollateralCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _StorageProviderCollateral.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StorageProviderCollateral *StorageProviderCollateralSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StorageProviderCollateral.Contract.SupportsInterface(&_StorageProviderCollateral.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StorageProviderCollateral *StorageProviderCollateralCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StorageProviderCollateral.Contract.SupportsInterface(&_StorageProviderCollateral.CallOpts, interfaceId)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_StorageProviderCollateral *StorageProviderCollateralCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StorageProviderCollateral.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_StorageProviderCollateral *StorageProviderCollateralSession) Version() (string, error) {
	return _StorageProviderCollateral.Contract.Version(&_StorageProviderCollateral.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_StorageProviderCollateral *StorageProviderCollateralCallerSession) Version() (string, error) {
	return _StorageProviderCollateral.Contract.Version(&_StorageProviderCollateral.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_StorageProviderCollateral *StorageProviderCollateralTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageProviderCollateral.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_StorageProviderCollateral *StorageProviderCollateralSession) Deposit() (*types.Transaction, error) {
	return _StorageProviderCollateral.Contract.Deposit(&_StorageProviderCollateral.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_StorageProviderCollateral *StorageProviderCollateralTransactorSession) Deposit() (*types.Transaction, error) {
	return _StorageProviderCollateral.Contract.Deposit(&_StorageProviderCollateral.TransactOpts)
}

// Fit is a paid mutator transaction binding the contract method 0xbea914e5.
//
// Solidity: function fit(uint64 _ownerId) returns()
func (_StorageProviderCollateral *StorageProviderCollateralTransactor) Fit(opts *bind.TransactOpts, _ownerId uint64) (*types.Transaction, error) {
	return _StorageProviderCollateral.contract.Transact(opts, "fit", _ownerId)
}

// Fit is a paid mutator transaction binding the contract method 0xbea914e5.
//
// Solidity: function fit(uint64 _ownerId) returns()
func (_StorageProviderCollateral *StorageProviderCollateralSession) Fit(_ownerId uint64) (*types.Transaction, error) {
	return _StorageProviderCollateral.Contract.Fit(&_StorageProviderCollateral.TransactOpts, _ownerId)
}

// Fit is a paid mutator transaction binding the contract method 0xbea914e5.
//
// Solidity: function fit(uint64 _ownerId) returns()
func (_StorageProviderCollateral *StorageProviderCollateralTransactorSession) Fit(_ownerId uint64) (*types.Transaction, error) {
	return _StorageProviderCollateral.Contract.Fit(&_StorageProviderCollateral.TransactOpts, _ownerId)
}

// GetDebt is a paid mutator transaction binding the contract method 0xc8025dda.
//
// Solidity: function getDebt(uint64 _ownerId) returns(uint256)
func (_StorageProviderCollateral *StorageProviderCollateralTransactor) GetDebt(opts *bind.TransactOpts, _ownerId uint64) (*types.Transaction, error) {
	return _StorageProviderCollateral.contract.Transact(opts, "getDebt", _ownerId)
}

// GetDebt is a paid mutator transaction binding the contract method 0xc8025dda.
//
// Solidity: function getDebt(uint64 _ownerId) returns(uint256)
func (_StorageProviderCollateral *StorageProviderCollateralSession) GetDebt(_ownerId uint64) (*types.Transaction, error) {
	return _StorageProviderCollateral.Contract.GetDebt(&_StorageProviderCollateral.TransactOpts, _ownerId)
}

// GetDebt is a paid mutator transaction binding the contract method 0xc8025dda.
//
// Solidity: function getDebt(uint64 _ownerId) returns(uint256)
func (_StorageProviderCollateral *StorageProviderCollateralTransactorSession) GetDebt(_ownerId uint64) (*types.Transaction, error) {
	return _StorageProviderCollateral.Contract.GetDebt(&_StorageProviderCollateral.TransactOpts, _ownerId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StorageProviderCollateral *StorageProviderCollateralTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StorageProviderCollateral.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StorageProviderCollateral *StorageProviderCollateralSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StorageProviderCollateral.Contract.GrantRole(&_StorageProviderCollateral.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StorageProviderCollateral *StorageProviderCollateralTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StorageProviderCollateral.Contract.GrantRole(&_StorageProviderCollateral.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address _wFIL, address _resolver, uint256 _baseRequirements) returns()
func (_StorageProviderCollateral *StorageProviderCollateralTransactor) Initialize(opts *bind.TransactOpts, _wFIL common.Address, _resolver common.Address, _baseRequirements *big.Int) (*types.Transaction, error) {
	return _StorageProviderCollateral.contract.Transact(opts, "initialize", _wFIL, _resolver, _baseRequirements)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address _wFIL, address _resolver, uint256 _baseRequirements) returns()
func (_StorageProviderCollateral *StorageProviderCollateralSession) Initialize(_wFIL common.Address, _resolver common.Address, _baseRequirements *big.Int) (*types.Transaction, error) {
	return _StorageProviderCollateral.Contract.Initialize(&_StorageProviderCollateral.TransactOpts, _wFIL, _resolver, _baseRequirements)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address _wFIL, address _resolver, uint256 _baseRequirements) returns()
func (_StorageProviderCollateral *StorageProviderCollateralTransactorSession) Initialize(_wFIL common.Address, _resolver common.Address, _baseRequirements *big.Int) (*types.Transaction, error) {
	return _StorageProviderCollateral.Contract.Initialize(&_StorageProviderCollateral.TransactOpts, _wFIL, _resolver, _baseRequirements)
}

// Lock is a paid mutator transaction binding the contract method 0x241c054a.
//
// Solidity: function lock(uint64 _ownerId, uint64 _minerId, uint256 _allocated) returns()
func (_StorageProviderCollateral *StorageProviderCollateralTransactor) Lock(opts *bind.TransactOpts, _ownerId uint64, _minerId uint64, _allocated *big.Int) (*types.Transaction, error) {
	return _StorageProviderCollateral.contract.Transact(opts, "lock", _ownerId, _minerId, _allocated)
}

// Lock is a paid mutator transaction binding the contract method 0x241c054a.
//
// Solidity: function lock(uint64 _ownerId, uint64 _minerId, uint256 _allocated) returns()
func (_StorageProviderCollateral *StorageProviderCollateralSession) Lock(_ownerId uint64, _minerId uint64, _allocated *big.Int) (*types.Transaction, error) {
	return _StorageProviderCollateral.Contract.Lock(&_StorageProviderCollateral.TransactOpts, _ownerId, _minerId, _allocated)
}

// Lock is a paid mutator transaction binding the contract method 0x241c054a.
//
// Solidity: function lock(uint64 _ownerId, uint64 _minerId, uint256 _allocated) returns()
func (_StorageProviderCollateral *StorageProviderCollateralTransactorSession) Lock(_ownerId uint64, _minerId uint64, _allocated *big.Int) (*types.Transaction, error) {
	return _StorageProviderCollateral.Contract.Lock(&_StorageProviderCollateral.TransactOpts, _ownerId, _minerId, _allocated)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StorageProviderCollateral *StorageProviderCollateralTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StorageProviderCollateral.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StorageProviderCollateral *StorageProviderCollateralSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StorageProviderCollateral.Contract.RenounceRole(&_StorageProviderCollateral.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StorageProviderCollateral *StorageProviderCollateralTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StorageProviderCollateral.Contract.RenounceRole(&_StorageProviderCollateral.TransactOpts, role, account)
}

// ReportRecovery is a paid mutator transaction binding the contract method 0x1917a6a7.
//
// Solidity: function reportRecovery(uint64 _ownerId) returns()
func (_StorageProviderCollateral *StorageProviderCollateralTransactor) ReportRecovery(opts *bind.TransactOpts, _ownerId uint64) (*types.Transaction, error) {
	return _StorageProviderCollateral.contract.Transact(opts, "reportRecovery", _ownerId)
}

// ReportRecovery is a paid mutator transaction binding the contract method 0x1917a6a7.
//
// Solidity: function reportRecovery(uint64 _ownerId) returns()
func (_StorageProviderCollateral *StorageProviderCollateralSession) ReportRecovery(_ownerId uint64) (*types.Transaction, error) {
	return _StorageProviderCollateral.Contract.ReportRecovery(&_StorageProviderCollateral.TransactOpts, _ownerId)
}

// ReportRecovery is a paid mutator transaction binding the contract method 0x1917a6a7.
//
// Solidity: function reportRecovery(uint64 _ownerId) returns()
func (_StorageProviderCollateral *StorageProviderCollateralTransactorSession) ReportRecovery(_ownerId uint64) (*types.Transaction, error) {
	return _StorageProviderCollateral.Contract.ReportRecovery(&_StorageProviderCollateral.TransactOpts, _ownerId)
}

// ReportSlashing is a paid mutator transaction binding the contract method 0xeeb77ba8.
//
// Solidity: function reportSlashing(uint64 _ownerId, uint256 _slashingAmt) returns()
func (_StorageProviderCollateral *StorageProviderCollateralTransactor) ReportSlashing(opts *bind.TransactOpts, _ownerId uint64, _slashingAmt *big.Int) (*types.Transaction, error) {
	return _StorageProviderCollateral.contract.Transact(opts, "reportSlashing", _ownerId, _slashingAmt)
}

// ReportSlashing is a paid mutator transaction binding the contract method 0xeeb77ba8.
//
// Solidity: function reportSlashing(uint64 _ownerId, uint256 _slashingAmt) returns()
func (_StorageProviderCollateral *StorageProviderCollateralSession) ReportSlashing(_ownerId uint64, _slashingAmt *big.Int) (*types.Transaction, error) {
	return _StorageProviderCollateral.Contract.ReportSlashing(&_StorageProviderCollateral.TransactOpts, _ownerId, _slashingAmt)
}

// ReportSlashing is a paid mutator transaction binding the contract method 0xeeb77ba8.
//
// Solidity: function reportSlashing(uint64 _ownerId, uint256 _slashingAmt) returns()
func (_StorageProviderCollateral *StorageProviderCollateralTransactorSession) ReportSlashing(_ownerId uint64, _slashingAmt *big.Int) (*types.Transaction, error) {
	return _StorageProviderCollateral.Contract.ReportSlashing(&_StorageProviderCollateral.TransactOpts, _ownerId, _slashingAmt)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StorageProviderCollateral *StorageProviderCollateralTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StorageProviderCollateral.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StorageProviderCollateral *StorageProviderCollateralSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StorageProviderCollateral.Contract.RevokeRole(&_StorageProviderCollateral.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StorageProviderCollateral *StorageProviderCollateralTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StorageProviderCollateral.Contract.RevokeRole(&_StorageProviderCollateral.TransactOpts, role, account)
}

// UpdateBaseCollateralRequirements is a paid mutator transaction binding the contract method 0xc7d3ebf8.
//
// Solidity: function updateBaseCollateralRequirements(uint256 requirements) returns()
func (_StorageProviderCollateral *StorageProviderCollateralTransactor) UpdateBaseCollateralRequirements(opts *bind.TransactOpts, requirements *big.Int) (*types.Transaction, error) {
	return _StorageProviderCollateral.contract.Transact(opts, "updateBaseCollateralRequirements", requirements)
}

// UpdateBaseCollateralRequirements is a paid mutator transaction binding the contract method 0xc7d3ebf8.
//
// Solidity: function updateBaseCollateralRequirements(uint256 requirements) returns()
func (_StorageProviderCollateral *StorageProviderCollateralSession) UpdateBaseCollateralRequirements(requirements *big.Int) (*types.Transaction, error) {
	return _StorageProviderCollateral.Contract.UpdateBaseCollateralRequirements(&_StorageProviderCollateral.TransactOpts, requirements)
}

// UpdateBaseCollateralRequirements is a paid mutator transaction binding the contract method 0xc7d3ebf8.
//
// Solidity: function updateBaseCollateralRequirements(uint256 requirements) returns()
func (_StorageProviderCollateral *StorageProviderCollateralTransactorSession) UpdateBaseCollateralRequirements(requirements *big.Int) (*types.Transaction, error) {
	return _StorageProviderCollateral.Contract.UpdateBaseCollateralRequirements(&_StorageProviderCollateral.TransactOpts, requirements)
}

// UpdateCollateralRequirements is a paid mutator transaction binding the contract method 0x335f1839.
//
// Solidity: function updateCollateralRequirements(uint64 _ownerId, uint256 requirements) returns()
func (_StorageProviderCollateral *StorageProviderCollateralTransactor) UpdateCollateralRequirements(opts *bind.TransactOpts, _ownerId uint64, requirements *big.Int) (*types.Transaction, error) {
	return _StorageProviderCollateral.contract.Transact(opts, "updateCollateralRequirements", _ownerId, requirements)
}

// UpdateCollateralRequirements is a paid mutator transaction binding the contract method 0x335f1839.
//
// Solidity: function updateCollateralRequirements(uint64 _ownerId, uint256 requirements) returns()
func (_StorageProviderCollateral *StorageProviderCollateralSession) UpdateCollateralRequirements(_ownerId uint64, requirements *big.Int) (*types.Transaction, error) {
	return _StorageProviderCollateral.Contract.UpdateCollateralRequirements(&_StorageProviderCollateral.TransactOpts, _ownerId, requirements)
}

// UpdateCollateralRequirements is a paid mutator transaction binding the contract method 0x335f1839.
//
// Solidity: function updateCollateralRequirements(uint64 _ownerId, uint256 requirements) returns()
func (_StorageProviderCollateral *StorageProviderCollateralTransactorSession) UpdateCollateralRequirements(_ownerId uint64, requirements *big.Int) (*types.Transaction, error) {
	return _StorageProviderCollateral.Contract.UpdateCollateralRequirements(&_StorageProviderCollateral.TransactOpts, _ownerId, requirements)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_StorageProviderCollateral *StorageProviderCollateralTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _StorageProviderCollateral.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_StorageProviderCollateral *StorageProviderCollateralSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _StorageProviderCollateral.Contract.UpgradeTo(&_StorageProviderCollateral.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_StorageProviderCollateral *StorageProviderCollateralTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _StorageProviderCollateral.Contract.UpgradeTo(&_StorageProviderCollateral.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_StorageProviderCollateral *StorageProviderCollateralTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _StorageProviderCollateral.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_StorageProviderCollateral *StorageProviderCollateralSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _StorageProviderCollateral.Contract.UpgradeToAndCall(&_StorageProviderCollateral.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_StorageProviderCollateral *StorageProviderCollateralTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _StorageProviderCollateral.Contract.UpgradeToAndCall(&_StorageProviderCollateral.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_StorageProviderCollateral *StorageProviderCollateralTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _StorageProviderCollateral.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_StorageProviderCollateral *StorageProviderCollateralSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _StorageProviderCollateral.Contract.Withdraw(&_StorageProviderCollateral.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_StorageProviderCollateral *StorageProviderCollateralTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _StorageProviderCollateral.Contract.Withdraw(&_StorageProviderCollateral.TransactOpts, _amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_StorageProviderCollateral *StorageProviderCollateralTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _StorageProviderCollateral.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_StorageProviderCollateral *StorageProviderCollateralSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _StorageProviderCollateral.Contract.Fallback(&_StorageProviderCollateral.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_StorageProviderCollateral *StorageProviderCollateralTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _StorageProviderCollateral.Contract.Fallback(&_StorageProviderCollateral.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StorageProviderCollateral *StorageProviderCollateralTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageProviderCollateral.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StorageProviderCollateral *StorageProviderCollateralSession) Receive() (*types.Transaction, error) {
	return _StorageProviderCollateral.Contract.Receive(&_StorageProviderCollateral.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StorageProviderCollateral *StorageProviderCollateralTransactorSession) Receive() (*types.Transaction, error) {
	return _StorageProviderCollateral.Contract.Receive(&_StorageProviderCollateral.TransactOpts)
}

// StorageProviderCollateralAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the StorageProviderCollateral contract.
type StorageProviderCollateralAdminChangedIterator struct {
	Event *StorageProviderCollateralAdminChanged // Event containing the contract specifics and raw log

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
func (it *StorageProviderCollateralAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProviderCollateralAdminChanged)
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
		it.Event = new(StorageProviderCollateralAdminChanged)
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
func (it *StorageProviderCollateralAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProviderCollateralAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProviderCollateralAdminChanged represents a AdminChanged event raised by the StorageProviderCollateral contract.
type StorageProviderCollateralAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*StorageProviderCollateralAdminChangedIterator, error) {

	logs, sub, err := _StorageProviderCollateral.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &StorageProviderCollateralAdminChangedIterator{contract: _StorageProviderCollateral.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *StorageProviderCollateralAdminChanged) (event.Subscription, error) {

	logs, sub, err := _StorageProviderCollateral.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProviderCollateralAdminChanged)
				if err := _StorageProviderCollateral.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) ParseAdminChanged(log types.Log) (*StorageProviderCollateralAdminChanged, error) {
	event := new(StorageProviderCollateralAdminChanged)
	if err := _StorageProviderCollateral.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageProviderCollateralBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the StorageProviderCollateral contract.
type StorageProviderCollateralBeaconUpgradedIterator struct {
	Event *StorageProviderCollateralBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *StorageProviderCollateralBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProviderCollateralBeaconUpgraded)
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
		it.Event = new(StorageProviderCollateralBeaconUpgraded)
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
func (it *StorageProviderCollateralBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProviderCollateralBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProviderCollateralBeaconUpgraded represents a BeaconUpgraded event raised by the StorageProviderCollateral contract.
type StorageProviderCollateralBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*StorageProviderCollateralBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _StorageProviderCollateral.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &StorageProviderCollateralBeaconUpgradedIterator{contract: _StorageProviderCollateral.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *StorageProviderCollateralBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _StorageProviderCollateral.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProviderCollateralBeaconUpgraded)
				if err := _StorageProviderCollateral.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) ParseBeaconUpgraded(log types.Log) (*StorageProviderCollateralBeaconUpgraded, error) {
	event := new(StorageProviderCollateralBeaconUpgraded)
	if err := _StorageProviderCollateral.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageProviderCollateralInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StorageProviderCollateral contract.
type StorageProviderCollateralInitializedIterator struct {
	Event *StorageProviderCollateralInitialized // Event containing the contract specifics and raw log

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
func (it *StorageProviderCollateralInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProviderCollateralInitialized)
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
		it.Event = new(StorageProviderCollateralInitialized)
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
func (it *StorageProviderCollateralInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProviderCollateralInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProviderCollateralInitialized represents a Initialized event raised by the StorageProviderCollateral contract.
type StorageProviderCollateralInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) FilterInitialized(opts *bind.FilterOpts) (*StorageProviderCollateralInitializedIterator, error) {

	logs, sub, err := _StorageProviderCollateral.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StorageProviderCollateralInitializedIterator{contract: _StorageProviderCollateral.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StorageProviderCollateralInitialized) (event.Subscription, error) {

	logs, sub, err := _StorageProviderCollateral.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProviderCollateralInitialized)
				if err := _StorageProviderCollateral.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) ParseInitialized(log types.Log) (*StorageProviderCollateralInitialized, error) {
	event := new(StorageProviderCollateralInitialized)
	if err := _StorageProviderCollateral.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageProviderCollateralReportRecoveryIterator is returned from FilterReportRecovery and is used to iterate over the raw logs and unpacked data for ReportRecovery events raised by the StorageProviderCollateral contract.
type StorageProviderCollateralReportRecoveryIterator struct {
	Event *StorageProviderCollateralReportRecovery // Event containing the contract specifics and raw log

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
func (it *StorageProviderCollateralReportRecoveryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProviderCollateralReportRecovery)
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
		it.Event = new(StorageProviderCollateralReportRecovery)
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
func (it *StorageProviderCollateralReportRecoveryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProviderCollateralReportRecoveryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProviderCollateralReportRecovery represents a ReportRecovery event raised by the StorageProviderCollateral contract.
type StorageProviderCollateralReportRecovery struct {
	OwnerId uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReportRecovery is a free log retrieval operation binding the contract event 0x4ae1ac1eeb2bb87a829856665535c45f6381998822857dde5d4573a1b4a4a9c1.
//
// Solidity: event ReportRecovery(uint64 ownerId)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) FilterReportRecovery(opts *bind.FilterOpts) (*StorageProviderCollateralReportRecoveryIterator, error) {

	logs, sub, err := _StorageProviderCollateral.contract.FilterLogs(opts, "ReportRecovery")
	if err != nil {
		return nil, err
	}
	return &StorageProviderCollateralReportRecoveryIterator{contract: _StorageProviderCollateral.contract, event: "ReportRecovery", logs: logs, sub: sub}, nil
}

// WatchReportRecovery is a free log subscription operation binding the contract event 0x4ae1ac1eeb2bb87a829856665535c45f6381998822857dde5d4573a1b4a4a9c1.
//
// Solidity: event ReportRecovery(uint64 ownerId)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) WatchReportRecovery(opts *bind.WatchOpts, sink chan<- *StorageProviderCollateralReportRecovery) (event.Subscription, error) {

	logs, sub, err := _StorageProviderCollateral.contract.WatchLogs(opts, "ReportRecovery")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProviderCollateralReportRecovery)
				if err := _StorageProviderCollateral.contract.UnpackLog(event, "ReportRecovery", log); err != nil {
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

// ParseReportRecovery is a log parse operation binding the contract event 0x4ae1ac1eeb2bb87a829856665535c45f6381998822857dde5d4573a1b4a4a9c1.
//
// Solidity: event ReportRecovery(uint64 ownerId)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) ParseReportRecovery(log types.Log) (*StorageProviderCollateralReportRecovery, error) {
	event := new(StorageProviderCollateralReportRecovery)
	if err := _StorageProviderCollateral.contract.UnpackLog(event, "ReportRecovery", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageProviderCollateralReportSlashingIterator is returned from FilterReportSlashing and is used to iterate over the raw logs and unpacked data for ReportSlashing events raised by the StorageProviderCollateral contract.
type StorageProviderCollateralReportSlashingIterator struct {
	Event *StorageProviderCollateralReportSlashing // Event containing the contract specifics and raw log

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
func (it *StorageProviderCollateralReportSlashingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProviderCollateralReportSlashing)
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
		it.Event = new(StorageProviderCollateralReportSlashing)
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
func (it *StorageProviderCollateralReportSlashingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProviderCollateralReportSlashingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProviderCollateralReportSlashing represents a ReportSlashing event raised by the StorageProviderCollateral contract.
type StorageProviderCollateralReportSlashing struct {
	OwnerId        uint64
	SlashingAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterReportSlashing is a free log retrieval operation binding the contract event 0xe920a3bd2d05623539eb004a3d191cb43a32f41b4dce001e9c4d9ee2e7a443e3.
//
// Solidity: event ReportSlashing(uint64 ownerId, uint256 slashingAmount)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) FilterReportSlashing(opts *bind.FilterOpts) (*StorageProviderCollateralReportSlashingIterator, error) {

	logs, sub, err := _StorageProviderCollateral.contract.FilterLogs(opts, "ReportSlashing")
	if err != nil {
		return nil, err
	}
	return &StorageProviderCollateralReportSlashingIterator{contract: _StorageProviderCollateral.contract, event: "ReportSlashing", logs: logs, sub: sub}, nil
}

// WatchReportSlashing is a free log subscription operation binding the contract event 0xe920a3bd2d05623539eb004a3d191cb43a32f41b4dce001e9c4d9ee2e7a443e3.
//
// Solidity: event ReportSlashing(uint64 ownerId, uint256 slashingAmount)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) WatchReportSlashing(opts *bind.WatchOpts, sink chan<- *StorageProviderCollateralReportSlashing) (event.Subscription, error) {

	logs, sub, err := _StorageProviderCollateral.contract.WatchLogs(opts, "ReportSlashing")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProviderCollateralReportSlashing)
				if err := _StorageProviderCollateral.contract.UnpackLog(event, "ReportSlashing", log); err != nil {
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

// ParseReportSlashing is a log parse operation binding the contract event 0xe920a3bd2d05623539eb004a3d191cb43a32f41b4dce001e9c4d9ee2e7a443e3.
//
// Solidity: event ReportSlashing(uint64 ownerId, uint256 slashingAmount)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) ParseReportSlashing(log types.Log) (*StorageProviderCollateralReportSlashing, error) {
	event := new(StorageProviderCollateralReportSlashing)
	if err := _StorageProviderCollateral.contract.UnpackLog(event, "ReportSlashing", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageProviderCollateralRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the StorageProviderCollateral contract.
type StorageProviderCollateralRoleAdminChangedIterator struct {
	Event *StorageProviderCollateralRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *StorageProviderCollateralRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProviderCollateralRoleAdminChanged)
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
		it.Event = new(StorageProviderCollateralRoleAdminChanged)
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
func (it *StorageProviderCollateralRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProviderCollateralRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProviderCollateralRoleAdminChanged represents a RoleAdminChanged event raised by the StorageProviderCollateral contract.
type StorageProviderCollateralRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*StorageProviderCollateralRoleAdminChangedIterator, error) {

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

	logs, sub, err := _StorageProviderCollateral.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &StorageProviderCollateralRoleAdminChangedIterator{contract: _StorageProviderCollateral.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *StorageProviderCollateralRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _StorageProviderCollateral.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProviderCollateralRoleAdminChanged)
				if err := _StorageProviderCollateral.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) ParseRoleAdminChanged(log types.Log) (*StorageProviderCollateralRoleAdminChanged, error) {
	event := new(StorageProviderCollateralRoleAdminChanged)
	if err := _StorageProviderCollateral.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageProviderCollateralRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the StorageProviderCollateral contract.
type StorageProviderCollateralRoleGrantedIterator struct {
	Event *StorageProviderCollateralRoleGranted // Event containing the contract specifics and raw log

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
func (it *StorageProviderCollateralRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProviderCollateralRoleGranted)
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
		it.Event = new(StorageProviderCollateralRoleGranted)
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
func (it *StorageProviderCollateralRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProviderCollateralRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProviderCollateralRoleGranted represents a RoleGranted event raised by the StorageProviderCollateral contract.
type StorageProviderCollateralRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StorageProviderCollateralRoleGrantedIterator, error) {

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

	logs, sub, err := _StorageProviderCollateral.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StorageProviderCollateralRoleGrantedIterator{contract: _StorageProviderCollateral.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *StorageProviderCollateralRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _StorageProviderCollateral.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProviderCollateralRoleGranted)
				if err := _StorageProviderCollateral.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) ParseRoleGranted(log types.Log) (*StorageProviderCollateralRoleGranted, error) {
	event := new(StorageProviderCollateralRoleGranted)
	if err := _StorageProviderCollateral.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageProviderCollateralRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the StorageProviderCollateral contract.
type StorageProviderCollateralRoleRevokedIterator struct {
	Event *StorageProviderCollateralRoleRevoked // Event containing the contract specifics and raw log

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
func (it *StorageProviderCollateralRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProviderCollateralRoleRevoked)
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
		it.Event = new(StorageProviderCollateralRoleRevoked)
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
func (it *StorageProviderCollateralRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProviderCollateralRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProviderCollateralRoleRevoked represents a RoleRevoked event raised by the StorageProviderCollateral contract.
type StorageProviderCollateralRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StorageProviderCollateralRoleRevokedIterator, error) {

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

	logs, sub, err := _StorageProviderCollateral.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StorageProviderCollateralRoleRevokedIterator{contract: _StorageProviderCollateral.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *StorageProviderCollateralRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _StorageProviderCollateral.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProviderCollateralRoleRevoked)
				if err := _StorageProviderCollateral.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) ParseRoleRevoked(log types.Log) (*StorageProviderCollateralRoleRevoked, error) {
	event := new(StorageProviderCollateralRoleRevoked)
	if err := _StorageProviderCollateral.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageProviderCollateralSetRegistryAddressIterator is returned from FilterSetRegistryAddress and is used to iterate over the raw logs and unpacked data for SetRegistryAddress events raised by the StorageProviderCollateral contract.
type StorageProviderCollateralSetRegistryAddressIterator struct {
	Event *StorageProviderCollateralSetRegistryAddress // Event containing the contract specifics and raw log

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
func (it *StorageProviderCollateralSetRegistryAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProviderCollateralSetRegistryAddress)
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
		it.Event = new(StorageProviderCollateralSetRegistryAddress)
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
func (it *StorageProviderCollateralSetRegistryAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProviderCollateralSetRegistryAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProviderCollateralSetRegistryAddress represents a SetRegistryAddress event raised by the StorageProviderCollateral contract.
type StorageProviderCollateralSetRegistryAddress struct {
	Registry common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetRegistryAddress is a free log retrieval operation binding the contract event 0x3e06c035fbf29a3956dd07cbab527255f029751c95add41cc26990b28ed7179b.
//
// Solidity: event SetRegistryAddress(address registry)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) FilterSetRegistryAddress(opts *bind.FilterOpts) (*StorageProviderCollateralSetRegistryAddressIterator, error) {

	logs, sub, err := _StorageProviderCollateral.contract.FilterLogs(opts, "SetRegistryAddress")
	if err != nil {
		return nil, err
	}
	return &StorageProviderCollateralSetRegistryAddressIterator{contract: _StorageProviderCollateral.contract, event: "SetRegistryAddress", logs: logs, sub: sub}, nil
}

// WatchSetRegistryAddress is a free log subscription operation binding the contract event 0x3e06c035fbf29a3956dd07cbab527255f029751c95add41cc26990b28ed7179b.
//
// Solidity: event SetRegistryAddress(address registry)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) WatchSetRegistryAddress(opts *bind.WatchOpts, sink chan<- *StorageProviderCollateralSetRegistryAddress) (event.Subscription, error) {

	logs, sub, err := _StorageProviderCollateral.contract.WatchLogs(opts, "SetRegistryAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProviderCollateralSetRegistryAddress)
				if err := _StorageProviderCollateral.contract.UnpackLog(event, "SetRegistryAddress", log); err != nil {
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

// ParseSetRegistryAddress is a log parse operation binding the contract event 0x3e06c035fbf29a3956dd07cbab527255f029751c95add41cc26990b28ed7179b.
//
// Solidity: event SetRegistryAddress(address registry)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) ParseSetRegistryAddress(log types.Log) (*StorageProviderCollateralSetRegistryAddress, error) {
	event := new(StorageProviderCollateralSetRegistryAddress)
	if err := _StorageProviderCollateral.contract.UnpackLog(event, "SetRegistryAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageProviderCollateralStorageProviderCollateralDepositIterator is returned from FilterStorageProviderCollateralDeposit and is used to iterate over the raw logs and unpacked data for StorageProviderCollateralDeposit events raised by the StorageProviderCollateral contract.
type StorageProviderCollateralStorageProviderCollateralDepositIterator struct {
	Event *StorageProviderCollateralStorageProviderCollateralDeposit // Event containing the contract specifics and raw log

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
func (it *StorageProviderCollateralStorageProviderCollateralDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProviderCollateralStorageProviderCollateralDeposit)
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
		it.Event = new(StorageProviderCollateralStorageProviderCollateralDeposit)
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
func (it *StorageProviderCollateralStorageProviderCollateralDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProviderCollateralStorageProviderCollateralDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProviderCollateralStorageProviderCollateralDeposit represents a StorageProviderCollateralDeposit event raised by the StorageProviderCollateral contract.
type StorageProviderCollateralStorageProviderCollateralDeposit struct {
	OwnerId uint64
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStorageProviderCollateralDeposit is a free log retrieval operation binding the contract event 0x73dee9c602a838e7139b8ea81aab0fa49d1da28733bb0481f9ea4b512bdfda7d.
//
// Solidity: event StorageProviderCollateralDeposit(uint64 _ownerId, uint256 amount)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) FilterStorageProviderCollateralDeposit(opts *bind.FilterOpts) (*StorageProviderCollateralStorageProviderCollateralDepositIterator, error) {

	logs, sub, err := _StorageProviderCollateral.contract.FilterLogs(opts, "StorageProviderCollateralDeposit")
	if err != nil {
		return nil, err
	}
	return &StorageProviderCollateralStorageProviderCollateralDepositIterator{contract: _StorageProviderCollateral.contract, event: "StorageProviderCollateralDeposit", logs: logs, sub: sub}, nil
}

// WatchStorageProviderCollateralDeposit is a free log subscription operation binding the contract event 0x73dee9c602a838e7139b8ea81aab0fa49d1da28733bb0481f9ea4b512bdfda7d.
//
// Solidity: event StorageProviderCollateralDeposit(uint64 _ownerId, uint256 amount)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) WatchStorageProviderCollateralDeposit(opts *bind.WatchOpts, sink chan<- *StorageProviderCollateralStorageProviderCollateralDeposit) (event.Subscription, error) {

	logs, sub, err := _StorageProviderCollateral.contract.WatchLogs(opts, "StorageProviderCollateralDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProviderCollateralStorageProviderCollateralDeposit)
				if err := _StorageProviderCollateral.contract.UnpackLog(event, "StorageProviderCollateralDeposit", log); err != nil {
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

// ParseStorageProviderCollateralDeposit is a log parse operation binding the contract event 0x73dee9c602a838e7139b8ea81aab0fa49d1da28733bb0481f9ea4b512bdfda7d.
//
// Solidity: event StorageProviderCollateralDeposit(uint64 _ownerId, uint256 amount)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) ParseStorageProviderCollateralDeposit(log types.Log) (*StorageProviderCollateralStorageProviderCollateralDeposit, error) {
	event := new(StorageProviderCollateralStorageProviderCollateralDeposit)
	if err := _StorageProviderCollateral.contract.UnpackLog(event, "StorageProviderCollateralDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageProviderCollateralStorageProviderCollateralRebalanceIterator is returned from FilterStorageProviderCollateralRebalance and is used to iterate over the raw logs and unpacked data for StorageProviderCollateralRebalance events raised by the StorageProviderCollateral contract.
type StorageProviderCollateralStorageProviderCollateralRebalanceIterator struct {
	Event *StorageProviderCollateralStorageProviderCollateralRebalance // Event containing the contract specifics and raw log

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
func (it *StorageProviderCollateralStorageProviderCollateralRebalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProviderCollateralStorageProviderCollateralRebalance)
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
		it.Event = new(StorageProviderCollateralStorageProviderCollateralRebalance)
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
func (it *StorageProviderCollateralStorageProviderCollateralRebalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProviderCollateralStorageProviderCollateralRebalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProviderCollateralStorageProviderCollateralRebalance represents a StorageProviderCollateralRebalance event raised by the StorageProviderCollateral contract.
type StorageProviderCollateralStorageProviderCollateralRebalance struct {
	OwnerId             uint64
	LockedCollateral    *big.Int
	AvailableCollateral *big.Int
	IsUnlock            bool
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterStorageProviderCollateralRebalance is a free log retrieval operation binding the contract event 0x772d8d1c74df48b2a983c30ed06fb65bb4819c7e61c7d97766a825bbee5a596f.
//
// Solidity: event StorageProviderCollateralRebalance(uint64 _ownerId, uint256 lockedCollateral, uint256 availableCollateral, bool isUnlock)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) FilterStorageProviderCollateralRebalance(opts *bind.FilterOpts) (*StorageProviderCollateralStorageProviderCollateralRebalanceIterator, error) {

	logs, sub, err := _StorageProviderCollateral.contract.FilterLogs(opts, "StorageProviderCollateralRebalance")
	if err != nil {
		return nil, err
	}
	return &StorageProviderCollateralStorageProviderCollateralRebalanceIterator{contract: _StorageProviderCollateral.contract, event: "StorageProviderCollateralRebalance", logs: logs, sub: sub}, nil
}

// WatchStorageProviderCollateralRebalance is a free log subscription operation binding the contract event 0x772d8d1c74df48b2a983c30ed06fb65bb4819c7e61c7d97766a825bbee5a596f.
//
// Solidity: event StorageProviderCollateralRebalance(uint64 _ownerId, uint256 lockedCollateral, uint256 availableCollateral, bool isUnlock)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) WatchStorageProviderCollateralRebalance(opts *bind.WatchOpts, sink chan<- *StorageProviderCollateralStorageProviderCollateralRebalance) (event.Subscription, error) {

	logs, sub, err := _StorageProviderCollateral.contract.WatchLogs(opts, "StorageProviderCollateralRebalance")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProviderCollateralStorageProviderCollateralRebalance)
				if err := _StorageProviderCollateral.contract.UnpackLog(event, "StorageProviderCollateralRebalance", log); err != nil {
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

// ParseStorageProviderCollateralRebalance is a log parse operation binding the contract event 0x772d8d1c74df48b2a983c30ed06fb65bb4819c7e61c7d97766a825bbee5a596f.
//
// Solidity: event StorageProviderCollateralRebalance(uint64 _ownerId, uint256 lockedCollateral, uint256 availableCollateral, bool isUnlock)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) ParseStorageProviderCollateralRebalance(log types.Log) (*StorageProviderCollateralStorageProviderCollateralRebalance, error) {
	event := new(StorageProviderCollateralStorageProviderCollateralRebalance)
	if err := _StorageProviderCollateral.contract.UnpackLog(event, "StorageProviderCollateralRebalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageProviderCollateralStorageProviderCollateralSlashIterator is returned from FilterStorageProviderCollateralSlash and is used to iterate over the raw logs and unpacked data for StorageProviderCollateralSlash events raised by the StorageProviderCollateral contract.
type StorageProviderCollateralStorageProviderCollateralSlashIterator struct {
	Event *StorageProviderCollateralStorageProviderCollateralSlash // Event containing the contract specifics and raw log

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
func (it *StorageProviderCollateralStorageProviderCollateralSlashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProviderCollateralStorageProviderCollateralSlash)
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
		it.Event = new(StorageProviderCollateralStorageProviderCollateralSlash)
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
func (it *StorageProviderCollateralStorageProviderCollateralSlashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProviderCollateralStorageProviderCollateralSlashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProviderCollateralStorageProviderCollateralSlash represents a StorageProviderCollateralSlash event raised by the StorageProviderCollateral contract.
type StorageProviderCollateralStorageProviderCollateralSlash struct {
	OwnerId     uint64
	SlashingAmt *big.Int
	Pool        common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStorageProviderCollateralSlash is a free log retrieval operation binding the contract event 0x243212650e99fb205ebebf8fcde07325988e08f3f0501f252c2dc0aaef6f39cb.
//
// Solidity: event StorageProviderCollateralSlash(uint64 _ownerId, uint256 slashingAmt, address pool)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) FilterStorageProviderCollateralSlash(opts *bind.FilterOpts) (*StorageProviderCollateralStorageProviderCollateralSlashIterator, error) {

	logs, sub, err := _StorageProviderCollateral.contract.FilterLogs(opts, "StorageProviderCollateralSlash")
	if err != nil {
		return nil, err
	}
	return &StorageProviderCollateralStorageProviderCollateralSlashIterator{contract: _StorageProviderCollateral.contract, event: "StorageProviderCollateralSlash", logs: logs, sub: sub}, nil
}

// WatchStorageProviderCollateralSlash is a free log subscription operation binding the contract event 0x243212650e99fb205ebebf8fcde07325988e08f3f0501f252c2dc0aaef6f39cb.
//
// Solidity: event StorageProviderCollateralSlash(uint64 _ownerId, uint256 slashingAmt, address pool)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) WatchStorageProviderCollateralSlash(opts *bind.WatchOpts, sink chan<- *StorageProviderCollateralStorageProviderCollateralSlash) (event.Subscription, error) {

	logs, sub, err := _StorageProviderCollateral.contract.WatchLogs(opts, "StorageProviderCollateralSlash")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProviderCollateralStorageProviderCollateralSlash)
				if err := _StorageProviderCollateral.contract.UnpackLog(event, "StorageProviderCollateralSlash", log); err != nil {
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

// ParseStorageProviderCollateralSlash is a log parse operation binding the contract event 0x243212650e99fb205ebebf8fcde07325988e08f3f0501f252c2dc0aaef6f39cb.
//
// Solidity: event StorageProviderCollateralSlash(uint64 _ownerId, uint256 slashingAmt, address pool)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) ParseStorageProviderCollateralSlash(log types.Log) (*StorageProviderCollateralStorageProviderCollateralSlash, error) {
	event := new(StorageProviderCollateralStorageProviderCollateralSlash)
	if err := _StorageProviderCollateral.contract.UnpackLog(event, "StorageProviderCollateralSlash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageProviderCollateralStorageProviderCollateralUpdateIterator is returned from FilterStorageProviderCollateralUpdate and is used to iterate over the raw logs and unpacked data for StorageProviderCollateralUpdate events raised by the StorageProviderCollateral contract.
type StorageProviderCollateralStorageProviderCollateralUpdateIterator struct {
	Event *StorageProviderCollateralStorageProviderCollateralUpdate // Event containing the contract specifics and raw log

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
func (it *StorageProviderCollateralStorageProviderCollateralUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProviderCollateralStorageProviderCollateralUpdate)
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
		it.Event = new(StorageProviderCollateralStorageProviderCollateralUpdate)
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
func (it *StorageProviderCollateralStorageProviderCollateralUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProviderCollateralStorageProviderCollateralUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProviderCollateralStorageProviderCollateralUpdate represents a StorageProviderCollateralUpdate event raised by the StorageProviderCollateral contract.
type StorageProviderCollateralStorageProviderCollateralUpdate struct {
	OwnerId          uint64
	PrevRequirements *big.Int
	Requirements     *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterStorageProviderCollateralUpdate is a free log retrieval operation binding the contract event 0x91147ad9b4ecd1955cddf8d3a9c61eb67a1b97c9411912f17865e356b304bd6d.
//
// Solidity: event StorageProviderCollateralUpdate(uint64 _ownerId, uint256 prevRequirements, uint256 requirements)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) FilterStorageProviderCollateralUpdate(opts *bind.FilterOpts) (*StorageProviderCollateralStorageProviderCollateralUpdateIterator, error) {

	logs, sub, err := _StorageProviderCollateral.contract.FilterLogs(opts, "StorageProviderCollateralUpdate")
	if err != nil {
		return nil, err
	}
	return &StorageProviderCollateralStorageProviderCollateralUpdateIterator{contract: _StorageProviderCollateral.contract, event: "StorageProviderCollateralUpdate", logs: logs, sub: sub}, nil
}

// WatchStorageProviderCollateralUpdate is a free log subscription operation binding the contract event 0x91147ad9b4ecd1955cddf8d3a9c61eb67a1b97c9411912f17865e356b304bd6d.
//
// Solidity: event StorageProviderCollateralUpdate(uint64 _ownerId, uint256 prevRequirements, uint256 requirements)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) WatchStorageProviderCollateralUpdate(opts *bind.WatchOpts, sink chan<- *StorageProviderCollateralStorageProviderCollateralUpdate) (event.Subscription, error) {

	logs, sub, err := _StorageProviderCollateral.contract.WatchLogs(opts, "StorageProviderCollateralUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProviderCollateralStorageProviderCollateralUpdate)
				if err := _StorageProviderCollateral.contract.UnpackLog(event, "StorageProviderCollateralUpdate", log); err != nil {
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

// ParseStorageProviderCollateralUpdate is a log parse operation binding the contract event 0x91147ad9b4ecd1955cddf8d3a9c61eb67a1b97c9411912f17865e356b304bd6d.
//
// Solidity: event StorageProviderCollateralUpdate(uint64 _ownerId, uint256 prevRequirements, uint256 requirements)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) ParseStorageProviderCollateralUpdate(log types.Log) (*StorageProviderCollateralStorageProviderCollateralUpdate, error) {
	event := new(StorageProviderCollateralStorageProviderCollateralUpdate)
	if err := _StorageProviderCollateral.contract.UnpackLog(event, "StorageProviderCollateralUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageProviderCollateralStorageProviderCollateralWithdrawIterator is returned from FilterStorageProviderCollateralWithdraw and is used to iterate over the raw logs and unpacked data for StorageProviderCollateralWithdraw events raised by the StorageProviderCollateral contract.
type StorageProviderCollateralStorageProviderCollateralWithdrawIterator struct {
	Event *StorageProviderCollateralStorageProviderCollateralWithdraw // Event containing the contract specifics and raw log

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
func (it *StorageProviderCollateralStorageProviderCollateralWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProviderCollateralStorageProviderCollateralWithdraw)
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
		it.Event = new(StorageProviderCollateralStorageProviderCollateralWithdraw)
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
func (it *StorageProviderCollateralStorageProviderCollateralWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProviderCollateralStorageProviderCollateralWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProviderCollateralStorageProviderCollateralWithdraw represents a StorageProviderCollateralWithdraw event raised by the StorageProviderCollateral contract.
type StorageProviderCollateralStorageProviderCollateralWithdraw struct {
	OwnerId uint64
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStorageProviderCollateralWithdraw is a free log retrieval operation binding the contract event 0xe4b3039f013ad4f4d16ea3537d0ee0ae75dad7fb7c97f1ff24fd9f10bd85b89b.
//
// Solidity: event StorageProviderCollateralWithdraw(uint64 _ownerId, uint256 amount)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) FilterStorageProviderCollateralWithdraw(opts *bind.FilterOpts) (*StorageProviderCollateralStorageProviderCollateralWithdrawIterator, error) {

	logs, sub, err := _StorageProviderCollateral.contract.FilterLogs(opts, "StorageProviderCollateralWithdraw")
	if err != nil {
		return nil, err
	}
	return &StorageProviderCollateralStorageProviderCollateralWithdrawIterator{contract: _StorageProviderCollateral.contract, event: "StorageProviderCollateralWithdraw", logs: logs, sub: sub}, nil
}

// WatchStorageProviderCollateralWithdraw is a free log subscription operation binding the contract event 0xe4b3039f013ad4f4d16ea3537d0ee0ae75dad7fb7c97f1ff24fd9f10bd85b89b.
//
// Solidity: event StorageProviderCollateralWithdraw(uint64 _ownerId, uint256 amount)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) WatchStorageProviderCollateralWithdraw(opts *bind.WatchOpts, sink chan<- *StorageProviderCollateralStorageProviderCollateralWithdraw) (event.Subscription, error) {

	logs, sub, err := _StorageProviderCollateral.contract.WatchLogs(opts, "StorageProviderCollateralWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProviderCollateralStorageProviderCollateralWithdraw)
				if err := _StorageProviderCollateral.contract.UnpackLog(event, "StorageProviderCollateralWithdraw", log); err != nil {
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

// ParseStorageProviderCollateralWithdraw is a log parse operation binding the contract event 0xe4b3039f013ad4f4d16ea3537d0ee0ae75dad7fb7c97f1ff24fd9f10bd85b89b.
//
// Solidity: event StorageProviderCollateralWithdraw(uint64 _ownerId, uint256 amount)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) ParseStorageProviderCollateralWithdraw(log types.Log) (*StorageProviderCollateralStorageProviderCollateralWithdraw, error) {
	event := new(StorageProviderCollateralStorageProviderCollateralWithdraw)
	if err := _StorageProviderCollateral.contract.UnpackLog(event, "StorageProviderCollateralWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageProviderCollateralUpdateBaseCollateralRequirementsIterator is returned from FilterUpdateBaseCollateralRequirements and is used to iterate over the raw logs and unpacked data for UpdateBaseCollateralRequirements events raised by the StorageProviderCollateral contract.
type StorageProviderCollateralUpdateBaseCollateralRequirementsIterator struct {
	Event *StorageProviderCollateralUpdateBaseCollateralRequirements // Event containing the contract specifics and raw log

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
func (it *StorageProviderCollateralUpdateBaseCollateralRequirementsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProviderCollateralUpdateBaseCollateralRequirements)
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
		it.Event = new(StorageProviderCollateralUpdateBaseCollateralRequirements)
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
func (it *StorageProviderCollateralUpdateBaseCollateralRequirementsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProviderCollateralUpdateBaseCollateralRequirementsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProviderCollateralUpdateBaseCollateralRequirements represents a UpdateBaseCollateralRequirements event raised by the StorageProviderCollateral contract.
type StorageProviderCollateralUpdateBaseCollateralRequirements struct {
	BaseCollateralRequirements *big.Int
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterUpdateBaseCollateralRequirements is a free log retrieval operation binding the contract event 0xe49b6558d30d6f2b9c3bb53bb6bedb636cf1dd53419057f0ab65812a7739639e.
//
// Solidity: event UpdateBaseCollateralRequirements(uint256 baseCollateralRequirements)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) FilterUpdateBaseCollateralRequirements(opts *bind.FilterOpts) (*StorageProviderCollateralUpdateBaseCollateralRequirementsIterator, error) {

	logs, sub, err := _StorageProviderCollateral.contract.FilterLogs(opts, "UpdateBaseCollateralRequirements")
	if err != nil {
		return nil, err
	}
	return &StorageProviderCollateralUpdateBaseCollateralRequirementsIterator{contract: _StorageProviderCollateral.contract, event: "UpdateBaseCollateralRequirements", logs: logs, sub: sub}, nil
}

// WatchUpdateBaseCollateralRequirements is a free log subscription operation binding the contract event 0xe49b6558d30d6f2b9c3bb53bb6bedb636cf1dd53419057f0ab65812a7739639e.
//
// Solidity: event UpdateBaseCollateralRequirements(uint256 baseCollateralRequirements)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) WatchUpdateBaseCollateralRequirements(opts *bind.WatchOpts, sink chan<- *StorageProviderCollateralUpdateBaseCollateralRequirements) (event.Subscription, error) {

	logs, sub, err := _StorageProviderCollateral.contract.WatchLogs(opts, "UpdateBaseCollateralRequirements")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProviderCollateralUpdateBaseCollateralRequirements)
				if err := _StorageProviderCollateral.contract.UnpackLog(event, "UpdateBaseCollateralRequirements", log); err != nil {
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

// ParseUpdateBaseCollateralRequirements is a log parse operation binding the contract event 0xe49b6558d30d6f2b9c3bb53bb6bedb636cf1dd53419057f0ab65812a7739639e.
//
// Solidity: event UpdateBaseCollateralRequirements(uint256 baseCollateralRequirements)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) ParseUpdateBaseCollateralRequirements(log types.Log) (*StorageProviderCollateralUpdateBaseCollateralRequirements, error) {
	event := new(StorageProviderCollateralUpdateBaseCollateralRequirements)
	if err := _StorageProviderCollateral.contract.UnpackLog(event, "UpdateBaseCollateralRequirements", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageProviderCollateralUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the StorageProviderCollateral contract.
type StorageProviderCollateralUpgradedIterator struct {
	Event *StorageProviderCollateralUpgraded // Event containing the contract specifics and raw log

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
func (it *StorageProviderCollateralUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProviderCollateralUpgraded)
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
		it.Event = new(StorageProviderCollateralUpgraded)
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
func (it *StorageProviderCollateralUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProviderCollateralUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProviderCollateralUpgraded represents a Upgraded event raised by the StorageProviderCollateral contract.
type StorageProviderCollateralUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*StorageProviderCollateralUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _StorageProviderCollateral.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &StorageProviderCollateralUpgradedIterator{contract: _StorageProviderCollateral.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *StorageProviderCollateralUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _StorageProviderCollateral.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProviderCollateralUpgraded)
				if err := _StorageProviderCollateral.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) ParseUpgraded(log types.Log) (*StorageProviderCollateralUpgraded, error) {
	event := new(StorageProviderCollateralUpgraded)
	if err := _StorageProviderCollateral.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
