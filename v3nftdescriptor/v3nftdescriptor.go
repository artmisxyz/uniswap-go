// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package v3nftdescriptor

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

// V3nftdescriptorMetaData contains all meta data concerning the V3nftdescriptor contract.
var V3nftdescriptorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_WETH9\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"priority\",\"type\":\"int256\"}],\"name\":\"UpdateTokenRatioPriority\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"WETH9\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"flipRatio\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"tokenRatioPriority\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractINonfungiblePositionManager\",\"name\":\"positionManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405161134638038061134683398101604081905261002f91610044565b60601b6001600160601b031916608052610072565b600060208284031215610055578081fd5b81516001600160a01b038116811461006b578182fd5b9392505050565b60805160601c6112b36100936000398060d1528061011652506112b36000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80634aa4a4fc146100515780637e5af7711461006f5780639d7b0ea81461008f578063e9dc6375146100af575b600080fd5b6100596100cf565b6040516100669190611105565b60405180910390f35b61008261007d366004610e0b565b6100f3565b6040516100669190611119565b6100a261009d366004610e4b565b610112565b6040516100669190611124565b6100c26100bd366004610e4b565b610257565b604051610066919061112d565b7f000000000000000000000000000000000000000000000000000000000000000081565b60006100ff8383610112565b6101098584610112565b13949350505050565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316836001600160a01b031614156101575750606319610251565b816001141561024d576001600160a01b03831673a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48141561018e575061012c610251565b6001600160a01b03831673dac17f958d2ee523a2206206994597c13d831ec714156101bb575060c8610251565b6001600160a01b038316736b175474e89094c44da98b954eedeac495271d0f14156101e857506064610251565b6001600160a01b038316738daebade922df735c38c80c7ebd708af50815faa1415610216575060c719610251565b6001600160a01b038316732260fac5e5542a773aa44fbcfedf7c193bc2c5991415610245575061012b19610251565b506000610251565b5060005b92915050565b60606000806000806000876001600160a01b03166399fbab88886040518263ffffffff1660e01b815260040161028d9190611124565b6101806040518083038186803b1580156102a657600080fd5b505afa1580156102ba573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102de9190610fd0565b5050505050965096509650965096505050600061039c896001600160a01b031663c45a01556040518163ffffffff1660e01b815260040160206040518083038186803b15801561032d57600080fd5b505afa158015610341573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103659190610de8565b6040518060600160405280896001600160a01b03168152602001886001600160a01b031681526020018762ffffff168152506106db565b905060006103ad878761007d6107d7565b9050600081156103bd57876103bf565b865b9050600082156103cf57876103d1565b885b90506000846001600160a01b0316633850c7bd6040518163ffffffff1660e01b815260040160e06040518083038186803b15801561040e57600080fd5b505afa158015610422573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104469190610f27565b50505050509150507342b24a95702b9986e82d421cc3568932790a48ec63c49917d7604051806101c001604052808f8152602001866001600160a01b03168152602001856001600160a01b031681526020016104a1876107db565b81526020016104af866107db565b8152602001866001600160a01b031663313ce5676040518163ffffffff1660e01b815260040160206040518083038186803b1580156104ed57600080fd5b505afa158015610501573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105259190610fb6565b60ff168152602001856001600160a01b031663313ce5676040518163ffffffff1660e01b815260040160206040518083038186803b15801561056657600080fd5b505afa15801561057a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061059e9190610fb6565b60ff16815260200187151581526020018a60020b81526020018960020b81526020018460020b8152602001886001600160a01b031663d0c93a7c6040518163ffffffff1660e01b815260040160206040518083038186803b15801561060257600080fd5b505afa158015610616573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061063a9190610e76565b60020b81526020018b62ffffff168152602001886001600160a01b03168152506040518263ffffffff1660e01b81526004016106769190611140565b60006040518083038186803b15801561068e57600080fd5b505af41580156106a2573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526106ca9190810190610e90565b9d9c50505050505050505050505050565b600081602001516001600160a01b031682600001516001600160a01b03161061070357600080fd5b50805160208083015160409384015184516001600160a01b0394851681850152939091168385015262ffffff166060808401919091528351808403820181526080840185528051908301207fff0000000000000000000000000000000000000000000000000000000000000060a085015294901b6bffffffffffffffffffffffff191660a183015260b58201939093527fe34f199b19b2b4f47f68442619d555527d244f78a3297ea89325f843f87b8b5460d5808301919091528251808303909101815260f5909101909152805191012090565b4690565b60606000610809837f95d89b410000000000000000000000000000000000000000000000000000000061082e565b90508051600014156108265761081e83610a83565b915050610829565b90505b919050565b60408051600481526024810182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000085161781529151815160609360009384936001600160a01b03891693919290918291908083835b602083106108c75780518252601f1990920191602091820191016108a8565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855afa9150503d8060008114610927576040519150601f19603f3d011682016040523d82523d6000602084013e61092c565b606091505b509150915081158061093d57508051155b1561095b576040518060200160405280600081525092505050610251565b80516020141561099357600081806020019051602081101561097c57600080fd5b5051905061098981610a90565b9350505050610251565b604081511115610a6b578080602001905160208110156109b257600080fd5b81019080805160405193929190846401000000008211156109d257600080fd5b9083019060208201858111156109e757600080fd5b8251640100000000811182820188101715610a0157600080fd5b82525081516020918201929091019080838360005b83811015610a2e578181015183820152602001610a16565b50505050905090810190601f168015610a5b5780820380516001836020036101000a031916815260200191505b5060405250505092505050610251565b50506040805160208101909152600081529392505050565b6060610826826006610bd0565b604080516020808252818301909252606091600091906020820181803683370190505090506000805b6020811015610b32576000858260208110610ad057fe5b1a60f81b90507fff00000000000000000000000000000000000000000000000000000000000000811615610b295780848481518110610b0b57fe5b60200101906001600160f81b031916908160001a9053506001909201915b50600101610ab9565b5060008167ffffffffffffffff81118015610b4c57600080fd5b506040519080825280601f01601f191660200182016040528015610b77576020820181803683370190505b50905060005b82811015610bc757838181518110610b9157fe5b602001015160f81c60f81b828281518110610ba857fe5b60200101906001600160f81b031916908160001a905350600101610b7d565b50949350505050565b606060028206158015610be35750600082115b8015610bf0575060288211155b610c5b57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f41646472657373537472696e675574696c3a20494e56414c49445f4c454e0000604482015290519081900360640190fd5b60008267ffffffffffffffff81118015610c7457600080fd5b506040519080825280601f01601f191660200182016040528015610c9f576020820181803683370190505b5090506001600160a01b03841660005b60028504811015610d4357600860138290030282901c600f600482901c1660f082168203610cdc82610d4d565b868560020281518110610ceb57fe5b60200101906001600160f81b031916908160001a905350610d0b81610d4d565b868560020260010181518110610d1d57fe5b60200101906001600160f81b031916908160001a9053505060019092019150610caf9050565b5090949350505050565b6000600a8260ff161015610d6857506030810160f81b610829565b506037810160f81b610829565b80516108298161128e565b8051600281900b811461082957600080fd5b80516fffffffffffffffffffffffffffffffff8116811461082957600080fd5b805161ffff8116811461082957600080fd5b805162ffffff8116811461082957600080fd5b805160ff8116811461082957600080fd5b600060208284031215610df9578081fd5b8151610e048161128e565b9392505050565b600080600060608486031215610e1f578182fd5b8335610e2a8161128e565b92506020840135610e3a8161128e565b929592945050506040919091013590565b60008060408385031215610e5d578182fd5b8235610e688161128e565b946020939093013593505050565b600060208284031215610e87578081fd5b610e0482610d80565b600060208284031215610ea1578081fd5b815167ffffffffffffffff80821115610eb8578283fd5b818401915084601f830112610ecb578283fd5b815181811115610ed757fe5b604051601f8201601f191681016020018381118282101715610ef557fe5b604052818152838201602001871015610f0c578485fd5b610f1d82602083016020870161125e565b9695505050505050565b600080600080600080600060e0888a031215610f41578283fd5b8751610f4c8161128e565b9650610f5a60208901610d80565b9550610f6860408901610db2565b9450610f7660608901610db2565b9350610f8460808901610db2565b9250610f9260a08901610dd7565b915060c08801518015158114610fa6578182fd5b8091505092959891949750929550565b600060208284031215610fc7578081fd5b610e0482610dd7565b6000806000806000806000806000806000806101808d8f031215610ff2578485fd5b8c516bffffffffffffffffffffffff8116811461100d578586fd5b9b5061101b60208e01610d75565b9a5061102960408e01610d75565b995061103760608e01610d75565b985061104560808e01610dc4565b975061105360a08e01610d80565b965061106160c08e01610d80565b955061106f60e08e01610d92565b94506101008d015193506101208d0151925061108e6101408e01610d92565b915061109d6101608e01610d92565b90509295989b509295989b509295989b565b6001600160a01b03169052565b15159052565b60020b9052565b600081518084526110e181602086016020860161125e565b601f01601f19169290920160200192915050565b62ffffff169052565b60ff169052565b6001600160a01b0391909116815260200190565b901515815260200190565b90815260200190565b600060208252610e0460208301846110c9565b60006020825282516020830152602083015161115f60408401826110af565b50604083015161117260608401826110af565b5060608301516101c080608085015261118f6101e08501836110c9565b91506080850151601f198584030160a08601526111ac83826110c9565b92505060a08501516111c160c08601826110fe565b5060c08501516111d460e08601826110fe565b5060e08501516101006111e9818701836110bc565b86015190506101206111fd868201836110c2565b8601519050610140611211868201836110c2565b8601519050610160611225868201836110c2565b8601519050610180611239868201836110c2565b86015190506101a061124d868201836110f5565b8601519050610d43858301826110af565b60005b83811015611279578181015183820152602001611261565b83811115611288576000848401525b50505050565b6001600160a01b03811681146112a357600080fd5b5056fea164736f6c6343000706000a000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
}

