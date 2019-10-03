// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mapping

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

// MappingABI is the input ABI used to generate the binding from.
const MappingABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"bytesTest\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_number\",\"type\":\"uint256\"},{\"name\":\"_phrase\",\"type\":\"string\"}],\"name\":\"setTestData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"addressTest\",\"outputs\":[{\"name\":\"number\",\"type\":\"uint256\"},{\"name\":\"phrase\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"bytes32\"},{\"name\":\"setBool\",\"type\":\"bool\"}],\"name\":\"setBytesBool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"testAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"num\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"saidWhat\",\"type\":\"string\"}],\"name\":\"LogSetData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"value\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"boolVal\",\"type\":\"bool\"}],\"name\":\"LogSetBool\",\"type\":\"event\"}]"

// MappingBin is the compiled bytecode used for deploying new contracts.
const MappingBin = `608060405234801561001057600080fd5b50604051610a19380380610a198339810180604052602081101561003357600080fd5b81019080805164010000000081111561004b57600080fd5b8281019050602081018481111561006157600080fd5b815185600182028301116401000000008211171561007e57600080fd5b505092919050505033600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600090805190602001906100dd9291906100e4565b5050610189565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061012557805160ff1916838001178555610153565b82800160010185558215610153579182015b82811115610152578251825591602001919060010190610137565b5b5090506101609190610164565b5090565b61018691905b8082111561018257600081600090555060010161016a565b5090565b90565b610881806101986000396000f3fe608060405234801561001057600080fd5b506004361061007f576000357c01000000000000000000000000000000000000000000000000000000009004806306fdde03146100845780631f504c5e1461010757806352b6b6e01461014d578063893d20e814610232578063ac59c9601461027c578063f23bc4f814610340575b600080fd5b61008c61037a565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156100cc5780820151818401526020810190506100b1565b50505050905090810190601f1680156100f95780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101336004803603602081101561011d57600080fd5b8101908080359060200190929190505050610418565b604051808215151515815260200191505060405180910390f35b6102306004803603606081101561016357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001906401000000008111156101aa57600080fd5b8201836020820111156101bc57600080fd5b803590602001918460018302840111640100000000831117156101de57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610438565b005b61023a610658565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6102be6004803603602081101561029257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610682565b6040518083815260200180602001828103825283818151815260200191508051906020019080838360005b838110156103045780820151818401526020810190506102e9565b50505050905090810190601f1680156103315780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b6103786004803603604081101561035657600080fd5b810190808035906020019092919080351515906020019092919050505061073e565b005b60008054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104105780601f106103e557610100808354040283529160200191610410565b820191906000526020600020905b8154815290600101906020018083116103f357829003601f168201915b505050505081565b60036020528060005260406000206000915054906101000a900460ff1681565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141515156104dd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600d8152602001807f6164647265737320656d7074790000000000000000000000000000000000000081525060200191505060405180910390fd5b81600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000018190555080600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101908051906020019061057a9291906107b0565b507f35e281515e5c697280749cf3bf65d1b3fa3774d620055b049f80ddc6ba1fcdfd838383604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156106175780820151818401526020810190506105fc565b50505050905090810190601f1680156106445780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a1505050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6002602052806000526040600020600091509050806000015490806001018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107345780601f1061070957610100808354040283529160200191610734565b820191906000526020600020905b81548152906001019060200180831161071757829003601f168201915b5050505050905082565b806003600084815260200190815260200160002060006101000a81548160ff0219169083151502179055507fe3a91865612aed23ede22b64c579423a007c7c863b8738c02e26ef7c2d72158e828260405180838152602001821515151581526020019250505060405180910390a15050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106107f157805160ff191683800117855561081f565b8280016001018555821561081f579182015b8281111561081e578251825591602001919060010190610803565b5b50905061082c9190610830565b5090565b61085291905b8082111561084e576000816000905550600101610836565b5090565b9056fea165627a7a72305820c6aec99e148ace2ddf7b215dbc35a9adc22c31d0b2b7e7d820e531d980a0040b0029`

