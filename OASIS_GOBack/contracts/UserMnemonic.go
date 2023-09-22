// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
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
)

// UserMnemonicABI is the input ABI used to generate the binding from.
const UserMnemonicABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"Sp1\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"Sp2\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"Sp3\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"forgetMnemonic\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"Sp1\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"Sp2\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"Sp3\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"setAuthenticationMetaInformation\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"Mnemonic\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"setMnemonic\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userMnemonic\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// UserMnemonicBin is the compiled bytecode used for deploying new contracts.
var UserMnemonicBin = "0x608060405234801561001057600080fd5b506118d0806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80633991621d146100515780636505280b1461010e578063b7c44dc114610262578063c78c6050146104e4575b600080fd5b6100936004803603602081101561006757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610766565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156100d35780820151818401526020810190506100b8565b50505050905090810190601f1680156101005780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101e76004803603604081101561012457600080fd5b810190808035906020019064010000000081111561014157600080fd5b82018360208201111561015357600080fd5b8035906020019184600183028401116401000000008311171561017557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610816565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561022757808201518184015260208101905061020c565b50505050905090810190601f1680156102545780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6104696004803603608081101561027857600080fd5b810190808035906020019064010000000081111561029557600080fd5b8201836020820111156102a757600080fd5b803590602001918460018302840111640100000000831117156102c957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561032c57600080fd5b82018360208201111561033e57600080fd5b8035906020019184600183028401116401000000008311171561036057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803590602001906401000000008111156103c357600080fd5b8201836020820111156103d557600080fd5b803590602001918460018302840111640100000000831117156103f757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610b25565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156104a957808201518184015260208101905061048e565b50505050905090810190601f1680156104d65780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6106eb600480360360808110156104fa57600080fd5b810190808035906020019064010000000081111561051757600080fd5b82018360208201111561052957600080fd5b8035906020019184600183028401116401000000008311171561054b57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803590602001906401000000008111156105ae57600080fd5b8201836020820111156105c057600080fd5b803590602001918460018302840111640100000000831117156105e257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561064557600080fd5b82018360208201111561065757600080fd5b8035906020019184600183028401116401000000008311171561067957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061122a565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561072b578082015181840152602081019050610710565b50505050905090810190601f1680156107585780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60006020528060005260406000206000915090508054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561080e5780601f106107e35761010080835404028352916020019161080e565b820191906000526020600020905b8154815290600101906020018083116107f157829003601f168201915b505050505081565b6060604051602001806000019050604051602081830303815290604052805190602001206000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060405160200180828054600181600116156101000203166002900480156108d65780601f106108b45761010080835404028352918201916108d6565b820191906000526020600020905b8154815290600101906020018083116108c2575b50509150506040516020818303038152906040528051906020012014610964576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260038152602001807f353030000000000000000000000000000000000000000000000000000000000081525060200191505060405180910390fd5b6040518060000190506040518091039020836040516020018082805190602001908083835b602083106109ac5780518252602082019150602081019050602083039250610989565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040528051906020012014158015610a225750600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b610a94576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260038152602001807f353035000000000000000000000000000000000000000000000000000000000081525060200191505060405180910390fd5b826000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209080519060200190610ae6929190611707565b506040518060400160405280600381526020017f3430300000000000000000000000000000000000000000000000000000000000815250905092915050565b606084848484600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600060038110610b7657fe5b016040516020018082805460018160011615610100020316600290048015610bd55780601f10610bb3576101008083540402835291820191610bd5565b820191906000526020600020905b815481529060010190602001808311610bc1575b505091505060405160208183030381529060405280519060200120846040516020018082805190602001908083835b60208310610c275780518252602082019150602081019050602083039250610c04565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405160208183030381529060405280519060200120148015610da35750600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600160038110610cb557fe5b016040516020018082805460018160011615610100020316600290048015610d145780601f10610cf2576101008083540402835291820191610d14565b820191906000526020600020905b815481529060010190602001808311610d00575b505091505060405160208183030381529060405280519060200120836040516020018082805190602001908083835b60208310610d665780518252602082019150602081019050602083039250610d43565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405160208183030381529060405280519060200120145b8015610ee35750600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600260038110610df557fe5b016040516020018082805460018160011615610100020316600290048015610e545780601f10610e32576101008083540402835291820191610e54565b820191906000526020600020905b815481529060010190602001808311610e40575b505091505060405160208183030381529060405280519060200120826040516020018082805190602001908083835b60208310610ea65780518252602082019150602081019050602083039250610e83565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405160208183030381529060405280519060200120145b610f55576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260038152602001807f353035000000000000000000000000000000000000000000000000000000000081525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff161415610ff8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260038152602001807f353035000000000000000000000000000000000000000000000000000000000081525060200191505060405180910390fd5b604051602001806000019050604051602081830303815290604052805190602001206000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060405160200180828054600181600116156101000203166002900480156110b65780601f106110945761010080835404028352918201916110b6565b820191906000526020600020905b8154815290600101906020018083116110a2575b5050915050604051602081830303815290604052805190602001201415611145576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260038152602001807f343034000000000000000000000000000000000000000000000000000000000081525060200191505060405180910390fd5b6000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156112175780601f106111ec57610100808354040283529160200191611217565b820191906000526020600020905b8154815290600101906020018083116111fa57829003601f168201915b5050505050945050505050949350505050565b6060604051602001806000019050604051602081830303815290604052805190602001206000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060405160200180828054600181600116156101000203166002900480156112ea5780601f106112c85761010080835404028352918201916112ea565b820191906000526020600020905b8154815290600101906020018083116112d6575b5050915050604051602081830303815290604052805190602001201415611379576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260038152602001807f343034000000000000000000000000000000000000000000000000000000000081525060200191505060405180910390fd5b6040518060000190506040518091039020856040516020018082805190602001908083835b602083106113c1578051825260208201915060208101905060208303925061139e565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051602081830303815290604052805190602001201415801561148b57506040518060000190506040518091039020846040516020018082805190602001908083835b6020831061144d578051825260208201915060208101905060208303925061142a565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040528051906020012014155b801561151857506040518060000190506040518091039020836040516020018082805190602001908083835b602083106114da57805182526020820191506020810190506020830392506114b7565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040528051906020012014155b61158a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260038152602001807f353035000000000000000000000000000000000000000000000000000000000081525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561162d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260038152602001807f353035000000000000000000000000000000000000000000000000000000000081525060200191505060405180910390fd5b611635611787565b858160006003811061164357fe5b6020020181905250848160016003811061165957fe5b6020020181905250838160026003811061166f57fe5b602002018190525080600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209060036116c59291906117ae565b506040518060400160405280600381526020017f3430300000000000000000000000000000000000000000000000000000000000815250915050949350505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061174857805160ff1916838001178555611776565b82800160010185558215611776579182015b8281111561177557825182559160200191906001019061175a565b5b5090506117839190611801565b5090565b60405180606001604052806003905b60608152602001906001900390816117965790505090565b82600381019282156117f0579160200282015b828111156117ef5782518290805190602001906117df929190611707565b50916020019190600101906117c1565b5b5090506117fd9190611826565b5090565b61182391905b8082111561181f576000816000905550600101611807565b5090565b90565b61184f91905b8082111561184b57600081816118429190611852565b5060010161182c565b5090565b90565b50805460018160011615610100020316600290046000825580601f106118785750611897565b601f0160209004906000526020600020908101906118969190611801565b5b5056fea26469706673582212206acc6205882379cc7155e3e336677261d57630360b8ff866f10ac5430ff2217e64736f6c634300060a0033"

