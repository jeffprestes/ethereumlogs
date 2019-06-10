// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package crud

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

// UserCrudABI is the input ABI used to generate the binding from.
const UserCrudABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"userAddress\",\"type\":\"address\"},{\"name\":\"userAge\",\"type\":\"uint256\"}],\"name\":\"updateUserAge\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"isUser\",\"outputs\":[{\"name\":\"isIndeed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"userAddress\",\"type\":\"address\"},{\"name\":\"userEmail\",\"type\":\"string\"},{\"name\":\"userAge\",\"type\":\"uint256\"}],\"name\":\"insertUser\",\"outputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"deleteUser\",\"outputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"userAddress\",\"type\":\"address\"},{\"name\":\"userEmail\",\"type\":\"string\"}],\"name\":\"updateUserEmail\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"getUser\",\"outputs\":[{\"name\":\"userEmail\",\"type\":\"string\"},{\"name\":\"userAge\",\"type\":\"uint256\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getUserAddressAtIndex\",\"outputs\":[{\"name\":\"userAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getUserCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// UserCrudBin is the compiled bytecode used for deploying new contracts.
const UserCrudBin = `0x608060405234801561001057600080fd5b506108fb806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c806364319ae61161005b57806364319ae6146101e35780636f77926b146102995780639eaa2b4b14610345578063b5cb15f71461037e57610088565b80632e071db31461008d5780634209fff1146100cd57806352b1e14a146100f35780635c60f226146101bd575b600080fd5b6100b9600480360360408110156100a357600080fd5b506001600160a01b038135169060200135610386565b604080519115158252519081900360200190f35b6100b9600480360360208110156100e357600080fd5b50356001600160a01b0316610401565b6101ab6004803603606081101561010957600080fd5b6001600160a01b03823516919081019060408101602082013564010000000081111561013457600080fd5b82018360208201111561014657600080fd5b8035906020019184600183028401116401000000008311171561016857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550509135925061045c915050565b60408051918252519081900360200190f35b6101ab600480360360208110156101d357600080fd5b50356001600160a01b0316610542565b6100b9600480360360408110156101f957600080fd5b6001600160a01b03823516919081019060408101602082013564010000000081111561022457600080fd5b82018360208201111561023657600080fd5b8035906020019184600183028401116401000000008311171561025857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610647945050505050565b6102bf600480360360208110156102af57600080fd5b50356001600160a01b03166106cc565b6040518080602001848152602001838152602001828103825285818151815260200191508051906020019080838360005b838110156103085781810151838201526020016102f0565b50505050905090810190601f1680156103355780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b6103626004803603602081101561035b57600080fd5b50356107dd565b604080516001600160a01b039092168252519081900360200190f35b6101ab610807565b600061039183610401565b6103da5760408051600160e51b62461bcd02815260206004820152600f6024820152600160881b6e75736572206e6f742065786973747302604482015290519081900360640190fd5b506001600160a01b0382166000908152602081905260409020600190810182905592915050565b60015460009061041357506000610457565b6001600160a01b03821660008181526020819052604090206002015460018054909190811061043e57fe5b6000918252602090912001546001600160a01b03161490505b919050565b600061046784610401565b156104ad5760408051600160e51b62461bcd02815260206004820152600b6024820152600160a81b6a757365722065786973747302604482015290519081900360640190fd5b6001600160a01b03841660009081526020818152604090912084516104d49286019061080e565b50506001600160a01b03929092166000818152602081905260409020600180820194909455835480850185557fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6810180546001600160a01b0319169093179092556002015550546000190190565b600061054d82610401565b6105965760408051600160e51b62461bcd02815260206004820152600f6024820152600160881b6e75736572206e6f742065786973747302604482015290519081900360640190fd5b6001600160a01b0382166000908152602081905260408120600201546001805491929160001981019081106105c757fe5b600091825260209091200154600180546001600160a01b0390921692508291849081106105f057fe5b600091825260208083209190910180546001600160a01b0319166001600160a01b0394851617905591831681529081905260409020600201829055600180549061063e90600019830161088c565b50909392505050565b600061065283610401565b61069b5760408051600160e51b62461bcd02815260206004820152600f6024820152600160881b6e75736572206e6f742065786973747302604482015290519081900360640190fd5b6001600160a01b03831660009081526020818152604090912083516106c29285019061080e565b5060019392505050565b60606000806106da84610401565b6107235760408051600160e51b62461bcd02815260206004820152600f6024820152600160881b6e75736572206e6f742065786973747302604482015290519081900360640190fd5b6001600160a01b03841660009081526020818152604091829020600180820154600280840154845487519481161561010002600019011691909104601f8101869004860284018601909652858352929490938591908301828280156107c95780601f1061079e576101008083540402835291602001916107c9565b820191906000526020600020905b8154815290600101906020018083116107ac57829003601f168201915b505050505092509250925092509193909250565b6000600182815481106107ec57fe5b6000918252602090912001546001600160a01b031692915050565b6001545b90565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061084f57805160ff191683800117855561087c565b8280016001018555821561087c579182015b8281111561087c578251825591602001919060010190610861565b506108889291506108b5565b5090565b8154818355818111156108b0576000838152602090206108b09181019083016108b5565b505050565b61080b91905b8082111561088857600081556001016108bb56fea165627a7a72305820bd0f82275b10bfac9b6fd2431d5edcbbfea7badc7b043ee3c14cca96e42f27e10029`

