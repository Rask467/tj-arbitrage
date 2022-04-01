// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sc 

import (
	"errors"
	"math/big"
	"strings"

	"github.com/ava-labs/coreth/accounts/abi"
	"github.com/ava-labs/coreth/accounts/abi/bind"
	"github.com/ava-labs/coreth/core/types"
	"github.com/ava-labs/coreth/interfaces"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = interfaces.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SCMetaData contains all meta data concerning the SC contract.
var SCMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"joeCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"pangolinCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV2Call\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SCABI is the input ABI used to generate the binding from.
// Deprecated: Use SCMetaData.ABI instead.
var SCABI = SCMetaData.ABI

// SC is an auto generated Go binding around an Ethereum contract.
type SC struct {
	SCCaller     // Read-only binding to the contract
	SCTransactor // Write-only binding to the contract
	SCFilterer   // Log filterer for contract events
}

// SCCaller is an auto generated read-only Go binding around an Ethereum contract.
type SCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SCFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SCSession struct {
	Contract     *SC               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SCCallerSession struct {
	Contract *SCCaller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SCTransactorSession struct {
	Contract     *SCTransactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SCRaw is an auto generated low-level Go binding around an Ethereum contract.
type SCRaw struct {
	Contract *SC // Generic contract binding to access the raw methods on
}

// SCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SCCallerRaw struct {
	Contract *SCCaller // Generic read-only contract binding to access the raw methods on
}

// SCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SCTransactorRaw struct {
	Contract *SCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSC creates a new instance of SC, bound to a specific deployed contract.
func NewSC(address common.Address, backend bind.ContractBackend) (*SC, error) {
	contract, err := bindSC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SC{SCCaller: SCCaller{contract: contract}, SCTransactor: SCTransactor{contract: contract}, SCFilterer: SCFilterer{contract: contract}}, nil
}

// NewSCCaller creates a new read-only instance of SC, bound to a specific deployed contract.
func NewSCCaller(address common.Address, caller bind.ContractCaller) (*SCCaller, error) {
	contract, err := bindSC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SCCaller{contract: contract}, nil
}

// NewSCTransactor creates a new write-only instance of SC, bound to a specific deployed contract.
func NewSCTransactor(address common.Address, transactor bind.ContractTransactor) (*SCTransactor, error) {
	contract, err := bindSC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SCTransactor{contract: contract}, nil
}

// NewSCFilterer creates a new log filterer instance of SC, bound to a specific deployed contract.
func NewSCFilterer(address common.Address, filterer bind.ContractFilterer) (*SCFilterer, error) {
	contract, err := bindSC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SCFilterer{contract: contract}, nil
}

// bindSC binds a generic wrapper to an already deployed contract.
func bindSC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SCABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SC *SCRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SC.Contract.SCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SC *SCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SC.Contract.SCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SC *SCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SC.Contract.SCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SC *SCCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SC *SCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SC *SCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SC.Contract.contract.Transact(opts, method, params...)
}

// JoeCall is a paid mutator transaction binding the contract method 0xee22dd87.
//
// Solidity: function joeCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_SC *SCTransactor) JoeCall(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _SC.contract.Transact(opts, "joeCall", sender, amount0, amount1, data)
}

// JoeCall is a paid mutator transaction binding the contract method 0xee22dd87.
//
// Solidity: function joeCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_SC *SCSession) JoeCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _SC.Contract.JoeCall(&_SC.TransactOpts, sender, amount0, amount1, data)
}

// JoeCall is a paid mutator transaction binding the contract method 0xee22dd87.
//
// Solidity: function joeCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_SC *SCTransactorSession) JoeCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _SC.Contract.JoeCall(&_SC.TransactOpts, sender, amount0, amount1, data)
}

// PangolinCall is a paid mutator transaction binding the contract method 0xf8890f8e.
//
// Solidity: function pangolinCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_SC *SCTransactor) PangolinCall(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _SC.contract.Transact(opts, "pangolinCall", sender, amount0, amount1, data)
}

// PangolinCall is a paid mutator transaction binding the contract method 0xf8890f8e.
//
// Solidity: function pangolinCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_SC *SCSession) PangolinCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _SC.Contract.PangolinCall(&_SC.TransactOpts, sender, amount0, amount1, data)
}

// PangolinCall is a paid mutator transaction binding the contract method 0xf8890f8e.
//
// Solidity: function pangolinCall(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_SC *SCTransactorSession) PangolinCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _SC.Contract.PangolinCall(&_SC.TransactOpts, sender, amount0, amount1, data)
}

// UniswapV2Call is a paid mutator transaction binding the contract method 0x10d1e85c.
//
// Solidity: function uniswapV2Call(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_SC *SCTransactor) UniswapV2Call(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _SC.contract.Transact(opts, "uniswapV2Call", sender, amount0, amount1, data)
}

// UniswapV2Call is a paid mutator transaction binding the contract method 0x10d1e85c.
//
// Solidity: function uniswapV2Call(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_SC *SCSession) UniswapV2Call(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _SC.Contract.UniswapV2Call(&_SC.TransactOpts, sender, amount0, amount1, data)
}

// UniswapV2Call is a paid mutator transaction binding the contract method 0x10d1e85c.
//
// Solidity: function uniswapV2Call(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_SC *SCTransactorSession) UniswapV2Call(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _SC.Contract.UniswapV2Call(&_SC.TransactOpts, sender, amount0, amount1, data)
}
