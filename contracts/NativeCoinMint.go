// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package NativeCoinMint

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
	_ = abi.ConvertType
)

// NativeCoinMintMetaData contains all meta data concerning the NativeCoinMint contract.
var NativeCoinMintMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506102948061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610029575f3560e01c806340c10f191461002d575b5f5ffd5b6100476004803603810190610042919061016d565b610049565b005b73020000000000000000000000000000000000000173ffffffffffffffffffffffffffffffffffffffff16634f5aaaba83670de0b6b3a76400008461008e91906101d8565b6040518363ffffffff1660e01b81526004016100ab929190610237565b5f604051808303815f87803b1580156100c2575f5ffd5b505af11580156100d4573d5f5f3e3d5ffd5b505050505050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610109826100e0565b9050919050565b610119816100ff565b8114610123575f5ffd5b50565b5f8135905061013481610110565b92915050565b5f819050919050565b61014c8161013a565b8114610156575f5ffd5b50565b5f8135905061016781610143565b92915050565b5f5f60408385031215610183576101826100dc565b5b5f61019085828601610126565b92505060206101a185828601610159565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6101e28261013a565b91506101ed8361013a565b92508282026101fb8161013a565b91508282048414831517610212576102116101ab565b5b5092915050565b610222816100ff565b82525050565b6102318161013a565b82525050565b5f60408201905061024a5f830185610219565b6102576020830184610228565b939250505056fea2646970667358221220a77010742f36c256c004e73d9de902e2e78eaeb78e3faf8e4fb99f300a36178a64736f6c634300081b0033",
}

// NativeCoinMintABI is the input ABI used to generate the binding from.
// Deprecated: Use NativeCoinMintMetaData.ABI instead.
var NativeCoinMintABI = NativeCoinMintMetaData.ABI

// NativeCoinMintBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NativeCoinMintMetaData.Bin instead.
var NativeCoinMintBin = NativeCoinMintMetaData.Bin

// DeployNativeCoinMint deploys a new Ethereum contract, binding an instance of NativeCoinMint to it.
func DeployNativeCoinMint(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NativeCoinMint, error) {
	parsed, err := NativeCoinMintMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NativeCoinMintBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NativeCoinMint{NativeCoinMintCaller: NativeCoinMintCaller{contract: contract}, NativeCoinMintTransactor: NativeCoinMintTransactor{contract: contract}, NativeCoinMintFilterer: NativeCoinMintFilterer{contract: contract}}, nil
}

// NativeCoinMint is an auto generated Go binding around an Ethereum contract.
type NativeCoinMint struct {
	NativeCoinMintCaller     // Read-only binding to the contract
	NativeCoinMintTransactor // Write-only binding to the contract
	NativeCoinMintFilterer   // Log filterer for contract events
}

// NativeCoinMintCaller is an auto generated read-only Go binding around an Ethereum contract.
type NativeCoinMintCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeCoinMintTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NativeCoinMintTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeCoinMintFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NativeCoinMintFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeCoinMintSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NativeCoinMintSession struct {
	Contract     *NativeCoinMint   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NativeCoinMintCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NativeCoinMintCallerSession struct {
	Contract *NativeCoinMintCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// NativeCoinMintTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NativeCoinMintTransactorSession struct {
	Contract     *NativeCoinMintTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// NativeCoinMintRaw is an auto generated low-level Go binding around an Ethereum contract.
type NativeCoinMintRaw struct {
	Contract *NativeCoinMint // Generic contract binding to access the raw methods on
}

// NativeCoinMintCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NativeCoinMintCallerRaw struct {
	Contract *NativeCoinMintCaller // Generic read-only contract binding to access the raw methods on
}

// NativeCoinMintTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NativeCoinMintTransactorRaw struct {
	Contract *NativeCoinMintTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNativeCoinMint creates a new instance of NativeCoinMint, bound to a specific deployed contract.
func NewNativeCoinMint(address common.Address, backend bind.ContractBackend) (*NativeCoinMint, error) {
	contract, err := bindNativeCoinMint(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NativeCoinMint{NativeCoinMintCaller: NativeCoinMintCaller{contract: contract}, NativeCoinMintTransactor: NativeCoinMintTransactor{contract: contract}, NativeCoinMintFilterer: NativeCoinMintFilterer{contract: contract}}, nil
}

// NewNativeCoinMintCaller creates a new read-only instance of NativeCoinMint, bound to a specific deployed contract.
func NewNativeCoinMintCaller(address common.Address, caller bind.ContractCaller) (*NativeCoinMintCaller, error) {
	contract, err := bindNativeCoinMint(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NativeCoinMintCaller{contract: contract}, nil
}

// NewNativeCoinMintTransactor creates a new write-only instance of NativeCoinMint, bound to a specific deployed contract.
func NewNativeCoinMintTransactor(address common.Address, transactor bind.ContractTransactor) (*NativeCoinMintTransactor, error) {
	contract, err := bindNativeCoinMint(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NativeCoinMintTransactor{contract: contract}, nil
}

// NewNativeCoinMintFilterer creates a new log filterer instance of NativeCoinMint, bound to a specific deployed contract.
func NewNativeCoinMintFilterer(address common.Address, filterer bind.ContractFilterer) (*NativeCoinMintFilterer, error) {
	contract, err := bindNativeCoinMint(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NativeCoinMintFilterer{contract: contract}, nil
}

// bindNativeCoinMint binds a generic wrapper to an already deployed contract.
func bindNativeCoinMint(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NativeCoinMintMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeCoinMint *NativeCoinMintRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeCoinMint.Contract.NativeCoinMintCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeCoinMint *NativeCoinMintRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeCoinMint.Contract.NativeCoinMintTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeCoinMint *NativeCoinMintRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeCoinMint.Contract.NativeCoinMintTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeCoinMint *NativeCoinMintCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeCoinMint.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeCoinMint *NativeCoinMintTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeCoinMint.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeCoinMint *NativeCoinMintTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeCoinMint.Contract.contract.Transact(opts, method, params...)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address addr, uint256 amount) returns()
func (_NativeCoinMint *NativeCoinMintTransactor) Mint(opts *bind.TransactOpts, addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeCoinMint.contract.Transact(opts, "mint", addr, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address addr, uint256 amount) returns()
func (_NativeCoinMint *NativeCoinMintSession) Mint(addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeCoinMint.Contract.Mint(&_NativeCoinMint.TransactOpts, addr, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address addr, uint256 amount) returns()
func (_NativeCoinMint *NativeCoinMintTransactorSession) Mint(addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeCoinMint.Contract.Mint(&_NativeCoinMint.TransactOpts, addr, amount)
}