// V3nftdescriptorABI is the input ABI used to generate the binding from.
// Deprecated: Use V3nftdescriptorMetaData.ABI instead.
var V3nftdescriptorABI = V3nftdescriptorMetaData.ABI

// V3nftdescriptorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use V3nftdescriptorMetaData.Bin instead.
var V3nftdescriptorBin = V3nftdescriptorMetaData.Bin

// DeployV3nftdescriptor deploys a new Ethereum contract, binding an instance of V3nftdescriptor to it.
func DeployV3nftdescriptor(auth *bind.TransactOpts, backend bind.ContractBackend, _WETH9 common.Address) (common.Address, *types.Transaction, *V3nftdescriptor, error) {
	parsed, err := V3nftdescriptorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(V3nftdescriptorBin), backend, _WETH9)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &V3nftdescriptor{V3nftdescriptorCaller: V3nftdescriptorCaller{contract: contract}, V3nftdescriptorTransactor: V3nftdescriptorTransactor{contract: contract}, V3nftdescriptorFilterer: V3nftdescriptorFilterer{contract: contract}}, nil
}

// V3nftdescriptor is an auto generated Go binding around an Ethereum contract.
type V3nftdescriptor struct {
	V3nftdescriptorCaller     // Read-only binding to the contract
	V3nftdescriptorTransactor // Write-only binding to the contract
	V3nftdescriptorFilterer   // Log filterer for contract events
}