// DeployUserMnemonic deploys a new contract, binding an instance of UserMnemonic to it.
func DeployUserMnemonic(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UserMnemonic, error) {
	parsed, err := abi.JSON(strings.NewReader(UserMnemonicABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UserMnemonicBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UserMnemonic{UserMnemonicCaller: UserMnemonicCaller{contract: contract}, UserMnemonicTransactor: UserMnemonicTransactor{contract: contract}, UserMnemonicFilterer: UserMnemonicFilterer{contract: contract}}, nil
}

func AsyncDeployUserMnemonic(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(UserMnemonicABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(UserMnemonicBin), backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// UserMnemonic is an auto generated Go binding around a Solidity contract.
type UserMnemonic struct {
	UserMnemonicCaller     // Read-only binding to the contract
	UserMnemonicTransactor // Write-only binding to the contract
	UserMnemonicFilterer   // Log filterer for contract events
}

// UserMnemonicCaller is an auto generated read-only Go binding around a Solidity contract.
type UserMnemonicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserMnemonicTransactor is an auto generated write-only Go binding around a Solidity contract.
type UserMnemonicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserMnemonicFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type UserMnemonicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserMnemonicSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type UserMnemonicSession struct {
	Contract     *UserMnemonic     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserMnemonicCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type UserMnemonicCallerSession struct {
	Contract *UserMnemonicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// UserMnemonicTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type UserMnemonicTransactorSession struct {
	Contract     *UserMnemonicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// UserMnemonicRaw is an auto generated low-level Go binding around a Solidity contract.
type UserMnemonicRaw struct {
	Contract *UserMnemonic // Generic contract binding to access the raw methods on
}

// UserMnemonicCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type UserMnemonicCallerRaw struct {
	Contract *UserMnemonicCaller // Generic read-only contract binding to access the raw methods on
}

// UserMnemonicTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type UserMnemonicTransactorRaw struct {
	Contract *UserMnemonicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUserMnemonic creates a new instance of UserMnemonic, bound to a specific deployed contract.
func NewUserMnemonic(address common.Address, backend bind.ContractBackend) (*UserMnemonic, error) {
	contract, err := bindUserMnemonic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UserMnemonic{UserMnemonicCaller: UserMnemonicCaller{contract: contract}, UserMnemonicTransactor: UserMnemonicTransactor{contract: contract}, UserMnemonicFilterer: UserMnemonicFilterer{contract: contract}}, nil
}

// NewUserMnemonicCaller creates a new read-only instance of UserMnemonic, bound to a specific deployed contract.
func NewUserMnemonicCaller(address common.Address, caller bind.ContractCaller) (*UserMnemonicCaller, error) {
	contract, err := bindUserMnemonic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UserMnemonicCaller{contract: contract}, nil
}

// NewUserMnemonicTransactor creates a new write-only instance of UserMnemonic, bound to a specific deployed contract.
func NewUserMnemonicTransactor(address common.Address, transactor bind.ContractTransactor) (*UserMnemonicTransactor, error) {
	contract, err := bindUserMnemonic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UserMnemonicTransactor{contract: contract}, nil
}

// NewUserMnemonicFilterer creates a new log filterer instance of UserMnemonic, bound to a specific deployed contract.
func NewUserMnemonicFilterer(address common.Address, filterer bind.ContractFilterer) (*UserMnemonicFilterer, error) {
	contract, err := bindUserMnemonic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UserMnemonicFilterer{contract: contract}, nil
}

// bindUserMnemonic binds a generic wrapper to an already deployed contract.
func bindUserMnemonic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UserMnemonicABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserMnemonic *UserMnemonicRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UserMnemonic.Contract.UserMnemonicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserMnemonic *UserMnemonicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _UserMnemonic.Contract.UserMnemonicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserMnemonic *UserMnemonicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _UserMnemonic.Contract.UserMnemonicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserMnemonic *UserMnemonicCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UserMnemonic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserMnemonic *UserMnemonicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _UserMnemonic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserMnemonic *UserMnemonicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _UserMnemonic.Contract.contract.Transact(opts, method, params...)
}

// ForgetMnemonic is a free data retrieval call binding the contract method 0xb7c44dc1.
//
// Solidity: function forgetMnemonic(string Sp1, string Sp2, string Sp3, address userAddress) constant returns(string)
func (_UserMnemonic *UserMnemonicCaller) ForgetMnemonic(opts *bind.CallOpts, Sp1 string, Sp2 string, Sp3 string, userAddress common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _UserMnemonic.contract.Call(opts, out, "forgetMnemonic", Sp1, Sp2, Sp3, userAddress)
	return *ret0, err
}

// ForgetMnemonic is a free data retrieval call binding the contract method 0xb7c44dc1.
//
// Solidity: function forgetMnemonic(string Sp1, string Sp2, string Sp3, address userAddress) constant returns(string)
func (_UserMnemonic *UserMnemonicSession) ForgetMnemonic(Sp1 string, Sp2 string, Sp3 string, userAddress common.Address) (string, error) {
	return _UserMnemonic.Contract.ForgetMnemonic(&_UserMnemonic.CallOpts, Sp1, Sp2, Sp3, userAddress)
}

// ForgetMnemonic is a free data retrieval call binding the contract method 0xb7c44dc1.
//
// Solidity: function forgetMnemonic(string Sp1, string Sp2, string Sp3, address userAddress) constant returns(string)
func (_UserMnemonic *UserMnemonicCallerSession) ForgetMnemonic(Sp1 string, Sp2 string, Sp3 string, userAddress common.Address) (string, error) {
	return _UserMnemonic.Contract.ForgetMnemonic(&_UserMnemonic.CallOpts, Sp1, Sp2, Sp3, userAddress)
}

// UserMnemonic is a free data retrieval call binding the contract method 0x3991621d.
//
// Solidity: function userMnemonic(address ) constant returns(string)
func (_UserMnemonic *UserMnemonicCaller) UserMnemonic(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _UserMnemonic.contract.Call(opts, out, "userMnemonic", arg0)
	return *ret0, err
}

// UserMnemonic is a free data retrieval call binding the contract method 0x3991621d.
//
// Solidity: function userMnemonic(address ) constant returns(string)
func (_UserMnemonic *UserMnemonicSession) UserMnemonic(arg0 common.Address) (string, error) {
	return _UserMnemonic.Contract.UserMnemonic(&_UserMnemonic.CallOpts, arg0)
}

// UserMnemonic is a free data retrieval call binding the contract method 0x3991621d.
//
// Solidity: function userMnemonic(address ) constant returns(string)
func (_UserMnemonic *UserMnemonicCallerSession) UserMnemonic(arg0 common.Address) (string, error) {
	return _UserMnemonic.Contract.UserMnemonic(&_UserMnemonic.CallOpts, arg0)
}

// SetAuthenticationMetaInformation is a paid mutator transaction binding the contract method 0xc78c6050.
//
// Solidity: function setAuthenticationMetaInformation(string Sp1, string Sp2, string Sp3, address userAddress) returns(string)
func (_UserMnemonic *UserMnemonicTransactor) SetAuthenticationMetaInformation(opts *bind.TransactOpts, Sp1 string, Sp2 string, Sp3 string, userAddress common.Address) (*types.Transaction, *types.Receipt, error) {
	return _UserMnemonic.contract.Transact(opts, "setAuthenticationMetaInformation", Sp1, Sp2, Sp3, userAddress)
}

func (_UserMnemonic *UserMnemonicTransactor) AsyncSetAuthenticationMetaInformation(handler func(*types.Receipt, error), opts *bind.TransactOpts, Sp1 string, Sp2 string, Sp3 string, userAddress common.Address) (*types.Transaction, error) {
	return _UserMnemonic.contract.AsyncTransact(opts, handler, "setAuthenticationMetaInformation", Sp1, Sp2, Sp3, userAddress)
}

// SetAuthenticationMetaInformation is a paid mutator transaction binding the contract method 0xc78c6050.
//
// Solidity: function setAuthenticationMetaInformation(string Sp1, string Sp2, string Sp3, address userAddress) returns(string)
func (_UserMnemonic *UserMnemonicSession) SetAuthenticationMetaInformation(Sp1 string, Sp2 string, Sp3 string, userAddress common.Address) (*types.Transaction, *types.Receipt, error) {
	return _UserMnemonic.Contract.SetAuthenticationMetaInformation(&_UserMnemonic.TransactOpts, Sp1, Sp2, Sp3, userAddress)
}

func (_UserMnemonic *UserMnemonicSession) AsyncSetAuthenticationMetaInformation(handler func(*types.Receipt, error), Sp1 string, Sp2 string, Sp3 string, userAddress common.Address) (*types.Transaction, error) {
	return _UserMnemonic.Contract.AsyncSetAuthenticationMetaInformation(handler, &_UserMnemonic.TransactOpts, Sp1, Sp2, Sp3, userAddress)
}

// SetAuthenticationMetaInformation is a paid mutator transaction binding the contract method 0xc78c6050.
//
// Solidity: function setAuthenticationMetaInformation(string Sp1, string Sp2, string Sp3, address userAddress) returns(string)
func (_UserMnemonic *UserMnemonicTransactorSession) SetAuthenticationMetaInformation(Sp1 string, Sp2 string, Sp3 string, userAddress common.Address) (*types.Transaction, *types.Receipt, error) {
	return _UserMnemonic.Contract.SetAuthenticationMetaInformation(&_UserMnemonic.TransactOpts, Sp1, Sp2, Sp3, userAddress)
}

func (_UserMnemonic *UserMnemonicTransactorSession) AsyncSetAuthenticationMetaInformation(handler func(*types.Receipt, error), Sp1 string, Sp2 string, Sp3 string, userAddress common.Address) (*types.Transaction, error) {
	return _UserMnemonic.Contract.AsyncSetAuthenticationMetaInformation(handler, &_UserMnemonic.TransactOpts, Sp1, Sp2, Sp3, userAddress)
}

// SetMnemonic is a paid mutator transaction binding the contract method 0x6505280b.
//
// Solidity: function setMnemonic(string Mnemonic, address userAddress) returns(string)
func (_UserMnemonic *UserMnemonicTransactor) SetMnemonic(opts *bind.TransactOpts, Mnemonic string, userAddress common.Address) (*types.Transaction, *types.Receipt, error) {
	return _UserMnemonic.contract.Transact(opts, "setMnemonic", Mnemonic, userAddress)
}

func (_UserMnemonic *UserMnemonicTransactor) AsyncSetMnemonic(handler func(*types.Receipt, error), opts *bind.TransactOpts, Mnemonic string, userAddress common.Address) (*types.Transaction, error) {
	return _UserMnemonic.contract.AsyncTransact(opts, handler, "setMnemonic", Mnemonic, userAddress)
}

// SetMnemonic is a paid mutator transaction binding the contract method 0x6505280b.
//
// Solidity: function setMnemonic(string Mnemonic, address userAddress) returns(string)
func (_UserMnemonic *UserMnemonicSession) SetMnemonic(Mnemonic string, userAddress common.Address) (*types.Transaction, *types.Receipt, error) {
	return _UserMnemonic.Contract.SetMnemonic(&_UserMnemonic.TransactOpts, Mnemonic, userAddress)
}

func (_UserMnemonic *UserMnemonicSession) AsyncSetMnemonic(handler func(*types.Receipt, error), Mnemonic string, userAddress common.Address) (*types.Transaction, error) {
	return _UserMnemonic.Contract.AsyncSetMnemonic(handler, &_UserMnemonic.TransactOpts, Mnemonic, userAddress)
}

// SetMnemonic is a paid mutator transaction binding the contract method 0x6505280b.
//
// Solidity: function setMnemonic(string Mnemonic, address userAddress) returns(string)
func (_UserMnemonic *UserMnemonicTransactorSession) SetMnemonic(Mnemonic string, userAddress common.Address) (*types.Transaction, *types.Receipt, error) {
	return _UserMnemonic.Contract.SetMnemonic(&_UserMnemonic.TransactOpts, Mnemonic, userAddress)
}

func (_UserMnemonic *UserMnemonicTransactorSession) AsyncSetMnemonic(handler func(*types.Receipt, error), Mnemonic string, userAddress common.Address) (*types.Transaction, error) {
	return _UserMnemonic.Contract.AsyncSetMnemonic(handler, &_UserMnemonic.TransactOpts, Mnemonic, userAddress)
}