// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package router_registry

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

// IRouterRegistryRouter is an auto generated low-level Go binding around an user-defined struct.
type IRouterRegistryRouter struct {
	Id       [32]byte
	Owner    common.Address
	Networks []*big.Int
	Endpoint string
}

// RouterRegistryMetaData contains all meta data concerning the RouterRegistry contract.
var RouterRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"RouterRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"RouterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"RouterUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROUTER_REGISTERER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROUTER_REMOVER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROUTER_UPDATE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRUSTED_FORWARDER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint24[]\",\"name\":\"networks\",\"type\":\"uint24[]\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"remove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"routerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"routers\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint24[]\",\"name\":\"networks\",\"type\":\"uint24[]\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"}],\"internalType\":\"structIRouterRegistry.Router\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"routersPaged\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint24[]\",\"name\":\"networks\",\"type\":\"uint24[]\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"}],\"internalType\":\"structIRouterRegistry.Router[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint24[]\",\"name\":\"networks\",\"type\":\"uint24[]\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RouterRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use RouterRegistryMetaData.ABI instead.
var RouterRegistryABI = RouterRegistryMetaData.ABI

// RouterRegistry is an auto generated Go binding around an Ethereum contract.
type RouterRegistry struct {
	RouterRegistryCaller     // Read-only binding to the contract
	RouterRegistryTransactor // Write-only binding to the contract
	RouterRegistryFilterer   // Log filterer for contract events
}

// RouterRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RouterRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RouterRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RouterRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RouterRegistrySession struct {
	Contract     *RouterRegistry   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RouterRegistryCallerSession struct {
	Contract *RouterRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// RouterRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RouterRegistryTransactorSession struct {
	Contract     *RouterRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// RouterRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RouterRegistryRaw struct {
	Contract *RouterRegistry // Generic contract binding to access the raw methods on
}

// RouterRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RouterRegistryCallerRaw struct {
	Contract *RouterRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RouterRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RouterRegistryTransactorRaw struct {
	Contract *RouterRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRouterRegistry creates a new instance of RouterRegistry, bound to a specific deployed contract.
func NewRouterRegistry(address common.Address, backend bind.ContractBackend) (*RouterRegistry, error) {
	contract, err := bindRouterRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RouterRegistry{RouterRegistryCaller: RouterRegistryCaller{contract: contract}, RouterRegistryTransactor: RouterRegistryTransactor{contract: contract}, RouterRegistryFilterer: RouterRegistryFilterer{contract: contract}}, nil
}

// NewRouterRegistryCaller creates a new read-only instance of RouterRegistry, bound to a specific deployed contract.
func NewRouterRegistryCaller(address common.Address, caller bind.ContractCaller) (*RouterRegistryCaller, error) {
	contract, err := bindRouterRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RouterRegistryCaller{contract: contract}, nil
}

// NewRouterRegistryTransactor creates a new write-only instance of RouterRegistry, bound to a specific deployed contract.
func NewRouterRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RouterRegistryTransactor, error) {
	contract, err := bindRouterRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RouterRegistryTransactor{contract: contract}, nil
}

// NewRouterRegistryFilterer creates a new log filterer instance of RouterRegistry, bound to a specific deployed contract.
func NewRouterRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RouterRegistryFilterer, error) {
	contract, err := bindRouterRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RouterRegistryFilterer{contract: contract}, nil
}

// bindRouterRegistry binds a generic wrapper to an already deployed contract.
func bindRouterRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RouterRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RouterRegistry *RouterRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RouterRegistry.Contract.RouterRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RouterRegistry *RouterRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RouterRegistry.Contract.RouterRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RouterRegistry *RouterRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RouterRegistry.Contract.RouterRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RouterRegistry *RouterRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RouterRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RouterRegistry *RouterRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RouterRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RouterRegistry *RouterRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RouterRegistry.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RouterRegistry *RouterRegistryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RouterRegistry.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RouterRegistry *RouterRegistrySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _RouterRegistry.Contract.DEFAULTADMINROLE(&_RouterRegistry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RouterRegistry *RouterRegistryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _RouterRegistry.Contract.DEFAULTADMINROLE(&_RouterRegistry.CallOpts)
}

