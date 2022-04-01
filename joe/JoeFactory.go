// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package joe

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

// JoeMetaData contains all meta data concerning the Joe contract.
var JoeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeToSetter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPairsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeToSetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pairCodeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeTo\",\"type\":\"address\"}],\"name\":\"setFeeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeToSetter\",\"type\":\"address\"}],\"name\":\"setFeeToSetter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_migrator\",\"type\":\"address\"}],\"name\":\"setMigrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// JoeABI is the input ABI used to generate the binding from.
// Deprecated: Use JoeMetaData.ABI instead.
var JoeABI = JoeMetaData.ABI

// Joe is an auto generated Go binding around an Ethereum contract.
type Joe struct {
	JoeCaller     // Read-only binding to the contract
	JoeTransactor // Write-only binding to the contract
	JoeFilterer   // Log filterer for contract events
}

// JoeCaller is an auto generated read-only Go binding around an Ethereum contract.
type JoeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JoeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JoeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JoeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JoeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JoeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JoeSession struct {
	Contract     *Joe              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JoeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JoeCallerSession struct {
	Contract *JoeCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// JoeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JoeTransactorSession struct {
	Contract     *JoeTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JoeRaw is an auto generated low-level Go binding around an Ethereum contract.
type JoeRaw struct {
	Contract *Joe // Generic contract binding to access the raw methods on
}

// JoeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JoeCallerRaw struct {
	Contract *JoeCaller // Generic read-only contract binding to access the raw methods on
}

// JoeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JoeTransactorRaw struct {
	Contract *JoeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJoe creates a new instance of Joe, bound to a specific deployed contract.
func NewJoe(address common.Address, backend bind.ContractBackend) (*Joe, error) {
	contract, err := bindJoe(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Joe{JoeCaller: JoeCaller{contract: contract}, JoeTransactor: JoeTransactor{contract: contract}, JoeFilterer: JoeFilterer{contract: contract}}, nil
}

// NewJoeCaller creates a new read-only instance of Joe, bound to a specific deployed contract.
func NewJoeCaller(address common.Address, caller bind.ContractCaller) (*JoeCaller, error) {
	contract, err := bindJoe(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JoeCaller{contract: contract}, nil
}

// NewJoeTransactor creates a new write-only instance of Joe, bound to a specific deployed contract.
func NewJoeTransactor(address common.Address, transactor bind.ContractTransactor) (*JoeTransactor, error) {
	contract, err := bindJoe(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JoeTransactor{contract: contract}, nil
}

// NewJoeFilterer creates a new log filterer instance of Joe, bound to a specific deployed contract.
func NewJoeFilterer(address common.Address, filterer bind.ContractFilterer) (*JoeFilterer, error) {
	contract, err := bindJoe(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JoeFilterer{contract: contract}, nil
}

// bindJoe binds a generic wrapper to an already deployed contract.
func bindJoe(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(JoeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Joe *JoeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Joe.Contract.JoeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Joe *JoeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Joe.Contract.JoeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Joe *JoeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Joe.Contract.JoeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Joe *JoeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Joe.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Joe *JoeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Joe.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Joe *JoeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Joe.Contract.contract.Transact(opts, method, params...)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_Joe *JoeCaller) AllPairs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Joe.contract.Call(opts, &out, "allPairs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_Joe *JoeSession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _Joe.Contract.AllPairs(&_Joe.CallOpts, arg0)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_Joe *JoeCallerSession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _Joe.Contract.AllPairs(&_Joe.CallOpts, arg0)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_Joe *JoeCaller) AllPairsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Joe.contract.Call(opts, &out, "allPairsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_Joe *JoeSession) AllPairsLength() (*big.Int, error) {
	return _Joe.Contract.AllPairsLength(&_Joe.CallOpts)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_Joe *JoeCallerSession) AllPairsLength() (*big.Int, error) {
	return _Joe.Contract.AllPairsLength(&_Joe.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_Joe *JoeCaller) FeeTo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Joe.contract.Call(opts, &out, "feeTo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_Joe *JoeSession) FeeTo() (common.Address, error) {
	return _Joe.Contract.FeeTo(&_Joe.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_Joe *JoeCallerSession) FeeTo() (common.Address, error) {
	return _Joe.Contract.FeeTo(&_Joe.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_Joe *JoeCaller) FeeToSetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Joe.contract.Call(opts, &out, "feeToSetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_Joe *JoeSession) FeeToSetter() (common.Address, error) {
	return _Joe.Contract.FeeToSetter(&_Joe.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_Joe *JoeCallerSession) FeeToSetter() (common.Address, error) {
	return _Joe.Contract.FeeToSetter(&_Joe.CallOpts)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_Joe *JoeCaller) GetPair(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Joe.contract.Call(opts, &out, "getPair", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_Joe *JoeSession) GetPair(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _Joe.Contract.GetPair(&_Joe.CallOpts, arg0, arg1)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_Joe *JoeCallerSession) GetPair(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _Joe.Contract.GetPair(&_Joe.CallOpts, arg0, arg1)
}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_Joe *JoeCaller) Migrator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Joe.contract.Call(opts, &out, "migrator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_Joe *JoeSession) Migrator() (common.Address, error) {
	return _Joe.Contract.Migrator(&_Joe.CallOpts)
}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_Joe *JoeCallerSession) Migrator() (common.Address, error) {
	return _Joe.Contract.Migrator(&_Joe.CallOpts)
}

// PairCodeHash is a free data retrieval call binding the contract method 0x9aab9248.
//
// Solidity: function pairCodeHash() pure returns(bytes32)
func (_Joe *JoeCaller) PairCodeHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Joe.contract.Call(opts, &out, "pairCodeHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PairCodeHash is a free data retrieval call binding the contract method 0x9aab9248.
//
// Solidity: function pairCodeHash() pure returns(bytes32)
func (_Joe *JoeSession) PairCodeHash() ([32]byte, error) {
	return _Joe.Contract.PairCodeHash(&_Joe.CallOpts)
}

// PairCodeHash is a free data retrieval call binding the contract method 0x9aab9248.
//
// Solidity: function pairCodeHash() pure returns(bytes32)
func (_Joe *JoeCallerSession) PairCodeHash() ([32]byte, error) {
	return _Joe.Contract.PairCodeHash(&_Joe.CallOpts)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_Joe *JoeTransactor) CreatePair(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _Joe.contract.Transact(opts, "createPair", tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_Joe *JoeSession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _Joe.Contract.CreatePair(&_Joe.TransactOpts, tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_Joe *JoeTransactorSession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _Joe.Contract.CreatePair(&_Joe.TransactOpts, tokenA, tokenB)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_Joe *JoeTransactor) SetFeeTo(opts *bind.TransactOpts, _feeTo common.Address) (*types.Transaction, error) {
	return _Joe.contract.Transact(opts, "setFeeTo", _feeTo)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_Joe *JoeSession) SetFeeTo(_feeTo common.Address) (*types.Transaction, error) {
	return _Joe.Contract.SetFeeTo(&_Joe.TransactOpts, _feeTo)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_Joe *JoeTransactorSession) SetFeeTo(_feeTo common.Address) (*types.Transaction, error) {
	return _Joe.Contract.SetFeeTo(&_Joe.TransactOpts, _feeTo)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_Joe *JoeTransactor) SetFeeToSetter(opts *bind.TransactOpts, _feeToSetter common.Address) (*types.Transaction, error) {
	return _Joe.contract.Transact(opts, "setFeeToSetter", _feeToSetter)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_Joe *JoeSession) SetFeeToSetter(_feeToSetter common.Address) (*types.Transaction, error) {
	return _Joe.Contract.SetFeeToSetter(&_Joe.TransactOpts, _feeToSetter)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_Joe *JoeTransactorSession) SetFeeToSetter(_feeToSetter common.Address) (*types.Transaction, error) {
	return _Joe.Contract.SetFeeToSetter(&_Joe.TransactOpts, _feeToSetter)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_Joe *JoeTransactor) SetMigrator(opts *bind.TransactOpts, _migrator common.Address) (*types.Transaction, error) {
	return _Joe.contract.Transact(opts, "setMigrator", _migrator)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_Joe *JoeSession) SetMigrator(_migrator common.Address) (*types.Transaction, error) {
	return _Joe.Contract.SetMigrator(&_Joe.TransactOpts, _migrator)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_Joe *JoeTransactorSession) SetMigrator(_migrator common.Address) (*types.Transaction, error) {
	return _Joe.Contract.SetMigrator(&_Joe.TransactOpts, _migrator)
}

// JoePairCreatedIterator is returned from FilterPairCreated and is used to iterate over the raw logs and unpacked data for PairCreated events raised by the Joe contract.
type JoePairCreatedIterator struct {
	Event *JoePairCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *JoePairCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JoePairCreated)
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
		it.Event = new(JoePairCreated)
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
func (it *JoePairCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JoePairCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JoePairCreated represents a PairCreated event raised by the Joe contract.
type JoePairCreated struct {
	Token0 common.Address
	Token1 common.Address
	Pair   common.Address
	Arg3   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPairCreated is a free log retrieval operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_Joe *JoeFilterer) FilterPairCreated(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address) (*JoePairCreatedIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _Joe.contract.FilterLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return &JoePairCreatedIterator{contract: _Joe.contract, event: "PairCreated", logs: logs, sub: sub}, nil
}

// WatchPairCreated is a free log subscription operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_Joe *JoeFilterer) WatchPairCreated(opts *bind.WatchOpts, sink chan<- *JoePairCreated, token0 []common.Address, token1 []common.Address) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _Joe.contract.WatchLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JoePairCreated)
				if err := _Joe.contract.UnpackLog(event, "PairCreated", log); err != nil {
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

// ParsePairCreated is a log parse operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_Joe *JoeFilterer) ParsePairCreated(log types.Log) (*JoePairCreated, error) {
	event := new(JoePairCreated)
	if err := _Joe.contract.UnpackLog(event, "PairCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
