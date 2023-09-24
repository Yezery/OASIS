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

// UsrMnemonicABI is the input ABI used to generate the binding from.
const UsrMnemonicABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"Sp1\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"Sp2\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"Sp3\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"forgetMnemonic\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"Sp1\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"Sp2\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"Sp3\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"NewMnemonic\",\"type\":\"string\"}],\"name\":\"reSetMnemonic\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"Sp1\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"Sp2\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"Sp3\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"setAuthenticationMetaInformation\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"Mnemonic\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"setMnemonic\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userMnemonic\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// UsrMnemonicBin is the compiled bytecode used for deploying new contracts.
var UsrMnemonicBin = "0x608060405234801561001057600080fd5b506120e1806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80633991621d1461005c57806350a8a354146101195780636505280b146103d1578063b7c44dc114610525578063c78c605014610746575b600080fd5b61009e6004803603602081101561007257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506109c8565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156100de5780820151818401526020810190506100c3565b50505050905090810190601f16801561010b5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6103b7600480360360a081101561012f57600080fd5b810190808035906020019064010000000081111561014c57600080fd5b82018360208201111561015e57600080fd5b8035906020019184600183028401116401000000008311171561018057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803590602001906401000000008111156101e357600080fd5b8201836020820111156101f557600080fd5b8035906020019184600183028401116401000000008311171561021757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561027a57600080fd5b82018360208201111561028c57600080fd5b803590602001918460018302840111640100000000831117156102ae57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019064010000000081111561033157600080fd5b82018360208201111561034357600080fd5b8035906020019184600183028401116401000000008311171561036557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610a78565b604051808215151515815260200191505060405180910390f35b6104aa600480360360408110156103e757600080fd5b810190808035906020019064010000000081111561040457600080fd5b82018360208201111561041657600080fd5b8035906020019184600183028401116401000000008311171561043857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506110fc565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156104ea5780820151818401526020810190506104cf565b50505050905090810190601f1680156105175780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61072c6004803603608081101561053b57600080fd5b810190808035906020019064010000000081111561055857600080fd5b82018360208201111561056a57600080fd5b8035906020019184600183028401116401000000008311171561058c57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803590602001906401000000008111156105ef57600080fd5b82018360208201111561060157600080fd5b8035906020019184600183028401116401000000008311171561062357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561068657600080fd5b82018360208201111561069857600080fd5b803590602001918460018302840111640100000000831117156106ba57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061140b565b604051808215151515815260200191505060405180910390f35b61094d6004803603608081101561075c57600080fd5b810190808035906020019064010000000081111561077957600080fd5b82018360208201111561078b57600080fd5b803590602001918460018302840111640100000000831117156107ad57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561081057600080fd5b82018360208201111561082257600080fd5b8035906020019184600183028401116401000000008311171561084457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803590602001906401000000008111156108a757600080fd5b8201836020820111156108b957600080fd5b803590602001918460018302840111640100000000831117156108db57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611a3b565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561098d578082015181840152602081019050610972565b50505050905090810190601f1680156109ba5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60006020528060005260406000206000915090508054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a705780601f10610a4557610100808354040283529160200191610a70565b820191906000526020600020905b815481529060010190602001808311610a5357829003601f168201915b505050505081565b600085858585600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600060038110610ac957fe5b016040516020018082805460018160011615610100020316600290048015610b285780601f10610b06576101008083540402835291820191610b28565b820191906000526020600020905b815481529060010190602001808311610b14575b505091505060405160208183030381529060405280519060200120846040516020018082805190602001908083835b60208310610b7a5780518252602082019150602081019050602083039250610b57565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405160208183030381529060405280519060200120148015610cf65750600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600160038110610c0857fe5b016040516020018082805460018160011615610100020316600290048015610c675780601f10610c45576101008083540402835291820191610c67565b820191906000526020600020905b815481529060010190602001808311610c53575b505091505060405160208183030381529060405280519060200120836040516020018082805190602001908083835b60208310610cb95780518252602082019150602081019050602083039250610c96565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405160208183030381529060405280519060200120145b8015610e365750600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600260038110610d4857fe5b016040516020018082805460018160011615610100020316600290048015610da75780601f10610d85576101008083540402835291820191610da7565b820191906000526020600020905b815481529060010190602001808311610d93575b505091505060405160208183030381529060405280519060200120826040516020018082805190602001908083835b60208310610df95780518252602082019150602081019050602083039250610dd6565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405160208183030381529060405280519060200120145b610ea8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260038152602001807f353035000000000000000000000000000000000000000000000000000000000081525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff161415610f4b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260038152602001807f353035000000000000000000000000000000000000000000000000000000000081525060200191505060405180910390fd5b604051602001806000019050604051602081830303815290604052805190602001206000808973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060405160200180828054600181600116156101000203166002900480156110095780601f10610fe7576101008083540402835291820191611009565b820191906000526020600020905b815481529060010190602001808311610ff5575b5050915050604051602081830303815290604052805190602001201415611098576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260038152602001807f343034000000000000000000000000000000000000000000000000000000000081525060200191505060405180910390fd5b856000808973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090805190602001906110ea929190611f18565b50600194505050505095945050505050565b6060604051602001806000019050604051602081830303815290604052805190602001206000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060405160200180828054600181600116156101000203166002900480156111bc5780601f1061119a5761010080835404028352918201916111bc565b820191906000526020600020905b8154815290600101906020018083116111a8575b5050915050604051602081830303815290604052805190602001201461124a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260038152602001807f353030000000000000000000000000000000000000000000000000000000000081525060200191505060405180910390fd5b6040518060000190506040518091039020836040516020018082805190602001908083835b60208310611292578051825260208201915060208101905060208303925061126f565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405160208183030381529060405280519060200120141580156113085750600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b61137a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260038152602001807f353035000000000000000000000000000000000000000000000000000000000081525060200191505060405180910390fd5b826000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090805190602001906113cc929190611f18565b506040518060400160405280600381526020017f3430300000000000000000000000000000000000000000000000000000000000815250905092915050565b600084848484600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006003811061145c57fe5b0160405160200180828054600181600116156101000203166002900480156114bb5780601f106114995761010080835404028352918201916114bb565b820191906000526020600020905b8154815290600101906020018083116114a7575b505091505060405160208183030381529060405280519060200120846040516020018082805190602001908083835b6020831061150d57805182526020820191506020810190506020830392506114ea565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051602081830303815290604052805190602001201480156116895750600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060016003811061159b57fe5b0160405160200180828054600181600116156101000203166002900480156115fa5780601f106115d85761010080835404028352918201916115fa565b820191906000526020600020905b8154815290600101906020018083116115e6575b505091505060405160208183030381529060405280519060200120836040516020018082805190602001908083835b6020831061164c5780518252602082019150602081019050602083039250611629565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405160208183030381529060405280519060200120145b80156117c95750600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206002600381106116db57fe5b01604051602001808280546001816001161561010002031660029004801561173a5780601f1061171857610100808354040283529182019161173a565b820191906000526020600020905b815481529060010190602001808311611726575b505091505060405160208183030381529060405280519060200120826040516020018082805190602001908083835b6020831061178c5780518252602082019150602081019050602083039250611769565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405160208183030381529060405280519060200120145b61183b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260038152602001807f353035000000000000000000000000000000000000000000000000000000000081525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff1614156118de576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260038152602001807f353035000000000000000000000000000000000000000000000000000000000081525060200191505060405180910390fd5b604051602001806000019050604051602081830303815290604052805190602001206000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020604051602001808280546001816001161561010002031660029004801561199c5780601f1061197a57610100808354040283529182019161199c565b820191906000526020600020905b815481529060010190602001808311611988575b5050915050604051602081830303815290604052805190602001201415611a2b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260038152602001807f343034000000000000000000000000000000000000000000000000000000000081525060200191505060405180910390fd5b6001945050505050949350505050565b6060604051602001806000019050604051602081830303815290604052805190602001206000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040516020018082805460018160011615610100020316600290048015611afb5780601f10611ad9576101008083540402835291820191611afb565b820191906000526020600020905b815481529060010190602001808311611ae7575b5050915050604051602081830303815290604052805190602001201415611b8a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260038152602001807f343034000000000000000000000000000000000000000000000000000000000081525060200191505060405180910390fd5b6040518060000190506040518091039020856040516020018082805190602001908083835b60208310611bd25780518252602082019150602081019050602083039250611baf565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040528051906020012014158015611c9c57506040518060000190506040518091039020846040516020018082805190602001908083835b60208310611c5e5780518252602082019150602081019050602083039250611c3b565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040528051906020012014155b8015611d2957506040518060000190506040518091039020836040516020018082805190602001908083835b60208310611ceb5780518252602082019150602081019050602083039250611cc8565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040528051906020012014155b611d9b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260038152602001807f353035000000000000000000000000000000000000000000000000000000000081525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611e3e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260038152602001807f353035000000000000000000000000000000000000000000000000000000000081525060200191505060405180910390fd5b611e46611f98565b8581600060038110611e5457fe5b60200201819052508481600160038110611e6a57fe5b60200201819052508381600260038110611e8057fe5b602002018190525080600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020906003611ed6929190611fbf565b506040518060400160405280600381526020017f3430300000000000000000000000000000000000000000000000000000000000815250915050949350505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611f5957805160ff1916838001178555611f87565b82800160010185558215611f87579182015b82811115611f86578251825591602001919060010190611f6b565b5b509050611f949190612012565b5090565b60405180606001604052806003905b6060815260200190600190039081611fa75790505090565b8260038101928215612001579160200282015b82811115612000578251829080519060200190611ff0929190611f18565b5091602001919060010190611fd2565b5b50905061200e9190612037565b5090565b61203491905b80821115612030576000816000905550600101612018565b5090565b90565b61206091905b8082111561205c57600081816120539190612063565b5060010161203d565b5090565b90565b50805460018160011615610100020316600290046000825580601f1061208957506120a8565b601f0160209004906000526020600020908101906120a79190612012565b5b5056fea26469706673582212201a40cd828442fed10ec581c4b0d0a810675127438598b88be07a887117190ee264736f6c634300060a0033"

