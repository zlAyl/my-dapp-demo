// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package count

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

// CountMetaData contains all meta data concerning the Count contract.
var CountMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b5060015f81905550610172806100235f395ff3fe608060405234801561000f575f5ffd5b5060043610610034575f3560e01c806306661abd14610038578063c1d88d5914610056575b5f5ffd5b610040610074565b60405161004d91906100af565b60405180910390f35b61005e610079565b60405161006b91906100af565b60405180910390f35b5f5481565b5f5f5f81548092919061008b906100f5565b91905055505f54905090565b5f819050919050565b6100a981610097565b82525050565b5f6020820190506100c25f8301846100a0565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6100ff82610097565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610131576101306100c8565b5b60018201905091905056fea26469706673582212201b03b3af3122beadea14b3dc3e0345e5c454c585352abe7c8e63e9e7e9308a9664736f6c634300081e0033",
}

// CountABI is the input ABI used to generate the binding from.
// Deprecated: Use CountMetaData.ABI instead.
var CountABI = CountMetaData.ABI

// CountBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CountMetaData.Bin instead.
var CountBin = CountMetaData.Bin

// DeployCount deploys a new Ethereum contract, binding an instance of Count to it.
func DeployCount(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Count, error) {
	parsed, err := CountMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CountBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Count{CountCaller: CountCaller{contract: contract}, CountTransactor: CountTransactor{contract: contract}, CountFilterer: CountFilterer{contract: contract}}, nil
}

// Count is an auto generated Go binding around an Ethereum contract.
type Count struct {
	CountCaller     // Read-only binding to the contract
	CountTransactor // Write-only binding to the contract
	CountFilterer   // Log filterer for contract events
}

// CountCaller is an auto generated read-only Go binding around an Ethereum contract.
type CountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CountSession struct {
	Contract     *Count            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CountCallerSession struct {
	Contract *CountCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CountTransactorSession struct {
	Contract     *CountTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CountRaw is an auto generated low-level Go binding around an Ethereum contract.
type CountRaw struct {
	Contract *Count // Generic contract binding to access the raw methods on
}

// CountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CountCallerRaw struct {
	Contract *CountCaller // Generic read-only contract binding to access the raw methods on
}

// CountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CountTransactorRaw struct {
	Contract *CountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCount creates a new instance of Count, bound to a specific deployed contract.
func NewCount(address common.Address, backend bind.ContractBackend) (*Count, error) {
	contract, err := bindCount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Count{CountCaller: CountCaller{contract: contract}, CountTransactor: CountTransactor{contract: contract}, CountFilterer: CountFilterer{contract: contract}}, nil
}

// NewCountCaller creates a new read-only instance of Count, bound to a specific deployed contract.
func NewCountCaller(address common.Address, caller bind.ContractCaller) (*CountCaller, error) {
	contract, err := bindCount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CountCaller{contract: contract}, nil
}

// NewCountTransactor creates a new write-only instance of Count, bound to a specific deployed contract.
func NewCountTransactor(address common.Address, transactor bind.ContractTransactor) (*CountTransactor, error) {
	contract, err := bindCount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CountTransactor{contract: contract}, nil
}

// NewCountFilterer creates a new log filterer instance of Count, bound to a specific deployed contract.
func NewCountFilterer(address common.Address, filterer bind.ContractFilterer) (*CountFilterer, error) {
	contract, err := bindCount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CountFilterer{contract: contract}, nil
}

// bindCount binds a generic wrapper to an already deployed contract.
func bindCount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CountMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Count *CountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Count.Contract.CountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Count *CountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Count.Contract.CountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Count *CountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Count.Contract.CountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Count *CountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Count.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Count *CountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Count.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Count *CountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Count.Contract.contract.Transact(opts, method, params...)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Count *CountCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Count.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Count *CountSession) Count() (*big.Int, error) {
	return _Count.Contract.Count(&_Count.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Count *CountCallerSession) Count() (*big.Int, error) {
	return _Count.Contract.Count(&_Count.CallOpts)
}

// AddCount is a paid mutator transaction binding the contract method 0xc1d88d59.
//
// Solidity: function AddCount() returns(uint256)
func (_Count *CountTransactor) AddCount(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Count.contract.Transact(opts, "AddCount")
}

// AddCount is a paid mutator transaction binding the contract method 0xc1d88d59.
//
// Solidity: function AddCount() returns(uint256)
func (_Count *CountSession) AddCount() (*types.Transaction, error) {
	return _Count.Contract.AddCount(&_Count.TransactOpts)
}

// AddCount is a paid mutator transaction binding the contract method 0xc1d88d59.
//
// Solidity: function AddCount() returns(uint256)
func (_Count *CountTransactorSession) AddCount() (*types.Transaction, error) {
	return _Count.Contract.AddCount(&_Count.TransactOpts)
}
