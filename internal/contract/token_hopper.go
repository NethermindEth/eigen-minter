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
	_ = abi.ConvertType
)

// ITokenHopperHopperConfiguration is an auto generated low-level Go binding around an user-defined struct.
type ITokenHopperHopperConfiguration struct {
	Token               common.Address
	StartTime           *big.Int
	CooldownSeconds     *big.Int
	ActionGenerator     common.Address
	DoesExpire          bool
	ExpirationTimestamp *big.Int
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cooldownSeconds\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"actionGenerator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"doesExpire\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structITokenHopper.HopperConfiguration\",\"name\":\"config\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCooldownHorizon\",\"type\":\"uint256\"}],\"name\":\"ButtonPressed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundsRetrieved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cooldownSeconds\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"actionGenerator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"doesExpire\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimestamp\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structITokenHopper.HopperConfiguration\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"HopperLoaded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"canPress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getHopperConfiguration\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cooldownSeconds\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"actionGenerator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"doesExpire\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structITokenHopper.HopperConfiguration\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isExpired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestPress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pressButton\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"retrieveFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// CanPress is a free data retrieval call binding the contract method 0x2d60458d.
//
// Solidity: function canPress() view returns(bool)
func (_Contract *ContractCaller) CanPress(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "canPress")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanPress is a free data retrieval call binding the contract method 0x2d60458d.
//
// Solidity: function canPress() view returns(bool)
func (_Contract *ContractSession) CanPress() (bool, error) {
	return _Contract.Contract.CanPress(&_Contract.CallOpts)
}

// CanPress is a free data retrieval call binding the contract method 0x2d60458d.
//
// Solidity: function canPress() view returns(bool)
func (_Contract *ContractCallerSession) CanPress() (bool, error) {
	return _Contract.Contract.CanPress(&_Contract.CallOpts)
}

// GetHopperConfiguration is a free data retrieval call binding the contract method 0x7233497a.
//
// Solidity: function getHopperConfiguration() view returns((address,uint256,uint256,address,bool,uint256))
func (_Contract *ContractCaller) GetHopperConfiguration(opts *bind.CallOpts) (ITokenHopperHopperConfiguration, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getHopperConfiguration")

	if err != nil {
		return *new(ITokenHopperHopperConfiguration), err
	}

	out0 := *abi.ConvertType(out[0], new(ITokenHopperHopperConfiguration)).(*ITokenHopperHopperConfiguration)

	return out0, err

}

// GetHopperConfiguration is a free data retrieval call binding the contract method 0x7233497a.
//
// Solidity: function getHopperConfiguration() view returns((address,uint256,uint256,address,bool,uint256))
func (_Contract *ContractSession) GetHopperConfiguration() (ITokenHopperHopperConfiguration, error) {
	return _Contract.Contract.GetHopperConfiguration(&_Contract.CallOpts)
}

// GetHopperConfiguration is a free data retrieval call binding the contract method 0x7233497a.
//
// Solidity: function getHopperConfiguration() view returns((address,uint256,uint256,address,bool,uint256))
func (_Contract *ContractCallerSession) GetHopperConfiguration() (ITokenHopperHopperConfiguration, error) {
	return _Contract.Contract.GetHopperConfiguration(&_Contract.CallOpts)
}

// IsExpired is a free data retrieval call binding the contract method 0x2f13b60c.
//
// Solidity: function isExpired() view returns(bool)
func (_Contract *ContractCaller) IsExpired(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isExpired")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExpired is a free data retrieval call binding the contract method 0x2f13b60c.
//
// Solidity: function isExpired() view returns(bool)
func (_Contract *ContractSession) IsExpired() (bool, error) {
	return _Contract.Contract.IsExpired(&_Contract.CallOpts)
}

// IsExpired is a free data retrieval call binding the contract method 0x2f13b60c.
//
// Solidity: function isExpired() view returns(bool)
func (_Contract *ContractCallerSession) IsExpired() (bool, error) {
	return _Contract.Contract.IsExpired(&_Contract.CallOpts)
}

// LatestPress is a free data retrieval call binding the contract method 0x9ab1a00b.
//
// Solidity: function latestPress() view returns(uint256)
func (_Contract *ContractCaller) LatestPress(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "latestPress")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestPress is a free data retrieval call binding the contract method 0x9ab1a00b.
//
// Solidity: function latestPress() view returns(uint256)
func (_Contract *ContractSession) LatestPress() (*big.Int, error) {
	return _Contract.Contract.LatestPress(&_Contract.CallOpts)
}

// LatestPress is a free data retrieval call binding the contract method 0x9ab1a00b.
//
// Solidity: function latestPress() view returns(uint256)
func (_Contract *ContractCallerSession) LatestPress() (*big.Int, error) {
	return _Contract.Contract.LatestPress(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// PressButton is a paid mutator transaction binding the contract method 0xc8b5da66.
//
// Solidity: function pressButton() returns()
func (_Contract *ContractTransactor) PressButton(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "pressButton")
}

// PressButton is a paid mutator transaction binding the contract method 0xc8b5da66.
//
// Solidity: function pressButton() returns()
func (_Contract *ContractSession) PressButton() (*types.Transaction, error) {
	return _Contract.Contract.PressButton(&_Contract.TransactOpts)
}

// PressButton is a paid mutator transaction binding the contract method 0xc8b5da66.
//
// Solidity: function pressButton() returns()
func (_Contract *ContractTransactorSession) PressButton() (*types.Transaction, error) {
	return _Contract.Contract.PressButton(&_Contract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// RetrieveFunds is a paid mutator transaction binding the contract method 0x61b20d8c.
//
// Solidity: function retrieveFunds() returns()
func (_Contract *ContractTransactor) RetrieveFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "retrieveFunds")
}

// RetrieveFunds is a paid mutator transaction binding the contract method 0x61b20d8c.
//
// Solidity: function retrieveFunds() returns()
func (_Contract *ContractSession) RetrieveFunds() (*types.Transaction, error) {
	return _Contract.Contract.RetrieveFunds(&_Contract.TransactOpts)
}

// RetrieveFunds is a paid mutator transaction binding the contract method 0x61b20d8c.
//
// Solidity: function retrieveFunds() returns()
func (_Contract *ContractTransactorSession) RetrieveFunds() (*types.Transaction, error) {
	return _Contract.Contract.RetrieveFunds(&_Contract.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// ContractButtonPressedIterator is returned from FilterButtonPressed and is used to iterate over the raw logs and unpacked data for ButtonPressed events raised by the Contract contract.
type ContractButtonPressedIterator struct {
	Event *ContractButtonPressed // Event containing the contract specifics and raw log

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
func (it *ContractButtonPressedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractButtonPressed)
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
		it.Event = new(ContractButtonPressed)
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
func (it *ContractButtonPressedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractButtonPressedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractButtonPressed represents a ButtonPressed event raised by the Contract contract.
type ContractButtonPressed struct {
	Caller             common.Address
	NewCooldownHorizon *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterButtonPressed is a free log retrieval operation binding the contract event 0xf1e2abbdd3969fea5f6950e78b803559d0587ff8da184624bdba51665307c161.
//
// Solidity: event ButtonPressed(address indexed caller, uint256 newCooldownHorizon)
func (_Contract *ContractFilterer) FilterButtonPressed(opts *bind.FilterOpts, caller []common.Address) (*ContractButtonPressedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ButtonPressed", callerRule)
	if err != nil {
		return nil, err
	}
	return &ContractButtonPressedIterator{contract: _Contract.contract, event: "ButtonPressed", logs: logs, sub: sub}, nil
}

// WatchButtonPressed is a free log subscription operation binding the contract event 0xf1e2abbdd3969fea5f6950e78b803559d0587ff8da184624bdba51665307c161.
//
// Solidity: event ButtonPressed(address indexed caller, uint256 newCooldownHorizon)
func (_Contract *ContractFilterer) WatchButtonPressed(opts *bind.WatchOpts, sink chan<- *ContractButtonPressed, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ButtonPressed", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractButtonPressed)
				if err := _Contract.contract.UnpackLog(event, "ButtonPressed", log); err != nil {
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

// ParseButtonPressed is a log parse operation binding the contract event 0xf1e2abbdd3969fea5f6950e78b803559d0587ff8da184624bdba51665307c161.
//
// Solidity: event ButtonPressed(address indexed caller, uint256 newCooldownHorizon)
func (_Contract *ContractFilterer) ParseButtonPressed(log types.Log) (*ContractButtonPressed, error) {
	event := new(ContractButtonPressed)
	if err := _Contract.contract.UnpackLog(event, "ButtonPressed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractFundsRetrievedIterator is returned from FilterFundsRetrieved and is used to iterate over the raw logs and unpacked data for FundsRetrieved events raised by the Contract contract.
type ContractFundsRetrievedIterator struct {
	Event *ContractFundsRetrieved // Event containing the contract specifics and raw log

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
func (it *ContractFundsRetrievedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractFundsRetrieved)
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
		it.Event = new(ContractFundsRetrieved)
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
func (it *ContractFundsRetrievedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractFundsRetrievedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractFundsRetrieved represents a FundsRetrieved event raised by the Contract contract.
type ContractFundsRetrieved struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFundsRetrieved is a free log retrieval operation binding the contract event 0xae2f9495e4eb15d60eced7f3c7944eb4558245d5d9089f1b32c17b493ed58306.
//
// Solidity: event FundsRetrieved(uint256 amount)
func (_Contract *ContractFilterer) FilterFundsRetrieved(opts *bind.FilterOpts) (*ContractFundsRetrievedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "FundsRetrieved")
	if err != nil {
		return nil, err
	}
	return &ContractFundsRetrievedIterator{contract: _Contract.contract, event: "FundsRetrieved", logs: logs, sub: sub}, nil
}

// WatchFundsRetrieved is a free log subscription operation binding the contract event 0xae2f9495e4eb15d60eced7f3c7944eb4558245d5d9089f1b32c17b493ed58306.
//
// Solidity: event FundsRetrieved(uint256 amount)
func (_Contract *ContractFilterer) WatchFundsRetrieved(opts *bind.WatchOpts, sink chan<- *ContractFundsRetrieved) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "FundsRetrieved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractFundsRetrieved)
				if err := _Contract.contract.UnpackLog(event, "FundsRetrieved", log); err != nil {
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

// ParseFundsRetrieved is a log parse operation binding the contract event 0xae2f9495e4eb15d60eced7f3c7944eb4558245d5d9089f1b32c17b493ed58306.
//
// Solidity: event FundsRetrieved(uint256 amount)
func (_Contract *ContractFilterer) ParseFundsRetrieved(log types.Log) (*ContractFundsRetrieved, error) {
	event := new(ContractFundsRetrieved)
	if err := _Contract.contract.UnpackLog(event, "FundsRetrieved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractHopperLoadedIterator is returned from FilterHopperLoaded and is used to iterate over the raw logs and unpacked data for HopperLoaded events raised by the Contract contract.
type ContractHopperLoadedIterator struct {
	Event *ContractHopperLoaded // Event containing the contract specifics and raw log

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
func (it *ContractHopperLoadedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractHopperLoaded)
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
		it.Event = new(ContractHopperLoaded)
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
func (it *ContractHopperLoadedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractHopperLoadedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractHopperLoaded represents a HopperLoaded event raised by the Contract contract.
type ContractHopperLoaded struct {
	Config ITokenHopperHopperConfiguration
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterHopperLoaded is a free log retrieval operation binding the contract event 0xb4ac6a42c48566440710983313cd5df20e819b8c753c9951d9e115d6c641e11a.
//
// Solidity: event HopperLoaded((address,uint256,uint256,address,bool,uint256) config)
func (_Contract *ContractFilterer) FilterHopperLoaded(opts *bind.FilterOpts) (*ContractHopperLoadedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "HopperLoaded")
	if err != nil {
		return nil, err
	}
	return &ContractHopperLoadedIterator{contract: _Contract.contract, event: "HopperLoaded", logs: logs, sub: sub}, nil
}

// WatchHopperLoaded is a free log subscription operation binding the contract event 0xb4ac6a42c48566440710983313cd5df20e819b8c753c9951d9e115d6c641e11a.
//
// Solidity: event HopperLoaded((address,uint256,uint256,address,bool,uint256) config)
func (_Contract *ContractFilterer) WatchHopperLoaded(opts *bind.WatchOpts, sink chan<- *ContractHopperLoaded) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "HopperLoaded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractHopperLoaded)
				if err := _Contract.contract.UnpackLog(event, "HopperLoaded", log); err != nil {
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

// ParseHopperLoaded is a log parse operation binding the contract event 0xb4ac6a42c48566440710983313cd5df20e819b8c753c9951d9e115d6c641e11a.
//
// Solidity: event HopperLoaded((address,uint256,uint256,address,bool,uint256) config)
func (_Contract *ContractFilterer) ParseHopperLoaded(log types.Log) (*ContractHopperLoaded, error) {
	event := new(ContractHopperLoaded)
	if err := _Contract.contract.UnpackLog(event, "HopperLoaded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contract contract.
type ContractOwnershipTransferredIterator struct {
	Event *ContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnershipTransferred)
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
		it.Event = new(ContractOwnershipTransferred)
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
func (it *ContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOwnershipTransferred represents a OwnershipTransferred event raised by the Contract contract.
type ContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOwnershipTransferredIterator{contract: _Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOwnershipTransferred)
				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOwnershipTransferred, error) {
	event := new(ContractOwnershipTransferred)
	if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
