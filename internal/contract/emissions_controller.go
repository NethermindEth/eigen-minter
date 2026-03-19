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

// IEmissionsControllerTypesDistribution is an auto generated low-level Go binding around an user-defined struct.
type IEmissionsControllerTypesDistribution struct {
	Weight                   uint64
	StartEpoch               uint64
	TotalEpochs              uint64
	DistributionType         uint8
	OperatorSet              OperatorSet
	StrategiesAndMultipliers [][]IRewardsCoordinatorTypesStrategyAndMultiplier
}

// IRewardsCoordinatorTypesStrategyAndMultiplier is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorTypesStrategyAndMultiplier struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEigen\",\"name\":\"eigen\",\"type\":\"address\"},{\"internalType\":\"contractIBackingEigen\",\"name\":\"backingEigen\",\"type\":\"address\"},{\"internalType\":\"contractIAllocationManager\",\"name\":\"allocationManager\",\"type\":\"address\"},{\"internalType\":\"contractIRewardsCoordinator\",\"name\":\"rewardsCoordinator\",\"type\":\"address\"},{\"internalType\":\"contractIPauserRegistry\",\"name\":\"pauserRegistry\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inflationRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cooldownSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calculationIntervalSeconds\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AllDistributionsMustBeProcessed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllDistributionsProcessed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerIsNotIncentiveCouncil\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotAddDisabledDistribution\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CurrentlyPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmissionsNotStarted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EpochLengthNotAlignedWithCalculationInterval\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InputAddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDistributionType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNewPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaliciousCallDetected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyUnpauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorSetNotRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RewardsSubmissionsCannotBeEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StartEpochMustBeInTheFuture\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StartTimeNotAlignedWithCalculationInterval\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TotalWeightExceedsMax\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"distributionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startEpoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"totalEpochs\",\"type\":\"uint64\"},{\"internalType\":\"enumIEmissionsControllerTypes.DistributionType\",\"name\":\"distributionType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"avs\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"}],\"internalType\":\"structOperatorSet\",\"name\":\"operatorSet\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"multiplier\",\"type\":\"uint96\"}],\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[][]\",\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[][]\"}],\"indexed\":false,\"internalType\":\"structIEmissionsControllerTypes.Distribution\",\"name\":\"distribution\",\"type\":\"tuple\"}],\"name\":\"DistributionAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"distributionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startEpoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"totalEpochs\",\"type\":\"uint64\"},{\"internalType\":\"enumIEmissionsControllerTypes.DistributionType\",\"name\":\"distributionType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"avs\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"}],\"internalType\":\"structOperatorSet\",\"name\":\"operatorSet\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"multiplier\",\"type\":\"uint96\"}],\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[][]\",\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[][]\"}],\"indexed\":false,\"internalType\":\"structIEmissionsControllerTypes.Distribution\",\"name\":\"distribution\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"DistributionProcessed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"distributionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startEpoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"totalEpochs\",\"type\":\"uint64\"},{\"internalType\":\"enumIEmissionsControllerTypes.DistributionType\",\"name\":\"distributionType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"avs\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"}],\"internalType\":\"structOperatorSet\",\"name\":\"operatorSet\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"multiplier\",\"type\":\"uint96\"}],\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[][]\",\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[][]\"}],\"indexed\":false,\"internalType\":\"structIEmissionsControllerTypes.Distribution\",\"name\":\"distribution\",\"type\":\"tuple\"}],\"name\":\"DistributionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"incentiveCouncil\",\"type\":\"address\"}],\"name\":\"IncentiveCouncilUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"incentiveCouncil\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Swept\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ALLOCATION_MANAGER\",\"outputs\":[{\"internalType\":\"contractIAllocationManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BACKING_EIGEN\",\"outputs\":[{\"internalType\":\"contractIBackingEigen\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EIGEN\",\"outputs\":[{\"internalType\":\"contractIEigen\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EMISSIONS_EPOCH_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EMISSIONS_INFLATION_RATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EMISSIONS_START_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_TOTAL_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARDS_COORDINATOR\",\"outputs\":[{\"internalType\":\"contractIRewardsCoordinator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startEpoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"totalEpochs\",\"type\":\"uint64\"},{\"internalType\":\"enumIEmissionsControllerTypes.DistributionType\",\"name\":\"distributionType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"avs\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"}],\"internalType\":\"structOperatorSet\",\"name\":\"operatorSet\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"multiplier\",\"type\":\"uint96\"}],\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[][]\",\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[][]\"}],\"internalType\":\"structIEmissionsControllerTypes.Distribution\",\"name\":\"distribution\",\"type\":\"tuple\"}],\"name\":\"addDistribution\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"distributionId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"distributionId\",\"type\":\"uint256\"}],\"name\":\"getDistribution\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startEpoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"totalEpochs\",\"type\":\"uint64\"},{\"internalType\":\"enumIEmissionsControllerTypes.DistributionType\",\"name\":\"distributionType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"avs\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"}],\"internalType\":\"structOperatorSet\",\"name\":\"operatorSet\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"multiplier\",\"type\":\"uint96\"}],\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[][]\",\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[][]\"}],\"internalType\":\"structIEmissionsControllerTypes.Distribution\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"getDistributions\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startEpoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"totalEpochs\",\"type\":\"uint64\"},{\"internalType\":\"enumIEmissionsControllerTypes.DistributionType\",\"name\":\"distributionType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"avs\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"}],\"internalType\":\"structOperatorSet\",\"name\":\"operatorSet\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"multiplier\",\"type\":\"uint96\"}],\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[][]\",\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[][]\"}],\"internalType\":\"structIEmissionsControllerTypes.Distribution[]\",\"name\":\"distributions\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalProcessableDistributions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incentiveCouncil\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialIncentiveCouncil\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialPausedStatus\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isButtonPressable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTimeButtonPressable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextTimeButtonPressable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"}],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauserRegistry\",\"outputs\":[{\"internalType\":\"contractIPauserRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"pressButton\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newIncentiveCouncil\",\"type\":\"address\"}],\"name\":\"setIncentiveCouncil\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sweep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalWeight\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"distributionId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startEpoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"totalEpochs\",\"type\":\"uint64\"},{\"internalType\":\"enumIEmissionsControllerTypes.DistributionType\",\"name\":\"distributionType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"avs\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"}],\"internalType\":\"structOperatorSet\",\"name\":\"operatorSet\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"multiplier\",\"type\":\"uint96\"}],\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[][]\",\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[][]\"}],\"internalType\":\"structIEmissionsControllerTypes.Distribution\",\"name\":\"distribution\",\"type\":\"tuple\"}],\"name\":\"updateDistribution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// ALLOCATIONMANAGER is a free data retrieval call binding the contract method 0x31232bc9.
//
// Solidity: function ALLOCATION_MANAGER() view returns(address)
func (_Contract *ContractCaller) ALLOCATIONMANAGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "ALLOCATION_MANAGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ALLOCATIONMANAGER is a free data retrieval call binding the contract method 0x31232bc9.
//
// Solidity: function ALLOCATION_MANAGER() view returns(address)
func (_Contract *ContractSession) ALLOCATIONMANAGER() (common.Address, error) {
	return _Contract.Contract.ALLOCATIONMANAGER(&_Contract.CallOpts)
}

