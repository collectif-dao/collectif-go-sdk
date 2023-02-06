// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PledgeOracle

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

// PledgeOracleMetaData contains all meta data concerning the PledgeOracle contract.
var PledgeOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"genesis\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"preCommitDeposit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"initialPledge\",\"type\":\"uint256\"}],\"name\":\"RecordUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"genesisEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getHistoricalInitialPledge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getHistoricalPreCommitDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getHistoricalRecord\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastInitialPledge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastPreCommitDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastRecord\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPledgeFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastEpochReport\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preCommitDeposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialPledge\",\"type\":\"uint256\"}],\"name\":\"updateRecord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PledgeOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use PledgeOracleMetaData.ABI instead.
var PledgeOracleABI = PledgeOracleMetaData.ABI

// PledgeOracle is an auto generated Go binding around an Ethereum contract.
type PledgeOracle struct {
	PledgeOracleCaller     // Read-only binding to the contract
	PledgeOracleTransactor // Write-only binding to the contract
	PledgeOracleFilterer   // Log filterer for contract events
}

// PledgeOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type PledgeOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PledgeOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PledgeOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PledgeOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PledgeOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PledgeOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PledgeOracleSession struct {
	Contract     *PledgeOracle     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PledgeOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PledgeOracleCallerSession struct {
	Contract *PledgeOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PledgeOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PledgeOracleTransactorSession struct {
	Contract     *PledgeOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PledgeOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type PledgeOracleRaw struct {
	Contract *PledgeOracle // Generic contract binding to access the raw methods on
}

// PledgeOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PledgeOracleCallerRaw struct {
	Contract *PledgeOracleCaller // Generic read-only contract binding to access the raw methods on
}

// PledgeOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PledgeOracleTransactorRaw struct {
	Contract *PledgeOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPledgeOracle creates a new instance of PledgeOracle, bound to a specific deployed contract.
func NewPledgeOracle(address common.Address, backend bind.ContractBackend) (*PledgeOracle, error) {
	contract, err := bindPledgeOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PledgeOracle{PledgeOracleCaller: PledgeOracleCaller{contract: contract}, PledgeOracleTransactor: PledgeOracleTransactor{contract: contract}, PledgeOracleFilterer: PledgeOracleFilterer{contract: contract}}, nil
}

// NewPledgeOracleCaller creates a new read-only instance of PledgeOracle, bound to a specific deployed contract.
func NewPledgeOracleCaller(address common.Address, caller bind.ContractCaller) (*PledgeOracleCaller, error) {
	contract, err := bindPledgeOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PledgeOracleCaller{contract: contract}, nil
}

// NewPledgeOracleTransactor creates a new write-only instance of PledgeOracle, bound to a specific deployed contract.
func NewPledgeOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*PledgeOracleTransactor, error) {
	contract, err := bindPledgeOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PledgeOracleTransactor{contract: contract}, nil
}

// NewPledgeOracleFilterer creates a new log filterer instance of PledgeOracle, bound to a specific deployed contract.
func NewPledgeOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*PledgeOracleFilterer, error) {
	contract, err := bindPledgeOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PledgeOracleFilterer{contract: contract}, nil
}

// bindPledgeOracle binds a generic wrapper to an already deployed contract.
func bindPledgeOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PledgeOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PledgeOracle *PledgeOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PledgeOracle.Contract.PledgeOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PledgeOracle *PledgeOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PledgeOracle.Contract.PledgeOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PledgeOracle *PledgeOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PledgeOracle.Contract.PledgeOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PledgeOracle *PledgeOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PledgeOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PledgeOracle *PledgeOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PledgeOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PledgeOracle *PledgeOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PledgeOracle.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PledgeOracle *PledgeOracleCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PledgeOracle.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PledgeOracle *PledgeOracleSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PledgeOracle.Contract.DEFAULTADMINROLE(&_PledgeOracle.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PledgeOracle *PledgeOracleCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PledgeOracle.Contract.DEFAULTADMINROLE(&_PledgeOracle.CallOpts)
}