// DeployUsrMnemonic deploys a new contract, binding an instance of UsrMnemonic to it.
func DeployUsrMnemonic(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UsrMnemonic, error) {
	parsed, err := abi.JSON(strings.NewReader(UsrMnemonicABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UsrMnemonicBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UsrMnemonic{UsrMnemonicCaller: UsrMnemonicCaller{contract: contract}, UsrMnemonicTransactor: UsrMnemonicTransactor{contract: contract}, UsrMnemonicFilterer: UsrMnemonicFilterer{contract: contract}}, nil
}

func AsyncDeployUsrMnemonic(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(UsrMnemonicABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(UsrMnemonicBin), backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// UsrMnemonic is an auto generated Go binding around a Solidity contract.
type UsrMnemonic struct {
	UsrMnemonicCaller     // Read-only binding to the contract
	UsrMnemonicTransactor // Write-only binding to the contract
	UsrMnemonicFilterer   // Log filterer for contract events
}

// UsrMnemonicCaller is an auto generated read-only Go binding around a Solidity contract.
type UsrMnemonicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsrMnemonicTransactor is an auto generated write-only Go binding around a Solidity contract.
type UsrMnemonicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsrMnemonicFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type UsrMnemonicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsrMnemonicSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type UsrMnemonicSession struct {
	Contract     *UsrMnemonic      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UsrMnemonicCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type UsrMnemonicCallerSession struct {
	Contract *UsrMnemonicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// UsrMnemonicTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type UsrMnemonicTransactorSession struct {
	Contract     *UsrMnemonicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// UsrMnemonicRaw is an auto generated low-level Go binding around a Solidity contract.
type UsrMnemonicRaw struct {
	Contract *UsrMnemonic // Generic contract binding to access the raw methods on
}

// UsrMnemonicCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type UsrMnemonicCallerRaw struct {
	Contract *UsrMnemonicCaller // Generic read-only contract binding to access the raw methods on
}

// UsrMnemonicTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type UsrMnemonicTransactorRaw struct {
	Contract *UsrMnemonicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUsrMnemonic creates a new instance of UsrMnemonic, bound to a specific deployed contract.
func NewUsrMnemonic(address common.Address, backend bind.ContractBackend) (*UsrMnemonic, error) {
	contract, err := bindUsrMnemonic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UsrMnemonic{UsrMnemonicCaller: UsrMnemonicCaller{contract: contract}, UsrMnemonicTransactor: UsrMnemonicTransactor{contract: contract}, UsrMnemonicFilterer: UsrMnemonicFilterer{contract: contract}}, nil
}

// NewUsrMnemonicCaller creates a new read-only instance of UsrMnemonic, bound to a specific deployed contract.
func NewUsrMnemonicCaller(address common.Address, caller bind.ContractCaller) (*UsrMnemonicCaller, error) {
	contract, err := bindUsrMnemonic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UsrMnemonicCaller{contract: contract}, nil
}

// NewUsrMnemonicTransactor creates a new write-only instance of UsrMnemonic, bound to a specific deployed contract.
func NewUsrMnemonicTransactor(address common.Address, transactor bind.ContractTransactor) (*UsrMnemonicTransactor, error) {
	contract, err := bindUsrMnemonic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UsrMnemonicTransactor{contract: contract}, nil
}

// NewUsrMnemonicFilterer creates a new log filterer instance of UsrMnemonic, bound to a specific deployed contract.
func NewUsrMnemonicFilterer(address common.Address, filterer bind.ContractFilterer) (*UsrMnemonicFilterer, error) {
	contract, err := bindUsrMnemonic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UsrMnemonicFilterer{contract: contract}, nil
}

// bindUsrMnemonic binds a generic wrapper to an already deployed contract.
func bindUsrMnemonic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UsrMnemonicABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UsrMnemonic *UsrMnemonicRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UsrMnemonic.Contract.UsrMnemonicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UsrMnemonic *UsrMnemonicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _UsrMnemonic.Contract.UsrMnemonicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UsrMnemonic *UsrMnemonicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _UsrMnemonic.Contract.UsrMnemonicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UsrMnemonic *UsrMnemonicCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UsrMnemonic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UsrMnemonic *UsrMnemonicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _UsrMnemonic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UsrMnemonic *UsrMnemonicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _UsrMnemonic.Contract.contract.Transact(opts, method, params...)
}

// ForgetMnemonic is a free data retrieval call binding the contract method 0xb7c44dc1.
//
// Solidity: function forgetMnemonic(string Sp1, string Sp2, string Sp3, address userAddress) constant returns(bool)
func (_UsrMnemonic *UsrMnemonicCaller) ForgetMnemonic(opts *bind.CallOpts, Sp1 string, Sp2 string, Sp3 string, userAddress common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _UsrMnemonic.contract.Call(opts, out, "forgetMnemonic", Sp1, Sp2, Sp3, userAddress)
	return *ret0, err
}

// ForgetMnemonic is a free data retrieval call binding the contract method 0xb7c44dc1.
//
// Solidity: function forgetMnemonic(string Sp1, string Sp2, string Sp3, address userAddress) constant returns(bool)
func (_UsrMnemonic *UsrMnemonicSession) ForgetMnemonic(Sp1 string, Sp2 string, Sp3 string, userAddress common.Address) (bool, error) {
	return _UsrMnemonic.Contract.ForgetMnemonic(&_UsrMnemonic.CallOpts, Sp1, Sp2, Sp3, userAddress)
}

// ForgetMnemonic is a free data retrieval call binding the contract method 0xb7c44dc1.
//
// Solidity: function forgetMnemonic(string Sp1, string Sp2, string Sp3, address userAddress) constant returns(bool)
func (_UsrMnemonic *UsrMnemonicCallerSession) ForgetMnemonic(Sp1 string, Sp2 string, Sp3 string, userAddress common.Address) (bool, error) {
	return _UsrMnemonic.Contract.ForgetMnemonic(&_UsrMnemonic.CallOpts, Sp1, Sp2, Sp3, userAddress)
}

// UserMnemonic is a free data retrieval call binding the contract method 0x3991621d.
//
// Solidity: function userMnemonic(address ) constant returns(string)
func (_UsrMnemonic *UsrMnemonicCaller) UserMnemonic(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _UsrMnemonic.contract.Call(opts, out, "userMnemonic", arg0)
	return *ret0, err
}

// UserMnemonic is a free data retrieval call binding the contract method 0x3991621d.
//
// Solidity: function userMnemonic(address ) constant returns(string)
func (_UsrMnemonic *UsrMnemonicSession) UserMnemonic(arg0 common.Address) (string, error) {
	return _UsrMnemonic.Contract.UserMnemonic(&_UsrMnemonic.CallOpts, arg0)
}

// UserMnemonic is a free data retrieval call binding the contract method 0x3991621d.
//
// Solidity: function userMnemonic(address ) constant returns(string)
func (_UsrMnemonic *UsrMnemonicCallerSession) UserMnemonic(arg0 common.Address) (string, error) {
	return _UsrMnemonic.Contract.UserMnemonic(&_UsrMnemonic.CallOpts, arg0)
}

// ReSetMnemonic is a paid mutator transaction binding the contract method 0x50a8a354.
//
// Solidity: function reSetMnemonic(string Sp1, string Sp2, string Sp3, address userAddress, string NewMnemonic) returns(bool)
func (_UsrMnemonic *UsrMnemonicTransactor) ReSetMnemonic(opts *bind.TransactOpts, Sp1 string, Sp2 string, Sp3 string, userAddress common.Address, NewMnemonic string) (*types.Transaction, *types.Receipt, error) {
	return _UsrMnemonic.contract.Transact(opts, "reSetMnemonic", Sp1, Sp2, Sp3, userAddress, NewMnemonic)
}

func (_UsrMnemonic *UsrMnemonicTransactor) AsyncReSetMnemonic(handler func(*types.Receipt, error), opts *bind.TransactOpts, Sp1 string, Sp2 string, Sp3 string, userAddress common.Address, NewMnemonic string) (*types.Transaction, error) {
	return _UsrMnemonic.contract.AsyncTransact(opts, handler, "reSetMnemonic", Sp1, Sp2, Sp3, userAddress, NewMnemonic)
}

// ReSetMnemonic is a paid mutator transaction binding the contract method 0x50a8a354.
//
// Solidity: function reSetMnemonic(string Sp1, string Sp2, string Sp3, address userAddress, string NewMnemonic) returns(bool)
func (_UsrMnemonic *UsrMnemonicSession) ReSetMnemonic(Sp1 string, Sp2 string, Sp3 string, userAddress common.Address, NewMnemonic string) (*types.Transaction, *types.Receipt, error) {
	return _UsrMnemonic.Contract.ReSetMnemonic(&_UsrMnemonic.TransactOpts, Sp1, Sp2, Sp3, userAddress, NewMnemonic)
}

func (_UsrMnemonic *UsrMnemonicSession) AsyncReSetMnemonic(handler func(*types.Receipt, error), Sp1 string, Sp2 string, Sp3 string, userAddress common.Address, NewMnemonic string) (*types.Transaction, error) {
	return _UsrMnemonic.Contract.AsyncReSetMnemonic(handler, &_UsrMnemonic.TransactOpts, Sp1, Sp2, Sp3, userAddress, NewMnemonic)
}

// ReSetMnemonic is a paid mutator transaction binding the contract method 0x50a8a354.
//
// Solidity: function reSetMnemonic(string Sp1, string Sp2, string Sp3, address userAddress, string NewMnemonic) returns(bool)
func (_UsrMnemonic *UsrMnemonicTransactorSession) ReSetMnemonic(Sp1 string, Sp2 string, Sp3 string, userAddress common.Address, NewMnemonic string) (*types.Transaction, *types.Receipt, error) {
	return _UsrMnemonic.Contract.ReSetMnemonic(&_UsrMnemonic.TransactOpts, Sp1, Sp2, Sp3, userAddress, NewMnemonic)
}

func (_UsrMnemonic *UsrMnemonicTransactorSession) AsyncReSetMnemonic(handler func(*types.Receipt, error), Sp1 string, Sp2 string, Sp3 string, userAddress common.Address, NewMnemonic string) (*types.Transaction, error) {
	return _UsrMnemonic.Contract.AsyncReSetMnemonic(handler, &_UsrMnemonic.TransactOpts, Sp1, Sp2, Sp3, userAddress, NewMnemonic)
}

// SetAuthenticationMetaInformation is a paid mutator transaction binding the contract method 0xc78c6050.
//
// Solidity: function setAuthenticationMetaInformation(string Sp1, string Sp2, string Sp3, address userAddress) returns(string)
func (_UsrMnemonic *UsrMnemonicTransactor) SetAuthenticationMetaInformation(opts *bind.TransactOpts, Sp1 string, Sp2 string, Sp3 string, userAddress common.Address) (*types.Transaction, *types.Receipt, error) {
	return _UsrMnemonic.contract.Transact(opts, "setAuthenticationMetaInformation", Sp1, Sp2, Sp3, userAddress)
}

func (_UsrMnemonic *UsrMnemonicTransactor) AsyncSetAuthenticationMetaInformation(handler func(*types.Receipt, error), opts *bind.TransactOpts, Sp1 string, Sp2 string, Sp3 string, userAddress common.Address) (*types.Transaction, error) {
	return _UsrMnemonic.contract.AsyncTransact(opts, handler, "setAuthenticationMetaInformation", Sp1, Sp2, Sp3, userAddress)
}

// SetAuthenticationMetaInformation is a paid mutator transaction binding the contract method 0xc78c6050.
//
// Solidity: function setAuthenticationMetaInformation(string Sp1, string Sp2, string Sp3, address userAddress) returns(string)
func (_UsrMnemonic *UsrMnemonicSession) SetAuthenticationMetaInformation(Sp1 string, Sp2 string, Sp3 string, userAddress common.Address) (*types.Transaction, *types.Receipt, error) {
	return _UsrMnemonic.Contract.SetAuthenticationMetaInformation(&_UsrMnemonic.TransactOpts, Sp1, Sp2, Sp3, userAddress)
}

func (_UsrMnemonic *UsrMnemonicSession) AsyncSetAuthenticationMetaInformation(handler func(*types.Receipt, error), Sp1 string, Sp2 string, Sp3 string, userAddress common.Address) (*types.Transaction, error) {
	return _UsrMnemonic.Contract.AsyncSetAuthenticationMetaInformation(handler, &_UsrMnemonic.TransactOpts, Sp1, Sp2, Sp3, userAddress)
}

// SetAuthenticationMetaInformation is a paid mutator transaction binding the contract method 0xc78c6050.
//
// Solidity: function setAuthenticationMetaInformation(string Sp1, string Sp2, string Sp3, address userAddress) returns(string)
func (_UsrMnemonic *UsrMnemonicTransactorSession) SetAuthenticationMetaInformation(Sp1 string, Sp2 string, Sp3 string, userAddress common.Address) (*types.Transaction, *types.Receipt, error) {
	return _UsrMnemonic.Contract.SetAuthenticationMetaInformation(&_UsrMnemonic.TransactOpts, Sp1, Sp2, Sp3, userAddress)
}

func (_UsrMnemonic *UsrMnemonicTransactorSession) AsyncSetAuthenticationMetaInformation(handler func(*types.Receipt, error), Sp1 string, Sp2 string, Sp3 string, userAddress common.Address) (*types.Transaction, error) {
	return _UsrMnemonic.Contract.AsyncSetAuthenticationMetaInformation(handler, &_UsrMnemonic.TransactOpts, Sp1, Sp2, Sp3, userAddress)
}

// SetMnemonic is a paid mutator transaction binding the contract method 0x6505280b.
//
// Solidity: function setMnemonic(string Mnemonic, address userAddress) returns(string)
func (_UsrMnemonic *UsrMnemonicTransactor) SetMnemonic(opts *bind.TransactOpts, Mnemonic string, userAddress common.Address) (*types.Transaction, *types.Receipt, error) {
	return _UsrMnemonic.contract.Transact(opts, "setMnemonic", Mnemonic, userAddress)
}

func (_UsrMnemonic *UsrMnemonicTransactor) AsyncSetMnemonic(handler func(*types.Receipt, error), opts *bind.TransactOpts, Mnemonic string, userAddress common.Address) (*types.Transaction, error) {
	return _UsrMnemonic.contract.AsyncTransact(opts, handler, "setMnemonic", Mnemonic, userAddress)
}

// SetMnemonic is a paid mutator transaction binding the contract method 0x6505280b.
//
// Solidity: function setMnemonic(string Mnemonic, address userAddress) returns(string)
func (_UsrMnemonic *UsrMnemonicSession) SetMnemonic(Mnemonic string, userAddress common.Address) (*types.Transaction, *types.Receipt, error) {
	return _UsrMnemonic.Contract.SetMnemonic(&_UsrMnemonic.TransactOpts, Mnemonic, userAddress)
}

func (_UsrMnemonic *UsrMnemonicSession) AsyncSetMnemonic(handler func(*types.Receipt, error), Mnemonic string, userAddress common.Address) (*types.Transaction, error) {
	return _UsrMnemonic.Contract.AsyncSetMnemonic(handler, &_UsrMnemonic.TransactOpts, Mnemonic, userAddress)
}

// SetMnemonic is a paid mutator transaction binding the contract method 0x6505280b.
//
// Solidity: function setMnemonic(string Mnemonic, address userAddress) returns(string)
func (_UsrMnemonic *UsrMnemonicTransactorSession) SetMnemonic(Mnemonic string, userAddress common.Address) (*types.Transaction, *types.Receipt, error) {
	return _UsrMnemonic.Contract.SetMnemonic(&_UsrMnemonic.TransactOpts, Mnemonic, userAddress)
}

func (_UsrMnemonic *UsrMnemonicTransactorSession) AsyncSetMnemonic(handler func(*types.Receipt, error), Mnemonic string, userAddress common.Address) (*types.Transaction, error) {
	return _UsrMnemonic.Contract.AsyncSetMnemonic(handler, &_UsrMnemonic.TransactOpts, Mnemonic, userAddress)
}