// ALLOCATIONMANAGER is a free data retrieval call binding the contract method 0x31232bc9.
//
// Solidity: function ALLOCATION_MANAGER() view returns(address)
func (_Contract *ContractCallerSession) ALLOCATIONMANAGER() (common.Address, error) {
	return _Contract.Contract.ALLOCATIONMANAGER(&_Contract.CallOpts)
}

// BACKINGEIGEN is a free data retrieval call binding the contract method 0xd455724e.
//
// Solidity: function BACKING_EIGEN() view returns(address)
func (_Contract *ContractCaller) BACKINGEIGEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "BACKING_EIGEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BACKINGEIGEN is a free data retrieval call binding the contract method 0xd455724e.
//
// Solidity: function BACKING_EIGEN() view returns(address)
func (_Contract *ContractSession) BACKINGEIGEN() (common.Address, error) {
	return _Contract.Contract.BACKINGEIGEN(&_Contract.CallOpts)
}

// BACKINGEIGEN is a free data retrieval call binding the contract method 0xd455724e.
//
// Solidity: function BACKING_EIGEN() view returns(address)
func (_Contract *ContractCallerSession) BACKINGEIGEN() (common.Address, error) {
	return _Contract.Contract.BACKINGEIGEN(&_Contract.CallOpts)
}

// EIGEN is a free data retrieval call binding the contract method 0xfdc371ce.
//
// Solidity: function EIGEN() view returns(address)
func (_Contract *ContractCaller) EIGEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "EIGEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EIGEN is a free data retrieval call binding the contract method 0xfdc371ce.
//
// Solidity: function EIGEN() view returns(address)
func (_Contract *ContractSession) EIGEN() (common.Address, error) {
	return _Contract.Contract.EIGEN(&_Contract.CallOpts)
}

// EIGEN is a free data retrieval call binding the contract method 0xfdc371ce.
//
// Solidity: function EIGEN() view returns(address)
func (_Contract *ContractCallerSession) EIGEN() (common.Address, error) {
	return _Contract.Contract.EIGEN(&_Contract.CallOpts)
}

// EMISSIONSEPOCHLENGTH is a free data retrieval call binding the contract method 0xc2f208e4.
//
// Solidity: function EMISSIONS_EPOCH_LENGTH() view returns(uint256)
func (_Contract *ContractCaller) EMISSIONSEPOCHLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "EMISSIONS_EPOCH_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EMISSIONSEPOCHLENGTH is a free data retrieval call binding the contract method 0xc2f208e4.
//
// Solidity: function EMISSIONS_EPOCH_LENGTH() view returns(uint256)
func (_Contract *ContractSession) EMISSIONSEPOCHLENGTH() (*big.Int, error) {
	return _Contract.Contract.EMISSIONSEPOCHLENGTH(&_Contract.CallOpts)
}

// EMISSIONSEPOCHLENGTH is a free data retrieval call binding the contract method 0xc2f208e4.
//
// Solidity: function EMISSIONS_EPOCH_LENGTH() view returns(uint256)
func (_Contract *ContractCallerSession) EMISSIONSEPOCHLENGTH() (*big.Int, error) {
	return _Contract.Contract.EMISSIONSEPOCHLENGTH(&_Contract.CallOpts)
}

// EMISSIONSINFLATIONRATE is a free data retrieval call binding the contract method 0x47a28ea2.
//
// Solidity: function EMISSIONS_INFLATION_RATE() view returns(uint256)
func (_Contract *ContractCaller) EMISSIONSINFLATIONRATE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "EMISSIONS_INFLATION_RATE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EMISSIONSINFLATIONRATE is a free data retrieval call binding the contract method 0x47a28ea2.
//
// Solidity: function EMISSIONS_INFLATION_RATE() view returns(uint256)
func (_Contract *ContractSession) EMISSIONSINFLATIONRATE() (*big.Int, error) {
	return _Contract.Contract.EMISSIONSINFLATIONRATE(&_Contract.CallOpts)
}

// EMISSIONSINFLATIONRATE is a free data retrieval call binding the contract method 0x47a28ea2.
//
// Solidity: function EMISSIONS_INFLATION_RATE() view returns(uint256)
func (_Contract *ContractCallerSession) EMISSIONSINFLATIONRATE() (*big.Int, error) {
	return _Contract.Contract.EMISSIONSINFLATIONRATE(&_Contract.CallOpts)
}