// V3nftdescriptorCaller is an auto generated read-only Go binding around an Ethereum contract.
type V3nftdescriptorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V3nftdescriptorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type V3nftdescriptorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V3nftdescriptorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type V3nftdescriptorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V3nftdescriptorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type V3nftdescriptorSession struct {
	Contract     *V3nftdescriptor  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// V3nftdescriptorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type V3nftdescriptorCallerSession struct {
	Contract *V3nftdescriptorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// V3nftdescriptorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type V3nftdescriptorTransactorSession struct {
	Contract     *V3nftdescriptorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// V3nftdescriptorRaw is an auto generated low-level Go binding around an Ethereum contract.
type V3nftdescriptorRaw struct {
	Contract *V3nftdescriptor // Generic contract binding to access the raw methods on
}

// V3nftdescriptorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type V3nftdescriptorCallerRaw struct {
	Contract *V3nftdescriptorCaller // Generic read-only contract binding to access the raw methods on
}

// V3nftdescriptorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type V3nftdescriptorTransactorRaw struct {
	Contract *V3nftdescriptorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewV3nftdescriptor creates a new instance of V3nftdescriptor, bound to a specific deployed contract.
func NewV3nftdescriptor(address common.Address, backend bind.ContractBackend) (*V3nftdescriptor, error) {
	contract, err := bindV3nftdescriptor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &V3nftdescriptor{V3nftdescriptorCaller: V3nftdescriptorCaller{contract: contract}, V3nftdescriptorTransactor: V3nftdescriptorTransactor{contract: contract}, V3nftdescriptorFilterer: V3nftdescriptorFilterer{contract: contract}}, nil
}

// NewV3nftdescriptorCaller creates a new read-only instance of V3nftdescriptor, bound to a specific deployed contract.
func NewV3nftdescriptorCaller(address common.Address, caller bind.ContractCaller) (*V3nftdescriptorCaller, error) {
	contract, err := bindV3nftdescriptor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &V3nftdescriptorCaller{contract: contract}, nil
}

// NewV3nftdescriptorTransactor creates a new write-only instance of V3nftdescriptor, bound to a specific deployed contract.
func NewV3nftdescriptorTransactor(address common.Address, transactor bind.ContractTransactor) (*V3nftdescriptorTransactor, error) {
	contract, err := bindV3nftdescriptor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &V3nftdescriptorTransactor{contract: contract}, nil
}

