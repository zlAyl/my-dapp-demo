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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nextCount\",\"type\":\"uint256\"}],\"name\":\"AddCount\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"addCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b5060015f819055506101c1806100235f395ff3fe608060405234801561000f575f5ffd5b5060043610610034575f3560e01c806306661abd146100385780636cefce6e14610056575b5f5ffd5b610040610074565b60405161004d91906100fe565b60405180910390f35b61005e610079565b60405161006b91906100fe565b60405180910390f35b5f5481565b5f5f5f81548092919061008b90610144565b91905055503373ffffffffffffffffffffffffffffffffffffffff167f6605ea7aa1f2d5ca5356fd6e346790585c2f60375bd0038ca74a9214004cd7145f546040516100d791906100fe565b60405180910390a25f54905090565b5f819050919050565b6100f8816100e6565b82525050565b5f6020820190506101115f8301846100ef565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61014e826100e6565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036101805761017f610117565b5b60018201905091905056fea2646970667358221220c1952bec46be712e9ce1a38e88bcb30cff5926c401ffd5975a162f1277dc7f5d64736f6c634300081e0033",
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

// AddCount is a paid mutator transaction binding the contract method 0x6cefce6e.
//
// Solidity: function addCount() returns(uint256)
func (_Count *CountTransactor) AddCount(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Count.contract.Transact(opts, "addCount")
}

// AddCount is a paid mutator transaction binding the contract method 0x6cefce6e.
//
// Solidity: function addCount() returns(uint256)
func (_Count *CountSession) AddCount() (*types.Transaction, error) {
	return _Count.Contract.AddCount(&_Count.TransactOpts)
}

// AddCount is a paid mutator transaction binding the contract method 0x6cefce6e.
//
// Solidity: function addCount() returns(uint256)
func (_Count *CountTransactorSession) AddCount() (*types.Transaction, error) {
	return _Count.Contract.AddCount(&_Count.TransactOpts)
}

// CountAddCountIterator is returned from FilterAddCount and is used to iterate over the raw logs and unpacked data for AddCount events raised by the Count contract.
type CountAddCountIterator struct {
	Event *CountAddCount // Event containing the contract specifics and raw log

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
func (it *CountAddCountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CountAddCount)
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
		it.Event = new(CountAddCount)
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
func (it *CountAddCountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CountAddCountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CountAddCount represents a AddCount event raised by the Count contract.
type CountAddCount struct {
	Sender    common.Address
	NextCount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAddCount is a free log retrieval operation binding the contract event 0x6605ea7aa1f2d5ca5356fd6e346790585c2f60375bd0038ca74a9214004cd714.
//
// Solidity: event AddCount(address indexed sender, uint256 nextCount)
func (_Count *CountFilterer) FilterAddCount(opts *bind.FilterOpts, sender []common.Address) (*CountAddCountIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Count.contract.FilterLogs(opts, "AddCount", senderRule)
	if err != nil {
		return nil, err
	}
	return &CountAddCountIterator{contract: _Count.contract, event: "AddCount", logs: logs, sub: sub}, nil
}

// WatchAddCount is a free log subscription operation binding the contract event 0x6605ea7aa1f2d5ca5356fd6e346790585c2f60375bd0038ca74a9214004cd714.
//
// Solidity: event AddCount(address indexed sender, uint256 nextCount)
func (_Count *CountFilterer) WatchAddCount(opts *bind.WatchOpts, sink chan<- *CountAddCount, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Count.contract.WatchLogs(opts, "AddCount", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CountAddCount)
				if err := _Count.contract.UnpackLog(event, "AddCount", log); err != nil {
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

// ParseAddCount is a log parse operation binding the contract event 0x6605ea7aa1f2d5ca5356fd6e346790585c2f60375bd0038ca74a9214004cd714.
//
// Solidity: event AddCount(address indexed sender, uint256 nextCount)
func (_Count *CountFilterer) ParseAddCount(log types.Log) (*CountAddCount, error) {
	event := new(CountAddCount)
	if err := _Count.contract.UnpackLog(event, "AddCount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