// EMISSIONSSTARTTIME is a free data retrieval call binding the contract method 0xc9d3eff9.
//
// Solidity: function EMISSIONS_START_TIME() view returns(uint256)
func (_Contract *ContractCaller) EMISSIONSSTARTTIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "EMISSIONS_START_TIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EMISSIONSSTARTTIME is a free data retrieval call binding the contract method 0xc9d3eff9.
//
// Solidity: function EMISSIONS_START_TIME() view returns(uint256)
func (_Contract *ContractSession) EMISSIONSSTARTTIME() (*big.Int, error) {
	return _Contract.Contract.EMISSIONSSTARTTIME(&_Contract.CallOpts)
}

// EMISSIONSSTARTTIME is a free data retrieval call binding the contract method 0xc9d3eff9.
//
// Solidity: function EMISSIONS_START_TIME() view returns(uint256)
func (_Contract *ContractCallerSession) EMISSIONSSTARTTIME() (*big.Int, error) {
	return _Contract.Contract.EMISSIONSSTARTTIME(&_Contract.CallOpts)
}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_Contract *ContractCaller) MAXTOTALWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "MAX_TOTAL_WEIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_Contract *ContractSession) MAXTOTALWEIGHT() (*big.Int, error) {
	return _Contract.Contract.MAXTOTALWEIGHT(&_Contract.CallOpts)
}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_Contract *ContractCallerSession) MAXTOTALWEIGHT() (*big.Int, error) {
	return _Contract.Contract.MAXTOTALWEIGHT(&_Contract.CallOpts)
}

// REWARDSCOORDINATOR is a free data retrieval call binding the contract method 0x71e2c264.
//
// Solidity: function REWARDS_COORDINATOR() view returns(address)
func (_Contract *ContractCaller) REWARDSCOORDINATOR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "REWARDS_COORDINATOR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// REWARDSCOORDINATOR is a free data retrieval call binding the contract method 0x71e2c264.
//
// Solidity: function REWARDS_COORDINATOR() view returns(address)
func (_Contract *ContractSession) REWARDSCOORDINATOR() (common.Address, error) {
	return _Contract.Contract.REWARDSCOORDINATOR(&_Contract.CallOpts)
}

// REWARDSCOORDINATOR is a free data retrieval call binding the contract method 0x71e2c264.
//
// Solidity: function REWARDS_COORDINATOR() view returns(address)
func (_Contract *ContractCallerSession) REWARDSCOORDINATOR() (common.Address, error) {
	return _Contract.Contract.REWARDSCOORDINATOR(&_Contract.CallOpts)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_Contract *ContractCaller) GetCurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getCurrentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_Contract *ContractSession) GetCurrentEpoch() (*big.Int, error) {
	return _Contract.Contract.GetCurrentEpoch(&_Contract.CallOpts)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_Contract *ContractCallerSession) GetCurrentEpoch() (*big.Int, error) {
	return _Contract.Contract.GetCurrentEpoch(&_Contract.CallOpts)
}

// GetDistribution is a free data retrieval call binding the contract method 0x3b345a87.
//
// Solidity: function getDistribution(uint256 distributionId) view returns((uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]))
func (_Contract *ContractCaller) GetDistribution(opts *bind.CallOpts, distributionId *big.Int) (IEmissionsControllerTypesDistribution, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getDistribution", distributionId)

	if err != nil {
		return *new(IEmissionsControllerTypesDistribution), err
	}

	out0 := *abi.ConvertType(out[0], new(IEmissionsControllerTypesDistribution)).(*IEmissionsControllerTypesDistribution)

	return out0, err

}

// GetDistribution is a free data retrieval call binding the contract method 0x3b345a87.
//
// Solidity: function getDistribution(uint256 distributionId) view returns((uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]))
func (_Contract *ContractSession) GetDistribution(distributionId *big.Int) (IEmissionsControllerTypesDistribution, error) {
	return _Contract.Contract.GetDistribution(&_Contract.CallOpts, distributionId)
}

// GetDistribution is a free data retrieval call binding the contract method 0x3b345a87.
//
// Solidity: function getDistribution(uint256 distributionId) view returns((uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]))
func (_Contract *ContractCallerSession) GetDistribution(distributionId *big.Int) (IEmissionsControllerTypesDistribution, error) {
	return _Contract.Contract.GetDistribution(&_Contract.CallOpts, distributionId)
}

// GetDistributions is a free data retrieval call binding the contract method 0x147a7a5b.
//
// Solidity: function getDistributions(uint256 start, uint256 length) view returns((uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][])[] distributions)
func (_Contract *ContractCaller) GetDistributions(opts *bind.CallOpts, start *big.Int, length *big.Int) ([]IEmissionsControllerTypesDistribution, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getDistributions", start, length)

	if err != nil {
		return *new([]IEmissionsControllerTypesDistribution), err
	}

	out0 := *abi.ConvertType(out[0], new([]IEmissionsControllerTypesDistribution)).(*[]IEmissionsControllerTypesDistribution)

	return out0, err

}

// GetDistributions is a free data retrieval call binding the contract method 0x147a7a5b.
//
// Solidity: function getDistributions(uint256 start, uint256 length) view returns((uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][])[] distributions)
func (_Contract *ContractSession) GetDistributions(start *big.Int, length *big.Int) ([]IEmissionsControllerTypesDistribution, error) {
	return _Contract.Contract.GetDistributions(&_Contract.CallOpts, start, length)
}

// GetDistributions is a free data retrieval call binding the contract method 0x147a7a5b.
//
// Solidity: function getDistributions(uint256 start, uint256 length) view returns((uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][])[] distributions)
func (_Contract *ContractCallerSession) GetDistributions(start *big.Int, length *big.Int) ([]IEmissionsControllerTypesDistribution, error) {
	return _Contract.Contract.GetDistributions(&_Contract.CallOpts, start, length)
}

