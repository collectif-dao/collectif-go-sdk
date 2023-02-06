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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIWETH9\",\"name\":\"_wFIL\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"provider\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StorageProviderCollateralDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"provider\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allocation\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockAmount\",\"type\":\"uint256\"}],\"name\":\"StorageProviderCollateralLock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"provider\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StorageProviderCollateralWithdraw\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"BASIS_POINTS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WFIL\",\"outputs\":[{\"internalType\":\"contractIWETH9\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateralRequirements\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"collaterals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"availableCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockedCollateral\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_provider\",\"type\":\"bytes\"}],\"name\":\"getAvailableCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_provider\",\"type\":\"bytes\"}],\"name\":\"getCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_provider\",\"type\":\"bytes\"}],\"name\":\"getLockedCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_provider\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_allocated\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractIStorageProviderRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
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

// BASISPOINTS is a free data retrieval call binding the contract method 0xe1f1c4a7.
//
// Solidity: function BASIS_POINTS() view returns(uint256)
func (_StorageProviderCollateral *StorageProviderCollateralCaller) BASISPOINTS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StorageProviderCollateral.contract.Call(opts, &out, "BASIS_POINTS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASISPOINTS is a free data retrieval call binding the contract method 0xe1f1c4a7.
//
// Solidity: function BASIS_POINTS() view returns(uint256)
func (_StorageProviderCollateral *StorageProviderCollateralSession) BASISPOINTS() (*big.Int, error) {
	return _StorageProviderCollateral.Contract.BASISPOINTS(&_StorageProviderCollateral.CallOpts)
}

// BASISPOINTS is a free data retrieval call binding the contract method 0xe1f1c4a7.
//
// Solidity: function BASIS_POINTS() view returns(uint256)
func (_StorageProviderCollateral *StorageProviderCollateralCallerSession) BASISPOINTS() (*big.Int, error) {
	return _StorageProviderCollateral.Contract.BASISPOINTS(&_StorageProviderCollateral.CallOpts)
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

// CollateralRequirements is a free data retrieval call binding the contract method 0xd4b8a554.
//
// Solidity: function collateralRequirements() view returns(uint256)
func (_StorageProviderCollateral *StorageProviderCollateralCaller) CollateralRequirements(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StorageProviderCollateral.contract.Call(opts, &out, "collateralRequirements")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollateralRequirements is a free data retrieval call binding the contract method 0xd4b8a554.
//
// Solidity: function collateralRequirements() view returns(uint256)
func (_StorageProviderCollateral *StorageProviderCollateralSession) CollateralRequirements() (*big.Int, error) {
	return _StorageProviderCollateral.Contract.CollateralRequirements(&_StorageProviderCollateral.CallOpts)
}

// CollateralRequirements is a free data retrieval call binding the contract method 0xd4b8a554.
//
// Solidity: function collateralRequirements() view returns(uint256)
func (_StorageProviderCollateral *StorageProviderCollateralCallerSession) CollateralRequirements() (*big.Int, error) {
	return _StorageProviderCollateral.Contract.CollateralRequirements(&_StorageProviderCollateral.CallOpts)
}

// Collaterals is a free data retrieval call binding the contract method 0xbf457113.
//
// Solidity: function collaterals(bytes ) view returns(uint256 availableCollateral, uint256 lockedCollateral)
func (_StorageProviderCollateral *StorageProviderCollateralCaller) Collaterals(opts *bind.CallOpts, arg0 []byte) (struct {
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

// Collaterals is a free data retrieval call binding the contract method 0xbf457113.
//
// Solidity: function collaterals(bytes ) view returns(uint256 availableCollateral, uint256 lockedCollateral)
func (_StorageProviderCollateral *StorageProviderCollateralSession) Collaterals(arg0 []byte) (struct {
	AvailableCollateral *big.Int
	LockedCollateral    *big.Int
}, error) {
	return _StorageProviderCollateral.Contract.Collaterals(&_StorageProviderCollateral.CallOpts, arg0)
}

// Collaterals is a free data retrieval call binding the contract method 0xbf457113.
//
// Solidity: function collaterals(bytes ) view returns(uint256 availableCollateral, uint256 lockedCollateral)
func (_StorageProviderCollateral *StorageProviderCollateralCallerSession) Collaterals(arg0 []byte) (struct {
	AvailableCollateral *big.Int
	LockedCollateral    *big.Int
}, error) {
	return _StorageProviderCollateral.Contract.Collaterals(&_StorageProviderCollateral.CallOpts, arg0)
}

// GetAvailableCollateral is a free data retrieval call binding the contract method 0x55d23449.
//
// Solidity: function getAvailableCollateral(bytes _provider) view returns(uint256)
func (_StorageProviderCollateral *StorageProviderCollateralCaller) GetAvailableCollateral(opts *bind.CallOpts, _provider []byte) (*big.Int, error) {
	var out []interface{}
	err := _StorageProviderCollateral.contract.Call(opts, &out, "getAvailableCollateral", _provider)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAvailableCollateral is a free data retrieval call binding the contract method 0x55d23449.
//
// Solidity: function getAvailableCollateral(bytes _provider) view returns(uint256)
func (_StorageProviderCollateral *StorageProviderCollateralSession) GetAvailableCollateral(_provider []byte) (*big.Int, error) {
	return _StorageProviderCollateral.Contract.GetAvailableCollateral(&_StorageProviderCollateral.CallOpts, _provider)
}

// GetAvailableCollateral is a free data retrieval call binding the contract method 0x55d23449.
//
// Solidity: function getAvailableCollateral(bytes _provider) view returns(uint256)
func (_StorageProviderCollateral *StorageProviderCollateralCallerSession) GetAvailableCollateral(_provider []byte) (*big.Int, error) {
	return _StorageProviderCollateral.Contract.GetAvailableCollateral(&_StorageProviderCollateral.CallOpts, _provider)
}

// GetCollateral is a free data retrieval call binding the contract method 0x47f5a64d.
//
// Solidity: function getCollateral(bytes _provider) view returns(uint256, uint256)
func (_StorageProviderCollateral *StorageProviderCollateralCaller) GetCollateral(opts *bind.CallOpts, _provider []byte) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _StorageProviderCollateral.contract.Call(opts, &out, "getCollateral", _provider)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetCollateral is a free data retrieval call binding the contract method 0x47f5a64d.
//
// Solidity: function getCollateral(bytes _provider) view returns(uint256, uint256)
func (_StorageProviderCollateral *StorageProviderCollateralSession) GetCollateral(_provider []byte) (*big.Int, *big.Int, error) {
	return _StorageProviderCollateral.Contract.GetCollateral(&_StorageProviderCollateral.CallOpts, _provider)
}

// GetCollateral is a free data retrieval call binding the contract method 0x47f5a64d.
//
// Solidity: function getCollateral(bytes _provider) view returns(uint256, uint256)
func (_StorageProviderCollateral *StorageProviderCollateralCallerSession) GetCollateral(_provider []byte) (*big.Int, *big.Int, error) {
	return _StorageProviderCollateral.Contract.GetCollateral(&_StorageProviderCollateral.CallOpts, _provider)
}

// GetLockedCollateral is a free data retrieval call binding the contract method 0x7d1d16dd.
//
// Solidity: function getLockedCollateral(bytes _provider) view returns(uint256)
func (_StorageProviderCollateral *StorageProviderCollateralCaller) GetLockedCollateral(opts *bind.CallOpts, _provider []byte) (*big.Int, error) {
	var out []interface{}
	err := _StorageProviderCollateral.contract.Call(opts, &out, "getLockedCollateral", _provider)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLockedCollateral is a free data retrieval call binding the contract method 0x7d1d16dd.
//
// Solidity: function getLockedCollateral(bytes _provider) view returns(uint256)
func (_StorageProviderCollateral *StorageProviderCollateralSession) GetLockedCollateral(_provider []byte) (*big.Int, error) {
	return _StorageProviderCollateral.Contract.GetLockedCollateral(&_StorageProviderCollateral.CallOpts, _provider)
}

// GetLockedCollateral is a free data retrieval call binding the contract method 0x7d1d16dd.
//
// Solidity: function getLockedCollateral(bytes _provider) view returns(uint256)
func (_StorageProviderCollateral *StorageProviderCollateralCallerSession) GetLockedCollateral(_provider []byte) (*big.Int, error) {
	return _StorageProviderCollateral.Contract.GetLockedCollateral(&_StorageProviderCollateral.CallOpts, _provider)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_StorageProviderCollateral *StorageProviderCollateralCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StorageProviderCollateral.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_StorageProviderCollateral *StorageProviderCollateralSession) Registry() (common.Address, error) {
	return _StorageProviderCollateral.Contract.Registry(&_StorageProviderCollateral.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_StorageProviderCollateral *StorageProviderCollateralCallerSession) Registry() (common.Address, error) {
	return _StorageProviderCollateral.Contract.Registry(&_StorageProviderCollateral.CallOpts)
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

// Lock is a paid mutator transaction binding the contract method 0x3babddcb.
//
// Solidity: function lock(bytes _provider, uint256 _allocated) returns()
func (_StorageProviderCollateral *StorageProviderCollateralTransactor) Lock(opts *bind.TransactOpts, _provider []byte, _allocated *big.Int) (*types.Transaction, error) {
	return _StorageProviderCollateral.contract.Transact(opts, "lock", _provider, _allocated)
}

// Lock is a paid mutator transaction binding the contract method 0x3babddcb.
//
// Solidity: function lock(bytes _provider, uint256 _allocated) returns()
func (_StorageProviderCollateral *StorageProviderCollateralSession) Lock(_provider []byte, _allocated *big.Int) (*types.Transaction, error) {
	return _StorageProviderCollateral.Contract.Lock(&_StorageProviderCollateral.TransactOpts, _provider, _allocated)
}

// Lock is a paid mutator transaction binding the contract method 0x3babddcb.
//
// Solidity: function lock(bytes _provider, uint256 _allocated) returns()
func (_StorageProviderCollateral *StorageProviderCollateralTransactorSession) Lock(_provider []byte, _allocated *big.Int) (*types.Transaction, error) {
	return _StorageProviderCollateral.Contract.Lock(&_StorageProviderCollateral.TransactOpts, _provider, _allocated)
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
	Provider []byte
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStorageProviderCollateralDeposit is a free log retrieval operation binding the contract event 0x5955619531a8c5a9d5f30513362c5c48b948cb2a10aa12401159edc7005cc53d.
//
// Solidity: event StorageProviderCollateralDeposit(bytes provider, uint256 amount)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) FilterStorageProviderCollateralDeposit(opts *bind.FilterOpts) (*StorageProviderCollateralStorageProviderCollateralDepositIterator, error) {

	logs, sub, err := _StorageProviderCollateral.contract.FilterLogs(opts, "StorageProviderCollateralDeposit")
	if err != nil {
		return nil, err
	}
	return &StorageProviderCollateralStorageProviderCollateralDepositIterator{contract: _StorageProviderCollateral.contract, event: "StorageProviderCollateralDeposit", logs: logs, sub: sub}, nil
}

// WatchStorageProviderCollateralDeposit is a free log subscription operation binding the contract event 0x5955619531a8c5a9d5f30513362c5c48b948cb2a10aa12401159edc7005cc53d.
//
// Solidity: event StorageProviderCollateralDeposit(bytes provider, uint256 amount)
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

// ParseStorageProviderCollateralDeposit is a log parse operation binding the contract event 0x5955619531a8c5a9d5f30513362c5c48b948cb2a10aa12401159edc7005cc53d.
//
// Solidity: event StorageProviderCollateralDeposit(bytes provider, uint256 amount)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) ParseStorageProviderCollateralDeposit(log types.Log) (*StorageProviderCollateralStorageProviderCollateralDeposit, error) {
	event := new(StorageProviderCollateralStorageProviderCollateralDeposit)
	if err := _StorageProviderCollateral.contract.UnpackLog(event, "StorageProviderCollateralDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageProviderCollateralStorageProviderCollateralLockIterator is returned from FilterStorageProviderCollateralLock and is used to iterate over the raw logs and unpacked data for StorageProviderCollateralLock events raised by the StorageProviderCollateral contract.
type StorageProviderCollateralStorageProviderCollateralLockIterator struct {
	Event *StorageProviderCollateralStorageProviderCollateralLock // Event containing the contract specifics and raw log

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
func (it *StorageProviderCollateralStorageProviderCollateralLockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageProviderCollateralStorageProviderCollateralLock)
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
		it.Event = new(StorageProviderCollateralStorageProviderCollateralLock)
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
func (it *StorageProviderCollateralStorageProviderCollateralLockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageProviderCollateralStorageProviderCollateralLockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageProviderCollateralStorageProviderCollateralLock represents a StorageProviderCollateralLock event raised by the StorageProviderCollateral contract.
type StorageProviderCollateralStorageProviderCollateralLock struct {
	Provider   []byte
	Allocation *big.Int
	LockAmount *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterStorageProviderCollateralLock is a free log retrieval operation binding the contract event 0xb4b28ce9e9594f5d63d437002b4b743fa08d88877f3f1721f41b8c4ba48f4710.
//
// Solidity: event StorageProviderCollateralLock(bytes provider, uint256 allocation, uint256 lockAmount)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) FilterStorageProviderCollateralLock(opts *bind.FilterOpts) (*StorageProviderCollateralStorageProviderCollateralLockIterator, error) {

	logs, sub, err := _StorageProviderCollateral.contract.FilterLogs(opts, "StorageProviderCollateralLock")
	if err != nil {
		return nil, err
	}
	return &StorageProviderCollateralStorageProviderCollateralLockIterator{contract: _StorageProviderCollateral.contract, event: "StorageProviderCollateralLock", logs: logs, sub: sub}, nil
}

// WatchStorageProviderCollateralLock is a free log subscription operation binding the contract event 0xb4b28ce9e9594f5d63d437002b4b743fa08d88877f3f1721f41b8c4ba48f4710.
//
// Solidity: event StorageProviderCollateralLock(bytes provider, uint256 allocation, uint256 lockAmount)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) WatchStorageProviderCollateralLock(opts *bind.WatchOpts, sink chan<- *StorageProviderCollateralStorageProviderCollateralLock) (event.Subscription, error) {

	logs, sub, err := _StorageProviderCollateral.contract.WatchLogs(opts, "StorageProviderCollateralLock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageProviderCollateralStorageProviderCollateralLock)
				if err := _StorageProviderCollateral.contract.UnpackLog(event, "StorageProviderCollateralLock", log); err != nil {
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

// ParseStorageProviderCollateralLock is a log parse operation binding the contract event 0xb4b28ce9e9594f5d63d437002b4b743fa08d88877f3f1721f41b8c4ba48f4710.
//
// Solidity: event StorageProviderCollateralLock(bytes provider, uint256 allocation, uint256 lockAmount)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) ParseStorageProviderCollateralLock(log types.Log) (*StorageProviderCollateralStorageProviderCollateralLock, error) {
	event := new(StorageProviderCollateralStorageProviderCollateralLock)
	if err := _StorageProviderCollateral.contract.UnpackLog(event, "StorageProviderCollateralLock", log); err != nil {
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
	Provider []byte
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStorageProviderCollateralWithdraw is a free log retrieval operation binding the contract event 0x826efc64d046a25932847d11880eb16b143a7f584a0ed6133d8488c40748c9fc.
//
// Solidity: event StorageProviderCollateralWithdraw(bytes provider, uint256 amount)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) FilterStorageProviderCollateralWithdraw(opts *bind.FilterOpts) (*StorageProviderCollateralStorageProviderCollateralWithdrawIterator, error) {

	logs, sub, err := _StorageProviderCollateral.contract.FilterLogs(opts, "StorageProviderCollateralWithdraw")
	if err != nil {
		return nil, err
	}
	return &StorageProviderCollateralStorageProviderCollateralWithdrawIterator{contract: _StorageProviderCollateral.contract, event: "StorageProviderCollateralWithdraw", logs: logs, sub: sub}, nil
}

// WatchStorageProviderCollateralWithdraw is a free log subscription operation binding the contract event 0x826efc64d046a25932847d11880eb16b143a7f584a0ed6133d8488c40748c9fc.
//
// Solidity: event StorageProviderCollateralWithdraw(bytes provider, uint256 amount)
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

// ParseStorageProviderCollateralWithdraw is a log parse operation binding the contract event 0x826efc64d046a25932847d11880eb16b143a7f584a0ed6133d8488c40748c9fc.
//
// Solidity: event StorageProviderCollateralWithdraw(bytes provider, uint256 amount)
func (_StorageProviderCollateral *StorageProviderCollateralFilterer) ParseStorageProviderCollateralWithdraw(log types.Log) (*StorageProviderCollateralStorageProviderCollateralWithdraw, error) {
	event := new(StorageProviderCollateralStorageProviderCollateralWithdraw)
	if err := _StorageProviderCollateral.contract.UnpackLog(event, "StorageProviderCollateralWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