// NewV3nftdescriptorFilterer creates a new log filterer instance of V3nftdescriptor, bound to a specific deployed contract.
func NewV3nftdescriptorFilterer(address common.Address, filterer bind.ContractFilterer) (*V3nftdescriptorFilterer, error) {
	contract, err := bindV3nftdescriptor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &V3nftdescriptorFilterer{contract: contract}, nil
}

// bindV3nftdescriptor binds a generic wrapper to an already deployed contract.
func bindV3nftdescriptor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(V3nftdescriptorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V3nftdescriptor *V3nftdescriptorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V3nftdescriptor.Contract.V3nftdescriptorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V3nftdescriptor *V3nftdescriptorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V3nftdescriptor.Contract.V3nftdescriptorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V3nftdescriptor *V3nftdescriptorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V3nftdescriptor.Contract.V3nftdescriptorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V3nftdescriptor *V3nftdescriptorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V3nftdescriptor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V3nftdescriptor *V3nftdescriptorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V3nftdescriptor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V3nftdescriptor *V3nftdescriptorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V3nftdescriptor.Contract.contract.Transact(opts, method, params...)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_V3nftdescriptor *V3nftdescriptorCaller) WETH9(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V3nftdescriptor.contract.Call(opts, &out, "WETH9")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_V3nftdescriptor *V3nftdescriptorSession) WETH9() (common.Address, error) {
	return _V3nftdescriptor.Contract.WETH9(&_V3nftdescriptor.CallOpts)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_V3nftdescriptor *V3nftdescriptorCallerSession) WETH9() (common.Address, error) {
	return _V3nftdescriptor.Contract.WETH9(&_V3nftdescriptor.CallOpts)
}

