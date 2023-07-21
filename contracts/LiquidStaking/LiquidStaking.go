// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package LiquidStaking

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

// LiquidStakingMetaData contains all meta data concerning the LiquidStaking contract.
var LiquidStakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ActiveSlashing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"errorCode\",\"type\":\"int256\"}],\"name\":\"ActorError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ActorNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllowanceUnderflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC4626Overflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC4626Underflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC4626ZeroShares\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailToCallActor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InactiveActor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAccess\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"InvalidCodec\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidResponseLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"NotEnoughBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"ownerId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"minerId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Pledge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PledgeRepayment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Stake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Unstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WFIL\",\"outputs\":[{\"internalType\":\"contractIWFIL\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUsageRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wFIL\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_resolver\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"pledge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"repayPledge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"restake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ownerId\",\"type\":\"uint64\"}],\"name\":\"totalFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalFilAvailable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalFilPledged\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"unstake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"unstakeAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// LiquidStakingABI is the input ABI used to generate the binding from.
// Deprecated: Use LiquidStakingMetaData.ABI instead.
var LiquidStakingABI = LiquidStakingMetaData.ABI

// LiquidStaking is an auto generated Go binding around an Ethereum contract.
type LiquidStaking struct {
	LiquidStakingCaller     // Read-only binding to the contract
	LiquidStakingTransactor // Write-only binding to the contract
	LiquidStakingFilterer   // Log filterer for contract events
}

// LiquidStakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type LiquidStakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidStakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LiquidStakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidStakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LiquidStakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidStakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LiquidStakingSession struct {
	Contract     *LiquidStaking    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LiquidStakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LiquidStakingCallerSession struct {
	Contract *LiquidStakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// LiquidStakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LiquidStakingTransactorSession struct {
	Contract     *LiquidStakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// LiquidStakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type LiquidStakingRaw struct {
	Contract *LiquidStaking // Generic contract binding to access the raw methods on
}

// LiquidStakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LiquidStakingCallerRaw struct {
	Contract *LiquidStakingCaller // Generic read-only contract binding to access the raw methods on
}

// LiquidStakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LiquidStakingTransactorRaw struct {
	Contract *LiquidStakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLiquidStaking creates a new instance of LiquidStaking, bound to a specific deployed contract.
func NewLiquidStaking(address common.Address, backend bind.ContractBackend) (*LiquidStaking, error) {
	contract, err := bindLiquidStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LiquidStaking{LiquidStakingCaller: LiquidStakingCaller{contract: contract}, LiquidStakingTransactor: LiquidStakingTransactor{contract: contract}, LiquidStakingFilterer: LiquidStakingFilterer{contract: contract}}, nil
}

// NewLiquidStakingCaller creates a new read-only instance of LiquidStaking, bound to a specific deployed contract.
func NewLiquidStakingCaller(address common.Address, caller bind.ContractCaller) (*LiquidStakingCaller, error) {
	contract, err := bindLiquidStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidStakingCaller{contract: contract}, nil
}

// NewLiquidStakingTransactor creates a new write-only instance of LiquidStaking, bound to a specific deployed contract.
func NewLiquidStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*LiquidStakingTransactor, error) {
	contract, err := bindLiquidStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidStakingTransactor{contract: contract}, nil
}

// NewLiquidStakingFilterer creates a new log filterer instance of LiquidStaking, bound to a specific deployed contract.
func NewLiquidStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*LiquidStakingFilterer, error) {
	contract, err := bindLiquidStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LiquidStakingFilterer{contract: contract}, nil
}

// bindLiquidStaking binds a generic wrapper to an already deployed contract.
func bindLiquidStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LiquidStakingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LiquidStaking *LiquidStakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LiquidStaking.Contract.LiquidStakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LiquidStaking *LiquidStakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidStaking.Contract.LiquidStakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LiquidStaking *LiquidStakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LiquidStaking.Contract.LiquidStakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LiquidStaking *LiquidStakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LiquidStaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LiquidStaking *LiquidStakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidStaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LiquidStaking *LiquidStakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LiquidStaking.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LiquidStaking *LiquidStakingCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LiquidStaking.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LiquidStaking *LiquidStakingSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _LiquidStaking.Contract.DEFAULTADMINROLE(&_LiquidStaking.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LiquidStaking *LiquidStakingCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _LiquidStaking.Contract.DEFAULTADMINROLE(&_LiquidStaking.CallOpts)
}

