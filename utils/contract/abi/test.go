// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// MystructMetaData contains all meta data concerning the Mystruct contract.
var MystructMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"get_owner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MystructABI is the input ABI used to generate the binding from.
// Deprecated: Use MystructMetaData.ABI instead.
var MystructABI = MystructMetaData.ABI

// Mystruct is an auto generated Go binding around an Ethereum contract.
type Mystruct struct {
	MystructCaller     // Read-only binding to the contract
	MystructTransactor // Write-only binding to the contract
	MystructFilterer   // Log filterer for contract events
}

// MystructCaller is an auto generated read-only Go binding around an Ethereum contract.
type MystructCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MystructTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MystructTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MystructFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MystructFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MystructSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MystructSession struct {
	Contract     *Mystruct         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MystructCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MystructCallerSession struct {
	Contract *MystructCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// MystructTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MystructTransactorSession struct {
	Contract     *MystructTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MystructRaw is an auto generated low-level Go binding around an Ethereum contract.
type MystructRaw struct {
	Contract *Mystruct // Generic contract binding to access the raw methods on
}

// MystructCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MystructCallerRaw struct {
	Contract *MystructCaller // Generic read-only contract binding to access the raw methods on
}

// MystructTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MystructTransactorRaw struct {
	Contract *MystructTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMystruct creates a new instance of Mystruct, bound to a specific deployed contract.
func NewMystruct(address common.Address, backend bind.ContractBackend) (*Mystruct, error) {
	contract, err := bindMystruct(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mystruct{MystructCaller: MystructCaller{contract: contract}, MystructTransactor: MystructTransactor{contract: contract}, MystructFilterer: MystructFilterer{contract: contract}}, nil
}

// NewMystructCaller creates a new read-only instance of Mystruct, bound to a specific deployed contract.
func NewMystructCaller(address common.Address, caller bind.ContractCaller) (*MystructCaller, error) {
	contract, err := bindMystruct(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MystructCaller{contract: contract}, nil
}

// NewMystructTransactor creates a new write-only instance of Mystruct, bound to a specific deployed contract.
func NewMystructTransactor(address common.Address, transactor bind.ContractTransactor) (*MystructTransactor, error) {
	contract, err := bindMystruct(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MystructTransactor{contract: contract}, nil
}

// NewMystructFilterer creates a new log filterer instance of Mystruct, bound to a specific deployed contract.
func NewMystructFilterer(address common.Address, filterer bind.ContractFilterer) (*MystructFilterer, error) {
	contract, err := bindMystruct(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MystructFilterer{contract: contract}, nil
}

// bindMystruct binds a generic wrapper to an already deployed contract.
func bindMystruct(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MystructABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mystruct *MystructRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mystruct.Contract.MystructCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mystruct *MystructRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mystruct.Contract.MystructTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mystruct *MystructRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mystruct.Contract.MystructTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mystruct *MystructCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mystruct.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mystruct *MystructTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mystruct.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mystruct *MystructTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mystruct.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mystruct *MystructCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mystruct.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mystruct *MystructSession) Owner() (common.Address, error) {
	return _Mystruct.Contract.Owner(&_Mystruct.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mystruct *MystructCallerSession) Owner() (common.Address, error) {
	return _Mystruct.Contract.Owner(&_Mystruct.CallOpts)
}

// GetOwner is a paid mutator transaction binding the contract method 0x0ac298dc.
//
// Solidity: function get_owner() returns()
func (_Mystruct *MystructTransactor) GetOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mystruct.contract.Transact(opts, "get_owner")
}

// GetOwner is a paid mutator transaction binding the contract method 0x0ac298dc.
//
// Solidity: function get_owner() returns()
func (_Mystruct *MystructSession) GetOwner() (*types.Transaction, error) {
	return _Mystruct.Contract.GetOwner(&_Mystruct.TransactOpts)
}

// GetOwner is a paid mutator transaction binding the contract method 0x0ac298dc.
//
// Solidity: function get_owner() returns()
func (_Mystruct *MystructTransactorSession) GetOwner() (*types.Transaction, error) {
	return _Mystruct.Contract.GetOwner(&_Mystruct.TransactOpts)
}