// DeployUserCrud deploys a new Ethereum contract, binding an instance of UserCrud to it.
func DeployUserCrud(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UserCrud, error) {
	parsed, err := abi.JSON(strings.NewReader(UserCrudABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UserCrudBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UserCrud{UserCrudCaller: UserCrudCaller{contract: contract}, UserCrudTransactor: UserCrudTransactor{contract: contract}, UserCrudFilterer: UserCrudFilterer{contract: contract}}, nil
}

// UserCrud is an auto generated Go binding around an Ethereum contract.
type UserCrud struct {
	UserCrudCaller     // Read-only binding to the contract
	UserCrudTransactor // Write-only binding to the contract
	UserCrudFilterer   // Log filterer for contract events
}

// UserCrudCaller is an auto generated read-only Go binding around an Ethereum contract.
type UserCrudCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserCrudTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UserCrudTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserCrudFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UserCrudFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserCrudSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UserCrudSession struct {
	Contract     *UserCrud         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserCrudCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UserCrudCallerSession struct {
	Contract *UserCrudCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// UserCrudTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UserCrudTransactorSession struct {
	Contract     *UserCrudTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// UserCrudRaw is an auto generated low-level Go binding around an Ethereum contract.
type UserCrudRaw struct {
	Contract *UserCrud // Generic contract binding to access the raw methods on
}

// UserCrudCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UserCrudCallerRaw struct {
	Contract *UserCrudCaller // Generic read-only contract binding to access the raw methods on
}

// UserCrudTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UserCrudTransactorRaw struct {
	Contract *UserCrudTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUserCrud creates a new instance of UserCrud, bound to a specific deployed contract.
func NewUserCrud(address common.Address, backend bind.ContractBackend) (*UserCrud, error) {
	contract, err := bindUserCrud(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UserCrud{UserCrudCaller: UserCrudCaller{contract: contract}, UserCrudTransactor: UserCrudTransactor{contract: contract}, UserCrudFilterer: UserCrudFilterer{contract: contract}}, nil
}

// NewUserCrudCaller creates a new read-only instance of UserCrud, bound to a specific deployed contract.
func NewUserCrudCaller(address common.Address, caller bind.ContractCaller) (*UserCrudCaller, error) {
	contract, err := bindUserCrud(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UserCrudCaller{contract: contract}, nil
}

// NewUserCrudTransactor creates a new write-only instance of UserCrud, bound to a specific deployed contract.
func NewUserCrudTransactor(address common.Address, transactor bind.ContractTransactor) (*UserCrudTransactor, error) {
	contract, err := bindUserCrud(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UserCrudTransactor{contract: contract}, nil
}

// NewUserCrudFilterer creates a new log filterer instance of UserCrud, bound to a specific deployed contract.
func NewUserCrudFilterer(address common.Address, filterer bind.ContractFilterer) (*UserCrudFilterer, error) {
	contract, err := bindUserCrud(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UserCrudFilterer{contract: contract}, nil
}

// bindUserCrud binds a generic wrapper to an already deployed contract.
func bindUserCrud(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UserCrudABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserCrud *UserCrudRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UserCrud.Contract.UserCrudCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserCrud *UserCrudRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserCrud.Contract.UserCrudTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserCrud *UserCrudRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserCrud.Contract.UserCrudTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserCrud *UserCrudCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UserCrud.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserCrud *UserCrudTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserCrud.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserCrud *UserCrudTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserCrud.Contract.contract.Transact(opts, method, params...)
}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address userAddress) constant returns(string userEmail, uint256 userAge, uint256 index)
func (_UserCrud *UserCrudCaller) GetUser(opts *bind.CallOpts, userAddress common.Address) (struct {
	UserEmail string
	UserAge   *big.Int
	Index     *big.Int
}, error) {
	ret := new(struct {
		UserEmail string
		UserAge   *big.Int
		Index     *big.Int
	})
	out := ret
	err := _UserCrud.contract.Call(opts, out, "getUser", userAddress)
	return *ret, err
}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address userAddress) constant returns(string userEmail, uint256 userAge, uint256 index)
func (_UserCrud *UserCrudSession) GetUser(userAddress common.Address) (struct {
	UserEmail string
	UserAge   *big.Int
	Index     *big.Int
}, error) {
	return _UserCrud.Contract.GetUser(&_UserCrud.CallOpts, userAddress)
}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address userAddress) constant returns(string userEmail, uint256 userAge, uint256 index)
func (_UserCrud *UserCrudCallerSession) GetUser(userAddress common.Address) (struct {
	UserEmail string
	UserAge   *big.Int
	Index     *big.Int
}, error) {
	return _UserCrud.Contract.GetUser(&_UserCrud.CallOpts, userAddress)
}

// GetUserAddressAtIndex is a free data retrieval call binding the contract method 0x9eaa2b4b.
//
// Solidity: function getUserAddressAtIndex(uint256 index) constant returns(address userAddress)
func (_UserCrud *UserCrudCaller) GetUserAddressAtIndex(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UserCrud.contract.Call(opts, out, "getUserAddressAtIndex", index)
	return *ret0, err
}

// GetUserAddressAtIndex is a free data retrieval call binding the contract method 0x9eaa2b4b.
//
// Solidity: function getUserAddressAtIndex(uint256 index) constant returns(address userAddress)
func (_UserCrud *UserCrudSession) GetUserAddressAtIndex(index *big.Int) (common.Address, error) {
	return _UserCrud.Contract.GetUserAddressAtIndex(&_UserCrud.CallOpts, index)
}

// GetUserAddressAtIndex is a free data retrieval call binding the contract method 0x9eaa2b4b.
//
// Solidity: function getUserAddressAtIndex(uint256 index) constant returns(address userAddress)
func (_UserCrud *UserCrudCallerSession) GetUserAddressAtIndex(index *big.Int) (common.Address, error) {
	return _UserCrud.Contract.GetUserAddressAtIndex(&_UserCrud.CallOpts, index)
}

// GetUserCount is a free data retrieval call binding the contract method 0xb5cb15f7.
//
// Solidity: function getUserCount() constant returns(uint256 count)
func (_UserCrud *UserCrudCaller) GetUserCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UserCrud.contract.Call(opts, out, "getUserCount")
	return *ret0, err
}

// GetUserCount is a free data retrieval call binding the contract method 0xb5cb15f7.
//
// Solidity: function getUserCount() constant returns(uint256 count)
func (_UserCrud *UserCrudSession) GetUserCount() (*big.Int, error) {
	return _UserCrud.Contract.GetUserCount(&_UserCrud.CallOpts)
}

// GetUserCount is a free data retrieval call binding the contract method 0xb5cb15f7.
//
// Solidity: function getUserCount() constant returns(uint256 count)
func (_UserCrud *UserCrudCallerSession) GetUserCount() (*big.Int, error) {
	return _UserCrud.Contract.GetUserCount(&_UserCrud.CallOpts)
}

// IsUser is a free data retrieval call binding the contract method 0x4209fff1.
//
// Solidity: function isUser(address userAddress) constant returns(bool isIndeed)
func (_UserCrud *UserCrudCaller) IsUser(opts *bind.CallOpts, userAddress common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _UserCrud.contract.Call(opts, out, "isUser", userAddress)
	return *ret0, err
}

// IsUser is a free data retrieval call binding the contract method 0x4209fff1.
//
// Solidity: function isUser(address userAddress) constant returns(bool isIndeed)
func (_UserCrud *UserCrudSession) IsUser(userAddress common.Address) (bool, error) {
	return _UserCrud.Contract.IsUser(&_UserCrud.CallOpts, userAddress)
}

// IsUser is a free data retrieval call binding the contract method 0x4209fff1.
//
// Solidity: function isUser(address userAddress) constant returns(bool isIndeed)
func (_UserCrud *UserCrudCallerSession) IsUser(userAddress common.Address) (bool, error) {
	return _UserCrud.Contract.IsUser(&_UserCrud.CallOpts, userAddress)
}

// DeleteUser is a paid mutator transaction binding the contract method 0x5c60f226.
//
// Solidity: function deleteUser(address userAddress) returns(uint256 index)
func (_UserCrud *UserCrudTransactor) DeleteUser(opts *bind.TransactOpts, userAddress common.Address) (*types.Transaction, error) {
	return _UserCrud.contract.Transact(opts, "deleteUser", userAddress)
}

// DeleteUser is a paid mutator transaction binding the contract method 0x5c60f226.
//
// Solidity: function deleteUser(address userAddress) returns(uint256 index)
func (_UserCrud *UserCrudSession) DeleteUser(userAddress common.Address) (*types.Transaction, error) {
	return _UserCrud.Contract.DeleteUser(&_UserCrud.TransactOpts, userAddress)
}

// DeleteUser is a paid mutator transaction binding the contract method 0x5c60f226.
//
// Solidity: function deleteUser(address userAddress) returns(uint256 index)
func (_UserCrud *UserCrudTransactorSession) DeleteUser(userAddress common.Address) (*types.Transaction, error) {
	return _UserCrud.Contract.DeleteUser(&_UserCrud.TransactOpts, userAddress)
}

// InsertUser is a paid mutator transaction binding the contract method 0x52b1e14a.
//
// Solidity: function insertUser(address userAddress, string userEmail, uint256 userAge) returns(uint256 index)
func (_UserCrud *UserCrudTransactor) InsertUser(opts *bind.TransactOpts, userAddress common.Address, userEmail string, userAge *big.Int) (*types.Transaction, error) {
	return _UserCrud.contract.Transact(opts, "insertUser", userAddress, userEmail, userAge)
}

// InsertUser is a paid mutator transaction binding the contract method 0x52b1e14a.
//
// Solidity: function insertUser(address userAddress, string userEmail, uint256 userAge) returns(uint256 index)
func (_UserCrud *UserCrudSession) InsertUser(userAddress common.Address, userEmail string, userAge *big.Int) (*types.Transaction, error) {
	return _UserCrud.Contract.InsertUser(&_UserCrud.TransactOpts, userAddress, userEmail, userAge)
}

// InsertUser is a paid mutator transaction binding the contract method 0x52b1e14a.
//
// Solidity: function insertUser(address userAddress, string userEmail, uint256 userAge) returns(uint256 index)
func (_UserCrud *UserCrudTransactorSession) InsertUser(userAddress common.Address, userEmail string, userAge *big.Int) (*types.Transaction, error) {
	return _UserCrud.Contract.InsertUser(&_UserCrud.TransactOpts, userAddress, userEmail, userAge)
}

// UpdateUserAge is a paid mutator transaction binding the contract method 0x2e071db3.
//
// Solidity: function updateUserAge(address userAddress, uint256 userAge) returns(bool success)
func (_UserCrud *UserCrudTransactor) UpdateUserAge(opts *bind.TransactOpts, userAddress common.Address, userAge *big.Int) (*types.Transaction, error) {
	return _UserCrud.contract.Transact(opts, "updateUserAge", userAddress, userAge)
}

// UpdateUserAge is a paid mutator transaction binding the contract method 0x2e071db3.
//
// Solidity: function updateUserAge(address userAddress, uint256 userAge) returns(bool success)
func (_UserCrud *UserCrudSession) UpdateUserAge(userAddress common.Address, userAge *big.Int) (*types.Transaction, error) {
	return _UserCrud.Contract.UpdateUserAge(&_UserCrud.TransactOpts, userAddress, userAge)
}

// UpdateUserAge is a paid mutator transaction binding the contract method 0x2e071db3.
//
// Solidity: function updateUserAge(address userAddress, uint256 userAge) returns(bool success)
func (_UserCrud *UserCrudTransactorSession) UpdateUserAge(userAddress common.Address, userAge *big.Int) (*types.Transaction, error) {
	return _UserCrud.Contract.UpdateUserAge(&_UserCrud.TransactOpts, userAddress, userAge)
}

// UpdateUserEmail is a paid mutator transaction binding the contract method 0x64319ae6.
//
// Solidity: function updateUserEmail(address userAddress, string userEmail) returns(bool success)
func (_UserCrud *UserCrudTransactor) UpdateUserEmail(opts *bind.TransactOpts, userAddress common.Address, userEmail string) (*types.Transaction, error) {
	return _UserCrud.contract.Transact(opts, "updateUserEmail", userAddress, userEmail)
}

// UpdateUserEmail is a paid mutator transaction binding the contract method 0x64319ae6.
//
// Solidity: function updateUserEmail(address userAddress, string userEmail) returns(bool success)
func (_UserCrud *UserCrudSession) UpdateUserEmail(userAddress common.Address, userEmail string) (*types.Transaction, error) {
	return _UserCrud.Contract.UpdateUserEmail(&_UserCrud.TransactOpts, userAddress, userEmail)
}

// UpdateUserEmail is a paid mutator transaction binding the contract method 0x64319ae6.
//
// Solidity: function updateUserEmail(address userAddress, string userEmail) returns(bool success)
func (_UserCrud *UserCrudTransactorSession) UpdateUserEmail(userAddress common.Address, userEmail string) (*types.Transaction, error) {
	return _UserCrud.Contract.UpdateUserEmail(&_UserCrud.TransactOpts, userAddress, userEmail)
}