// FlipRatio is a free data retrieval call binding the contract method 0x7e5af771.
//
// Solidity: function flipRatio(address token0, address token1, uint256 chainId) view returns(bool)
func (_V3nftdescriptor *V3nftdescriptorCaller) FlipRatio(opts *bind.CallOpts, token0 common.Address, token1 common.Address, chainId *big.Int) (bool, error) {
	var out []interface{}
	err := _V3nftdescriptor.contract.Call(opts, &out, "flipRatio", token0, token1, chainId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FlipRatio is a free data retrieval call binding the contract method 0x7e5af771.
//
// Solidity: function flipRatio(address token0, address token1, uint256 chainId) view returns(bool)
func (_V3nftdescriptor *V3nftdescriptorSession) FlipRatio(token0 common.Address, token1 common.Address, chainId *big.Int) (bool, error) {
	return _V3nftdescriptor.Contract.FlipRatio(&_V3nftdescriptor.CallOpts, token0, token1, chainId)
}

// FlipRatio is a free data retrieval call binding the contract method 0x7e5af771.
//
// Solidity: function flipRatio(address token0, address token1, uint256 chainId) view returns(bool)
func (_V3nftdescriptor *V3nftdescriptorCallerSession) FlipRatio(token0 common.Address, token1 common.Address, chainId *big.Int) (bool, error) {
	return _V3nftdescriptor.Contract.FlipRatio(&_V3nftdescriptor.CallOpts, token0, token1, chainId)
}

// TokenRatioPriority is a free data retrieval call binding the contract method 0x9d7b0ea8.
//
// Solidity: function tokenRatioPriority(address token, uint256 chainId) view returns(int256)
func (_V3nftdescriptor *V3nftdescriptorCaller) TokenRatioPriority(opts *bind.CallOpts, token common.Address, chainId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _V3nftdescriptor.contract.Call(opts, &out, "tokenRatioPriority", token, chainId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenRatioPriority is a free data retrieval call binding the contract method 0x9d7b0ea8.
//
// Solidity: function tokenRatioPriority(address token, uint256 chainId) view returns(int256)
func (_V3nftdescriptor *V3nftdescriptorSession) TokenRatioPriority(token common.Address, chainId *big.Int) (*big.Int, error) {
	return _V3nftdescriptor.Contract.TokenRatioPriority(&_V3nftdescriptor.CallOpts, token, chainId)
}

// TokenRatioPriority is a free data retrieval call binding the contract method 0x9d7b0ea8.
//
// Solidity: function tokenRatioPriority(address token, uint256 chainId) view returns(int256)
func (_V3nftdescriptor *V3nftdescriptorCallerSession) TokenRatioPriority(token common.Address, chainId *big.Int) (*big.Int, error) {
	return _V3nftdescriptor.Contract.TokenRatioPriority(&_V3nftdescriptor.CallOpts, token, chainId)
}

// TokenURI is a free data retrieval call binding the contract method 0xe9dc6375.
//
// Solidity: function tokenURI(address positionManager, uint256 tokenId) view returns(string)
func (_V3nftdescriptor *V3nftdescriptorCaller) TokenURI(opts *bind.CallOpts, positionManager common.Address, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _V3nftdescriptor.contract.Call(opts, &out, "tokenURI", positionManager, tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xe9dc6375.
//
// Solidity: function tokenURI(address positionManager, uint256 tokenId) view returns(string)
func (_V3nftdescriptor *V3nftdescriptorSession) TokenURI(positionManager common.Address, tokenId *big.Int) (string, error) {
	return _V3nftdescriptor.Contract.TokenURI(&_V3nftdescriptor.CallOpts, positionManager, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xe9dc6375.
//
// Solidity: function tokenURI(address positionManager, uint256 tokenId) view returns(string)
func (_V3nftdescriptor *V3nftdescriptorCallerSession) TokenURI(positionManager common.Address, tokenId *big.Int) (string, error) {
	return _V3nftdescriptor.Contract.TokenURI(&_V3nftdescriptor.CallOpts, positionManager, tokenId)
}

// V3nftdescriptorUpdateTokenRatioPriorityIterator is returned from FilterUpdateTokenRatioPriority and is used to iterate over the raw logs and unpacked data for UpdateTokenRatioPriority events raised by the V3nftdescriptor contract.
type V3nftdescriptorUpdateTokenRatioPriorityIterator struct {
	Event *V3nftdescriptorUpdateTokenRatioPriority // Event containing the contract specifics and raw log

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
func (it *V3nftdescriptorUpdateTokenRatioPriorityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V3nftdescriptorUpdateTokenRatioPriority)
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
		it.Event = new(V3nftdescriptorUpdateTokenRatioPriority)
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
func (it *V3nftdescriptorUpdateTokenRatioPriorityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V3nftdescriptorUpdateTokenRatioPriorityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V3nftdescriptorUpdateTokenRatioPriority represents a UpdateTokenRatioPriority event raised by the V3nftdescriptor contract.
type V3nftdescriptorUpdateTokenRatioPriority struct {
	Token    common.Address
	Priority *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateTokenRatioPriority is a free log retrieval operation binding the contract event 0x5fd68dea96a9ad4e6dc9cfad3617d7eb84c2070e6352848fb091a254f4629e92.
//
// Solidity: event UpdateTokenRatioPriority(address token, int256 priority)
func (_V3nftdescriptor *V3nftdescriptorFilterer) FilterUpdateTokenRatioPriority(opts *bind.FilterOpts) (*V3nftdescriptorUpdateTokenRatioPriorityIterator, error) {

	logs, sub, err := _V3nftdescriptor.contract.FilterLogs(opts, "UpdateTokenRatioPriority")
	if err != nil {
		return nil, err
	}
	return &V3nftdescriptorUpdateTokenRatioPriorityIterator{contract: _V3nftdescriptor.contract, event: "UpdateTokenRatioPriority", logs: logs, sub: sub}, nil
}

// WatchUpdateTokenRatioPriority is a free log subscription operation binding the contract event 0x5fd68dea96a9ad4e6dc9cfad3617d7eb84c2070e6352848fb091a254f4629e92.
//
// Solidity: event UpdateTokenRatioPriority(address token, int256 priority)
func (_V3nftdescriptor *V3nftdescriptorFilterer) WatchUpdateTokenRatioPriority(opts *bind.WatchOpts, sink chan<- *V3nftdescriptorUpdateTokenRatioPriority) (event.Subscription, error) {

	logs, sub, err := _V3nftdescriptor.contract.WatchLogs(opts, "UpdateTokenRatioPriority")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V3nftdescriptorUpdateTokenRatioPriority)
				if err := _V3nftdescriptor.contract.UnpackLog(event, "UpdateTokenRatioPriority", log); err != nil {
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

// ParseUpdateTokenRatioPriority is a log parse operation binding the contract event 0x5fd68dea96a9ad4e6dc9cfad3617d7eb84c2070e6352848fb091a254f4629e92.
//
// Solidity: event UpdateTokenRatioPriority(address token, int256 priority)
func (_V3nftdescriptor *V3nftdescriptorFilterer) ParseUpdateTokenRatioPriority(log types.Log) (*V3nftdescriptorUpdateTokenRatioPriority, error) {
	event := new(V3nftdescriptorUpdateTokenRatioPriority)
	if err := _V3nftdescriptor.contract.UnpackLog(event, "UpdateTokenRatioPriority", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