// WFIL is a free data retrieval call binding the contract method 0x09448a86.
//
// Solidity: function WFIL() view returns(address)
func (_LiquidStaking *LiquidStakingCaller) WFIL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidStaking.contract.Call(opts, &out, "WFIL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WFIL is a free data retrieval call binding the contract method 0x09448a86.
//
// Solidity: function WFIL() view returns(address)
func (_LiquidStaking *LiquidStakingSession) WFIL() (common.Address, error) {
	return _LiquidStaking.Contract.WFIL(&_LiquidStaking.CallOpts)
}

// WFIL is a free data retrieval call binding the contract method 0x09448a86.
//
// Solidity: function WFIL() view returns(address)
func (_LiquidStaking *LiquidStakingCallerSession) WFIL() (common.Address, error) {
	return _LiquidStaking.Contract.WFIL(&_LiquidStaking.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LiquidStaking *LiquidStakingCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquidStaking.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LiquidStaking *LiquidStakingSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _LiquidStaking.Contract.Allowance(&_LiquidStaking.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LiquidStaking *LiquidStakingCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _LiquidStaking.Contract.Allowance(&_LiquidStaking.CallOpts, owner, spender)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_LiquidStaking *LiquidStakingCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidStaking.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_LiquidStaking *LiquidStakingSession) Asset() (common.Address, error) {
	return _LiquidStaking.Contract.Asset(&_LiquidStaking.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_LiquidStaking *LiquidStakingCallerSession) Asset() (common.Address, error) {
	return _LiquidStaking.Contract.Asset(&_LiquidStaking.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LiquidStaking *LiquidStakingCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquidStaking.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LiquidStaking *LiquidStakingSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _LiquidStaking.Contract.BalanceOf(&_LiquidStaking.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LiquidStaking *LiquidStakingCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _LiquidStaking.Contract.BalanceOf(&_LiquidStaking.CallOpts, account)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256 assets)
func (_LiquidStaking *LiquidStakingCaller) ConvertToAssets(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LiquidStaking.contract.Call(opts, &out, "convertToAssets", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256 assets)
func (_LiquidStaking *LiquidStakingSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _LiquidStaking.Contract.ConvertToAssets(&_LiquidStaking.CallOpts, shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256 assets)
func (_LiquidStaking *LiquidStakingCallerSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _LiquidStaking.Contract.ConvertToAssets(&_LiquidStaking.CallOpts, shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256 shares)
func (_LiquidStaking *LiquidStakingCaller) ConvertToShares(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LiquidStaking.contract.Call(opts, &out, "convertToShares", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256 shares)
func (_LiquidStaking *LiquidStakingSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _LiquidStaking.Contract.ConvertToShares(&_LiquidStaking.CallOpts, assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256 shares)
func (_LiquidStaking *LiquidStakingCallerSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _LiquidStaking.Contract.ConvertToShares(&_LiquidStaking.CallOpts, assets)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LiquidStaking *LiquidStakingCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LiquidStaking.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LiquidStaking *LiquidStakingSession) Decimals() (uint8, error) {
	return _LiquidStaking.Contract.Decimals(&_LiquidStaking.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LiquidStaking *LiquidStakingCallerSession) Decimals() (uint8, error) {
	return _LiquidStaking.Contract.Decimals(&_LiquidStaking.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_LiquidStaking *LiquidStakingCaller) GetImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidStaking.contract.Call(opts, &out, "getImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_LiquidStaking *LiquidStakingSession) GetImplementation() (common.Address, error) {
	return _LiquidStaking.Contract.GetImplementation(&_LiquidStaking.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_LiquidStaking *LiquidStakingCallerSession) GetImplementation() (common.Address, error) {
	return _LiquidStaking.Contract.GetImplementation(&_LiquidStaking.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LiquidStaking *LiquidStakingCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _LiquidStaking.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LiquidStaking *LiquidStakingSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _LiquidStaking.Contract.GetRoleAdmin(&_LiquidStaking.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LiquidStaking *LiquidStakingCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _LiquidStaking.Contract.GetRoleAdmin(&_LiquidStaking.CallOpts, role)
}

// GetUsageRatio is a free data retrieval call binding the contract method 0x44775506.
//
// Solidity: function getUsageRatio() view returns(uint256)
func (_LiquidStaking *LiquidStakingCaller) GetUsageRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquidStaking.contract.Call(opts, &out, "getUsageRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUsageRatio is a free data retrieval call binding the contract method 0x44775506.
//
// Solidity: function getUsageRatio() view returns(uint256)
func (_LiquidStaking *LiquidStakingSession) GetUsageRatio() (*big.Int, error) {
	return _LiquidStaking.Contract.GetUsageRatio(&_LiquidStaking.CallOpts)
}

// GetUsageRatio is a free data retrieval call binding the contract method 0x44775506.
//
// Solidity: function getUsageRatio() view returns(uint256)
func (_LiquidStaking *LiquidStakingCallerSession) GetUsageRatio() (*big.Int, error) {
	return _LiquidStaking.Contract.GetUsageRatio(&_LiquidStaking.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LiquidStaking *LiquidStakingCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _LiquidStaking.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LiquidStaking *LiquidStakingSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _LiquidStaking.Contract.HasRole(&_LiquidStaking.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LiquidStaking *LiquidStakingCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _LiquidStaking.Contract.HasRole(&_LiquidStaking.CallOpts, role, account)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_LiquidStaking *LiquidStakingCaller) MaxDeposit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquidStaking.contract.Call(opts, &out, "maxDeposit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_LiquidStaking *LiquidStakingSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _LiquidStaking.Contract.MaxDeposit(&_LiquidStaking.CallOpts, arg0)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_LiquidStaking *LiquidStakingCallerSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _LiquidStaking.Contract.MaxDeposit(&_LiquidStaking.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_LiquidStaking *LiquidStakingCaller) MaxMint(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquidStaking.contract.Call(opts, &out, "maxMint", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_LiquidStaking *LiquidStakingSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _LiquidStaking.Contract.MaxMint(&_LiquidStaking.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_LiquidStaking *LiquidStakingCallerSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _LiquidStaking.Contract.MaxMint(&_LiquidStaking.CallOpts, arg0)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_LiquidStaking *LiquidStakingCaller) MaxRedeem(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquidStaking.contract.Call(opts, &out, "maxRedeem", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_LiquidStaking *LiquidStakingSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _LiquidStaking.Contract.MaxRedeem(&_LiquidStaking.CallOpts, owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_LiquidStaking *LiquidStakingCallerSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _LiquidStaking.Contract.MaxRedeem(&_LiquidStaking.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_LiquidStaking *LiquidStakingCaller) MaxWithdraw(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquidStaking.contract.Call(opts, &out, "maxWithdraw", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_LiquidStaking *LiquidStakingSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _LiquidStaking.Contract.MaxWithdraw(&_LiquidStaking.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_LiquidStaking *LiquidStakingCallerSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _LiquidStaking.Contract.MaxWithdraw(&_LiquidStaking.CallOpts, owner)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LiquidStaking *LiquidStakingCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LiquidStaking.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LiquidStaking *LiquidStakingSession) Name() (string, error) {
	return _LiquidStaking.Contract.Name(&_LiquidStaking.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LiquidStaking *LiquidStakingCallerSession) Name() (string, error) {
	return _LiquidStaking.Contract.Name(&_LiquidStaking.CallOpts)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_LiquidStaking *LiquidStakingCaller) PreviewDeposit(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LiquidStaking.contract.Call(opts, &out, "previewDeposit", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_LiquidStaking *LiquidStakingSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _LiquidStaking.Contract.PreviewDeposit(&_LiquidStaking.CallOpts, assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_LiquidStaking *LiquidStakingCallerSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _LiquidStaking.Contract.PreviewDeposit(&_LiquidStaking.CallOpts, assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_LiquidStaking *LiquidStakingCaller) PreviewMint(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LiquidStaking.contract.Call(opts, &out, "previewMint", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_LiquidStaking *LiquidStakingSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _LiquidStaking.Contract.PreviewMint(&_LiquidStaking.CallOpts, shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_LiquidStaking *LiquidStakingCallerSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _LiquidStaking.Contract.PreviewMint(&_LiquidStaking.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_LiquidStaking *LiquidStakingCaller) PreviewRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LiquidStaking.contract.Call(opts, &out, "previewRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_LiquidStaking *LiquidStakingSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _LiquidStaking.Contract.PreviewRedeem(&_LiquidStaking.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_LiquidStaking *LiquidStakingCallerSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _LiquidStaking.Contract.PreviewRedeem(&_LiquidStaking.CallOpts, shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_LiquidStaking *LiquidStakingCaller) PreviewWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LiquidStaking.contract.Call(opts, &out, "previewWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_LiquidStaking *LiquidStakingSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _LiquidStaking.Contract.PreviewWithdraw(&_LiquidStaking.CallOpts, assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_LiquidStaking *LiquidStakingCallerSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _LiquidStaking.Contract.PreviewWithdraw(&_LiquidStaking.CallOpts, assets)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LiquidStaking *LiquidStakingCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LiquidStaking.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LiquidStaking *LiquidStakingSession) ProxiableUUID() ([32]byte, error) {
	return _LiquidStaking.Contract.ProxiableUUID(&_LiquidStaking.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LiquidStaking *LiquidStakingCallerSession) ProxiableUUID() ([32]byte, error) {
	return _LiquidStaking.Contract.ProxiableUUID(&_LiquidStaking.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LiquidStaking *LiquidStakingCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LiquidStaking.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LiquidStaking *LiquidStakingSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LiquidStaking.Contract.SupportsInterface(&_LiquidStaking.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LiquidStaking *LiquidStakingCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LiquidStaking.Contract.SupportsInterface(&_LiquidStaking.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LiquidStaking *LiquidStakingCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LiquidStaking.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LiquidStaking *LiquidStakingSession) Symbol() (string, error) {
	return _LiquidStaking.Contract.Symbol(&_LiquidStaking.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LiquidStaking *LiquidStakingCallerSession) Symbol() (string, error) {
	return _LiquidStaking.Contract.Symbol(&_LiquidStaking.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_LiquidStaking *LiquidStakingCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquidStaking.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_LiquidStaking *LiquidStakingSession) TotalAssets() (*big.Int, error) {
	return _LiquidStaking.Contract.TotalAssets(&_LiquidStaking.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_LiquidStaking *LiquidStakingCallerSession) TotalAssets() (*big.Int, error) {
	return _LiquidStaking.Contract.TotalAssets(&_LiquidStaking.CallOpts)
}

// TotalFees is a free data retrieval call binding the contract method 0xa83a2e65.
//
// Solidity: function totalFees(uint64 _ownerId) view returns(uint256)
func (_LiquidStaking *LiquidStakingCaller) TotalFees(opts *bind.CallOpts, _ownerId uint64) (*big.Int, error) {
	var out []interface{}
	err := _LiquidStaking.contract.Call(opts, &out, "totalFees", _ownerId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalFees is a free data retrieval call binding the contract method 0xa83a2e65.
//
// Solidity: function totalFees(uint64 _ownerId) view returns(uint256)
func (_LiquidStaking *LiquidStakingSession) TotalFees(_ownerId uint64) (*big.Int, error) {
	return _LiquidStaking.Contract.TotalFees(&_LiquidStaking.CallOpts, _ownerId)
}

// TotalFees is a free data retrieval call binding the contract method 0xa83a2e65.
//
// Solidity: function totalFees(uint64 _ownerId) view returns(uint256)
func (_LiquidStaking *LiquidStakingCallerSession) TotalFees(_ownerId uint64) (*big.Int, error) {
	return _LiquidStaking.Contract.TotalFees(&_LiquidStaking.CallOpts, _ownerId)
}

// TotalFilAvailable is a free data retrieval call binding the contract method 0x655cbba7.
//
// Solidity: function totalFilAvailable() view returns(uint256)
func (_LiquidStaking *LiquidStakingCaller) TotalFilAvailable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquidStaking.contract.Call(opts, &out, "totalFilAvailable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalFilAvailable is a free data retrieval call binding the contract method 0x655cbba7.
//
// Solidity: function totalFilAvailable() view returns(uint256)
func (_LiquidStaking *LiquidStakingSession) TotalFilAvailable() (*big.Int, error) {
	return _LiquidStaking.Contract.TotalFilAvailable(&_LiquidStaking.CallOpts)
}

// TotalFilAvailable is a free data retrieval call binding the contract method 0x655cbba7.
//
// Solidity: function totalFilAvailable() view returns(uint256)
func (_LiquidStaking *LiquidStakingCallerSession) TotalFilAvailable() (*big.Int, error) {
	return _LiquidStaking.Contract.TotalFilAvailable(&_LiquidStaking.CallOpts)
}

// TotalFilPledged is a free data retrieval call binding the contract method 0xad05dc33.
//
// Solidity: function totalFilPledged() view returns(uint256)
func (_LiquidStaking *LiquidStakingCaller) TotalFilPledged(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquidStaking.contract.Call(opts, &out, "totalFilPledged")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalFilPledged is a free data retrieval call binding the contract method 0xad05dc33.
//
// Solidity: function totalFilPledged() view returns(uint256)
func (_LiquidStaking *LiquidStakingSession) TotalFilPledged() (*big.Int, error) {
	return _LiquidStaking.Contract.TotalFilPledged(&_LiquidStaking.CallOpts)
}

// TotalFilPledged is a free data retrieval call binding the contract method 0xad05dc33.
//
// Solidity: function totalFilPledged() view returns(uint256)
func (_LiquidStaking *LiquidStakingCallerSession) TotalFilPledged() (*big.Int, error) {
	return _LiquidStaking.Contract.TotalFilPledged(&_LiquidStaking.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LiquidStaking *LiquidStakingCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquidStaking.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LiquidStaking *LiquidStakingSession) TotalSupply() (*big.Int, error) {
	return _LiquidStaking.Contract.TotalSupply(&_LiquidStaking.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LiquidStaking *LiquidStakingCallerSession) TotalSupply() (*big.Int, error) {
	return _LiquidStaking.Contract.TotalSupply(&_LiquidStaking.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_LiquidStaking *LiquidStakingCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LiquidStaking.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_LiquidStaking *LiquidStakingSession) Version() (string, error) {
	return _LiquidStaking.Contract.Version(&_LiquidStaking.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_LiquidStaking *LiquidStakingCallerSession) Version() (string, error) {
	return _LiquidStaking.Contract.Version(&_LiquidStaking.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LiquidStaking *LiquidStakingTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LiquidStaking.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LiquidStaking *LiquidStakingSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LiquidStaking.Contract.Approve(&_LiquidStaking.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LiquidStaking *LiquidStakingTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LiquidStaking.Contract.Approve(&_LiquidStaking.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_LiquidStaking *LiquidStakingTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LiquidStaking.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_LiquidStaking *LiquidStakingSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LiquidStaking.Contract.DecreaseAllowance(&_LiquidStaking.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_LiquidStaking *LiquidStakingTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LiquidStaking.Contract.DecreaseAllowance(&_LiquidStaking.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_LiquidStaking *LiquidStakingTransactor) Deposit(opts *bind.TransactOpts, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _LiquidStaking.contract.Transact(opts, "deposit", assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_LiquidStaking *LiquidStakingSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _LiquidStaking.Contract.Deposit(&_LiquidStaking.TransactOpts, assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_LiquidStaking *LiquidStakingTransactorSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _LiquidStaking.Contract.Deposit(&_LiquidStaking.TransactOpts, assets, receiver)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LiquidStaking *LiquidStakingTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LiquidStaking.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LiquidStaking *LiquidStakingSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LiquidStaking.Contract.GrantRole(&_LiquidStaking.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LiquidStaking *LiquidStakingTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LiquidStaking.Contract.GrantRole(&_LiquidStaking.TransactOpts, role, account)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_LiquidStaking *LiquidStakingTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LiquidStaking.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_LiquidStaking *LiquidStakingSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LiquidStaking.Contract.IncreaseAllowance(&_LiquidStaking.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_LiquidStaking *LiquidStakingTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LiquidStaking.Contract.IncreaseAllowance(&_LiquidStaking.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _wFIL, address _resolver) returns()
func (_LiquidStaking *LiquidStakingTransactor) Initialize(opts *bind.TransactOpts, _wFIL common.Address, _resolver common.Address) (*types.Transaction, error) {
	return _LiquidStaking.contract.Transact(opts, "initialize", _wFIL, _resolver)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _wFIL, address _resolver) returns()
func (_LiquidStaking *LiquidStakingSession) Initialize(_wFIL common.Address, _resolver common.Address) (*types.Transaction, error) {
	return _LiquidStaking.Contract.Initialize(&_LiquidStaking.TransactOpts, _wFIL, _resolver)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _wFIL, address _resolver) returns()
func (_LiquidStaking *LiquidStakingTransactorSession) Initialize(_wFIL common.Address, _resolver common.Address) (*types.Transaction, error) {
	return _LiquidStaking.Contract.Initialize(&_LiquidStaking.TransactOpts, _wFIL, _resolver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_LiquidStaking *LiquidStakingTransactor) Mint(opts *bind.TransactOpts, shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _LiquidStaking.contract.Transact(opts, "mint", shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_LiquidStaking *LiquidStakingSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _LiquidStaking.Contract.Mint(&_LiquidStaking.TransactOpts, shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_LiquidStaking *LiquidStakingTransactorSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _LiquidStaking.Contract.Mint(&_LiquidStaking.TransactOpts, shares, receiver)
}

// Pledge is a paid mutator transaction binding the contract method 0x7326c9c0.
//
// Solidity: function pledge(uint256 amount) returns()
func (_LiquidStaking *LiquidStakingTransactor) Pledge(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _LiquidStaking.contract.Transact(opts, "pledge", amount)
}

// Pledge is a paid mutator transaction binding the contract method 0x7326c9c0.
//
// Solidity: function pledge(uint256 amount) returns()
func (_LiquidStaking *LiquidStakingSession) Pledge(amount *big.Int) (*types.Transaction, error) {
	return _LiquidStaking.Contract.Pledge(&_LiquidStaking.TransactOpts, amount)
}

// Pledge is a paid mutator transaction binding the contract method 0x7326c9c0.
//
// Solidity: function pledge(uint256 amount) returns()
func (_LiquidStaking *LiquidStakingTransactorSession) Pledge(amount *big.Int) (*types.Transaction, error) {
	return _LiquidStaking.Contract.Pledge(&_LiquidStaking.TransactOpts, amount)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_LiquidStaking *LiquidStakingTransactor) Redeem(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _LiquidStaking.contract.Transact(opts, "redeem", shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_LiquidStaking *LiquidStakingSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _LiquidStaking.Contract.Redeem(&_LiquidStaking.TransactOpts, shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_LiquidStaking *LiquidStakingTransactorSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _LiquidStaking.Contract.Redeem(&_LiquidStaking.TransactOpts, shares, receiver, owner)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_LiquidStaking *LiquidStakingTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LiquidStaking.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_LiquidStaking *LiquidStakingSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LiquidStaking.Contract.RenounceRole(&_LiquidStaking.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_LiquidStaking *LiquidStakingTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LiquidStaking.Contract.RenounceRole(&_LiquidStaking.TransactOpts, role, account)
}

// RepayPledge is a paid mutator transaction binding the contract method 0xd2a0460c.
//
// Solidity: function repayPledge(uint256 amount) returns()
func (_LiquidStaking *LiquidStakingTransactor) RepayPledge(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _LiquidStaking.contract.Transact(opts, "repayPledge", amount)
}

// RepayPledge is a paid mutator transaction binding the contract method 0xd2a0460c.
//
// Solidity: function repayPledge(uint256 amount) returns()
func (_LiquidStaking *LiquidStakingSession) RepayPledge(amount *big.Int) (*types.Transaction, error) {
	return _LiquidStaking.Contract.RepayPledge(&_LiquidStaking.TransactOpts, amount)
}

// RepayPledge is a paid mutator transaction binding the contract method 0xd2a0460c.
//
// Solidity: function repayPledge(uint256 amount) returns()
func (_LiquidStaking *LiquidStakingTransactorSession) RepayPledge(amount *big.Int) (*types.Transaction, error) {
	return _LiquidStaking.Contract.RepayPledge(&_LiquidStaking.TransactOpts, amount)
}

// Restake is a paid mutator transaction binding the contract method 0x413a885b.
//
// Solidity: function restake(uint256 assets, address receiver) returns(uint256 shares)
func (_LiquidStaking *LiquidStakingTransactor) Restake(opts *bind.TransactOpts, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _LiquidStaking.contract.Transact(opts, "restake", assets, receiver)
}

// Restake is a paid mutator transaction binding the contract method 0x413a885b.
//
// Solidity: function restake(uint256 assets, address receiver) returns(uint256 shares)
func (_LiquidStaking *LiquidStakingSession) Restake(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _LiquidStaking.Contract.Restake(&_LiquidStaking.TransactOpts, assets, receiver)
}

// Restake is a paid mutator transaction binding the contract method 0x413a885b.
//
// Solidity: function restake(uint256 assets, address receiver) returns(uint256 shares)
func (_LiquidStaking *LiquidStakingTransactorSession) Restake(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _LiquidStaking.Contract.Restake(&_LiquidStaking.TransactOpts, assets, receiver)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LiquidStaking *LiquidStakingTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LiquidStaking.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LiquidStaking *LiquidStakingSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LiquidStaking.Contract.RevokeRole(&_LiquidStaking.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LiquidStaking *LiquidStakingTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LiquidStaking.Contract.RevokeRole(&_LiquidStaking.TransactOpts, role, account)
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() payable returns(uint256 shares)
func (_LiquidStaking *LiquidStakingTransactor) Stake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidStaking.contract.Transact(opts, "stake")
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() payable returns(uint256 shares)
func (_LiquidStaking *LiquidStakingSession) Stake() (*types.Transaction, error) {
	return _LiquidStaking.Contract.Stake(&_LiquidStaking.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() payable returns(uint256 shares)
func (_LiquidStaking *LiquidStakingTransactorSession) Stake() (*types.Transaction, error) {
	return _LiquidStaking.Contract.Stake(&_LiquidStaking.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_LiquidStaking *LiquidStakingTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LiquidStaking.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_LiquidStaking *LiquidStakingSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LiquidStaking.Contract.Transfer(&_LiquidStaking.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_LiquidStaking *LiquidStakingTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LiquidStaking.Contract.Transfer(&_LiquidStaking.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_LiquidStaking *LiquidStakingTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LiquidStaking.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_LiquidStaking *LiquidStakingSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LiquidStaking.Contract.TransferFrom(&_LiquidStaking.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_LiquidStaking *LiquidStakingTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LiquidStaking.Contract.TransferFrom(&_LiquidStaking.TransactOpts, from, to, amount)
}

// Unstake is a paid mutator transaction binding the contract method 0x8381e182.
//
// Solidity: function unstake(uint256 shares, address owner) returns(uint256 assets)
func (_LiquidStaking *LiquidStakingTransactor) Unstake(opts *bind.TransactOpts, shares *big.Int, owner common.Address) (*types.Transaction, error) {
	return _LiquidStaking.contract.Transact(opts, "unstake", shares, owner)
}

// Unstake is a paid mutator transaction binding the contract method 0x8381e182.
//
// Solidity: function unstake(uint256 shares, address owner) returns(uint256 assets)
func (_LiquidStaking *LiquidStakingSession) Unstake(shares *big.Int, owner common.Address) (*types.Transaction, error) {
	return _LiquidStaking.Contract.Unstake(&_LiquidStaking.TransactOpts, shares, owner)
}

// Unstake is a paid mutator transaction binding the contract method 0x8381e182.
//
// Solidity: function unstake(uint256 shares, address owner) returns(uint256 assets)
func (_LiquidStaking *LiquidStakingTransactorSession) Unstake(shares *big.Int, owner common.Address) (*types.Transaction, error) {
	return _LiquidStaking.Contract.Unstake(&_LiquidStaking.TransactOpts, shares, owner)
}

// UnstakeAssets is a paid mutator transaction binding the contract method 0xf79126d6.
//
// Solidity: function unstakeAssets(uint256 assets, address owner) returns(uint256 shares)
func (_LiquidStaking *LiquidStakingTransactor) UnstakeAssets(opts *bind.TransactOpts, assets *big.Int, owner common.Address) (*types.Transaction, error) {
	return _LiquidStaking.contract.Transact(opts, "unstakeAssets", assets, owner)
}

// UnstakeAssets is a paid mutator transaction binding the contract method 0xf79126d6.
//
// Solidity: function unstakeAssets(uint256 assets, address owner) returns(uint256 shares)
func (_LiquidStaking *LiquidStakingSession) UnstakeAssets(assets *big.Int, owner common.Address) (*types.Transaction, error) {
	return _LiquidStaking.Contract.UnstakeAssets(&_LiquidStaking.TransactOpts, assets, owner)
}

// UnstakeAssets is a paid mutator transaction binding the contract method 0xf79126d6.
//
// Solidity: function unstakeAssets(uint256 assets, address owner) returns(uint256 shares)
func (_LiquidStaking *LiquidStakingTransactorSession) UnstakeAssets(assets *big.Int, owner common.Address) (*types.Transaction, error) {
	return _LiquidStaking.Contract.UnstakeAssets(&_LiquidStaking.TransactOpts, assets, owner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_LiquidStaking *LiquidStakingTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _LiquidStaking.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_LiquidStaking *LiquidStakingSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _LiquidStaking.Contract.UpgradeTo(&_LiquidStaking.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_LiquidStaking *LiquidStakingTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _LiquidStaking.Contract.UpgradeTo(&_LiquidStaking.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LiquidStaking *LiquidStakingTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LiquidStaking.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LiquidStaking *LiquidStakingSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LiquidStaking.Contract.UpgradeToAndCall(&_LiquidStaking.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LiquidStaking *LiquidStakingTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LiquidStaking.Contract.UpgradeToAndCall(&_LiquidStaking.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_LiquidStaking *LiquidStakingTransactor) Withdraw(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _LiquidStaking.contract.Transact(opts, "withdraw", assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_LiquidStaking *LiquidStakingSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _LiquidStaking.Contract.Withdraw(&_LiquidStaking.TransactOpts, assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_LiquidStaking *LiquidStakingTransactorSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _LiquidStaking.Contract.Withdraw(&_LiquidStaking.TransactOpts, assets, receiver, owner)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_LiquidStaking *LiquidStakingTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _LiquidStaking.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_LiquidStaking *LiquidStakingSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _LiquidStaking.Contract.Fallback(&_LiquidStaking.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_LiquidStaking *LiquidStakingTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _LiquidStaking.Contract.Fallback(&_LiquidStaking.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LiquidStaking *LiquidStakingTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidStaking.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LiquidStaking *LiquidStakingSession) Receive() (*types.Transaction, error) {
	return _LiquidStaking.Contract.Receive(&_LiquidStaking.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LiquidStaking *LiquidStakingTransactorSession) Receive() (*types.Transaction, error) {
	return _LiquidStaking.Contract.Receive(&_LiquidStaking.TransactOpts)
}

// LiquidStakingAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the LiquidStaking contract.
type LiquidStakingAdminChangedIterator struct {
	Event *LiquidStakingAdminChanged // Event containing the contract specifics and raw log

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
func (it *LiquidStakingAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidStakingAdminChanged)
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
		it.Event = new(LiquidStakingAdminChanged)
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
func (it *LiquidStakingAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidStakingAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidStakingAdminChanged represents a AdminChanged event raised by the LiquidStaking contract.
type LiquidStakingAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_LiquidStaking *LiquidStakingFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*LiquidStakingAdminChangedIterator, error) {

	logs, sub, err := _LiquidStaking.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &LiquidStakingAdminChangedIterator{contract: _LiquidStaking.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_LiquidStaking *LiquidStakingFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *LiquidStakingAdminChanged) (event.Subscription, error) {

	logs, sub, err := _LiquidStaking.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidStakingAdminChanged)
				if err := _LiquidStaking.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_LiquidStaking *LiquidStakingFilterer) ParseAdminChanged(log types.Log) (*LiquidStakingAdminChanged, error) {
	event := new(LiquidStakingAdminChanged)
	if err := _LiquidStaking.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidStakingApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the LiquidStaking contract.
type LiquidStakingApprovalIterator struct {
	Event *LiquidStakingApproval // Event containing the contract specifics and raw log

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
func (it *LiquidStakingApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidStakingApproval)
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
		it.Event = new(LiquidStakingApproval)
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
func (it *LiquidStakingApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidStakingApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidStakingApproval represents a Approval event raised by the LiquidStaking contract.
type LiquidStakingApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LiquidStaking *LiquidStakingFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*LiquidStakingApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LiquidStaking.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &LiquidStakingApprovalIterator{contract: _LiquidStaking.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LiquidStaking *LiquidStakingFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LiquidStakingApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LiquidStaking.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidStakingApproval)
				if err := _LiquidStaking.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LiquidStaking *LiquidStakingFilterer) ParseApproval(log types.Log) (*LiquidStakingApproval, error) {
	event := new(LiquidStakingApproval)
	if err := _LiquidStaking.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidStakingBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the LiquidStaking contract.
type LiquidStakingBeaconUpgradedIterator struct {
	Event *LiquidStakingBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *LiquidStakingBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidStakingBeaconUpgraded)
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
		it.Event = new(LiquidStakingBeaconUpgraded)
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
func (it *LiquidStakingBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidStakingBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidStakingBeaconUpgraded represents a BeaconUpgraded event raised by the LiquidStaking contract.
type LiquidStakingBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_LiquidStaking *LiquidStakingFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*LiquidStakingBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _LiquidStaking.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &LiquidStakingBeaconUpgradedIterator{contract: _LiquidStaking.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_LiquidStaking *LiquidStakingFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *LiquidStakingBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _LiquidStaking.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidStakingBeaconUpgraded)
				if err := _LiquidStaking.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_LiquidStaking *LiquidStakingFilterer) ParseBeaconUpgraded(log types.Log) (*LiquidStakingBeaconUpgraded, error) {
	event := new(LiquidStakingBeaconUpgraded)
	if err := _LiquidStaking.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidStakingDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the LiquidStaking contract.
type LiquidStakingDepositIterator struct {
	Event *LiquidStakingDeposit // Event containing the contract specifics and raw log

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
func (it *LiquidStakingDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidStakingDeposit)
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
		it.Event = new(LiquidStakingDeposit)
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
func (it *LiquidStakingDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidStakingDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidStakingDeposit represents a Deposit event raised by the LiquidStaking contract.
type LiquidStakingDeposit struct {
	Sender common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_LiquidStaking *LiquidStakingFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, owner []common.Address) (*LiquidStakingDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _LiquidStaking.contract.FilterLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &LiquidStakingDepositIterator{contract: _LiquidStaking.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_LiquidStaking *LiquidStakingFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *LiquidStakingDeposit, sender []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _LiquidStaking.contract.WatchLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidStakingDeposit)
				if err := _LiquidStaking.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_LiquidStaking *LiquidStakingFilterer) ParseDeposit(log types.Log) (*LiquidStakingDeposit, error) {
	event := new(LiquidStakingDeposit)
	if err := _LiquidStaking.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidStakingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the LiquidStaking contract.
type LiquidStakingInitializedIterator struct {
	Event *LiquidStakingInitialized // Event containing the contract specifics and raw log

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
func (it *LiquidStakingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidStakingInitialized)
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
		it.Event = new(LiquidStakingInitialized)
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
func (it *LiquidStakingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidStakingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidStakingInitialized represents a Initialized event raised by the LiquidStaking contract.
type LiquidStakingInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LiquidStaking *LiquidStakingFilterer) FilterInitialized(opts *bind.FilterOpts) (*LiquidStakingInitializedIterator, error) {

	logs, sub, err := _LiquidStaking.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LiquidStakingInitializedIterator{contract: _LiquidStaking.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LiquidStaking *LiquidStakingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LiquidStakingInitialized) (event.Subscription, error) {

	logs, sub, err := _LiquidStaking.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidStakingInitialized)
				if err := _LiquidStaking.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_LiquidStaking *LiquidStakingFilterer) ParseInitialized(log types.Log) (*LiquidStakingInitialized, error) {
	event := new(LiquidStakingInitialized)
	if err := _LiquidStaking.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidStakingPledgeIterator is returned from FilterPledge and is used to iterate over the raw logs and unpacked data for Pledge events raised by the LiquidStaking contract.
type LiquidStakingPledgeIterator struct {
	Event *LiquidStakingPledge // Event containing the contract specifics and raw log

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
func (it *LiquidStakingPledgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidStakingPledge)
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
		it.Event = new(LiquidStakingPledge)
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
func (it *LiquidStakingPledgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidStakingPledgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidStakingPledge represents a Pledge event raised by the LiquidStaking contract.
type LiquidStakingPledge struct {
	OwnerId uint64
	MinerId uint64
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPledge is a free log retrieval operation binding the contract event 0x15453566d67df7c06fb80d23b2fad96f5f14cefd285a97fa5a2c96bfdfed5ad5.
//
// Solidity: event Pledge(uint64 ownerId, uint64 minerId, uint256 amount)
func (_LiquidStaking *LiquidStakingFilterer) FilterPledge(opts *bind.FilterOpts) (*LiquidStakingPledgeIterator, error) {

	logs, sub, err := _LiquidStaking.contract.FilterLogs(opts, "Pledge")
	if err != nil {
		return nil, err
	}
	return &LiquidStakingPledgeIterator{contract: _LiquidStaking.contract, event: "Pledge", logs: logs, sub: sub}, nil
}

// WatchPledge is a free log subscription operation binding the contract event 0x15453566d67df7c06fb80d23b2fad96f5f14cefd285a97fa5a2c96bfdfed5ad5.
//
// Solidity: event Pledge(uint64 ownerId, uint64 minerId, uint256 amount)
func (_LiquidStaking *LiquidStakingFilterer) WatchPledge(opts *bind.WatchOpts, sink chan<- *LiquidStakingPledge) (event.Subscription, error) {

	logs, sub, err := _LiquidStaking.contract.WatchLogs(opts, "Pledge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidStakingPledge)
				if err := _LiquidStaking.contract.UnpackLog(event, "Pledge", log); err != nil {
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

// ParsePledge is a log parse operation binding the contract event 0x15453566d67df7c06fb80d23b2fad96f5f14cefd285a97fa5a2c96bfdfed5ad5.
//
// Solidity: event Pledge(uint64 ownerId, uint64 minerId, uint256 amount)
func (_LiquidStaking *LiquidStakingFilterer) ParsePledge(log types.Log) (*LiquidStakingPledge, error) {
	event := new(LiquidStakingPledge)
	if err := _LiquidStaking.contract.UnpackLog(event, "Pledge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidStakingPledgeRepaymentIterator is returned from FilterPledgeRepayment and is used to iterate over the raw logs and unpacked data for PledgeRepayment events raised by the LiquidStaking contract.
type LiquidStakingPledgeRepaymentIterator struct {
	Event *LiquidStakingPledgeRepayment // Event containing the contract specifics and raw log

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
func (it *LiquidStakingPledgeRepaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidStakingPledgeRepayment)
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
		it.Event = new(LiquidStakingPledgeRepayment)
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
func (it *LiquidStakingPledgeRepaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidStakingPledgeRepaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidStakingPledgeRepayment represents a PledgeRepayment event raised by the LiquidStaking contract.
type LiquidStakingPledgeRepayment struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPledgeRepayment is a free log retrieval operation binding the contract event 0x455b377f92ceb4087ffafa7984aa323a5ffeb499175000372e06198c0e45b9a4.
//
// Solidity: event PledgeRepayment(uint256 amount)
func (_LiquidStaking *LiquidStakingFilterer) FilterPledgeRepayment(opts *bind.FilterOpts) (*LiquidStakingPledgeRepaymentIterator, error) {

	logs, sub, err := _LiquidStaking.contract.FilterLogs(opts, "PledgeRepayment")
	if err != nil {
		return nil, err
	}
	return &LiquidStakingPledgeRepaymentIterator{contract: _LiquidStaking.contract, event: "PledgeRepayment", logs: logs, sub: sub}, nil
}

// WatchPledgeRepayment is a free log subscription operation binding the contract event 0x455b377f92ceb4087ffafa7984aa323a5ffeb499175000372e06198c0e45b9a4.
//
// Solidity: event PledgeRepayment(uint256 amount)
func (_LiquidStaking *LiquidStakingFilterer) WatchPledgeRepayment(opts *bind.WatchOpts, sink chan<- *LiquidStakingPledgeRepayment) (event.Subscription, error) {

	logs, sub, err := _LiquidStaking.contract.WatchLogs(opts, "PledgeRepayment")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidStakingPledgeRepayment)
				if err := _LiquidStaking.contract.UnpackLog(event, "PledgeRepayment", log); err != nil {
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

// ParsePledgeRepayment is a log parse operation binding the contract event 0x455b377f92ceb4087ffafa7984aa323a5ffeb499175000372e06198c0e45b9a4.
//
// Solidity: event PledgeRepayment(uint256 amount)
func (_LiquidStaking *LiquidStakingFilterer) ParsePledgeRepayment(log types.Log) (*LiquidStakingPledgeRepayment, error) {
	event := new(LiquidStakingPledgeRepayment)
	if err := _LiquidStaking.contract.UnpackLog(event, "PledgeRepayment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidStakingRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the LiquidStaking contract.
type LiquidStakingRoleAdminChangedIterator struct {
	Event *LiquidStakingRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *LiquidStakingRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidStakingRoleAdminChanged)
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
		it.Event = new(LiquidStakingRoleAdminChanged)
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
func (it *LiquidStakingRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidStakingRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidStakingRoleAdminChanged represents a RoleAdminChanged event raised by the LiquidStaking contract.
type LiquidStakingRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LiquidStaking *LiquidStakingFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*LiquidStakingRoleAdminChangedIterator, error) {

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

	logs, sub, err := _LiquidStaking.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &LiquidStakingRoleAdminChangedIterator{contract: _LiquidStaking.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LiquidStaking *LiquidStakingFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *LiquidStakingRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _LiquidStaking.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidStakingRoleAdminChanged)
				if err := _LiquidStaking.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_LiquidStaking *LiquidStakingFilterer) ParseRoleAdminChanged(log types.Log) (*LiquidStakingRoleAdminChanged, error) {
	event := new(LiquidStakingRoleAdminChanged)
	if err := _LiquidStaking.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidStakingRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the LiquidStaking contract.
type LiquidStakingRoleGrantedIterator struct {
	Event *LiquidStakingRoleGranted // Event containing the contract specifics and raw log

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
func (it *LiquidStakingRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidStakingRoleGranted)
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
		it.Event = new(LiquidStakingRoleGranted)
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
func (it *LiquidStakingRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidStakingRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidStakingRoleGranted represents a RoleGranted event raised by the LiquidStaking contract.
type LiquidStakingRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LiquidStaking *LiquidStakingFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LiquidStakingRoleGrantedIterator, error) {

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

	logs, sub, err := _LiquidStaking.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LiquidStakingRoleGrantedIterator{contract: _LiquidStaking.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LiquidStaking *LiquidStakingFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *LiquidStakingRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _LiquidStaking.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidStakingRoleGranted)
				if err := _LiquidStaking.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_LiquidStaking *LiquidStakingFilterer) ParseRoleGranted(log types.Log) (*LiquidStakingRoleGranted, error) {
	event := new(LiquidStakingRoleGranted)
	if err := _LiquidStaking.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidStakingRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the LiquidStaking contract.
type LiquidStakingRoleRevokedIterator struct {
	Event *LiquidStakingRoleRevoked // Event containing the contract specifics and raw log

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
func (it *LiquidStakingRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidStakingRoleRevoked)
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
		it.Event = new(LiquidStakingRoleRevoked)
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
func (it *LiquidStakingRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidStakingRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidStakingRoleRevoked represents a RoleRevoked event raised by the LiquidStaking contract.
type LiquidStakingRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LiquidStaking *LiquidStakingFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LiquidStakingRoleRevokedIterator, error) {

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

	logs, sub, err := _LiquidStaking.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LiquidStakingRoleRevokedIterator{contract: _LiquidStaking.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LiquidStaking *LiquidStakingFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *LiquidStakingRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _LiquidStaking.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidStakingRoleRevoked)
				if err := _LiquidStaking.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_LiquidStaking *LiquidStakingFilterer) ParseRoleRevoked(log types.Log) (*LiquidStakingRoleRevoked, error) {
	event := new(LiquidStakingRoleRevoked)
	if err := _LiquidStaking.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidStakingStakeIterator is returned from FilterStake and is used to iterate over the raw logs and unpacked data for Stake events raised by the LiquidStaking contract.
type LiquidStakingStakeIterator struct {
	Event *LiquidStakingStake // Event containing the contract specifics and raw log

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
func (it *LiquidStakingStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidStakingStake)
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
		it.Event = new(LiquidStakingStake)
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
func (it *LiquidStakingStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidStakingStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidStakingStake represents a Stake event raised by the LiquidStaking contract.
type LiquidStakingStake struct {
	User   common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStake is a free log retrieval operation binding the contract event 0x63602d0ecc7b3a0ef7ff1a116e23056662d64280355ba8031b6d0d767c4b4458.
//
// Solidity: event Stake(address indexed user, address indexed owner, uint256 assets, uint256 shares)
func (_LiquidStaking *LiquidStakingFilterer) FilterStake(opts *bind.FilterOpts, user []common.Address, owner []common.Address) (*LiquidStakingStakeIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _LiquidStaking.contract.FilterLogs(opts, "Stake", userRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &LiquidStakingStakeIterator{contract: _LiquidStaking.contract, event: "Stake", logs: logs, sub: sub}, nil
}

// WatchStake is a free log subscription operation binding the contract event 0x63602d0ecc7b3a0ef7ff1a116e23056662d64280355ba8031b6d0d767c4b4458.
//
// Solidity: event Stake(address indexed user, address indexed owner, uint256 assets, uint256 shares)
func (_LiquidStaking *LiquidStakingFilterer) WatchStake(opts *bind.WatchOpts, sink chan<- *LiquidStakingStake, user []common.Address, owner []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _LiquidStaking.contract.WatchLogs(opts, "Stake", userRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidStakingStake)
				if err := _LiquidStaking.contract.UnpackLog(event, "Stake", log); err != nil {
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

// ParseStake is a log parse operation binding the contract event 0x63602d0ecc7b3a0ef7ff1a116e23056662d64280355ba8031b6d0d767c4b4458.
//
// Solidity: event Stake(address indexed user, address indexed owner, uint256 assets, uint256 shares)
func (_LiquidStaking *LiquidStakingFilterer) ParseStake(log types.Log) (*LiquidStakingStake, error) {
	event := new(LiquidStakingStake)
	if err := _LiquidStaking.contract.UnpackLog(event, "Stake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidStakingTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the LiquidStaking contract.
type LiquidStakingTransferIterator struct {
	Event *LiquidStakingTransfer // Event containing the contract specifics and raw log

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
func (it *LiquidStakingTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidStakingTransfer)
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
		it.Event = new(LiquidStakingTransfer)
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
func (it *LiquidStakingTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidStakingTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidStakingTransfer represents a Transfer event raised by the LiquidStaking contract.
type LiquidStakingTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LiquidStaking *LiquidStakingFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LiquidStakingTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LiquidStaking.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LiquidStakingTransferIterator{contract: _LiquidStaking.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LiquidStaking *LiquidStakingFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LiquidStakingTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LiquidStaking.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidStakingTransfer)
				if err := _LiquidStaking.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LiquidStaking *LiquidStakingFilterer) ParseTransfer(log types.Log) (*LiquidStakingTransfer, error) {
	event := new(LiquidStakingTransfer)
	if err := _LiquidStaking.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidStakingUnstakedIterator is returned from FilterUnstaked and is used to iterate over the raw logs and unpacked data for Unstaked events raised by the LiquidStaking contract.
type LiquidStakingUnstakedIterator struct {
	Event *LiquidStakingUnstaked // Event containing the contract specifics and raw log

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
func (it *LiquidStakingUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidStakingUnstaked)
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
		it.Event = new(LiquidStakingUnstaked)
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
func (it *LiquidStakingUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidStakingUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidStakingUnstaked represents a Unstaked event raised by the LiquidStaking contract.
type LiquidStakingUnstaked struct {
	User   common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnstaked is a free log retrieval operation binding the contract event 0x06cc7e90b4f2b554a9614b0caa84f909f3498c820ae47c731f490c28c07f7d3b.
//
// Solidity: event Unstaked(address indexed user, address indexed owner, uint256 assets, uint256 shares)
func (_LiquidStaking *LiquidStakingFilterer) FilterUnstaked(opts *bind.FilterOpts, user []common.Address, owner []common.Address) (*LiquidStakingUnstakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _LiquidStaking.contract.FilterLogs(opts, "Unstaked", userRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &LiquidStakingUnstakedIterator{contract: _LiquidStaking.contract, event: "Unstaked", logs: logs, sub: sub}, nil
}

// WatchUnstaked is a free log subscription operation binding the contract event 0x06cc7e90b4f2b554a9614b0caa84f909f3498c820ae47c731f490c28c07f7d3b.
//
// Solidity: event Unstaked(address indexed user, address indexed owner, uint256 assets, uint256 shares)
func (_LiquidStaking *LiquidStakingFilterer) WatchUnstaked(opts *bind.WatchOpts, sink chan<- *LiquidStakingUnstaked, user []common.Address, owner []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _LiquidStaking.contract.WatchLogs(opts, "Unstaked", userRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidStakingUnstaked)
				if err := _LiquidStaking.contract.UnpackLog(event, "Unstaked", log); err != nil {
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

// ParseUnstaked is a log parse operation binding the contract event 0x06cc7e90b4f2b554a9614b0caa84f909f3498c820ae47c731f490c28c07f7d3b.
//
// Solidity: event Unstaked(address indexed user, address indexed owner, uint256 assets, uint256 shares)
func (_LiquidStaking *LiquidStakingFilterer) ParseUnstaked(log types.Log) (*LiquidStakingUnstaked, error) {
	event := new(LiquidStakingUnstaked)
	if err := _LiquidStaking.contract.UnpackLog(event, "Unstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidStakingUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the LiquidStaking contract.
type LiquidStakingUpgradedIterator struct {
	Event *LiquidStakingUpgraded // Event containing the contract specifics and raw log

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
func (it *LiquidStakingUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidStakingUpgraded)
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
		it.Event = new(LiquidStakingUpgraded)
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
func (it *LiquidStakingUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidStakingUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidStakingUpgraded represents a Upgraded event raised by the LiquidStaking contract.
type LiquidStakingUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_LiquidStaking *LiquidStakingFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*LiquidStakingUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _LiquidStaking.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &LiquidStakingUpgradedIterator{contract: _LiquidStaking.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_LiquidStaking *LiquidStakingFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *LiquidStakingUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _LiquidStaking.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidStakingUpgraded)
				if err := _LiquidStaking.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_LiquidStaking *LiquidStakingFilterer) ParseUpgraded(log types.Log) (*LiquidStakingUpgraded, error) {
	event := new(LiquidStakingUpgraded)
	if err := _LiquidStaking.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidStakingWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the LiquidStaking contract.
type LiquidStakingWithdrawIterator struct {
	Event *LiquidStakingWithdraw // Event containing the contract specifics and raw log

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
func (it *LiquidStakingWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidStakingWithdraw)
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
		it.Event = new(LiquidStakingWithdraw)
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
func (it *LiquidStakingWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidStakingWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidStakingWithdraw represents a Withdraw event raised by the LiquidStaking contract.
type LiquidStakingWithdraw struct {
	Sender   common.Address
	Receiver common.Address
	Owner    common.Address
	Assets   *big.Int
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_LiquidStaking *LiquidStakingFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address, owner []common.Address) (*LiquidStakingWithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _LiquidStaking.contract.FilterLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &LiquidStakingWithdrawIterator{contract: _LiquidStaking.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_LiquidStaking *LiquidStakingFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *LiquidStakingWithdraw, sender []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _LiquidStaking.contract.WatchLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidStakingWithdraw)
				if err := _LiquidStaking.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_LiquidStaking *LiquidStakingFilterer) ParseWithdraw(log types.Log) (*LiquidStakingWithdraw, error) {
	event := new(LiquidStakingWithdraw)
	if err := _LiquidStaking.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