// ROUTERREGISTERERROLE is a free data retrieval call binding the contract method 0x1916b5f7.
//
// Solidity: function ROUTER_REGISTERER_ROLE() view returns(bytes32)
func (_RouterRegistry *RouterRegistryCaller) ROUTERREGISTERERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RouterRegistry.contract.Call(opts, &out, "ROUTER_REGISTERER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ROUTERREGISTERERROLE is a free data retrieval call binding the contract method 0x1916b5f7.
//
// Solidity: function ROUTER_REGISTERER_ROLE() view returns(bytes32)
func (_RouterRegistry *RouterRegistrySession) ROUTERREGISTERERROLE() ([32]byte, error) {
	return _RouterRegistry.Contract.ROUTERREGISTERERROLE(&_RouterRegistry.CallOpts)
}

// ROUTERREGISTERERROLE is a free data retrieval call binding the contract method 0x1916b5f7.
//
// Solidity: function ROUTER_REGISTERER_ROLE() view returns(bytes32)
func (_RouterRegistry *RouterRegistryCallerSession) ROUTERREGISTERERROLE() ([32]byte, error) {
	return _RouterRegistry.Contract.ROUTERREGISTERERROLE(&_RouterRegistry.CallOpts)
}

// ROUTERREMOVERROLE is a free data retrieval call binding the contract method 0x7a97d02d.
//
// Solidity: function ROUTER_REMOVER_ROLE() view returns(bytes32)
func (_RouterRegistry *RouterRegistryCaller) ROUTERREMOVERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RouterRegistry.contract.Call(opts, &out, "ROUTER_REMOVER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ROUTERREMOVERROLE is a free data retrieval call binding the contract method 0x7a97d02d.
//
// Solidity: function ROUTER_REMOVER_ROLE() view returns(bytes32)
func (_RouterRegistry *RouterRegistrySession) ROUTERREMOVERROLE() ([32]byte, error) {
	return _RouterRegistry.Contract.ROUTERREMOVERROLE(&_RouterRegistry.CallOpts)
}

// ROUTERREMOVERROLE is a free data retrieval call binding the contract method 0x7a97d02d.
//
// Solidity: function ROUTER_REMOVER_ROLE() view returns(bytes32)
func (_RouterRegistry *RouterRegistryCallerSession) ROUTERREMOVERROLE() ([32]byte, error) {
	return _RouterRegistry.Contract.ROUTERREMOVERROLE(&_RouterRegistry.CallOpts)
}

// ROUTERUPDATEROLE is a free data retrieval call binding the contract method 0x5487d3ce.
//
// Solidity: function ROUTER_UPDATE_ROLE() view returns(bytes32)
func (_RouterRegistry *RouterRegistryCaller) ROUTERUPDATEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RouterRegistry.contract.Call(opts, &out, "ROUTER_UPDATE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ROUTERUPDATEROLE is a free data retrieval call binding the contract method 0x5487d3ce.
//
// Solidity: function ROUTER_UPDATE_ROLE() view returns(bytes32)
func (_RouterRegistry *RouterRegistrySession) ROUTERUPDATEROLE() ([32]byte, error) {
	return _RouterRegistry.Contract.ROUTERUPDATEROLE(&_RouterRegistry.CallOpts)
}

// ROUTERUPDATEROLE is a free data retrieval call binding the contract method 0x5487d3ce.
//
// Solidity: function ROUTER_UPDATE_ROLE() view returns(bytes32)
func (_RouterRegistry *RouterRegistryCallerSession) ROUTERUPDATEROLE() ([32]byte, error) {
	return _RouterRegistry.Contract.ROUTERUPDATEROLE(&_RouterRegistry.CallOpts)
}

// TRUSTEDFORWARDERROLE is a free data retrieval call binding the contract method 0x00cba943.
//
// Solidity: function TRUSTED_FORWARDER_ROLE() view returns(bytes32)
func (_RouterRegistry *RouterRegistryCaller) TRUSTEDFORWARDERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RouterRegistry.contract.Call(opts, &out, "TRUSTED_FORWARDER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TRUSTEDFORWARDERROLE is a free data retrieval call binding the contract method 0x00cba943.
//
// Solidity: function TRUSTED_FORWARDER_ROLE() view returns(bytes32)
func (_RouterRegistry *RouterRegistrySession) TRUSTEDFORWARDERROLE() ([32]byte, error) {
	return _RouterRegistry.Contract.TRUSTEDFORWARDERROLE(&_RouterRegistry.CallOpts)
}

// TRUSTEDFORWARDERROLE is a free data retrieval call binding the contract method 0x00cba943.
//
// Solidity: function TRUSTED_FORWARDER_ROLE() view returns(bytes32)
func (_RouterRegistry *RouterRegistryCallerSession) TRUSTEDFORWARDERROLE() ([32]byte, error) {
	return _RouterRegistry.Contract.TRUSTEDFORWARDERROLE(&_RouterRegistry.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RouterRegistry *RouterRegistryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _RouterRegistry.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RouterRegistry *RouterRegistrySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _RouterRegistry.Contract.GetRoleAdmin(&_RouterRegistry.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RouterRegistry *RouterRegistryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _RouterRegistry.Contract.GetRoleAdmin(&_RouterRegistry.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RouterRegistry *RouterRegistryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _RouterRegistry.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RouterRegistry *RouterRegistrySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _RouterRegistry.Contract.HasRole(&_RouterRegistry.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RouterRegistry *RouterRegistryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _RouterRegistry.Contract.HasRole(&_RouterRegistry.CallOpts, role, account)
}

// RouterCount is a free data retrieval call binding the contract method 0x8e67e049.
//
// Solidity: function routerCount() view returns(uint256)
func (_RouterRegistry *RouterRegistryCaller) RouterCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RouterRegistry.contract.Call(opts, &out, "routerCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RouterCount is a free data retrieval call binding the contract method 0x8e67e049.
//
// Solidity: function routerCount() view returns(uint256)
func (_RouterRegistry *RouterRegistrySession) RouterCount() (*big.Int, error) {
	return _RouterRegistry.Contract.RouterCount(&_RouterRegistry.CallOpts)
}

// RouterCount is a free data retrieval call binding the contract method 0x8e67e049.
//
// Solidity: function routerCount() view returns(uint256)
func (_RouterRegistry *RouterRegistryCallerSession) RouterCount() (*big.Int, error) {
	return _RouterRegistry.Contract.RouterCount(&_RouterRegistry.CallOpts)
}

// Routers is a free data retrieval call binding the contract method 0xaa1fce69.
//
// Solidity: function routers(bytes32 id) view returns((bytes32,address,uint24[],string))
func (_RouterRegistry *RouterRegistryCaller) Routers(opts *bind.CallOpts, id [32]byte) (IRouterRegistryRouter, error) {
	var out []interface{}
	err := _RouterRegistry.contract.Call(opts, &out, "routers", id)

	if err != nil {
		return *new(IRouterRegistryRouter), err
	}

	out0 := *abi.ConvertType(out[0], new(IRouterRegistryRouter)).(*IRouterRegistryRouter)

	return out0, err

}

// Routers is a free data retrieval call binding the contract method 0xaa1fce69.
//
// Solidity: function routers(bytes32 id) view returns((bytes32,address,uint24[],string))
func (_RouterRegistry *RouterRegistrySession) Routers(id [32]byte) (IRouterRegistryRouter, error) {
	return _RouterRegistry.Contract.Routers(&_RouterRegistry.CallOpts, id)
}

// Routers is a free data retrieval call binding the contract method 0xaa1fce69.
//
// Solidity: function routers(bytes32 id) view returns((bytes32,address,uint24[],string))
func (_RouterRegistry *RouterRegistryCallerSession) Routers(id [32]byte) (IRouterRegistryRouter, error) {
	return _RouterRegistry.Contract.Routers(&_RouterRegistry.CallOpts, id)
}

// RoutersPaged is a free data retrieval call binding the contract method 0x5c6201eb.
//
// Solidity: function routersPaged(uint256 start, uint256 end) view returns((bytes32,address,uint24[],string)[])
func (_RouterRegistry *RouterRegistryCaller) RoutersPaged(opts *bind.CallOpts, start *big.Int, end *big.Int) ([]IRouterRegistryRouter, error) {
	var out []interface{}
	err := _RouterRegistry.contract.Call(opts, &out, "routersPaged", start, end)

	if err != nil {
		return *new([]IRouterRegistryRouter), err
	}

	out0 := *abi.ConvertType(out[0], new([]IRouterRegistryRouter)).(*[]IRouterRegistryRouter)

	return out0, err

}

// RoutersPaged is a free data retrieval call binding the contract method 0x5c6201eb.
//
// Solidity: function routersPaged(uint256 start, uint256 end) view returns((bytes32,address,uint24[],string)[])
func (_RouterRegistry *RouterRegistrySession) RoutersPaged(start *big.Int, end *big.Int) ([]IRouterRegistryRouter, error) {
	return _RouterRegistry.Contract.RoutersPaged(&_RouterRegistry.CallOpts, start, end)
}

// RoutersPaged is a free data retrieval call binding the contract method 0x5c6201eb.
//
// Solidity: function routersPaged(uint256 start, uint256 end) view returns((bytes32,address,uint24[],string)[])
func (_RouterRegistry *RouterRegistryCallerSession) RoutersPaged(start *big.Int, end *big.Int) ([]IRouterRegistryRouter, error) {
	return _RouterRegistry.Contract.RoutersPaged(&_RouterRegistry.CallOpts, start, end)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RouterRegistry *RouterRegistryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _RouterRegistry.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RouterRegistry *RouterRegistrySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RouterRegistry.Contract.SupportsInterface(&_RouterRegistry.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RouterRegistry *RouterRegistryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RouterRegistry.Contract.SupportsInterface(&_RouterRegistry.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RouterRegistry *RouterRegistryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RouterRegistry.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RouterRegistry *RouterRegistrySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RouterRegistry.Contract.GrantRole(&_RouterRegistry.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RouterRegistry *RouterRegistryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RouterRegistry.Contract.GrantRole(&_RouterRegistry.TransactOpts, role, account)
}

// Register is a paid mutator transaction binding the contract method 0x5f46e4b3.
//
// Solidity: function register(bytes32 id, address owner, uint24[] networks, string endpoint) returns()
func (_RouterRegistry *RouterRegistryTransactor) Register(opts *bind.TransactOpts, id [32]byte, owner common.Address, networks []*big.Int, endpoint string) (*types.Transaction, error) {
	return _RouterRegistry.contract.Transact(opts, "register", id, owner, networks, endpoint)
}

// Register is a paid mutator transaction binding the contract method 0x5f46e4b3.
//
// Solidity: function register(bytes32 id, address owner, uint24[] networks, string endpoint) returns()
func (_RouterRegistry *RouterRegistrySession) Register(id [32]byte, owner common.Address, networks []*big.Int, endpoint string) (*types.Transaction, error) {
	return _RouterRegistry.Contract.Register(&_RouterRegistry.TransactOpts, id, owner, networks, endpoint)
}

// Register is a paid mutator transaction binding the contract method 0x5f46e4b3.
//
// Solidity: function register(bytes32 id, address owner, uint24[] networks, string endpoint) returns()
func (_RouterRegistry *RouterRegistryTransactorSession) Register(id [32]byte, owner common.Address, networks []*big.Int, endpoint string) (*types.Transaction, error) {
	return _RouterRegistry.Contract.Register(&_RouterRegistry.TransactOpts, id, owner, networks, endpoint)
}

// Remove is a paid mutator transaction binding the contract method 0x2874528e.
//
// Solidity: function remove(bytes32 id, address owner) returns()
func (_RouterRegistry *RouterRegistryTransactor) Remove(opts *bind.TransactOpts, id [32]byte, owner common.Address) (*types.Transaction, error) {
	return _RouterRegistry.contract.Transact(opts, "remove", id, owner)
}

// Remove is a paid mutator transaction binding the contract method 0x2874528e.
//
// Solidity: function remove(bytes32 id, address owner) returns()
func (_RouterRegistry *RouterRegistrySession) Remove(id [32]byte, owner common.Address) (*types.Transaction, error) {
	return _RouterRegistry.Contract.Remove(&_RouterRegistry.TransactOpts, id, owner)
}

// Remove is a paid mutator transaction binding the contract method 0x2874528e.
//
// Solidity: function remove(bytes32 id, address owner) returns()
func (_RouterRegistry *RouterRegistryTransactorSession) Remove(id [32]byte, owner common.Address) (*types.Transaction, error) {
	return _RouterRegistry.Contract.Remove(&_RouterRegistry.TransactOpts, id, owner)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RouterRegistry *RouterRegistryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RouterRegistry.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RouterRegistry *RouterRegistrySession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RouterRegistry.Contract.RenounceRole(&_RouterRegistry.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RouterRegistry *RouterRegistryTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RouterRegistry.Contract.RenounceRole(&_RouterRegistry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RouterRegistry *RouterRegistryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RouterRegistry.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RouterRegistry *RouterRegistrySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RouterRegistry.Contract.RevokeRole(&_RouterRegistry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RouterRegistry *RouterRegistryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RouterRegistry.Contract.RevokeRole(&_RouterRegistry.TransactOpts, role, account)
}

// Update is a paid mutator transaction binding the contract method 0x5bb5b1e6.
//
// Solidity: function update(bytes32 id, address owner, uint24[] networks, string endpoint) returns()
func (_RouterRegistry *RouterRegistryTransactor) Update(opts *bind.TransactOpts, id [32]byte, owner common.Address, networks []*big.Int, endpoint string) (*types.Transaction, error) {
	return _RouterRegistry.contract.Transact(opts, "update", id, owner, networks, endpoint)
}

// Update is a paid mutator transaction binding the contract method 0x5bb5b1e6.
//
// Solidity: function update(bytes32 id, address owner, uint24[] networks, string endpoint) returns()
func (_RouterRegistry *RouterRegistrySession) Update(id [32]byte, owner common.Address, networks []*big.Int, endpoint string) (*types.Transaction, error) {
	return _RouterRegistry.Contract.Update(&_RouterRegistry.TransactOpts, id, owner, networks, endpoint)
}

// Update is a paid mutator transaction binding the contract method 0x5bb5b1e6.
//
// Solidity: function update(bytes32 id, address owner, uint24[] networks, string endpoint) returns()
func (_RouterRegistry *RouterRegistryTransactorSession) Update(id [32]byte, owner common.Address, networks []*big.Int, endpoint string) (*types.Transaction, error) {
	return _RouterRegistry.Contract.Update(&_RouterRegistry.TransactOpts, id, owner, networks, endpoint)
}

// RouterRegistryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the RouterRegistry contract.
type RouterRegistryRoleAdminChangedIterator struct {
	Event *RouterRegistryRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *RouterRegistryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterRegistryRoleAdminChanged)
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
		it.Event = new(RouterRegistryRoleAdminChanged)
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
func (it *RouterRegistryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterRegistryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterRegistryRoleAdminChanged represents a RoleAdminChanged event raised by the RouterRegistry contract.
type RouterRegistryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_RouterRegistry *RouterRegistryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*RouterRegistryRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _RouterRegistry.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &RouterRegistryRoleAdminChangedIterator{contract: _RouterRegistry.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_RouterRegistry *RouterRegistryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *RouterRegistryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _RouterRegistry.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterRegistryRoleAdminChanged)
				if err := _RouterRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_RouterRegistry *RouterRegistryFilterer) ParseRoleAdminChanged(log types.Log) (*RouterRegistryRoleAdminChanged, error) {
	event := new(RouterRegistryRoleAdminChanged)
	if err := _RouterRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterRegistryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the RouterRegistry contract.
type RouterRegistryRoleGrantedIterator struct {
	Event *RouterRegistryRoleGranted // Event containing the contract specifics and raw log

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
func (it *RouterRegistryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterRegistryRoleGranted)
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
		it.Event = new(RouterRegistryRoleGranted)
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
func (it *RouterRegistryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterRegistryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterRegistryRoleGranted represents a RoleGranted event raised by the RouterRegistry contract.
type RouterRegistryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_RouterRegistry *RouterRegistryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RouterRegistryRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RouterRegistry.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RouterRegistryRoleGrantedIterator{contract: _RouterRegistry.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_RouterRegistry *RouterRegistryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *RouterRegistryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RouterRegistry.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterRegistryRoleGranted)
				if err := _RouterRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_RouterRegistry *RouterRegistryFilterer) ParseRoleGranted(log types.Log) (*RouterRegistryRoleGranted, error) {
	event := new(RouterRegistryRoleGranted)
	if err := _RouterRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterRegistryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the RouterRegistry contract.
type RouterRegistryRoleRevokedIterator struct {
	Event *RouterRegistryRoleRevoked // Event containing the contract specifics and raw log

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
func (it *RouterRegistryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterRegistryRoleRevoked)
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
		it.Event = new(RouterRegistryRoleRevoked)
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
func (it *RouterRegistryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterRegistryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterRegistryRoleRevoked represents a RoleRevoked event raised by the RouterRegistry contract.
type RouterRegistryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_RouterRegistry *RouterRegistryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RouterRegistryRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RouterRegistry.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RouterRegistryRoleRevokedIterator{contract: _RouterRegistry.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_RouterRegistry *RouterRegistryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *RouterRegistryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RouterRegistry.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterRegistryRoleRevoked)
				if err := _RouterRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_RouterRegistry *RouterRegistryFilterer) ParseRoleRevoked(log types.Log) (*RouterRegistryRoleRevoked, error) {
	event := new(RouterRegistryRoleRevoked)
	if err := _RouterRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterRegistryRouterRegisteredIterator is returned from FilterRouterRegistered and is used to iterate over the raw logs and unpacked data for RouterRegistered events raised by the RouterRegistry contract.
type RouterRegistryRouterRegisteredIterator struct {
	Event *RouterRegistryRouterRegistered // Event containing the contract specifics and raw log

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
func (it *RouterRegistryRouterRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterRegistryRouterRegistered)
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
		it.Event = new(RouterRegistryRouterRegistered)
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
func (it *RouterRegistryRouterRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterRegistryRouterRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterRegistryRouterRegistered represents a RouterRegistered event raised by the RouterRegistry contract.
type RouterRegistryRouterRegistered struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRouterRegistered is a free log retrieval operation binding the contract event 0xc8c04b68c2c10643ec3b3d87b29ee62eda2d6e5c929cd2aba791fe90aecc8c1b.
//
// Solidity: event RouterRegistered(bytes32 indexed id)
func (_RouterRegistry *RouterRegistryFilterer) FilterRouterRegistered(opts *bind.FilterOpts, id [][32]byte) (*RouterRegistryRouterRegisteredIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _RouterRegistry.contract.FilterLogs(opts, "RouterRegistered", idRule)
	if err != nil {
		return nil, err
	}
	return &RouterRegistryRouterRegisteredIterator{contract: _RouterRegistry.contract, event: "RouterRegistered", logs: logs, sub: sub}, nil
}

// WatchRouterRegistered is a free log subscription operation binding the contract event 0xc8c04b68c2c10643ec3b3d87b29ee62eda2d6e5c929cd2aba791fe90aecc8c1b.
//
// Solidity: event RouterRegistered(bytes32 indexed id)
func (_RouterRegistry *RouterRegistryFilterer) WatchRouterRegistered(opts *bind.WatchOpts, sink chan<- *RouterRegistryRouterRegistered, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _RouterRegistry.contract.WatchLogs(opts, "RouterRegistered", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterRegistryRouterRegistered)
				if err := _RouterRegistry.contract.UnpackLog(event, "RouterRegistered", log); err != nil {
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

// ParseRouterRegistered is a log parse operation binding the contract event 0xc8c04b68c2c10643ec3b3d87b29ee62eda2d6e5c929cd2aba791fe90aecc8c1b.
//
// Solidity: event RouterRegistered(bytes32 indexed id)
func (_RouterRegistry *RouterRegistryFilterer) ParseRouterRegistered(log types.Log) (*RouterRegistryRouterRegistered, error) {
	event := new(RouterRegistryRouterRegistered)
	if err := _RouterRegistry.contract.UnpackLog(event, "RouterRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterRegistryRouterRemovedIterator is returned from FilterRouterRemoved and is used to iterate over the raw logs and unpacked data for RouterRemoved events raised by the RouterRegistry contract.
type RouterRegistryRouterRemovedIterator struct {
	Event *RouterRegistryRouterRemoved // Event containing the contract specifics and raw log

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
func (it *RouterRegistryRouterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterRegistryRouterRemoved)
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
		it.Event = new(RouterRegistryRouterRemoved)
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
func (it *RouterRegistryRouterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterRegistryRouterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterRegistryRouterRemoved represents a RouterRemoved event raised by the RouterRegistry contract.
type RouterRegistryRouterRemoved struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRouterRemoved is a free log retrieval operation binding the contract event 0x813b69b7582e7b6cbe5fbfb9fbaceeef78bb34c4297b68f138b189e96eebeead.
//
// Solidity: event RouterRemoved(bytes32 indexed id)
func (_RouterRegistry *RouterRegistryFilterer) FilterRouterRemoved(opts *bind.FilterOpts, id [][32]byte) (*RouterRegistryRouterRemovedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _RouterRegistry.contract.FilterLogs(opts, "RouterRemoved", idRule)
	if err != nil {
		return nil, err
	}
	return &RouterRegistryRouterRemovedIterator{contract: _RouterRegistry.contract, event: "RouterRemoved", logs: logs, sub: sub}, nil
}

// WatchRouterRemoved is a free log subscription operation binding the contract event 0x813b69b7582e7b6cbe5fbfb9fbaceeef78bb34c4297b68f138b189e96eebeead.
//
// Solidity: event RouterRemoved(bytes32 indexed id)
func (_RouterRegistry *RouterRegistryFilterer) WatchRouterRemoved(opts *bind.WatchOpts, sink chan<- *RouterRegistryRouterRemoved, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _RouterRegistry.contract.WatchLogs(opts, "RouterRemoved", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterRegistryRouterRemoved)
				if err := _RouterRegistry.contract.UnpackLog(event, "RouterRemoved", log); err != nil {
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

// ParseRouterRemoved is a log parse operation binding the contract event 0x813b69b7582e7b6cbe5fbfb9fbaceeef78bb34c4297b68f138b189e96eebeead.
//
// Solidity: event RouterRemoved(bytes32 indexed id)
func (_RouterRegistry *RouterRegistryFilterer) ParseRouterRemoved(log types.Log) (*RouterRegistryRouterRemoved, error) {
	event := new(RouterRegistryRouterRemoved)
	if err := _RouterRegistry.contract.UnpackLog(event, "RouterRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterRegistryRouterUpdatedIterator is returned from FilterRouterUpdated and is used to iterate over the raw logs and unpacked data for RouterUpdated events raised by the RouterRegistry contract.
type RouterRegistryRouterUpdatedIterator struct {
	Event *RouterRegistryRouterUpdated // Event containing the contract specifics and raw log

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
func (it *RouterRegistryRouterUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterRegistryRouterUpdated)
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
		it.Event = new(RouterRegistryRouterUpdated)
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
func (it *RouterRegistryRouterUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterRegistryRouterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterRegistryRouterUpdated represents a RouterUpdated event raised by the RouterRegistry contract.
type RouterRegistryRouterUpdated struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRouterUpdated is a free log retrieval operation binding the contract event 0x8af614f5e1bf8bc42df40f6ef41468c010e2f999f579ba0ec3658405ced7451c.
//
// Solidity: event RouterUpdated(bytes32 indexed id)
func (_RouterRegistry *RouterRegistryFilterer) FilterRouterUpdated(opts *bind.FilterOpts, id [][32]byte) (*RouterRegistryRouterUpdatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _RouterRegistry.contract.FilterLogs(opts, "RouterUpdated", idRule)
	if err != nil {
		return nil, err
	}
	return &RouterRegistryRouterUpdatedIterator{contract: _RouterRegistry.contract, event: "RouterUpdated", logs: logs, sub: sub}, nil
}

// WatchRouterUpdated is a free log subscription operation binding the contract event 0x8af614f5e1bf8bc42df40f6ef41468c010e2f999f579ba0ec3658405ced7451c.
//
// Solidity: event RouterUpdated(bytes32 indexed id)
func (_RouterRegistry *RouterRegistryFilterer) WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *RouterRegistryRouterUpdated, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _RouterRegistry.contract.WatchLogs(opts, "RouterUpdated", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterRegistryRouterUpdated)
				if err := _RouterRegistry.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
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

// ParseRouterUpdated is a log parse operation binding the contract event 0x8af614f5e1bf8bc42df40f6ef41468c010e2f999f579ba0ec3658405ced7451c.
//
// Solidity: event RouterUpdated(bytes32 indexed id)
func (_RouterRegistry *RouterRegistryFilterer) ParseRouterUpdated(log types.Log) (*RouterRegistryRouterUpdated, error) {
	event := new(RouterRegistryRouterUpdated)
	if err := _RouterRegistry.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