// GenesisEpoch is a free data retrieval call binding the contract method 0xb70e6be6.
//
// Solidity: function genesisEpoch() view returns(uint256)
func (_PledgeOracle *PledgeOracleCaller) GenesisEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PledgeOracle.contract.Call(opts, &out, "genesisEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GenesisEpoch is a free data retrieval call binding the contract method 0xb70e6be6.
//
// Solidity: function genesisEpoch() view returns(uint256)
func (_PledgeOracle *PledgeOracleSession) GenesisEpoch() (*big.Int, error) {
	return _PledgeOracle.Contract.GenesisEpoch(&_PledgeOracle.CallOpts)
}

// GenesisEpoch is a free data retrieval call binding the contract method 0xb70e6be6.
//
// Solidity: function genesisEpoch() view returns(uint256)
func (_PledgeOracle *PledgeOracleCallerSession) GenesisEpoch() (*big.Int, error) {
	return _PledgeOracle.Contract.GenesisEpoch(&_PledgeOracle.CallOpts)
}

// GetHistoricalInitialPledge is a free data retrieval call binding the contract method 0x8dc7a237.
//
// Solidity: function getHistoricalInitialPledge(uint256 epoch) view returns(uint256)
func (_PledgeOracle *PledgeOracleCaller) GetHistoricalInitialPledge(opts *bind.CallOpts, epoch *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PledgeOracle.contract.Call(opts, &out, "getHistoricalInitialPledge", epoch)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetHistoricalInitialPledge is a free data retrieval call binding the contract method 0x8dc7a237.
//
// Solidity: function getHistoricalInitialPledge(uint256 epoch) view returns(uint256)
func (_PledgeOracle *PledgeOracleSession) GetHistoricalInitialPledge(epoch *big.Int) (*big.Int, error) {
	return _PledgeOracle.Contract.GetHistoricalInitialPledge(&_PledgeOracle.CallOpts, epoch)
}

// GetHistoricalInitialPledge is a free data retrieval call binding the contract method 0x8dc7a237.
//
// Solidity: function getHistoricalInitialPledge(uint256 epoch) view returns(uint256)
func (_PledgeOracle *PledgeOracleCallerSession) GetHistoricalInitialPledge(epoch *big.Int) (*big.Int, error) {
	return _PledgeOracle.Contract.GetHistoricalInitialPledge(&_PledgeOracle.CallOpts, epoch)
}

// GetHistoricalPreCommitDeposit is a free data retrieval call binding the contract method 0xdf08bfba.
//
// Solidity: function getHistoricalPreCommitDeposit(uint256 epoch) view returns(uint256)
func (_PledgeOracle *PledgeOracleCaller) GetHistoricalPreCommitDeposit(opts *bind.CallOpts, epoch *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PledgeOracle.contract.Call(opts, &out, "getHistoricalPreCommitDeposit", epoch)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetHistoricalPreCommitDeposit is a free data retrieval call binding the contract method 0xdf08bfba.
//
// Solidity: function getHistoricalPreCommitDeposit(uint256 epoch) view returns(uint256)
func (_PledgeOracle *PledgeOracleSession) GetHistoricalPreCommitDeposit(epoch *big.Int) (*big.Int, error) {
	return _PledgeOracle.Contract.GetHistoricalPreCommitDeposit(&_PledgeOracle.CallOpts, epoch)
}

// GetHistoricalPreCommitDeposit is a free data retrieval call binding the contract method 0xdf08bfba.
//
// Solidity: function getHistoricalPreCommitDeposit(uint256 epoch) view returns(uint256)
func (_PledgeOracle *PledgeOracleCallerSession) GetHistoricalPreCommitDeposit(epoch *big.Int) (*big.Int, error) {
	return _PledgeOracle.Contract.GetHistoricalPreCommitDeposit(&_PledgeOracle.CallOpts, epoch)
}

// GetHistoricalRecord is a free data retrieval call binding the contract method 0x51628a81.
//
// Solidity: function getHistoricalRecord(uint256 epoch) view returns(uint256, uint256)
func (_PledgeOracle *PledgeOracleCaller) GetHistoricalRecord(opts *bind.CallOpts, epoch *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _PledgeOracle.contract.Call(opts, &out, "getHistoricalRecord", epoch)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetHistoricalRecord is a free data retrieval call binding the contract method 0x51628a81.
//
// Solidity: function getHistoricalRecord(uint256 epoch) view returns(uint256, uint256)
func (_PledgeOracle *PledgeOracleSession) GetHistoricalRecord(epoch *big.Int) (*big.Int, *big.Int, error) {
	return _PledgeOracle.Contract.GetHistoricalRecord(&_PledgeOracle.CallOpts, epoch)
}

// GetHistoricalRecord is a free data retrieval call binding the contract method 0x51628a81.
//
// Solidity: function getHistoricalRecord(uint256 epoch) view returns(uint256, uint256)
func (_PledgeOracle *PledgeOracleCallerSession) GetHistoricalRecord(epoch *big.Int) (*big.Int, *big.Int, error) {
	return _PledgeOracle.Contract.GetHistoricalRecord(&_PledgeOracle.CallOpts, epoch)
}

// GetLastInitialPledge is a free data retrieval call binding the contract method 0x62f89112.
//
// Solidity: function getLastInitialPledge() view returns(uint256)
func (_PledgeOracle *PledgeOracleCaller) GetLastInitialPledge(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PledgeOracle.contract.Call(opts, &out, "getLastInitialPledge")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastInitialPledge is a free data retrieval call binding the contract method 0x62f89112.
//
// Solidity: function getLastInitialPledge() view returns(uint256)
func (_PledgeOracle *PledgeOracleSession) GetLastInitialPledge() (*big.Int, error) {
	return _PledgeOracle.Contract.GetLastInitialPledge(&_PledgeOracle.CallOpts)
}

// GetLastInitialPledge is a free data retrieval call binding the contract method 0x62f89112.
//
// Solidity: function getLastInitialPledge() view returns(uint256)
func (_PledgeOracle *PledgeOracleCallerSession) GetLastInitialPledge() (*big.Int, error) {
	return _PledgeOracle.Contract.GetLastInitialPledge(&_PledgeOracle.CallOpts)
}

// GetLastPreCommitDeposit is a free data retrieval call binding the contract method 0xbeb55976.
//
// Solidity: function getLastPreCommitDeposit() view returns(uint256)
func (_PledgeOracle *PledgeOracleCaller) GetLastPreCommitDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PledgeOracle.contract.Call(opts, &out, "getLastPreCommitDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastPreCommitDeposit is a free data retrieval call binding the contract method 0xbeb55976.
//
// Solidity: function getLastPreCommitDeposit() view returns(uint256)
func (_PledgeOracle *PledgeOracleSession) GetLastPreCommitDeposit() (*big.Int, error) {
	return _PledgeOracle.Contract.GetLastPreCommitDeposit(&_PledgeOracle.CallOpts)
}

// GetLastPreCommitDeposit is a free data retrieval call binding the contract method 0xbeb55976.
//
// Solidity: function getLastPreCommitDeposit() view returns(uint256)
func (_PledgeOracle *PledgeOracleCallerSession) GetLastPreCommitDeposit() (*big.Int, error) {
	return _PledgeOracle.Contract.GetLastPreCommitDeposit(&_PledgeOracle.CallOpts)
}

// GetLastRecord is a free data retrieval call binding the contract method 0x09fe2f0e.
//
// Solidity: function getLastRecord() view returns(uint256, uint256)
func (_PledgeOracle *PledgeOracleCaller) GetLastRecord(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _PledgeOracle.contract.Call(opts, &out, "getLastRecord")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetLastRecord is a free data retrieval call binding the contract method 0x09fe2f0e.
//
// Solidity: function getLastRecord() view returns(uint256, uint256)
func (_PledgeOracle *PledgeOracleSession) GetLastRecord() (*big.Int, *big.Int, error) {
	return _PledgeOracle.Contract.GetLastRecord(&_PledgeOracle.CallOpts)
}

// GetLastRecord is a free data retrieval call binding the contract method 0x09fe2f0e.
//
// Solidity: function getLastRecord() view returns(uint256, uint256)
func (_PledgeOracle *PledgeOracleCallerSession) GetLastRecord() (*big.Int, *big.Int, error) {
	return _PledgeOracle.Contract.GetLastRecord(&_PledgeOracle.CallOpts)
}

// GetPledgeFees is a free data retrieval call binding the contract method 0xb61a2231.
//
// Solidity: function getPledgeFees() view returns(uint256)
func (_PledgeOracle *PledgeOracleCaller) GetPledgeFees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PledgeOracle.contract.Call(opts, &out, "getPledgeFees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPledgeFees is a free data retrieval call binding the contract method 0xb61a2231.
//
// Solidity: function getPledgeFees() view returns(uint256)
func (_PledgeOracle *PledgeOracleSession) GetPledgeFees() (*big.Int, error) {
	return _PledgeOracle.Contract.GetPledgeFees(&_PledgeOracle.CallOpts)
}

// GetPledgeFees is a free data retrieval call binding the contract method 0xb61a2231.
//
// Solidity: function getPledgeFees() view returns(uint256)
func (_PledgeOracle *PledgeOracleCallerSession) GetPledgeFees() (*big.Int, error) {
	return _PledgeOracle.Contract.GetPledgeFees(&_PledgeOracle.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PledgeOracle *PledgeOracleCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _PledgeOracle.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PledgeOracle *PledgeOracleSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PledgeOracle.Contract.GetRoleAdmin(&_PledgeOracle.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PledgeOracle *PledgeOracleCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PledgeOracle.Contract.GetRoleAdmin(&_PledgeOracle.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PledgeOracle *PledgeOracleCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _PledgeOracle.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PledgeOracle *PledgeOracleSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PledgeOracle.Contract.HasRole(&_PledgeOracle.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PledgeOracle *PledgeOracleCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PledgeOracle.Contract.HasRole(&_PledgeOracle.CallOpts, role, account)
}

// LastEpochReport is a free data retrieval call binding the contract method 0xc4cfd19c.
//
// Solidity: function lastEpochReport() view returns(uint256)
func (_PledgeOracle *PledgeOracleCaller) LastEpochReport(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PledgeOracle.contract.Call(opts, &out, "lastEpochReport")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastEpochReport is a free data retrieval call binding the contract method 0xc4cfd19c.
//
// Solidity: function lastEpochReport() view returns(uint256)
func (_PledgeOracle *PledgeOracleSession) LastEpochReport() (*big.Int, error) {
	return _PledgeOracle.Contract.LastEpochReport(&_PledgeOracle.CallOpts)
}

// LastEpochReport is a free data retrieval call binding the contract method 0xc4cfd19c.
//
// Solidity: function lastEpochReport() view returns(uint256)
func (_PledgeOracle *PledgeOracleCallerSession) LastEpochReport() (*big.Int, error) {
	return _PledgeOracle.Contract.LastEpochReport(&_PledgeOracle.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PledgeOracle *PledgeOracleCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _PledgeOracle.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PledgeOracle *PledgeOracleSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PledgeOracle.Contract.SupportsInterface(&_PledgeOracle.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PledgeOracle *PledgeOracleCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PledgeOracle.Contract.SupportsInterface(&_PledgeOracle.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PledgeOracle *PledgeOracleTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PledgeOracle.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PledgeOracle *PledgeOracleSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PledgeOracle.Contract.GrantRole(&_PledgeOracle.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PledgeOracle *PledgeOracleTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PledgeOracle.Contract.GrantRole(&_PledgeOracle.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PledgeOracle *PledgeOracleTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PledgeOracle.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PledgeOracle *PledgeOracleSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PledgeOracle.Contract.RenounceRole(&_PledgeOracle.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PledgeOracle *PledgeOracleTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PledgeOracle.Contract.RenounceRole(&_PledgeOracle.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PledgeOracle *PledgeOracleTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PledgeOracle.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PledgeOracle *PledgeOracleSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PledgeOracle.Contract.RevokeRole(&_PledgeOracle.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PledgeOracle *PledgeOracleTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PledgeOracle.Contract.RevokeRole(&_PledgeOracle.TransactOpts, role, account)
}

// UpdateRecord is a paid mutator transaction binding the contract method 0x79d41df0.
//
// Solidity: function updateRecord(uint256 epoch, uint256 preCommitDeposit, uint256 initialPledge) returns()
func (_PledgeOracle *PledgeOracleTransactor) UpdateRecord(opts *bind.TransactOpts, epoch *big.Int, preCommitDeposit *big.Int, initialPledge *big.Int) (*types.Transaction, error) {
	return _PledgeOracle.contract.Transact(opts, "updateRecord", epoch, preCommitDeposit, initialPledge)
}

// UpdateRecord is a paid mutator transaction binding the contract method 0x79d41df0.
//
// Solidity: function updateRecord(uint256 epoch, uint256 preCommitDeposit, uint256 initialPledge) returns()
func (_PledgeOracle *PledgeOracleSession) UpdateRecord(epoch *big.Int, preCommitDeposit *big.Int, initialPledge *big.Int) (*types.Transaction, error) {
	return _PledgeOracle.Contract.UpdateRecord(&_PledgeOracle.TransactOpts, epoch, preCommitDeposit, initialPledge)
}

// UpdateRecord is a paid mutator transaction binding the contract method 0x79d41df0.
//
// Solidity: function updateRecord(uint256 epoch, uint256 preCommitDeposit, uint256 initialPledge) returns()
func (_PledgeOracle *PledgeOracleTransactorSession) UpdateRecord(epoch *big.Int, preCommitDeposit *big.Int, initialPledge *big.Int) (*types.Transaction, error) {
	return _PledgeOracle.Contract.UpdateRecord(&_PledgeOracle.TransactOpts, epoch, preCommitDeposit, initialPledge)
}

// PledgeOracleRecordUpdatedIterator is returned from FilterRecordUpdated and is used to iterate over the raw logs and unpacked data for RecordUpdated events raised by the PledgeOracle contract.
type PledgeOracleRecordUpdatedIterator struct {
	Event *PledgeOracleRecordUpdated // Event containing the contract specifics and raw log

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
func (it *PledgeOracleRecordUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgeOracleRecordUpdated)
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
		it.Event = new(PledgeOracleRecordUpdated)
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
func (it *PledgeOracleRecordUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgeOracleRecordUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgeOracleRecordUpdated represents a RecordUpdated event raised by the PledgeOracle contract.
type PledgeOracleRecordUpdated struct {
	Epoch            *big.Int
	PreCommitDeposit *big.Int
	InitialPledge    *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRecordUpdated is a free log retrieval operation binding the contract event 0x05cdc323f67bbb6f92fe5df2d12590c3a4f836c7c71f0d013f530ceddeed3755.
//
// Solidity: event RecordUpdated(uint256 epoch, uint256 preCommitDeposit, uint256 initialPledge)
func (_PledgeOracle *PledgeOracleFilterer) FilterRecordUpdated(opts *bind.FilterOpts) (*PledgeOracleRecordUpdatedIterator, error) {

	logs, sub, err := _PledgeOracle.contract.FilterLogs(opts, "RecordUpdated")
	if err != nil {
		return nil, err
	}
	return &PledgeOracleRecordUpdatedIterator{contract: _PledgeOracle.contract, event: "RecordUpdated", logs: logs, sub: sub}, nil
}

// WatchRecordUpdated is a free log subscription operation binding the contract event 0x05cdc323f67bbb6f92fe5df2d12590c3a4f836c7c71f0d013f530ceddeed3755.
//
// Solidity: event RecordUpdated(uint256 epoch, uint256 preCommitDeposit, uint256 initialPledge)
func (_PledgeOracle *PledgeOracleFilterer) WatchRecordUpdated(opts *bind.WatchOpts, sink chan<- *PledgeOracleRecordUpdated) (event.Subscription, error) {

	logs, sub, err := _PledgeOracle.contract.WatchLogs(opts, "RecordUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgeOracleRecordUpdated)
				if err := _PledgeOracle.contract.UnpackLog(event, "RecordUpdated", log); err != nil {
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

// ParseRecordUpdated is a log parse operation binding the contract event 0x05cdc323f67bbb6f92fe5df2d12590c3a4f836c7c71f0d013f530ceddeed3755.
//
// Solidity: event RecordUpdated(uint256 epoch, uint256 preCommitDeposit, uint256 initialPledge)
func (_PledgeOracle *PledgeOracleFilterer) ParseRecordUpdated(log types.Log) (*PledgeOracleRecordUpdated, error) {
	event := new(PledgeOracleRecordUpdated)
	if err := _PledgeOracle.contract.UnpackLog(event, "RecordUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgeOracleRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the PledgeOracle contract.
type PledgeOracleRoleAdminChangedIterator struct {
	Event *PledgeOracleRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *PledgeOracleRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgeOracleRoleAdminChanged)
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
		it.Event = new(PledgeOracleRoleAdminChanged)
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
func (it *PledgeOracleRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgeOracleRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgeOracleRoleAdminChanged represents a RoleAdminChanged event raised by the PledgeOracle contract.
type PledgeOracleRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PledgeOracle *PledgeOracleFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PledgeOracleRoleAdminChangedIterator, error) {

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

	logs, sub, err := _PledgeOracle.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &PledgeOracleRoleAdminChangedIterator{contract: _PledgeOracle.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PledgeOracle *PledgeOracleFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PledgeOracleRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _PledgeOracle.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgeOracleRoleAdminChanged)
				if err := _PledgeOracle.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_PledgeOracle *PledgeOracleFilterer) ParseRoleAdminChanged(log types.Log) (*PledgeOracleRoleAdminChanged, error) {
	event := new(PledgeOracleRoleAdminChanged)
	if err := _PledgeOracle.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgeOracleRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the PledgeOracle contract.
type PledgeOracleRoleGrantedIterator struct {
	Event *PledgeOracleRoleGranted // Event containing the contract specifics and raw log

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
func (it *PledgeOracleRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgeOracleRoleGranted)
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
		it.Event = new(PledgeOracleRoleGranted)
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
func (it *PledgeOracleRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgeOracleRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgeOracleRoleGranted represents a RoleGranted event raised by the PledgeOracle contract.
type PledgeOracleRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PledgeOracle *PledgeOracleFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PledgeOracleRoleGrantedIterator, error) {

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

	logs, sub, err := _PledgeOracle.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PledgeOracleRoleGrantedIterator{contract: _PledgeOracle.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PledgeOracle *PledgeOracleFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PledgeOracleRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PledgeOracle.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgeOracleRoleGranted)
				if err := _PledgeOracle.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_PledgeOracle *PledgeOracleFilterer) ParseRoleGranted(log types.Log) (*PledgeOracleRoleGranted, error) {
	event := new(PledgeOracleRoleGranted)
	if err := _PledgeOracle.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgeOracleRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the PledgeOracle contract.
type PledgeOracleRoleRevokedIterator struct {
	Event *PledgeOracleRoleRevoked // Event containing the contract specifics and raw log

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
func (it *PledgeOracleRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgeOracleRoleRevoked)
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
		it.Event = new(PledgeOracleRoleRevoked)
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
func (it *PledgeOracleRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgeOracleRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgeOracleRoleRevoked represents a RoleRevoked event raised by the PledgeOracle contract.
type PledgeOracleRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PledgeOracle *PledgeOracleFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PledgeOracleRoleRevokedIterator, error) {

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

	logs, sub, err := _PledgeOracle.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PledgeOracleRoleRevokedIterator{contract: _PledgeOracle.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PledgeOracle *PledgeOracleFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PledgeOracleRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PledgeOracle.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgeOracleRoleRevoked)
				if err := _PledgeOracle.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_PledgeOracle *PledgeOracleFilterer) ParseRoleRevoked(log types.Log) (*PledgeOracleRoleRevoked, error) {
	event := new(PledgeOracleRoleRevoked)
	if err := _PledgeOracle.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