// DeployMapping deploys a new Ethereum contract, binding an instance of Mapping to it.
func DeployMapping(auth *bind.TransactOpts, backend bind.ContractBackend, _name string) (common.Address, *types.Transaction, *Mapping, error) {
	parsed, err := abi.JSON(strings.NewReader(MappingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MappingBin), backend, _name)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Mapping{MappingCaller: MappingCaller{contract: contract}, MappingTransactor: MappingTransactor{contract: contract}, MappingFilterer: MappingFilterer{contract: contract}}, nil
}

// Mapping is an auto generated Go binding around an Ethereum contract.
type Mapping struct {
	MappingCaller     // Read-only binding to the contract
	MappingTransactor // Write-only binding to the contract
	MappingFilterer   // Log filterer for contract events
}

// MappingCaller is an auto generated read-only Go binding around an Ethereum contract.
type MappingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MappingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MappingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MappingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MappingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MappingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MappingSession struct {
	Contract     *Mapping          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MappingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MappingCallerSession struct {
	Contract *MappingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// MappingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MappingTransactorSession struct {
	Contract     *MappingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MappingRaw is an auto generated low-level Go binding around an Ethereum contract.
type MappingRaw struct {
	Contract *Mapping // Generic contract binding to access the raw methods on
}

// MappingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MappingCallerRaw struct {
	Contract *MappingCaller // Generic read-only contract binding to access the raw methods on
}

// MappingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MappingTransactorRaw struct {
	Contract *MappingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMapping creates a new instance of Mapping, bound to a specific deployed contract.
func NewMapping(address common.Address, backend bind.ContractBackend) (*Mapping, error) {
	contract, err := bindMapping(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mapping{MappingCaller: MappingCaller{contract: contract}, MappingTransactor: MappingTransactor{contract: contract}, MappingFilterer: MappingFilterer{contract: contract}}, nil
}

// NewMappingCaller creates a new read-only instance of Mapping, bound to a specific deployed contract.
func NewMappingCaller(address common.Address, caller bind.ContractCaller) (*MappingCaller, error) {
	contract, err := bindMapping(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MappingCaller{contract: contract}, nil
}

// NewMappingTransactor creates a new write-only instance of Mapping, bound to a specific deployed contract.
func NewMappingTransactor(address common.Address, transactor bind.ContractTransactor) (*MappingTransactor, error) {
	contract, err := bindMapping(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MappingTransactor{contract: contract}, nil
}

// NewMappingFilterer creates a new log filterer instance of Mapping, bound to a specific deployed contract.
func NewMappingFilterer(address common.Address, filterer bind.ContractFilterer) (*MappingFilterer, error) {
	contract, err := bindMapping(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MappingFilterer{contract: contract}, nil
}

// bindMapping binds a generic wrapper to an already deployed contract.
func bindMapping(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MappingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mapping *MappingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Mapping.Contract.MappingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mapping *MappingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mapping.Contract.MappingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mapping *MappingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mapping.Contract.MappingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mapping *MappingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Mapping.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mapping *MappingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mapping.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mapping *MappingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mapping.Contract.contract.Transact(opts, method, params...)
}

// AddressTest is a free data retrieval call binding the contract method 0xac59c960.
//
// Solidity: function addressTest(address ) constant returns(uint256 number, string phrase)
func (_Mapping *MappingCaller) AddressTest(opts *bind.CallOpts, arg0 common.Address) (struct {
	Number *big.Int
	Phrase string
}, error) {
	ret := new(struct {
		Number *big.Int
		Phrase string
	})
	out := ret
	err := _Mapping.contract.Call(opts, out, "addressTest", arg0)
	return *ret, err
}

// AddressTest is a free data retrieval call binding the contract method 0xac59c960.
//
// Solidity: function addressTest(address ) constant returns(uint256 number, string phrase)
func (_Mapping *MappingSession) AddressTest(arg0 common.Address) (struct {
	Number *big.Int
	Phrase string
}, error) {
	return _Mapping.Contract.AddressTest(&_Mapping.CallOpts, arg0)
}

// AddressTest is a free data retrieval call binding the contract method 0xac59c960.
//
// Solidity: function addressTest(address ) constant returns(uint256 number, string phrase)
func (_Mapping *MappingCallerSession) AddressTest(arg0 common.Address) (struct {
	Number *big.Int
	Phrase string
}, error) {
	return _Mapping.Contract.AddressTest(&_Mapping.CallOpts, arg0)
}

// BytesTest is a free data retrieval call binding the contract method 0x1f504c5e.
//
// Solidity: function bytesTest(bytes32 ) constant returns(bool)
func (_Mapping *MappingCaller) BytesTest(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Mapping.contract.Call(opts, out, "bytesTest", arg0)
	return *ret0, err
}

// BytesTest is a free data retrieval call binding the contract method 0x1f504c5e.
//
// Solidity: function bytesTest(bytes32 ) constant returns(bool)
func (_Mapping *MappingSession) BytesTest(arg0 [32]byte) (bool, error) {
	return _Mapping.Contract.BytesTest(&_Mapping.CallOpts, arg0)
}

// BytesTest is a free data retrieval call binding the contract method 0x1f504c5e.
//
// Solidity: function bytesTest(bytes32 ) constant returns(bool)
func (_Mapping *MappingCallerSession) BytesTest(arg0 [32]byte) (bool, error) {
	return _Mapping.Contract.BytesTest(&_Mapping.CallOpts, arg0)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Mapping *MappingCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Mapping.contract.Call(opts, out, "getOwner")
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Mapping *MappingSession) GetOwner() (common.Address, error) {
	return _Mapping.Contract.GetOwner(&_Mapping.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Mapping *MappingCallerSession) GetOwner() (common.Address, error) {
	return _Mapping.Contract.GetOwner(&_Mapping.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Mapping *MappingCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Mapping.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Mapping *MappingSession) Name() (string, error) {
	return _Mapping.Contract.Name(&_Mapping.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Mapping *MappingCallerSession) Name() (string, error) {
	return _Mapping.Contract.Name(&_Mapping.CallOpts)
}

// SetBytesBool is a paid mutator transaction binding the contract method 0xf23bc4f8.
//
// Solidity: function setBytesBool(bytes32 value, bool setBool) returns()
func (_Mapping *MappingTransactor) SetBytesBool(opts *bind.TransactOpts, value [32]byte, setBool bool) (*types.Transaction, error) {
	return _Mapping.contract.Transact(opts, "setBytesBool", value, setBool)
}

// SetBytesBool is a paid mutator transaction binding the contract method 0xf23bc4f8.
//
// Solidity: function setBytesBool(bytes32 value, bool setBool) returns()
func (_Mapping *MappingSession) SetBytesBool(value [32]byte, setBool bool) (*types.Transaction, error) {
	return _Mapping.Contract.SetBytesBool(&_Mapping.TransactOpts, value, setBool)
}

// SetBytesBool is a paid mutator transaction binding the contract method 0xf23bc4f8.
//
// Solidity: function setBytesBool(bytes32 value, bool setBool) returns()
func (_Mapping *MappingTransactorSession) SetBytesBool(value [32]byte, setBool bool) (*types.Transaction, error) {
	return _Mapping.Contract.SetBytesBool(&_Mapping.TransactOpts, value, setBool)
}

// SetTestData is a paid mutator transaction binding the contract method 0x52b6b6e0.
//
// Solidity: function setTestData(address _address, uint256 _number, string _phrase) returns()
func (_Mapping *MappingTransactor) SetTestData(opts *bind.TransactOpts, _address common.Address, _number *big.Int, _phrase string) (*types.Transaction, error) {
	return _Mapping.contract.Transact(opts, "setTestData", _address, _number, _phrase)
}

// SetTestData is a paid mutator transaction binding the contract method 0x52b6b6e0.
//
// Solidity: function setTestData(address _address, uint256 _number, string _phrase) returns()
func (_Mapping *MappingSession) SetTestData(_address common.Address, _number *big.Int, _phrase string) (*types.Transaction, error) {
	return _Mapping.Contract.SetTestData(&_Mapping.TransactOpts, _address, _number, _phrase)
}

// SetTestData is a paid mutator transaction binding the contract method 0x52b6b6e0.
//
// Solidity: function setTestData(address _address, uint256 _number, string _phrase) returns()
func (_Mapping *MappingTransactorSession) SetTestData(_address common.Address, _number *big.Int, _phrase string) (*types.Transaction, error) {
	return _Mapping.Contract.SetTestData(&_Mapping.TransactOpts, _address, _number, _phrase)
}

// MappingLogSetBoolIterator is returned from FilterLogSetBool and is used to iterate over the raw logs and unpacked data for LogSetBool events raised by the Mapping contract.
type MappingLogSetBoolIterator struct {
	Event *MappingLogSetBool // Event containing the contract specifics and raw log

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
func (it *MappingLogSetBoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MappingLogSetBool)
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
		it.Event = new(MappingLogSetBool)
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
func (it *MappingLogSetBoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MappingLogSetBoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MappingLogSetBool represents a LogSetBool event raised by the Mapping contract.
type MappingLogSetBool struct {
	Value   [32]byte
	BoolVal bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogSetBool is a free log retrieval operation binding the contract event 0xe3a91865612aed23ede22b64c579423a007c7c863b8738c02e26ef7c2d72158e.
//
// Solidity: event LogSetBool(bytes32 value, bool boolVal)
func (_Mapping *MappingFilterer) FilterLogSetBool(opts *bind.FilterOpts) (*MappingLogSetBoolIterator, error) {

	logs, sub, err := _Mapping.contract.FilterLogs(opts, "LogSetBool")
	if err != nil {
		return nil, err
	}
	return &MappingLogSetBoolIterator{contract: _Mapping.contract, event: "LogSetBool", logs: logs, sub: sub}, nil
}

// WatchLogSetBool is a free log subscription operation binding the contract event 0xe3a91865612aed23ede22b64c579423a007c7c863b8738c02e26ef7c2d72158e.
//
// Solidity: event LogSetBool(bytes32 value, bool boolVal)
func (_Mapping *MappingFilterer) WatchLogSetBool(opts *bind.WatchOpts, sink chan<- *MappingLogSetBool) (event.Subscription, error) {

	logs, sub, err := _Mapping.contract.WatchLogs(opts, "LogSetBool")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MappingLogSetBool)
				if err := _Mapping.contract.UnpackLog(event, "LogSetBool", log); err != nil {
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

// MappingLogSetDataIterator is returned from FilterLogSetData and is used to iterate over the raw logs and unpacked data for LogSetData events raised by the Mapping contract.
type MappingLogSetDataIterator struct {
	Event *MappingLogSetData // Event containing the contract specifics and raw log

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
func (it *MappingLogSetDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MappingLogSetData)
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
		it.Event = new(MappingLogSetData)
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
func (it *MappingLogSetDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MappingLogSetDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MappingLogSetData represents a LogSetData event raised by the Mapping contract.
type MappingLogSetData struct {
	TestAddr common.Address
	Num      *big.Int
	SaidWhat string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogSetData is a free log retrieval operation binding the contract event 0x35e281515e5c697280749cf3bf65d1b3fa3774d620055b049f80ddc6ba1fcdfd.
//
// Solidity: event LogSetData(address testAddr, uint256 num, string saidWhat)
func (_Mapping *MappingFilterer) FilterLogSetData(opts *bind.FilterOpts) (*MappingLogSetDataIterator, error) {

	logs, sub, err := _Mapping.contract.FilterLogs(opts, "LogSetData")
	if err != nil {
		return nil, err
	}
	return &MappingLogSetDataIterator{contract: _Mapping.contract, event: "LogSetData", logs: logs, sub: sub}, nil
}

// WatchLogSetData is a free log subscription operation binding the contract event 0x35e281515e5c697280749cf3bf65d1b3fa3774d620055b049f80ddc6ba1fcdfd.
//
// Solidity: event LogSetData(address testAddr, uint256 num, string saidWhat)
func (_Mapping *MappingFilterer) WatchLogSetData(opts *bind.WatchOpts, sink chan<- *MappingLogSetData) (event.Subscription, error) {

	logs, sub, err := _Mapping.contract.WatchLogs(opts, "LogSetData")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MappingLogSetData)
				if err := _Mapping.contract.UnpackLog(event, "LogSetData", log); err != nil {
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
