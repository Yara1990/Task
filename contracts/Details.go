// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SampelABI is the input ABI used to generate the binding from.
const SampelABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"getUser\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"setUser\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SampelBin is the compiled bytecode used for deploying new contracts.
const SampelBin = `0x608060405234801561001057600080fd5b506102f3806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806331feb6711461003b578063d0020fad146100fd575b600080fd5b6100e16004803603602081101561005157600080fd5b81019060208101813564010000000081111561006c57600080fd5b82018360208201111561007e57600080fd5b803590602001918460018302840111640100000000831117156100a057600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506101a3945050505050565b604080516001600160a01b039092168252519081900360200190f35b6100e16004803603602081101561011357600080fd5b81019060208101813564010000000081111561012e57600080fd5b82018360208201111561014057600080fd5b8035906020019184600183028401116401000000008311171561016257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610213945050505050565b600080826040518082805190602001908083835b602083106101d65780518252601f1990920191602091820191016101b7565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220546001600160a01b0316949350505050565b6000336000836040518082805190602001908083835b602083106102485780518252601f199092019160209182019101610229565b51815160001960209485036101000a01908116901991909116179052920194855250604051938490038101842080546001600160a01b03969096166001600160a01b031990961695909517909455505083516000928592918291840190808383602083106101d65780518252601f1990920191602091820191016101b756fea165627a7a72305820a0d09d4aaaa9ad8327eef14c185adfed15a9ebe9f07956c06771b16661994eeb0029`

// DeploySampel deploys a new Ethereum contract, binding an instance of Sampel to it.
func DeploySampel(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Sampel, error) {
	parsed, err := abi.JSON(strings.NewReader(SampelABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SampelBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Sampel{SampelCaller: SampelCaller{contract: contract}, SampelTransactor: SampelTransactor{contract: contract}, SampelFilterer: SampelFilterer{contract: contract}}, nil
}

// Sampel is an auto generated Go binding around an Ethereum contract.
type Sampel struct {
	SampelCaller     // Read-only binding to the contract
	SampelTransactor // Write-only binding to the contract
	SampelFilterer   // Log filterer for contract events
}

// SampelCaller is an auto generated read-only Go binding around an Ethereum contract.
type SampelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SampelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SampelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SampelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SampelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SampelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SampelSession struct {
	Contract     *Sampel           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SampelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SampelCallerSession struct {
	Contract *SampelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SampelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SampelTransactorSession struct {
	Contract     *SampelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SampelRaw is an auto generated low-level Go binding around an Ethereum contract.
type SampelRaw struct {
	Contract *Sampel // Generic contract binding to access the raw methods on
}

// SampelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SampelCallerRaw struct {
	Contract *SampelCaller // Generic read-only contract binding to access the raw methods on
}

// SampelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SampelTransactorRaw struct {
	Contract *SampelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSampel creates a new instance of Sampel, bound to a specific deployed contract.
func NewSampel(address common.Address, backend bind.ContractBackend) (*Sampel, error) {
	contract, err := bindSampel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sampel{SampelCaller: SampelCaller{contract: contract}, SampelTransactor: SampelTransactor{contract: contract}, SampelFilterer: SampelFilterer{contract: contract}}, nil
}

// NewSampelCaller creates a new read-only instance of Sampel, bound to a specific deployed contract.
func NewSampelCaller(address common.Address, caller bind.ContractCaller) (*SampelCaller, error) {
	contract, err := bindSampel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SampelCaller{contract: contract}, nil
}

// NewSampelTransactor creates a new write-only instance of Sampel, bound to a specific deployed contract.
func NewSampelTransactor(address common.Address, transactor bind.ContractTransactor) (*SampelTransactor, error) {
	contract, err := bindSampel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SampelTransactor{contract: contract}, nil
}

// NewSampelFilterer creates a new log filterer instance of Sampel, bound to a specific deployed contract.
func NewSampelFilterer(address common.Address, filterer bind.ContractFilterer) (*SampelFilterer, error) {
	contract, err := bindSampel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SampelFilterer{contract: contract}, nil
}

// bindSampel binds a generic wrapper to an already deployed contract.
func bindSampel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SampelABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sampel *SampelRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Sampel.Contract.SampelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sampel *SampelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sampel.Contract.SampelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sampel *SampelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sampel.Contract.SampelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sampel *SampelCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Sampel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sampel *SampelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sampel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sampel *SampelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sampel.Contract.contract.Transact(opts, method, params...)
}

// GetUser is a free data retrieval call binding the contract method 0x31feb671.
//
// Solidity: function getUser(string _name) constant returns(address)
func (_Sampel *SampelCaller) GetUser(opts *bind.CallOpts, _name string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Sampel.contract.Call(opts, out, "getUser", _name)
	return *ret0, err
}

// GetUser is a free data retrieval call binding the contract method 0x31feb671.
//
// Solidity: function getUser(string _name) constant returns(address)
func (_Sampel *SampelSession) GetUser(_name string) (common.Address, error) {
	return _Sampel.Contract.GetUser(&_Sampel.CallOpts, _name)
}

// GetUser is a free data retrieval call binding the contract method 0x31feb671.
//
// Solidity: function getUser(string _name) constant returns(address)
func (_Sampel *SampelCallerSession) GetUser(_name string) (common.Address, error) {
	return _Sampel.Contract.GetUser(&_Sampel.CallOpts, _name)
}

// SetUser is a paid mutator transaction binding the contract method 0xd0020fad.
//
// Solidity: function setUser(string _name) returns(address)
func (_Sampel *SampelTransactor) SetUser(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _Sampel.contract.Transact(opts, "setUser", _name)
}

// SetUser is a paid mutator transaction binding the contract method 0xd0020fad.
//
// Solidity: function setUser(string _name) returns(address)
func (_Sampel *SampelSession) SetUser(_name string) (*types.Transaction, error) {
	return _Sampel.Contract.SetUser(&_Sampel.TransactOpts, _name)
}

// SetUser is a paid mutator transaction binding the contract method 0xd0020fad.
//
// Solidity: function setUser(string _name) returns(address)
func (_Sampel *SampelTransactorSession) SetUser(_name string) (*types.Transaction, error) {
	return _Sampel.Contract.SetUser(&_Sampel.TransactOpts, _name)
}