// GetTotalProcessableDistributions is a free data retrieval call binding the contract method 0xbe851337.
//
// Solidity: function getTotalProcessableDistributions() view returns(uint256)
func (_Contract *ContractCaller) GetTotalProcessableDistributions(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getTotalProcessableDistributions")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalProcessableDistributions is a free data retrieval call binding the contract method 0xbe851337.
//
// Solidity: function getTotalProcessableDistributions() view returns(uint256)
func (_Contract *ContractSession) GetTotalProcessableDistributions() (*big.Int, error) {
	return _Contract.Contract.GetTotalProcessableDistributions(&_Contract.CallOpts)
}

// GetTotalProcessableDistributions is a free data retrieval call binding the contract method 0xbe851337.
//
// Solidity: function getTotalProcessableDistributions() view returns(uint256)
func (_Contract *ContractCallerSession) GetTotalProcessableDistributions() (*big.Int, error) {
	return _Contract.Contract.GetTotalProcessableDistributions(&_Contract.CallOpts)
}

// IncentiveCouncil is a free data retrieval call binding the contract method 0xc44cb727.
//
// Solidity: function incentiveCouncil() view returns(address)
func (_Contract *ContractCaller) IncentiveCouncil(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "incentiveCouncil")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IncentiveCouncil is a free data retrieval call binding the contract method 0xc44cb727.
//
// Solidity: function incentiveCouncil() view returns(address)
func (_Contract *ContractSession) IncentiveCouncil() (common.Address, error) {
	return _Contract.Contract.IncentiveCouncil(&_Contract.CallOpts)
}

// IncentiveCouncil is a free data retrieval call binding the contract method 0xc44cb727.
//
// Solidity: function incentiveCouncil() view returns(address)
func (_Contract *ContractCallerSession) IncentiveCouncil() (common.Address, error) {
	return _Contract.Contract.IncentiveCouncil(&_Contract.CallOpts)
}

// IsButtonPressable is a free data retrieval call binding the contract method 0xd8393150.
//
// Solidity: function isButtonPressable() view returns(bool)
func (_Contract *ContractCaller) IsButtonPressable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isButtonPressable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsButtonPressable is a free data retrieval call binding the contract method 0xd8393150.
//
// Solidity: function isButtonPressable() view returns(bool)
func (_Contract *ContractSession) IsButtonPressable() (bool, error) {
	return _Contract.Contract.IsButtonPressable(&_Contract.CallOpts)
}

// IsButtonPressable is a free data retrieval call binding the contract method 0xd8393150.
//
// Solidity: function isButtonPressable() view returns(bool)
func (_Contract *ContractCallerSession) IsButtonPressable() (bool, error) {
	return _Contract.Contract.IsButtonPressable(&_Contract.CallOpts)
}

// LastTimeButtonPressable is a free data retrieval call binding the contract method 0xd44b1c9e.
//
// Solidity: function lastTimeButtonPressable() view returns(uint256)
func (_Contract *ContractCaller) LastTimeButtonPressable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "lastTimeButtonPressable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTimeButtonPressable is a free data retrieval call binding the contract method 0xd44b1c9e.
//
// Solidity: function lastTimeButtonPressable() view returns(uint256)
func (_Contract *ContractSession) LastTimeButtonPressable() (*big.Int, error) {
	return _Contract.Contract.LastTimeButtonPressable(&_Contract.CallOpts)
}

// LastTimeButtonPressable is a free data retrieval call binding the contract method 0xd44b1c9e.
//
// Solidity: function lastTimeButtonPressable() view returns(uint256)
func (_Contract *ContractCallerSession) LastTimeButtonPressable() (*big.Int, error) {
	return _Contract.Contract.LastTimeButtonPressable(&_Contract.CallOpts)
}

// NextTimeButtonPressable is a free data retrieval call binding the contract method 0xf769479f.
//
// Solidity: function nextTimeButtonPressable() view returns(uint256)
func (_Contract *ContractCaller) NextTimeButtonPressable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "nextTimeButtonPressable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextTimeButtonPressable is a free data retrieval call binding the contract method 0xf769479f.
//
// Solidity: function nextTimeButtonPressable() view returns(uint256)
func (_Contract *ContractSession) NextTimeButtonPressable() (*big.Int, error) {
	return _Contract.Contract.NextTimeButtonPressable(&_Contract.CallOpts)
}

// NextTimeButtonPressable is a free data retrieval call binding the contract method 0xf769479f.
//
// Solidity: function nextTimeButtonPressable() view returns(uint256)
func (_Contract *ContractCallerSession) NextTimeButtonPressable() (*big.Int, error) {
	return _Contract.Contract.NextTimeButtonPressable(&_Contract.CallOpts)
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

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_Contract *ContractCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_Contract *ContractSession) Paused(index uint8) (bool, error) {
	return _Contract.Contract.Paused(&_Contract.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_Contract *ContractCallerSession) Paused(index uint8) (bool, error) {
	return _Contract.Contract.Paused(&_Contract.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_Contract *ContractCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_Contract *ContractSession) Paused0() (*big.Int, error) {
	return _Contract.Contract.Paused0(&_Contract.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_Contract *ContractCallerSession) Paused0() (*big.Int, error) {
	return _Contract.Contract.Paused0(&_Contract.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_Contract *ContractCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_Contract *ContractSession) PauserRegistry() (common.Address, error) {
	return _Contract.Contract.PauserRegistry(&_Contract.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_Contract *ContractCallerSession) PauserRegistry() (common.Address, error) {
	return _Contract.Contract.PauserRegistry(&_Contract.CallOpts)
}

// TotalWeight is a free data retrieval call binding the contract method 0x96c82e57.
//
// Solidity: function totalWeight() view returns(uint16)
func (_Contract *ContractCaller) TotalWeight(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "totalWeight")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// TotalWeight is a free data retrieval call binding the contract method 0x96c82e57.
//
// Solidity: function totalWeight() view returns(uint16)
func (_Contract *ContractSession) TotalWeight() (uint16, error) {
	return _Contract.Contract.TotalWeight(&_Contract.CallOpts)
}

// TotalWeight is a free data retrieval call binding the contract method 0x96c82e57.
//
// Solidity: function totalWeight() view returns(uint16)
func (_Contract *ContractCallerSession) TotalWeight() (uint16, error) {
	return _Contract.Contract.TotalWeight(&_Contract.CallOpts)
}

// AddDistribution is a paid mutator transaction binding the contract method 0xcd1e341b.
//
// Solidity: function addDistribution((uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]) distribution) returns(uint256 distributionId)
func (_Contract *ContractTransactor) AddDistribution(opts *bind.TransactOpts, distribution IEmissionsControllerTypesDistribution) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "addDistribution", distribution)
}

// AddDistribution is a paid mutator transaction binding the contract method 0xcd1e341b.
//
// Solidity: function addDistribution((uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]) distribution) returns(uint256 distributionId)
func (_Contract *ContractSession) AddDistribution(distribution IEmissionsControllerTypesDistribution) (*types.Transaction, error) {
	return _Contract.Contract.AddDistribution(&_Contract.TransactOpts, distribution)
}

// AddDistribution is a paid mutator transaction binding the contract method 0xcd1e341b.
//
// Solidity: function addDistribution((uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]) distribution) returns(uint256 distributionId)
func (_Contract *ContractTransactorSession) AddDistribution(distribution IEmissionsControllerTypesDistribution) (*types.Transaction, error) {
	return _Contract.Contract.AddDistribution(&_Contract.TransactOpts, distribution)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address initialIncentiveCouncil, uint256 initialPausedStatus) returns()
func (_Contract *ContractTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, initialIncentiveCouncil common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "initialize", initialOwner, initialIncentiveCouncil, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address initialIncentiveCouncil, uint256 initialPausedStatus) returns()
func (_Contract *ContractSession) Initialize(initialOwner common.Address, initialIncentiveCouncil common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, initialOwner, initialIncentiveCouncil, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address initialIncentiveCouncil, uint256 initialPausedStatus) returns()
func (_Contract *ContractTransactorSession) Initialize(initialOwner common.Address, initialIncentiveCouncil common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, initialOwner, initialIncentiveCouncil, initialPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_Contract *ContractTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_Contract *ContractSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Pause(&_Contract.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_Contract *ContractTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Pause(&_Contract.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_Contract *ContractTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_Contract *ContractSession) PauseAll() (*types.Transaction, error) {
	return _Contract.Contract.PauseAll(&_Contract.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_Contract *ContractTransactorSession) PauseAll() (*types.Transaction, error) {
	return _Contract.Contract.PauseAll(&_Contract.TransactOpts)
}

// PressButton is a paid mutator transaction binding the contract method 0x400efa85.
//
// Solidity: function pressButton(uint256 length) returns()
func (_Contract *ContractTransactor) PressButton(opts *bind.TransactOpts, length *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "pressButton", length)
}

// PressButton is a paid mutator transaction binding the contract method 0x400efa85.
//
// Solidity: function pressButton(uint256 length) returns()
func (_Contract *ContractSession) PressButton(length *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PressButton(&_Contract.TransactOpts, length)
}

// PressButton is a paid mutator transaction binding the contract method 0x400efa85.
//
// Solidity: function pressButton(uint256 length) returns()
func (_Contract *ContractTransactorSession) PressButton(length *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PressButton(&_Contract.TransactOpts, length)
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

// SetIncentiveCouncil is a paid mutator transaction binding the contract method 0xc695acdb.
//
// Solidity: function setIncentiveCouncil(address newIncentiveCouncil) returns()
func (_Contract *ContractTransactor) SetIncentiveCouncil(opts *bind.TransactOpts, newIncentiveCouncil common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setIncentiveCouncil", newIncentiveCouncil)
}

// SetIncentiveCouncil is a paid mutator transaction binding the contract method 0xc695acdb.
//
// Solidity: function setIncentiveCouncil(address newIncentiveCouncil) returns()
func (_Contract *ContractSession) SetIncentiveCouncil(newIncentiveCouncil common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetIncentiveCouncil(&_Contract.TransactOpts, newIncentiveCouncil)
}

// SetIncentiveCouncil is a paid mutator transaction binding the contract method 0xc695acdb.
//
// Solidity: function setIncentiveCouncil(address newIncentiveCouncil) returns()
func (_Contract *ContractTransactorSession) SetIncentiveCouncil(newIncentiveCouncil common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetIncentiveCouncil(&_Contract.TransactOpts, newIncentiveCouncil)
}

// Sweep is a paid mutator transaction binding the contract method 0x35faa416.
//
// Solidity: function sweep() returns()
func (_Contract *ContractTransactor) Sweep(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "sweep")
}

// Sweep is a paid mutator transaction binding the contract method 0x35faa416.
//
// Solidity: function sweep() returns()
func (_Contract *ContractSession) Sweep() (*types.Transaction, error) {
	return _Contract.Contract.Sweep(&_Contract.TransactOpts)
}

// Sweep is a paid mutator transaction binding the contract method 0x35faa416.
//
// Solidity: function sweep() returns()
func (_Contract *ContractTransactorSession) Sweep() (*types.Transaction, error) {
	return _Contract.Contract.Sweep(&_Contract.TransactOpts)
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

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_Contract *ContractTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_Contract *ContractSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Unpause(&_Contract.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_Contract *ContractTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Unpause(&_Contract.TransactOpts, newPausedStatus)
}

// UpdateDistribution is a paid mutator transaction binding the contract method 0x44a32028.
//
// Solidity: function updateDistribution(uint256 distributionId, (uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]) distribution) returns()
func (_Contract *ContractTransactor) UpdateDistribution(opts *bind.TransactOpts, distributionId *big.Int, distribution IEmissionsControllerTypesDistribution) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateDistribution", distributionId, distribution)
}

// UpdateDistribution is a paid mutator transaction binding the contract method 0x44a32028.
//
// Solidity: function updateDistribution(uint256 distributionId, (uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]) distribution) returns()
func (_Contract *ContractSession) UpdateDistribution(distributionId *big.Int, distribution IEmissionsControllerTypesDistribution) (*types.Transaction, error) {
	return _Contract.Contract.UpdateDistribution(&_Contract.TransactOpts, distributionId, distribution)
}

// UpdateDistribution is a paid mutator transaction binding the contract method 0x44a32028.
//
// Solidity: function updateDistribution(uint256 distributionId, (uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]) distribution) returns()
func (_Contract *ContractTransactorSession) UpdateDistribution(distributionId *big.Int, distribution IEmissionsControllerTypesDistribution) (*types.Transaction, error) {
	return _Contract.Contract.UpdateDistribution(&_Contract.TransactOpts, distributionId, distribution)
}

// ContractDistributionAddedIterator is returned from FilterDistributionAdded and is used to iterate over the raw logs and unpacked data for DistributionAdded events raised by the Contract contract.
type ContractDistributionAddedIterator struct {
	Event *ContractDistributionAdded // Event containing the contract specifics and raw log

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
func (it *ContractDistributionAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDistributionAdded)
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
		it.Event = new(ContractDistributionAdded)
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
func (it *ContractDistributionAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDistributionAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDistributionAdded represents a DistributionAdded event raised by the Contract contract.
type ContractDistributionAdded struct {
	DistributionId *big.Int
	Epoch          *big.Int
	Distribution   IEmissionsControllerTypesDistribution
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDistributionAdded is a free log retrieval operation binding the contract event 0x006f7ba35643ecf5852cfe01b66220b1fe04a4cd4866923d5f3e66c7fcb390ef.
//
// Solidity: event DistributionAdded(uint256 indexed distributionId, uint256 indexed epoch, (uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]) distribution)
func (_Contract *ContractFilterer) FilterDistributionAdded(opts *bind.FilterOpts, distributionId []*big.Int, epoch []*big.Int) (*ContractDistributionAddedIterator, error) {

	var distributionIdRule []interface{}
	for _, distributionIdItem := range distributionId {
		distributionIdRule = append(distributionIdRule, distributionIdItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "DistributionAdded", distributionIdRule, epochRule)
	if err != nil {
		return nil, err
	}
	return &ContractDistributionAddedIterator{contract: _Contract.contract, event: "DistributionAdded", logs: logs, sub: sub}, nil
}

// WatchDistributionAdded is a free log subscription operation binding the contract event 0x006f7ba35643ecf5852cfe01b66220b1fe04a4cd4866923d5f3e66c7fcb390ef.
//
// Solidity: event DistributionAdded(uint256 indexed distributionId, uint256 indexed epoch, (uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]) distribution)
func (_Contract *ContractFilterer) WatchDistributionAdded(opts *bind.WatchOpts, sink chan<- *ContractDistributionAdded, distributionId []*big.Int, epoch []*big.Int) (event.Subscription, error) {

	var distributionIdRule []interface{}
	for _, distributionIdItem := range distributionId {
		distributionIdRule = append(distributionIdRule, distributionIdItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "DistributionAdded", distributionIdRule, epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDistributionAdded)
				if err := _Contract.contract.UnpackLog(event, "DistributionAdded", log); err != nil {
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

// ParseDistributionAdded is a log parse operation binding the contract event 0x006f7ba35643ecf5852cfe01b66220b1fe04a4cd4866923d5f3e66c7fcb390ef.
//
// Solidity: event DistributionAdded(uint256 indexed distributionId, uint256 indexed epoch, (uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]) distribution)
func (_Contract *ContractFilterer) ParseDistributionAdded(log types.Log) (*ContractDistributionAdded, error) {
	event := new(ContractDistributionAdded)
	if err := _Contract.contract.UnpackLog(event, "DistributionAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDistributionProcessedIterator is returned from FilterDistributionProcessed and is used to iterate over the raw logs and unpacked data for DistributionProcessed events raised by the Contract contract.
type ContractDistributionProcessedIterator struct {
	Event *ContractDistributionProcessed // Event containing the contract specifics and raw log

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
func (it *ContractDistributionProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDistributionProcessed)
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
		it.Event = new(ContractDistributionProcessed)
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
func (it *ContractDistributionProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDistributionProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDistributionProcessed represents a DistributionProcessed event raised by the Contract contract.
type ContractDistributionProcessed struct {
	DistributionId *big.Int
	Epoch          *big.Int
	Distribution   IEmissionsControllerTypesDistribution
	Success        bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDistributionProcessed is a free log retrieval operation binding the contract event 0xba5d66336bc9a4f8459242151a4da4d5020ac581243d98403bb55f7f348e071b.
//
// Solidity: event DistributionProcessed(uint256 indexed distributionId, uint256 indexed epoch, (uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]) distribution, bool success)
func (_Contract *ContractFilterer) FilterDistributionProcessed(opts *bind.FilterOpts, distributionId []*big.Int, epoch []*big.Int) (*ContractDistributionProcessedIterator, error) {

	var distributionIdRule []interface{}
	for _, distributionIdItem := range distributionId {
		distributionIdRule = append(distributionIdRule, distributionIdItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "DistributionProcessed", distributionIdRule, epochRule)
	if err != nil {
		return nil, err
	}
	return &ContractDistributionProcessedIterator{contract: _Contract.contract, event: "DistributionProcessed", logs: logs, sub: sub}, nil
}

// WatchDistributionProcessed is a free log subscription operation binding the contract event 0xba5d66336bc9a4f8459242151a4da4d5020ac581243d98403bb55f7f348e071b.
//
// Solidity: event DistributionProcessed(uint256 indexed distributionId, uint256 indexed epoch, (uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]) distribution, bool success)
func (_Contract *ContractFilterer) WatchDistributionProcessed(opts *bind.WatchOpts, sink chan<- *ContractDistributionProcessed, distributionId []*big.Int, epoch []*big.Int) (event.Subscription, error) {

	var distributionIdRule []interface{}
	for _, distributionIdItem := range distributionId {
		distributionIdRule = append(distributionIdRule, distributionIdItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "DistributionProcessed", distributionIdRule, epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDistributionProcessed)
				if err := _Contract.contract.UnpackLog(event, "DistributionProcessed", log); err != nil {
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

// ParseDistributionProcessed is a log parse operation binding the contract event 0xba5d66336bc9a4f8459242151a4da4d5020ac581243d98403bb55f7f348e071b.
//
// Solidity: event DistributionProcessed(uint256 indexed distributionId, uint256 indexed epoch, (uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]) distribution, bool success)
func (_Contract *ContractFilterer) ParseDistributionProcessed(log types.Log) (*ContractDistributionProcessed, error) {
	event := new(ContractDistributionProcessed)
	if err := _Contract.contract.UnpackLog(event, "DistributionProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDistributionUpdatedIterator is returned from FilterDistributionUpdated and is used to iterate over the raw logs and unpacked data for DistributionUpdated events raised by the Contract contract.
type ContractDistributionUpdatedIterator struct {
	Event *ContractDistributionUpdated // Event containing the contract specifics and raw log

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
func (it *ContractDistributionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDistributionUpdated)
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
		it.Event = new(ContractDistributionUpdated)
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
func (it *ContractDistributionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDistributionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDistributionUpdated represents a DistributionUpdated event raised by the Contract contract.
type ContractDistributionUpdated struct {
	DistributionId *big.Int
	Epoch          *big.Int
	Distribution   IEmissionsControllerTypesDistribution
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDistributionUpdated is a free log retrieval operation binding the contract event 0x548fb50d6be978df2bacbf48c6840e4a4743672408921282117f3f00555b2b4c.
//
// Solidity: event DistributionUpdated(uint256 indexed distributionId, uint256 indexed epoch, (uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]) distribution)
func (_Contract *ContractFilterer) FilterDistributionUpdated(opts *bind.FilterOpts, distributionId []*big.Int, epoch []*big.Int) (*ContractDistributionUpdatedIterator, error) {

	var distributionIdRule []interface{}
	for _, distributionIdItem := range distributionId {
		distributionIdRule = append(distributionIdRule, distributionIdItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "DistributionUpdated", distributionIdRule, epochRule)
	if err != nil {
		return nil, err
	}
	return &ContractDistributionUpdatedIterator{contract: _Contract.contract, event: "DistributionUpdated", logs: logs, sub: sub}, nil
}

// WatchDistributionUpdated is a free log subscription operation binding the contract event 0x548fb50d6be978df2bacbf48c6840e4a4743672408921282117f3f00555b2b4c.
//
// Solidity: event DistributionUpdated(uint256 indexed distributionId, uint256 indexed epoch, (uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]) distribution)
func (_Contract *ContractFilterer) WatchDistributionUpdated(opts *bind.WatchOpts, sink chan<- *ContractDistributionUpdated, distributionId []*big.Int, epoch []*big.Int) (event.Subscription, error) {

	var distributionIdRule []interface{}
	for _, distributionIdItem := range distributionId {
		distributionIdRule = append(distributionIdRule, distributionIdItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "DistributionUpdated", distributionIdRule, epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDistributionUpdated)
				if err := _Contract.contract.UnpackLog(event, "DistributionUpdated", log); err != nil {
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

// ParseDistributionUpdated is a log parse operation binding the contract event 0x548fb50d6be978df2bacbf48c6840e4a4743672408921282117f3f00555b2b4c.
//
// Solidity: event DistributionUpdated(uint256 indexed distributionId, uint256 indexed epoch, (uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]) distribution)
func (_Contract *ContractFilterer) ParseDistributionUpdated(log types.Log) (*ContractDistributionUpdated, error) {
	event := new(ContractDistributionUpdated)
	if err := _Contract.contract.UnpackLog(event, "DistributionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncentiveCouncilUpdatedIterator is returned from FilterIncentiveCouncilUpdated and is used to iterate over the raw logs and unpacked data for IncentiveCouncilUpdated events raised by the Contract contract.
type ContractIncentiveCouncilUpdatedIterator struct {
	Event *ContractIncentiveCouncilUpdated // Event containing the contract specifics and raw log

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
func (it *ContractIncentiveCouncilUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncentiveCouncilUpdated)
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
		it.Event = new(ContractIncentiveCouncilUpdated)
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
func (it *ContractIncentiveCouncilUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncentiveCouncilUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncentiveCouncilUpdated represents a IncentiveCouncilUpdated event raised by the Contract contract.
type ContractIncentiveCouncilUpdated struct {
	IncentiveCouncil common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterIncentiveCouncilUpdated is a free log retrieval operation binding the contract event 0x8befac6896b786e67b23cefc473bfabd36e7fc013125c883dfeec8e3a9636216.
//
// Solidity: event IncentiveCouncilUpdated(address indexed incentiveCouncil)
func (_Contract *ContractFilterer) FilterIncentiveCouncilUpdated(opts *bind.FilterOpts, incentiveCouncil []common.Address) (*ContractIncentiveCouncilUpdatedIterator, error) {

	var incentiveCouncilRule []interface{}
	for _, incentiveCouncilItem := range incentiveCouncil {
		incentiveCouncilRule = append(incentiveCouncilRule, incentiveCouncilItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "IncentiveCouncilUpdated", incentiveCouncilRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncentiveCouncilUpdatedIterator{contract: _Contract.contract, event: "IncentiveCouncilUpdated", logs: logs, sub: sub}, nil
}

// WatchIncentiveCouncilUpdated is a free log subscription operation binding the contract event 0x8befac6896b786e67b23cefc473bfabd36e7fc013125c883dfeec8e3a9636216.
//
// Solidity: event IncentiveCouncilUpdated(address indexed incentiveCouncil)
func (_Contract *ContractFilterer) WatchIncentiveCouncilUpdated(opts *bind.WatchOpts, sink chan<- *ContractIncentiveCouncilUpdated, incentiveCouncil []common.Address) (event.Subscription, error) {

	var incentiveCouncilRule []interface{}
	for _, incentiveCouncilItem := range incentiveCouncil {
		incentiveCouncilRule = append(incentiveCouncilRule, incentiveCouncilItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "IncentiveCouncilUpdated", incentiveCouncilRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncentiveCouncilUpdated)
				if err := _Contract.contract.UnpackLog(event, "IncentiveCouncilUpdated", log); err != nil {
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

// ParseIncentiveCouncilUpdated is a log parse operation binding the contract event 0x8befac6896b786e67b23cefc473bfabd36e7fc013125c883dfeec8e3a9636216.
//
// Solidity: event IncentiveCouncilUpdated(address indexed incentiveCouncil)
func (_Contract *ContractFilterer) ParseIncentiveCouncilUpdated(log types.Log) (*ContractIncentiveCouncilUpdated, error) {
	event := new(ContractIncentiveCouncilUpdated)
	if err := _Contract.contract.UnpackLog(event, "IncentiveCouncilUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Contract contract.
type ContractInitializedIterator struct {
	Event *ContractInitialized // Event containing the contract specifics and raw log

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
func (it *ContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractInitialized)
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
		it.Event = new(ContractInitialized)
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
func (it *ContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractInitialized represents a Initialized event raised by the Contract contract.
type ContractInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Contract *ContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractInitializedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractInitializedIterator{contract: _Contract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Contract *ContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractInitialized) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractInitialized)
				if err := _Contract.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Contract *ContractFilterer) ParseInitialized(log types.Log) (*ContractInitialized, error) {
	event := new(ContractInitialized)
	if err := _Contract.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ContractPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Contract contract.
type ContractPausedIterator struct {
	Event *ContractPaused // Event containing the contract specifics and raw log

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
func (it *ContractPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPaused)
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
		it.Event = new(ContractPaused)
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
func (it *ContractPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPaused represents a Paused event raised by the Contract contract.
type ContractPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_Contract *ContractFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractPausedIterator{contract: _Contract.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_Contract *ContractFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPaused)
				if err := _Contract.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_Contract *ContractFilterer) ParsePaused(log types.Log) (*ContractPaused, error) {
	event := new(ContractPaused)
	if err := _Contract.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSweptIterator is returned from FilterSwept and is used to iterate over the raw logs and unpacked data for Swept events raised by the Contract contract.
type ContractSweptIterator struct {
	Event *ContractSwept // Event containing the contract specifics and raw log

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
func (it *ContractSweptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSwept)
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
		it.Event = new(ContractSwept)
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
func (it *ContractSweptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSweptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSwept represents a Swept event raised by the Contract contract.
type ContractSwept struct {
	IncentiveCouncil common.Address
	Amount           *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSwept is a free log retrieval operation binding the contract event 0xc36b5179cb9c303b200074996eab2b3473eac370fdd7eba3bec636fe35109696.
//
// Solidity: event Swept(address indexed incentiveCouncil, uint256 amount)
func (_Contract *ContractFilterer) FilterSwept(opts *bind.FilterOpts, incentiveCouncil []common.Address) (*ContractSweptIterator, error) {

	var incentiveCouncilRule []interface{}
	for _, incentiveCouncilItem := range incentiveCouncil {
		incentiveCouncilRule = append(incentiveCouncilRule, incentiveCouncilItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Swept", incentiveCouncilRule)
	if err != nil {
		return nil, err
	}
	return &ContractSweptIterator{contract: _Contract.contract, event: "Swept", logs: logs, sub: sub}, nil
}

// WatchSwept is a free log subscription operation binding the contract event 0xc36b5179cb9c303b200074996eab2b3473eac370fdd7eba3bec636fe35109696.
//
// Solidity: event Swept(address indexed incentiveCouncil, uint256 amount)
func (_Contract *ContractFilterer) WatchSwept(opts *bind.WatchOpts, sink chan<- *ContractSwept, incentiveCouncil []common.Address) (event.Subscription, error) {

	var incentiveCouncilRule []interface{}
	for _, incentiveCouncilItem := range incentiveCouncil {
		incentiveCouncilRule = append(incentiveCouncilRule, incentiveCouncilItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Swept", incentiveCouncilRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSwept)
				if err := _Contract.contract.UnpackLog(event, "Swept", log); err != nil {
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

// ParseSwept is a log parse operation binding the contract event 0xc36b5179cb9c303b200074996eab2b3473eac370fdd7eba3bec636fe35109696.
//
// Solidity: event Swept(address indexed incentiveCouncil, uint256 amount)
func (_Contract *ContractFilterer) ParseSwept(log types.Log) (*ContractSwept, error) {
	event := new(ContractSwept)
	if err := _Contract.contract.UnpackLog(event, "Swept", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Contract contract.
type ContractUnpausedIterator struct {
	Event *ContractUnpaused // Event containing the contract specifics and raw log

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
func (it *ContractUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUnpaused)
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
		it.Event = new(ContractUnpaused)
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
func (it *ContractUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUnpaused represents a Unpaused event raised by the Contract contract.
type ContractUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_Contract *ContractFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractUnpausedIterator{contract: _Contract.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_Contract *ContractFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUnpaused)
				if err := _Contract.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_Contract *ContractFilterer) ParseUnpaused(log types.Log) (*ContractUnpaused, error) {
	event := new(ContractUnpaused)
	if err := _Contract.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
