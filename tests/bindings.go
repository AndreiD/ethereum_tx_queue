// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// HELLOMetaData contains all meta data concerning the HELLO contract.
var HELLOMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"getMessage\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newMessage\",\"type\":\"string\"}],\"name\":\"setMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// HELLOABI is the input ABI used to generate the binding from.
// Deprecated: Use HELLOMetaData.ABI instead.
var HELLOABI = HELLOMetaData.ABI

// HELLO is an auto generated Go binding around an Ethereum contract.
type HELLO struct {
	HELLOCaller     // Read-only binding to the contract
	HELLOTransactor // Write-only binding to the contract
	HELLOFilterer   // Log filterer for contract events
}

// HELLOCaller is an auto generated read-only Go binding around an Ethereum contract.
type HELLOCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HELLOTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HELLOTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HELLOFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HELLOFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HELLOSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HELLOSession struct {
	Contract     *HELLO            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HELLOCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HELLOCallerSession struct {
	Contract *HELLOCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HELLOTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HELLOTransactorSession struct {
	Contract     *HELLOTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HELLORaw is an auto generated low-level Go binding around an Ethereum contract.
type HELLORaw struct {
	Contract *HELLO // Generic contract binding to access the raw methods on
}

// HELLOCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HELLOCallerRaw struct {
	Contract *HELLOCaller // Generic read-only contract binding to access the raw methods on
}

// HELLOTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HELLOTransactorRaw struct {
	Contract *HELLOTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHELLO creates a new instance of HELLO, bound to a specific deployed contract.
func NewHELLO(address common.Address, backend bind.ContractBackend) (*HELLO, error) {
	contract, err := bindHELLO(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HELLO{HELLOCaller: HELLOCaller{contract: contract}, HELLOTransactor: HELLOTransactor{contract: contract}, HELLOFilterer: HELLOFilterer{contract: contract}}, nil
}

// NewHELLOCaller creates a new read-only instance of HELLO, bound to a specific deployed contract.
func NewHELLOCaller(address common.Address, caller bind.ContractCaller) (*HELLOCaller, error) {
	contract, err := bindHELLO(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HELLOCaller{contract: contract}, nil
}

// NewHELLOTransactor creates a new write-only instance of HELLO, bound to a specific deployed contract.
func NewHELLOTransactor(address common.Address, transactor bind.ContractTransactor) (*HELLOTransactor, error) {
	contract, err := bindHELLO(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HELLOTransactor{contract: contract}, nil
}

// NewHELLOFilterer creates a new log filterer instance of HELLO, bound to a specific deployed contract.
func NewHELLOFilterer(address common.Address, filterer bind.ContractFilterer) (*HELLOFilterer, error) {
	contract, err := bindHELLO(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HELLOFilterer{contract: contract}, nil
}

// bindHELLO binds a generic wrapper to an already deployed contract.
func bindHELLO(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HELLOABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HELLO *HELLORaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HELLO.Contract.HELLOCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HELLO *HELLORaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HELLO.Contract.HELLOTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HELLO *HELLORaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HELLO.Contract.HELLOTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HELLO *HELLOCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HELLO.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HELLO *HELLOTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HELLO.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HELLO *HELLOTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HELLO.Contract.contract.Transact(opts, method, params...)
}

// GetMessage is a free data retrieval call binding the contract method 0xce6d41de.
//
// Solidity: function getMessage() view returns(string)
func (_HELLO *HELLOCaller) GetMessage(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _HELLO.contract.Call(opts, &out, "getMessage")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetMessage is a free data retrieval call binding the contract method 0xce6d41de.
//
// Solidity: function getMessage() view returns(string)
func (_HELLO *HELLOSession) GetMessage() (string, error) {
	return _HELLO.Contract.GetMessage(&_HELLO.CallOpts)
}

// GetMessage is a free data retrieval call binding the contract method 0xce6d41de.
//
// Solidity: function getMessage() view returns(string)
func (_HELLO *HELLOCallerSession) GetMessage() (string, error) {
	return _HELLO.Contract.GetMessage(&_HELLO.CallOpts)
}

// SetMessage is a paid mutator transaction binding the contract method 0x368b8772.
//
// Solidity: function setMessage(string newMessage) returns()
func (_HELLO *HELLOTransactor) SetMessage(opts *bind.TransactOpts, newMessage string) (*types.Transaction, error) {
	return _HELLO.contract.Transact(opts, "setMessage", newMessage)
}

// SetMessage is a paid mutator transaction binding the contract method 0x368b8772.
//
// Solidity: function setMessage(string newMessage) returns()
func (_HELLO *HELLOSession) SetMessage(newMessage string) (*types.Transaction, error) {
	return _HELLO.Contract.SetMessage(&_HELLO.TransactOpts, newMessage)
}

// SetMessage is a paid mutator transaction binding the contract method 0x368b8772.
//
// Solidity: function setMessage(string newMessage) returns()
func (_HELLO *HELLOTransactorSession) SetMessage(newMessage string) (*types.Transaction, error) {
	return _HELLO.Contract.SetMessage(&_HELLO.TransactOpts, newMessage)
}
