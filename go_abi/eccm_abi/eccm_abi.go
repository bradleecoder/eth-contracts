// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eccm_abi

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContextABI is the input ABI used to generate the binding from.
const ContextABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// ECCUtilsABI is the input ABI used to generate the binding from.
const ECCUtilsABI = "[]"

// ECCUtilsBin is the compiled bytecode used for deploying new contracts.
var ECCUtilsBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7230582083218d3dd34146722a575b40b4b02ff7021dc4c59580e29b5d814a30d3b9612064736f6c634300050a0032"

// DeployECCUtils deploys a new Ethereum contract, binding an instance of ECCUtils to it.
func DeployECCUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ECCUtils, error) {
	parsed, err := abi.JSON(strings.NewReader(ECCUtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ECCUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ECCUtils{ECCUtilsCaller: ECCUtilsCaller{contract: contract}, ECCUtilsTransactor: ECCUtilsTransactor{contract: contract}, ECCUtilsFilterer: ECCUtilsFilterer{contract: contract}}, nil
}

// ECCUtils is an auto generated Go binding around an Ethereum contract.
type ECCUtils struct {
	ECCUtilsCaller     // Read-only binding to the contract
	ECCUtilsTransactor // Write-only binding to the contract
	ECCUtilsFilterer   // Log filterer for contract events
}

// ECCUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ECCUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECCUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ECCUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECCUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ECCUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECCUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ECCUtilsSession struct {
	Contract     *ECCUtils         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECCUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ECCUtilsCallerSession struct {
	Contract *ECCUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ECCUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ECCUtilsTransactorSession struct {
	Contract     *ECCUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ECCUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ECCUtilsRaw struct {
	Contract *ECCUtils // Generic contract binding to access the raw methods on
}

// ECCUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ECCUtilsCallerRaw struct {
	Contract *ECCUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// ECCUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ECCUtilsTransactorRaw struct {
	Contract *ECCUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewECCUtils creates a new instance of ECCUtils, bound to a specific deployed contract.
func NewECCUtils(address common.Address, backend bind.ContractBackend) (*ECCUtils, error) {
	contract, err := bindECCUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECCUtils{ECCUtilsCaller: ECCUtilsCaller{contract: contract}, ECCUtilsTransactor: ECCUtilsTransactor{contract: contract}, ECCUtilsFilterer: ECCUtilsFilterer{contract: contract}}, nil
}

// NewECCUtilsCaller creates a new read-only instance of ECCUtils, bound to a specific deployed contract.
func NewECCUtilsCaller(address common.Address, caller bind.ContractCaller) (*ECCUtilsCaller, error) {
	contract, err := bindECCUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ECCUtilsCaller{contract: contract}, nil
}

// NewECCUtilsTransactor creates a new write-only instance of ECCUtils, bound to a specific deployed contract.
func NewECCUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*ECCUtilsTransactor, error) {
	contract, err := bindECCUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ECCUtilsTransactor{contract: contract}, nil
}

// NewECCUtilsFilterer creates a new log filterer instance of ECCUtils, bound to a specific deployed contract.
func NewECCUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*ECCUtilsFilterer, error) {
	contract, err := bindECCUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ECCUtilsFilterer{contract: contract}, nil
}

// bindECCUtils binds a generic wrapper to an already deployed contract.
func bindECCUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ECCUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECCUtils *ECCUtilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ECCUtils.Contract.ECCUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECCUtils *ECCUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECCUtils.Contract.ECCUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECCUtils *ECCUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECCUtils.Contract.ECCUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECCUtils *ECCUtilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ECCUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECCUtils *ECCUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECCUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECCUtils *ECCUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECCUtils.Contract.contract.Transact(opts, method, params...)
}

// EthCrossChainManagerABI is the input ABI used to generate the binding from.
const EthCrossChainManagerABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"EthCrossChainDataAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"rawHeader\",\"type\":\"bytes\"},{\"name\":\"pubKeyList\",\"type\":\"bytes\"},{\"name\":\"sigList\",\"type\":\"bytes\"}],\"name\":\"changeBookKeeper\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"rawHeader\",\"type\":\"bytes\"},{\"name\":\"pubKeyList\",\"type\":\"bytes\"}],\"name\":\"initGenesisBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newChainId\",\"type\":\"uint64\"}],\"name\":\"setChainId\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newEthCrossChainManagerAddress\",\"type\":\"address\"}],\"name\":\"upgradeToNew\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"chainId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"toChainId\",\"type\":\"uint64\"},{\"name\":\"toContract\",\"type\":\"bytes\"},{\"name\":\"method\",\"type\":\"bytes\"},{\"name\":\"txData\",\"type\":\"bytes\"}],\"name\":\"crossChain\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\"},{\"name\":\"rawHeader\",\"type\":\"bytes\"},{\"name\":\"headerProof\",\"type\":\"bytes\"},{\"name\":\"curRawHeader\",\"type\":\"bytes\"},{\"name\":\"headerSig\",\"type\":\"bytes\"}],\"name\":\"verifyHeaderAndExecuteTx\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_eccd\",\"type\":\"address\"},{\"name\":\"_chainId\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"height\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"rawHeader\",\"type\":\"bytes\"}],\"name\":\"InitGenesisBlockEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"height\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"rawHeader\",\"type\":\"bytes\"}],\"name\":\"ChangeBookKeeperEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"txId\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"proxyOrAssetContract\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"toChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"toContract\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"rawdata\",\"type\":\"bytes\"}],\"name\":\"CrossChainEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"fromChainID\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"toContract\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"crossChainTxHash\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"fromChainTxHash\",\"type\":\"bytes\"}],\"name\":\"VerifyHeaderAndExecuteTxEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// EthCrossChainManagerFuncSigs maps the 4-byte function signature to its string representation.
var EthCrossChainManagerFuncSigs = map[string]string{
	"00ba1694": "EthCrossChainDataAddress()",
	"9a8a0592": "chainId()",
	"29dcf4ab": "changeBookKeeper(bytes,bytes,bytes)",
	"bd5cf625": "crossChain(uint64,bytes,bytes,bytes)",
	"34a773eb": "initGenesisBlock(bytes,bytes)",
	"8f32d59b": "isOwner()",
	"8da5cb5b": "owner()",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"715018a6": "renounceOwnership()",
	"6f31031d": "setChainId(uint64)",
	"f2fde38b": "transferOwnership(address)",
	"3f4ba83a": "unpause()",
	"7e724ff3": "upgradeToNew(address)",
	"d450e04c": "verifyHeaderAndExecuteTx(bytes,bytes,bytes,bytes,bytes)",
}

// EthCrossChainManagerBin is the compiled bytecode used for deploying new contracts.
var EthCrossChainManagerBin = "0x60806040523480156200001157600080fd5b50604051620048f1380380620048f1833981810160405260408110156200003757600080fd5b50805160209091015181816000620000576001600160e01b036200010b16565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3506000805460ff60a01b19169055600180546001600160401b039092167401000000000000000000000000000000000000000002600160a01b600160e01b03196001600160a01b039094166001600160a01b03199093169290921792909216179055506200010f9050565b3390565b6147d2806200011f6000396000f3fe608060405234801561001057600080fd5b50600436106100f45760003560e01c80637e724ff3116100975780639a8a0592116100665780639a8a059214610486578063bd5cf625146104aa578063d450e04c146105c8578063f2fde38b14610880576100f4565b80637e724ff3146104485780638456cb591461046e5780638da5cb5b146104765780638f32d59b1461047e576100f4565b80633f4ba83a116100d35780633f4ba83a146104085780635c975abb146104105780636f31031d14610418578063715018a61461043e576100f4565b8062ba1694146100f957806329dcf4ab1461011d57806334a773eb146102df575b600080fd5b6101016108a6565b604080516001600160a01b039092168252519081900360200190f35b6102cb6004803603606081101561013357600080fd5b810190602081018135600160201b81111561014d57600080fd5b82018360208201111561015f57600080fd5b803590602001918460018302840111600160201b8311171561018057600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156101d257600080fd5b8201836020820111156101e457600080fd5b803590602001918460018302840111600160201b8311171561020557600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561025757600080fd5b82018360208201111561026957600080fd5b803590602001918460018302840111600160201b8311171561028a57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506108b5945050505050565b604080519115158252519081900360200190f35b6102cb600480360360408110156102f557600080fd5b810190602081018135600160201b81111561030f57600080fd5b82018360208201111561032157600080fd5b803590602001918460018302840111600160201b8311171561034257600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561039457600080fd5b8201836020820111156103a657600080fd5b803590602001918460018302840111600160201b831117156103c757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610e74945050505050565b6102cb6112e4565b6102cb611463565b6102cb6004803603602081101561042e57600080fd5b50356001600160401b0316611473565b61044661153c565b005b6102cb6004803603602081101561045e57600080fd5b50356001600160a01b03166115cd565b6102cb6116dc565b610101611851565b6102cb611860565b61048e611884565b604080516001600160401b039092168252519081900360200190f35b6102cb600480360360808110156104c057600080fd5b6001600160401b038235169190810190604081016020820135600160201b8111156104ea57600080fd5b8201836020820111156104fc57600080fd5b803590602001918460018302840111600160201b8311171561051d57600080fd5b919390929091602081019035600160201b81111561053a57600080fd5b82018360208201111561054c57600080fd5b803590602001918460018302840111600160201b8311171561056d57600080fd5b919390929091602081019035600160201b81111561058a57600080fd5b82018360208201111561059c57600080fd5b803590602001918460018302840111600160201b831117156105bd57600080fd5b50909250905061189a565b6102cb600480360360a08110156105de57600080fd5b810190602081018135600160201b8111156105f857600080fd5b82018360208201111561060a57600080fd5b803590602001918460018302840111600160201b8311171561062b57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561067d57600080fd5b82018360208201111561068f57600080fd5b803590602001918460018302840111600160201b831117156106b057600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561070257600080fd5b82018360208201111561071457600080fd5b803590602001918460018302840111600160201b8311171561073557600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561078757600080fd5b82018360208201111561079957600080fd5b803590602001918460018302840111600160201b831117156107ba57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561080c57600080fd5b82018360208201111561081e57600080fd5b803590602001918460018302840111600160201b8311171561083f57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506119fb945050505050565b6104466004803603602081101561089657600080fd5b50356001600160a01b03166120e8565b6001546001600160a01b031681565b60008054600160a01b900460ff1615610908576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b610910614229565b6109198561213b565b90506000600160009054906101000a90046001600160a01b031690506000816001600160a01b0316635ac407906040518163ffffffff1660e01b815260040160206040518083038186803b15801561097057600080fd5b505afa158015610984573d6000803e3d6000fd5b505050506040513d602081101561099a57600080fd5b5051606084015163ffffffff91821692501681106109e95760405162461bcd60e51b815260040180806020018281038252603e815260200180614446603e913960400191505060405180910390fd5b6101408301516bffffffffffffffffffffffff1916610a395760405162461bcd60e51b81526004018080602001828103825260258152602001806143186025913960400191505060405180910390fd5b6060610b04836001600160a01b03166369d480746040518163ffffffff1660e01b815260040160006040518083038186803b158015610a7757600080fd5b505afa158015610a8b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015610ab457600080fd5b810190808051600160201b811115610acb57600080fd5b82016020810184811115610ade57600080fd5b8151600160201b811182820187101715610af757600080fd5b505092919050505061225a565b8051909150610b1f898884600360001986015b04850361230d565b610b70576040805162461bcd60e51b815260206004820152601860248201527f566572696679207369676e6174757265206661696c6564210000000000000000604482015290519081900360640190fd5b60006060610b7d8a61250c565b91509150816001600160601b0319168761014001516001600160601b03191614610be4576040805162461bcd60e51b815260206004820152601360248201527213995e1d109bdbdad95c9cc81a5b1b1959d85b606a1b604482015290519081900360640190fd5b856001600160a01b0316638a8bd17f88606001516040518263ffffffff1660e01b8152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b158015610c3a57600080fd5b505af1158015610c4e573d6000803e3d6000fd5b505050506040513d6020811015610c6457600080fd5b5051610ca15760405162461bcd60e51b815260040180806020018281038252602d815260200180614560602d913960400191505060405180910390fd5b856001600160a01b03166341973cd9610cb9836125eb565b6040518263ffffffff1660e01b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610d05578181015183820152602001610ced565b50505050905090810190601f168015610d325780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b158015610d5157600080fd5b505af1158015610d65573d6000803e3d6000fd5b505050506040513d6020811015610d7b57600080fd5b5051610db85760405162461bcd60e51b815260040180806020018281038252603b8152602001806144af603b913960400191505060405180910390fd5b7fe60d33488cba3977bf65766cd2f8ac9617f64bf3b3198aff6240ce5c7d43b69087606001518c604051808363ffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610e28578181015183820152602001610e10565b50505050905090810190601f168015610e555780820380516001836020036101000a031916815260200191505b50935050505060405180910390a15060019a9950505050505050505050565b60008054600160a01b900460ff1615610ec7576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b60015460408051631a75201d60e21b815290516001600160a01b039092169182916369d48074916004808301926000929190829003018186803b158015610f0d57600080fd5b505afa158015610f21573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015610f4a57600080fd5b810190808051600160201b811115610f6157600080fd5b82016020810184811115610f7457600080fd5b8151600160201b811182820187101715610f8d57600080fd5b505051159250610fd19150505760405162461bcd60e51b81526004018080602001828103825260388152602001806146456038913960400191505060405180910390fd5b610fd9614229565b610fe28561213b565b905060006060610ff18661250c565b91509150816001600160601b0319168361014001516001600160601b03191614611058576040805162461bcd60e51b815260206004820152601360248201527213995e1d109bdbdad95c9cc81a5b1b1959d85b606a1b604482015290519081900360640190fd5b836001600160a01b0316638a8bd17f84606001516040518263ffffffff1660e01b8152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b1580156110ae57600080fd5b505af11580156110c2573d6000803e3d6000fd5b505050506040513d60208110156110d857600080fd5b50516111155760405162461bcd60e51b81526004018080602001828103825260438152602001806143af6043913960600191505060405180910390fd5b836001600160a01b03166341973cd961112d836125eb565b6040518263ffffffff1660e01b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611179578181015183820152602001611161565b50505050905090810190601f1680156111a65780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b1580156111c557600080fd5b505af11580156111d9573d6000803e3d6000fd5b505050506040513d60208110156111ef57600080fd5b505161122c5760405162461bcd60e51b815260040180806020018281038252604381526020018061467d6043913960600191505060405180910390fd5b7ff01968fc3a2655cf1b5144cb32de6dc898f91b9239c103744e8457152ab2fbde836060015188604051808363ffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561129c578181015183820152602001611284565b50505050905090810190601f1680156112c95780820380516001836020036101000a031916815260200191505b50935050505060405180910390a15060019695505050505050565b60006112ee611860565b61132d576040805162461bcd60e51b815260206004820181905260248201526000805160206144ea833981519152604482015290519081900360640190fd5b611335611463565b15611342576113426126f2565b60015460408051635c975abb60e01b815290516001600160a01b03909216918291635c975abb916004808301926020929190829003018186803b15801561138857600080fd5b505afa15801561139c573d6000803e3d6000fd5b505050506040513d60208110156113b257600080fd5b50511561145b57806001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156113f457600080fd5b505af1158015611408573d6000803e3d6000fd5b505050506040513d602081101561141e57600080fd5b505161145b5760405162461bcd60e51b815260040180806020018281038252602981526020018061433d6029913960400191505060405180910390fd5b600191505090565b600054600160a01b900460ff1690565b60008054600160a01b900460ff166114c9576040805162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b604482015290519081900360640190fd5b6114d1611860565b611510576040805162461bcd60e51b815260206004820181905260248201526000805160206144ea833981519152604482015290519081900360640190fd5b506001805467ffffffffffffffff60a01b1916600160a01b6001600160401b038416021781555b919050565b611544611860565b611583576040805162461bcd60e51b815260206004820181905260248201526000805160206144ea833981519152604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b60008054600160a01b900460ff16611623576040805162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b604482015290519081900360640190fd5b61162b611860565b61166a576040805162461bcd60e51b815260206004820181905260248201526000805160206144ea833981519152604482015290519081900360640190fd5b6001546040805163f2fde38b60e01b81526001600160a01b03858116600483015291519190921691829163f2fde38b9160248082019260009290919082900301818387803b1580156116bb57600080fd5b505af11580156116cf573d6000803e3d6000fd5b5060019695505050505050565b60006116e6611860565b611725576040805162461bcd60e51b815260206004820181905260248201526000805160206144ea833981519152604482015290519081900360640190fd5b61172d611463565b6117395761173961279a565b60015460408051635c975abb60e01b815290516001600160a01b03909216918291635c975abb916004808301926020929190829003018186803b15801561177f57600080fd5b505afa158015611793573d6000803e3d6000fd5b505050506040513d60208110156117a957600080fd5b505161145b57806001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156117ea57600080fd5b505af11580156117fe573d6000803e3d6000fd5b505050506040513d602081101561181457600080fd5b505161145b5760405162461bcd60e51b81526004018080602001828103825260278152602001806142f16027913960400191505060405180910390fd5b6000546001600160a01b031690565b600080546001600160a01b0316611875612824565b6001600160a01b031614905090565b600154600160a01b90046001600160401b031681565b60008054600160a01b900460ff16156118ed576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b60015460408051600162c2db5f60e01b0319815290516001600160a01b0390921691600091839163ff3d24a191600480820192602092909190829003018186803b15801561193a57600080fd5b505afa15801561194e573d6000803e3d6000fd5b505050506040513d602081101561196457600080fd5b50519050606061197382612828565b90506060611980826128a1565b6040516020018082805190602001908083835b602083106119b25780518252601f199092019160209182019101611993565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405160208183030381529060405290506001945050505050979650505050505050565b60008054600160a01b900460ff1615611a4e576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b611a56614229565b611a5f8661213b565b90506000600160009054906101000a90046001600160a01b031690506060611ab9826001600160a01b03166369d480746040518163ffffffff1660e01b815260040160006040518083038186803b158015610a7757600080fd5b90506000826001600160a01b0316635ac407906040518163ffffffff1660e01b815260040160206040518083038186803b158015611af657600080fd5b505afa158015611b0a573d6000803e3d6000fd5b505050506040513d6020811015611b2057600080fd5b50518251606086015163ffffffff92831693509091168211611b8e57611b4e8a888560036000198601610b17565b611b895760405162461bcd60e51b815260040180806020018281038252602a815260200180614774602a913960400191505060405180910390fd5b611c68565b611ba088888560036000198601610b17565b611bdb5760405162461bcd60e51b815260040180806020018281038252603881526020018061458d6038913960400191505060405180910390fd5b611be3614229565b611bec8961213b565b90506060611bff8b836101000151612967565b9050611c0a81612a8d565b611c138d612aed565b14611c65576040805162461bcd60e51b815260206004820152601b60248201527f766572696679206865616465722070726f6f66206661696c6564210000000000604482015290519081900360640190fd5b50505b6060611c788c8760e00151612967565b9050611c82614284565b611c8b82612c1e565b9050856001600160a01b0316630586763c8260200151611cae8460000151612a8d565b6040518363ffffffff1660e01b815260040180836001600160401b03166001600160401b031681526020018281526020019250505060206040518083038186803b158015611cfb57600080fd5b505afa158015611d0f573d6000803e3d6000fd5b505050506040513d6020811015611d2557600080fd5b505115611d635760405162461bcd60e51b81526004018080602001828103825260228152602001806146e26022913960400191505060405180910390fd5b856001600160a01b031663e90bfdcf8260200151611d848460000151612a8d565b6040518363ffffffff1660e01b815260040180836001600160401b03166001600160401b0316815260200182815260200192505050602060405180830381600087803b158015611dd357600080fd5b505af1158015611de7573d6000803e3d6000fd5b505050506040513d6020811015611dfd57600080fd5b5051611e50576040805162461bcd60e51b815260206004820181905260248201527f536176652063726f7373636861696e207478206578697374206661696c656421604482015290519081900360640190fd5b6001546040820151606001516001600160401b03908116600160a01b9092041614611eac5760405162461bcd60e51b815260040180806020018281038252602a815260200180614704602a913960400191505060405180910390fd5b6000611ebf826040015160800151612cfa565b9050611eea81836040015160a00151846040015160c001518560400151604001518660200151612d44565b611f3b576040805162461bcd60e51b815260206004820152601d60248201527f457865637574652043726f7373436861696e205478206661696c656421000000604482015290519081900360640190fd5b7f8a4a2663ce60ce4955c595da2894de0415240f1ace024cfbff85f513b656bdae8260200151836040015160800151846000015185604001516000015160405180856001600160401b03166001600160401b03168152602001806020018060200180602001848103845287818151815260200191508051906020019080838360005b83811015611fd5578181015183820152602001611fbd565b50505050905090810190601f1680156120025780820380516001836020036101000a031916815260200191505b50848103835286518152865160209182019188019080838360005b8381101561203557818101518382015260200161201d565b50505050905090810190601f1680156120625780820380516001836020036101000a031916815260200191505b50848103825285518152855160209182019187019080838360005b8381101561209557818101518382015260200161207d565b50505050905090810190601f1680156120c25780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a15060019d9c50505050505050505050505050565b6120f0611860565b61212f576040805162461bcd60e51b815260206004820181905260248201526000805160206144ea833981519152604482015290519081900360640190fd5b61213881613118565b50565b612143614229565b61214b614229565b600061215784826131b8565b63ffffffff9091168352905061216d8482613260565b6001600160401b03909116602084015290506121898482613306565b60a0840191909152905061219d8482613306565b60c084019190915290506121b18482613306565b60e084019190915290506121c58482613306565b61010084019190915290506121da84826131b8565b63ffffffff909116604084015290506121f384826131b8565b63ffffffff9091166060840152905061220c8482613260565b6001600160401b03909116608084015290506122288482613380565b610120840191909152905061223d8482613453565b506bffffffffffffffffffffffff19166101408301525092915050565b60606000806122698482613260565b80935081925050506060816001600160401b03166040519080825280602002602001820160405280156122a6578160200160208202803883390190505b509050606060005b836001600160401b0316811015612302576122c98786613380565b955091506122d682612cfa565b8382815181106122e257fe5b6001600160a01b03909216602092830291909101909101526001016122ae565b509095945050505050565b60008061231986612aed565b90506000604186518161232857fe5b049050606081604051908082528060200260200182016040528015612357578160200160208202803883390190505b50905060008080805b858110156124f15761237f61237a8c6041840260206134bb565b612a8d565b935061239661237a8c6041840260200160206134bb565b92508a60418202604001815181106123aa57fe5b602001015160f81c60f81b60f81c601b0191506001600288604051602001808281526020019150506040516020818303038152906040526040518082805190602001908083835b602083106124105780518252601f1990920191602091820191016123f1565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa15801561244f573d6000803e3d6000fd5b5050506040513d602081101561246457600080fd5b5051604080516000815260208181018084529390935260ff8616818301526060810188905260808101879052905160a08083019392601f198301929081900390910190855afa1580156124bb573d6000803e3d6000fd5b505050602060405103518582815181106124d157fe5b6001600160a01b0390921660209283029190910190910152600101612360565b506124fd89858a61353b565b9b9a5050505050505050505050565b60006060604383518161251b57fe5b061561256e576040805162461bcd60e51b815260206004820152601b60248201527f5f7075624b65794c697374206c656e67746820696c6c6567616c210000000000604482015290519081900360640190fd5b6000604384518161257b57fe5b04905060018110156125cd576040805162461bcd60e51b8152602060048201526016602482015275746f6f2073686f7274205f7075624b65794c6973742160501b604482015290519081900360640190fd5b6125e18160036000198201048303866135dc565b9250925050915091565b8051606090816125fa82613929565b905060005b828110156126ea578161262c61262787848151811061261a57fe5b602002602001015161396c565b6128a1565b6040516020018083805190602001908083835b6020831061265e5780518252601f19909201916020918201910161263f565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106126a65780518252601f199092019160209182019101612687565b6001836020036101000a03801982511681845116808217855250505050505090500192505050604051602081830303815290604052915080806001019150506125ff565b509392505050565b600054600160a01b900460ff16612747576040805162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b604482015290519081900360640190fd5b6000805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa61277d612824565b604080516001600160a01b039092168252519081900360200190a1565b600054600160a01b900460ff16156127ec576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b6000805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861277d5b3390565b60606001600160ff1b03821115612886576040805162461bcd60e51b815260206004820152601760248201527f56616c75652065786365656473207468652072616e6765000000000000000000604482015290519081900360640190fd5b60405190506020815281602082015260408101604052919050565b80516060906128af81613987565b836040516020018083805190602001908083835b602083106128e25780518252601f1990920191602091820191016128c3565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b6020831061292a5780518252601f19909201916020918201910161290b565b6001836020036101000a03801982511681845116808217855250505050505090500192505050604051602081830303815290604052915050919050565b60606000816129768583613380565b92509050600061298582613acd565b905060006021848851038161299657fe5b049050600080805b83811015612a41576129b08a88613b74565b975091506129be8a88613306565b975092506001600160f81b031982166129e2576129db8386613bf2565b9450612a39565b600160f81b6001600160f81b031983161415612a02576129db8584613bf2565b60405162461bcd60e51b815260040180806020018281038252602e81526020018061450a602e913960400191505060405180910390fd5b60010161299e565b50878414612a805760405162461bcd60e51b81526004018080602001828103825260318152602001806144156031913960400191505060405180910390fd5b5092979650505050505050565b60008151602014612ae5576040805162461bcd60e51b815260206004820152601760248201527f6279746573206c656e677468206973206e6f742033322e000000000000000000604482015290519081900360640190fd5b506020015190565b6000600280836040518082805190602001908083835b60208310612b225780518252601f199092019160209182019101612b03565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015612b61573d6000803e3d6000fd5b5050506040513d6020811015612b7657600080fd5b50516040805160208181019390935281518082038401815290820191829052805190928291908401908083835b60208310612bc25780518252601f199092019160209182019101612ba3565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015612c01573d6000803e3d6000fd5b5050506040513d6020811015612c1657600080fd5b505192915050565b612c26614284565b612c2e614284565b6000612c3a8482613380565b9083529050612c498482613260565b6001600160401b0390911660208401529050612c636142aa565b612c6d8583613380565b9082529150612c7c8583613380565b60208301919091529150612c908583613380565b60408301919091529150612ca48583613260565b6001600160401b0390911660608301529150612cc08583613380565b60808301919091529150612cd48583613380565b60a08301919091529150612ce88583613380565b5060c082015260408301525092915050565b60008151601414612d3c5760405162461bcd60e51b81526004018080602001828103825260238152602001806143666023913960400191505060405180910390fd5b506014015190565b6000612d4f86613cb6565b612d8a5760405162461bcd60e51b81526004018080602001828103825260288152602001806145386028913960400191505060405180910390fd5b60606000876001600160a01b0316876040516020018082805190602001908083835b60208310612dcb5780518252601f199092019160209182019101612dac565b51815160001960209485036101000a01908116901991909116179052732862797465732c62797465732c75696e7436342960601b9390910192835260408051600b19818603018152601485019091528051908201206001600160401b038b1660748501526060603485019081528d5160948601528d519196508d95508c948c945090928392605483019260b4019188019080838360005b83811015612e7a578181015183820152602001612e62565b50505050905090810190601f168015612ea75780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b83811015612eda578181015183820152602001612ec2565b50505050905090810190601f168015612f075780820380516001836020036101000a031916815260200191505b509550505050505060405160208183030381529060405260405160200180836001600160e01b0319166001600160e01b031916815260040182805190602001908083835b60208310612f6a5780518252601f199092019160209182019101612f4b565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040526040518082805190602001908083835b60208310612fce5780518252601f199092019160209182019101612faf565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114613030576040519150601f19603f3d011682016040523d82523d6000602084013e613035565b606091505b509250905060018115151461307b5760405162461bcd60e51b815260040180806020018281038252602b815260200180614484602b913960400191505060405180910390fd5b81516130b85760405162461bcd60e51b81526004018080602001828103825260278152602001806145c56027913960400191505060405180910390fd5b60006130c583601f613cf2565b5090506001811515146131095760405162461bcd60e51b81526004018080602001828103825260378152602001806145ec6037913960400191505060405180910390fd5b50600198975050505050505050565b6001600160a01b03811661315d5760405162461bcd60e51b81526004018080602001828103825260268152602001806143896026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b600080835183600401111580156131d157508260040183105b61320c5760405162461bcd60e51b815260040180806020018281038252602281526020018061472e6022913960400191505060405180910390fd5b600060405160046000600182038760208a0101515b838310156132415780821a83860153600183019250600182039150613221565b50505080820160405260200390035192505050600482015b9250929050565b6000808351836008011115801561327957508260080183105b6132b45760405162461bcd60e51b81526004018080602001828103825260228152602001806146c06022913960400191505060405180910390fd5b600060405160086000600182038760208a0101515b838310156132e95780821a838601536001830192506001820391506132c9565b505050808201604052602003900351956008949094019450505050565b6000808351836020011115801561331f57508260200183105b613370576040805162461bcd60e51b815260206004820181905260248201527f4e657874486173682c206f66667365742065786365656473206d6178696d756d604482015290519081900360640190fd5b5050602091810182015192910190565b606060008061338f8585613de3565b8651909550909150818501118015906133a9575080840184105b6133e45760405162461bcd60e51b81526004018080602001828103825260248152602001806147506024913960400191505060405180910390fd5b6060811580156133ff57604051915060208201604052613449565b6040519150601f8316801560200281840101848101888315602002848c0101015b81831015613438578051835260209283019201613420565b5050848452601f01601f1916604052505b5095930193505050565b6000808351836014011115801561346c57508260140183105b6134a75760405162461bcd60e51b81526004018080602001828103825260238152602001806143f26023913960400191505060405180910390fd5b505081810160200151601482019250929050565b6060818301845110156134cd57600080fd5b6060821580156134e857604051915060208201604052613532565b6040519150601f8416801560200281840101858101878315602002848b0101015b81831015613521578051835260209283019201613509565b5050858452601f01601f1916604052505b50949350505050565b600080805b84518110156135d05760005b86518110156135c75786818151811061356157fe5b60200260200101516001600160a01b031686838151811061357e57fe5b60200260200101516001600160a01b031614156135bf5782806001019350508681815181106135a957fe5b6020026020010160006001600160a01b03168152505b60010161354c565b50600101613540565b50909111159392505050565b60006060806135ea86613ffc565b9050606086604051908082528060200260200182016040528015613618578160200160208202803883390190505b50905060006060815b8981101561373e57613638886043830260436134bb565b9150846136476126278461403f565b6040516020018083805190602001908083835b602083106136795780518252601f19909201916020918201910161365a565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106136c15780518252601f1990920191602091820191016136a2565b6001836020036101000a03801982511681845116808217855250505050505090500192505050604051602081830303815290604052945061370582600360406134bb565b8051906020012092508260001c84828151811061371e57fe5b6001600160a01b0390921660209283029190910190910152600101613621565b508361374989613ffc565b6040516020018083805190602001908083835b6020831061377b5780518252601f19909201916020918201910161375c565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106137c35780518252601f1990920191602091820191016137a4565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040529350600060036002866040518082805190602001908083835b602083106138305780518252601f199092019160209182019101613811565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa15801561386f573d6000803e3d6000fd5b5050506040513d602081101561388457600080fd5b50516040805160208181019390935281518082038401815290820191829052805190928291908401908083835b602083106138d05780518252601f1990920191602091820191016138b1565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa15801561390f573d6000803e3d6000fd5b50506040515160601b9b949a509398505050505050505050565b6040516008808252606091906000601f5b8282101561395c5785811a82602086010153600191909101906000190161393a565b5050506028810160405292915050565b604080516014815260609290921b6020830152818101905290565b606060fd826001600160401b031610156139ab576139a48261412d565b9050611537565b61ffff826001600160401b031611613a89576139ca60fd60f81b614149565b6139d383613ffc565b6040516020018083805190602001908083835b60208310613a055780518252601f1990920191602091820191016139e6565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b60208310613a4d5780518252601f199092019160209182019101613a2e565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040529050611537565b63ffffffff826001600160401b031611613ab357613aaa607f60f91b614149565b6139d38361415d565b613ac46001600160f81b0319614149565b6139d383613929565b6040516000602080830182815284519293600293859387939260210191908401908083835b60208310613b115780518252601f199092019160209182019101613af2565b6001836020036101000a0380198251168184511680821785525050505050509050019250505060405160208183030381529060405260405180828051906020019080838360208310612bc25780518252601f199092019160209182019101612ba3565b60008083518360010111158015613b8d57508260010183105b613bde576040805162461bcd60e51b815260206004820181905260248201527f4e657874427974652c204f66667365742065786365656473206d6178696d756d604482015290519081900360640190fd5b505081810160200151600182019250929050565b60408051600160f81b60208083019190915260218201859052604180830185905283518084039091018152606190920192839052815160009360029392909182918401908083835b60208310613c595780518252601f199092019160209182019101613c3a565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015613c98573d6000803e3d6000fd5b5050506040513d6020811015613cad57600080fd5b50519392505050565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a4708115801590613cea5750808214155b949350505050565b60008083518360010111158015613d0b57508260010183105b613d53576040805162461bcd60e51b815260206004820152601460248201527313d9999cd95d08195e18d959591cc81b1a5b5a5d60621b604482015290519081900360640190fd5b838301602001516000600160f81b6001600160f81b031983161415613d7a57506001613dd5565b6001600160f81b03198216613d9157506000613dd5565b6040805162461bcd60e51b81526020600482015260146024820152732732bc3a2137b7b6103b30b63ab29032b93937b960611b604482015290519081900360640190fd5b956001949094019450505050565b6000806000613df28585613b74565b94509050600060fd60f81b6001600160f81b031983161415613e9057613e1886866141a0565b955061ffff16905060fd8110801590613e33575061ffff8111155b613e84576040805162461bcd60e51b815260206004820152601f60248201527f4e65787455696e7431362c2076616c7565206f7574736964652072616e676500604482015290519081900360640190fd5b92508391506132599050565b607f60f91b6001600160f81b031983161415613f2057613eb086866131b8565b955063ffffffff16905061ffff81118015613ecf575063ffffffff8111155b613e84576040805162461bcd60e51b815260206004820181905260248201527f4e65787456617255696e742c2076616c7565206f7574736964652072616e6765604482015290519081900360640190fd5b6001600160f81b03198083161415613fa157613f3c8686613260565b95506001600160401b0316905063ffffffff8111613e84576040805162461bcd60e51b815260206004820181905260248201527f4e65787456617255696e742c2076616c7565206f7574736964652072616e6765604482015290519081900360640190fd5b5060f881901c60fd8110613e84576040805162461bcd60e51b815260206004820181905260248201527f4e65787456617255696e742c2076616c7565206f7574736964652072616e6765604482015290519081900360640190fd5b6040516002808252606091906000601f5b8282101561402f5785811a82602086010153600191909101906000190161400d565b5050506022810160405292915050565b6060604382511015614098576040805162461bcd60e51b815260206004820152601760248201527f6b6579206c656e67676820697320746f6f2073686f7274000000000000000000604482015290519081900360640190fd5b6140a582600060236134bb565b90506002826042815181106140b657fe5b016020015160f81c816140c557fe5b0660ff16600014156140ff57600260f81b816002815181106140e357fe5b60200101906001600160f81b031916908160001a905350611537565b600360f81b8160028151811061411157fe5b60200101906001600160f81b031916908160001a905350919050565b604080516001815260f89290921b602083015260218201905290565b60606141578260f81c61412d565b92915050565b6040516004808252606091906000601f5b828210156141905785811a82602086010153600191909101906000190161416e565b5050506024810160405292915050565b600080835183600201111580156141b957508260020183105b6141f45760405162461bcd60e51b81526004018080602001828103825260228152602001806146236022913960400191505060405180910390fd5b6000604051846020870101518060011a82538060001a6001830153506002818101604052601d19909101519694019450505050565b604080516101608101825260008082526020820181905291810182905260608082018390526080820183905260a0820183905260c0820183905260e08201839052610100820183905261012082015261014081019190915290565b604080516101208101825260608152600060208201529081016142a56142aa565b905290565b6040518060e0016040528060608152602001606081526020016060815260200160006001600160401b03168152602001606081526020016060815260200160608152509056fe70617573652045746843726f7373436861696e4461746120636f6e7472616374206661696c6564546865206e657874426f6f6b4b6565706572206f662068656164657220697320656d707479756e70617573652045746843726f7373436861696e4461746120636f6e7472616374206661696c65646279746573206c656e67746820646f6573206e6f74206d6174636820616464726573734f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573735361766520506f6c7920636861696e2063757272656e742065706f63682073746172742068656967687420746f204461746120636f6e7472616374206661696c6564214e657874427974657332302c206f66667365742065786365656473206d6178696d756d6d65726b6c6550726f76652c2065787065637420726f6f74206973206e6f7420657175616c2061637475616c20726f6f7454686520686569676874206f6620686561646572206973206c6f776572207468616e2063757272656e742065706f6368207374617274206865696768742145746843726f7373436861696e2063616c6c20627573696e65737320636f6e7472616374206661696c65645361766520506f6c7920636861696e20626f6f6b206b65657065727320627974657320746f204461746120636f6e7472616374206661696c6564214f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726d65726b6c6550726f76652c204e6578744279746520666f7220706f736974696f6e20696e666f206661696c65645468652070617373656420696e2061646472657373206973206e6f74206120636f6e74726163742153617665204d43204c617465737448656967687420746f204461746120636f6e7472616374206661696c65642156657269667920706f6c7920636861696e2063757272656e742065706f636820686561646572207369676e6174757265206661696c6564214e6f2072657475726e2076616c75652066726f6d20627573696e65737320636f6e74726163742145746843726f7373436861696e2063616c6c20627573696e65737320636f6e74726163742072657475726e206973206e6f7420747275654e65787455696e7431362c206f66667365742065786365656473206d6178696d756d45746843726f7373436861696e4461746120636f6e74726163742068617320616c7265616479206265656e20696e697469616c697a6564215361766520506f6c7920636861696e2063757272656e742065706f636820626f6f6b206b65657065727320746f204461746120636f6e7472616374206661696c6564214e65787455696e7436342c206f66667365742065786365656473206d6178696d756d746865207472616e73616374696f6e20686173206265656e2065786563757465642154686973205478206973206e6f742061696d696e6720617420457468657265756d206e6574776f726b214e65787455696e7433322c206f66667365742065786365656473206d6178696d756d4e65787456617242797465732c206f66667365742065786365656473206d6178696d756d56657269667920706f6c7920636861696e20686561646572207369676e6174757265206661696c656421a265627a7a72305820f2a631e794c8db93904015a5d4fa919d5e7dc6da1ed2ad66edbfb4989f377b0864736f6c634300050a0032"

// DeployEthCrossChainManager deploys a new Ethereum contract, binding an instance of EthCrossChainManager to it.
func DeployEthCrossChainManager(auth *bind.TransactOpts, backend bind.ContractBackend, _eccd common.Address, _chainId uint64) (common.Address, *types.Transaction, *EthCrossChainManager, error) {
	parsed, err := abi.JSON(strings.NewReader(EthCrossChainManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EthCrossChainManagerBin), backend, _eccd, _chainId)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EthCrossChainManager{EthCrossChainManagerCaller: EthCrossChainManagerCaller{contract: contract}, EthCrossChainManagerTransactor: EthCrossChainManagerTransactor{contract: contract}, EthCrossChainManagerFilterer: EthCrossChainManagerFilterer{contract: contract}}, nil
}

// EthCrossChainManager is an auto generated Go binding around an Ethereum contract.
type EthCrossChainManager struct {
	EthCrossChainManagerCaller     // Read-only binding to the contract
	EthCrossChainManagerTransactor // Write-only binding to the contract
	EthCrossChainManagerFilterer   // Log filterer for contract events
}

// EthCrossChainManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthCrossChainManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthCrossChainManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthCrossChainManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthCrossChainManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthCrossChainManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthCrossChainManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthCrossChainManagerSession struct {
	Contract     *EthCrossChainManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// EthCrossChainManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthCrossChainManagerCallerSession struct {
	Contract *EthCrossChainManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// EthCrossChainManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthCrossChainManagerTransactorSession struct {
	Contract     *EthCrossChainManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// EthCrossChainManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthCrossChainManagerRaw struct {
	Contract *EthCrossChainManager // Generic contract binding to access the raw methods on
}

// EthCrossChainManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthCrossChainManagerCallerRaw struct {
	Contract *EthCrossChainManagerCaller // Generic read-only contract binding to access the raw methods on
}

// EthCrossChainManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthCrossChainManagerTransactorRaw struct {
	Contract *EthCrossChainManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthCrossChainManager creates a new instance of EthCrossChainManager, bound to a specific deployed contract.
func NewEthCrossChainManager(address common.Address, backend bind.ContractBackend) (*EthCrossChainManager, error) {
	contract, err := bindEthCrossChainManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthCrossChainManager{EthCrossChainManagerCaller: EthCrossChainManagerCaller{contract: contract}, EthCrossChainManagerTransactor: EthCrossChainManagerTransactor{contract: contract}, EthCrossChainManagerFilterer: EthCrossChainManagerFilterer{contract: contract}}, nil
}

// NewEthCrossChainManagerCaller creates a new read-only instance of EthCrossChainManager, bound to a specific deployed contract.
func NewEthCrossChainManagerCaller(address common.Address, caller bind.ContractCaller) (*EthCrossChainManagerCaller, error) {
	contract, err := bindEthCrossChainManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthCrossChainManagerCaller{contract: contract}, nil
}

// NewEthCrossChainManagerTransactor creates a new write-only instance of EthCrossChainManager, bound to a specific deployed contract.
func NewEthCrossChainManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*EthCrossChainManagerTransactor, error) {
	contract, err := bindEthCrossChainManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthCrossChainManagerTransactor{contract: contract}, nil
}

// NewEthCrossChainManagerFilterer creates a new log filterer instance of EthCrossChainManager, bound to a specific deployed contract.
func NewEthCrossChainManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*EthCrossChainManagerFilterer, error) {
	contract, err := bindEthCrossChainManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthCrossChainManagerFilterer{contract: contract}, nil
}

// bindEthCrossChainManager binds a generic wrapper to an already deployed contract.
func bindEthCrossChainManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthCrossChainManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthCrossChainManager *EthCrossChainManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EthCrossChainManager.Contract.EthCrossChainManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthCrossChainManager *EthCrossChainManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.EthCrossChainManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthCrossChainManager *EthCrossChainManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.EthCrossChainManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthCrossChainManager *EthCrossChainManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EthCrossChainManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthCrossChainManager *EthCrossChainManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthCrossChainManager *EthCrossChainManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.contract.Transact(opts, method, params...)
}

// EthCrossChainDataAddress is a free data retrieval call binding the contract method 0x00ba1694.
//
// Solidity: function EthCrossChainDataAddress() view returns(address)
func (_EthCrossChainManager *EthCrossChainManagerCaller) EthCrossChainDataAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EthCrossChainManager.contract.Call(opts, out, "EthCrossChainDataAddress")
	return *ret0, err
}

// EthCrossChainDataAddress is a free data retrieval call binding the contract method 0x00ba1694.
//
// Solidity: function EthCrossChainDataAddress() view returns(address)
func (_EthCrossChainManager *EthCrossChainManagerSession) EthCrossChainDataAddress() (common.Address, error) {
	return _EthCrossChainManager.Contract.EthCrossChainDataAddress(&_EthCrossChainManager.CallOpts)
}

// EthCrossChainDataAddress is a free data retrieval call binding the contract method 0x00ba1694.
//
// Solidity: function EthCrossChainDataAddress() view returns(address)
func (_EthCrossChainManager *EthCrossChainManagerCallerSession) EthCrossChainDataAddress() (common.Address, error) {
	return _EthCrossChainManager.Contract.EthCrossChainDataAddress(&_EthCrossChainManager.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint64)
func (_EthCrossChainManager *EthCrossChainManagerCaller) ChainId(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _EthCrossChainManager.contract.Call(opts, out, "chainId")
	return *ret0, err
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint64)
func (_EthCrossChainManager *EthCrossChainManagerSession) ChainId() (uint64, error) {
	return _EthCrossChainManager.Contract.ChainId(&_EthCrossChainManager.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint64)
func (_EthCrossChainManager *EthCrossChainManagerCallerSession) ChainId() (uint64, error) {
	return _EthCrossChainManager.Contract.ChainId(&_EthCrossChainManager.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EthCrossChainManager.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerSession) IsOwner() (bool, error) {
	return _EthCrossChainManager.Contract.IsOwner(&_EthCrossChainManager.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerCallerSession) IsOwner() (bool, error) {
	return _EthCrossChainManager.Contract.IsOwner(&_EthCrossChainManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthCrossChainManager *EthCrossChainManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EthCrossChainManager.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthCrossChainManager *EthCrossChainManagerSession) Owner() (common.Address, error) {
	return _EthCrossChainManager.Contract.Owner(&_EthCrossChainManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthCrossChainManager *EthCrossChainManagerCallerSession) Owner() (common.Address, error) {
	return _EthCrossChainManager.Contract.Owner(&_EthCrossChainManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EthCrossChainManager.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerSession) Paused() (bool, error) {
	return _EthCrossChainManager.Contract.Paused(&_EthCrossChainManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerCallerSession) Paused() (bool, error) {
	return _EthCrossChainManager.Contract.Paused(&_EthCrossChainManager.CallOpts)
}

// ChangeBookKeeper is a paid mutator transaction binding the contract method 0x29dcf4ab.
//
// Solidity: function changeBookKeeper(bytes rawHeader, bytes pubKeyList, bytes sigList) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactor) ChangeBookKeeper(opts *bind.TransactOpts, rawHeader []byte, pubKeyList []byte, sigList []byte) (*types.Transaction, error) {
	return _EthCrossChainManager.contract.Transact(opts, "changeBookKeeper", rawHeader, pubKeyList, sigList)
}

// ChangeBookKeeper is a paid mutator transaction binding the contract method 0x29dcf4ab.
//
// Solidity: function changeBookKeeper(bytes rawHeader, bytes pubKeyList, bytes sigList) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerSession) ChangeBookKeeper(rawHeader []byte, pubKeyList []byte, sigList []byte) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.ChangeBookKeeper(&_EthCrossChainManager.TransactOpts, rawHeader, pubKeyList, sigList)
}

// ChangeBookKeeper is a paid mutator transaction binding the contract method 0x29dcf4ab.
//
// Solidity: function changeBookKeeper(bytes rawHeader, bytes pubKeyList, bytes sigList) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactorSession) ChangeBookKeeper(rawHeader []byte, pubKeyList []byte, sigList []byte) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.ChangeBookKeeper(&_EthCrossChainManager.TransactOpts, rawHeader, pubKeyList, sigList)
}

// CrossChain is a paid mutator transaction binding the contract method 0xbd5cf625.
//
// Solidity: function crossChain(uint64 toChainId, bytes toContract, bytes method, bytes txData) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactor) CrossChain(opts *bind.TransactOpts, toChainId uint64, toContract []byte, method []byte, txData []byte) (*types.Transaction, error) {
	return _EthCrossChainManager.contract.Transact(opts, "crossChain", toChainId, toContract, method, txData)
}

// CrossChain is a paid mutator transaction binding the contract method 0xbd5cf625.
//
// Solidity: function crossChain(uint64 toChainId, bytes toContract, bytes method, bytes txData) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerSession) CrossChain(toChainId uint64, toContract []byte, method []byte, txData []byte) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.CrossChain(&_EthCrossChainManager.TransactOpts, toChainId, toContract, method, txData)
}

// CrossChain is a paid mutator transaction binding the contract method 0xbd5cf625.
//
// Solidity: function crossChain(uint64 toChainId, bytes toContract, bytes method, bytes txData) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactorSession) CrossChain(toChainId uint64, toContract []byte, method []byte, txData []byte) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.CrossChain(&_EthCrossChainManager.TransactOpts, toChainId, toContract, method, txData)
}

// InitGenesisBlock is a paid mutator transaction binding the contract method 0x34a773eb.
//
// Solidity: function initGenesisBlock(bytes rawHeader, bytes pubKeyList) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactor) InitGenesisBlock(opts *bind.TransactOpts, rawHeader []byte, pubKeyList []byte) (*types.Transaction, error) {
	return _EthCrossChainManager.contract.Transact(opts, "initGenesisBlock", rawHeader, pubKeyList)
}

// InitGenesisBlock is a paid mutator transaction binding the contract method 0x34a773eb.
//
// Solidity: function initGenesisBlock(bytes rawHeader, bytes pubKeyList) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerSession) InitGenesisBlock(rawHeader []byte, pubKeyList []byte) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.InitGenesisBlock(&_EthCrossChainManager.TransactOpts, rawHeader, pubKeyList)
}

// InitGenesisBlock is a paid mutator transaction binding the contract method 0x34a773eb.
//
// Solidity: function initGenesisBlock(bytes rawHeader, bytes pubKeyList) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactorSession) InitGenesisBlock(rawHeader []byte, pubKeyList []byte) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.InitGenesisBlock(&_EthCrossChainManager.TransactOpts, rawHeader, pubKeyList)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthCrossChainManager.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerSession) Pause() (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.Pause(&_EthCrossChainManager.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactorSession) Pause() (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.Pause(&_EthCrossChainManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EthCrossChainManager *EthCrossChainManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthCrossChainManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EthCrossChainManager *EthCrossChainManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.RenounceOwnership(&_EthCrossChainManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EthCrossChainManager *EthCrossChainManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.RenounceOwnership(&_EthCrossChainManager.TransactOpts)
}

// SetChainId is a paid mutator transaction binding the contract method 0x6f31031d.
//
// Solidity: function setChainId(uint64 _newChainId) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactor) SetChainId(opts *bind.TransactOpts, _newChainId uint64) (*types.Transaction, error) {
	return _EthCrossChainManager.contract.Transact(opts, "setChainId", _newChainId)
}

// SetChainId is a paid mutator transaction binding the contract method 0x6f31031d.
//
// Solidity: function setChainId(uint64 _newChainId) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerSession) SetChainId(_newChainId uint64) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.SetChainId(&_EthCrossChainManager.TransactOpts, _newChainId)
}

// SetChainId is a paid mutator transaction binding the contract method 0x6f31031d.
//
// Solidity: function setChainId(uint64 _newChainId) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactorSession) SetChainId(_newChainId uint64) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.SetChainId(&_EthCrossChainManager.TransactOpts, _newChainId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EthCrossChainManager *EthCrossChainManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EthCrossChainManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EthCrossChainManager *EthCrossChainManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.TransferOwnership(&_EthCrossChainManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EthCrossChainManager *EthCrossChainManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.TransferOwnership(&_EthCrossChainManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthCrossChainManager.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerSession) Unpause() (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.Unpause(&_EthCrossChainManager.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactorSession) Unpause() (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.Unpause(&_EthCrossChainManager.TransactOpts)
}

// UpgradeToNew is a paid mutator transaction binding the contract method 0x7e724ff3.
//
// Solidity: function upgradeToNew(address newEthCrossChainManagerAddress) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactor) UpgradeToNew(opts *bind.TransactOpts, newEthCrossChainManagerAddress common.Address) (*types.Transaction, error) {
	return _EthCrossChainManager.contract.Transact(opts, "upgradeToNew", newEthCrossChainManagerAddress)
}

// UpgradeToNew is a paid mutator transaction binding the contract method 0x7e724ff3.
//
// Solidity: function upgradeToNew(address newEthCrossChainManagerAddress) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerSession) UpgradeToNew(newEthCrossChainManagerAddress common.Address) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.UpgradeToNew(&_EthCrossChainManager.TransactOpts, newEthCrossChainManagerAddress)
}

// UpgradeToNew is a paid mutator transaction binding the contract method 0x7e724ff3.
//
// Solidity: function upgradeToNew(address newEthCrossChainManagerAddress) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactorSession) UpgradeToNew(newEthCrossChainManagerAddress common.Address) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.UpgradeToNew(&_EthCrossChainManager.TransactOpts, newEthCrossChainManagerAddress)
}

// VerifyHeaderAndExecuteTx is a paid mutator transaction binding the contract method 0xd450e04c.
//
// Solidity: function verifyHeaderAndExecuteTx(bytes proof, bytes rawHeader, bytes headerProof, bytes curRawHeader, bytes headerSig) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactor) VerifyHeaderAndExecuteTx(opts *bind.TransactOpts, proof []byte, rawHeader []byte, headerProof []byte, curRawHeader []byte, headerSig []byte) (*types.Transaction, error) {
	return _EthCrossChainManager.contract.Transact(opts, "verifyHeaderAndExecuteTx", proof, rawHeader, headerProof, curRawHeader, headerSig)
}

// VerifyHeaderAndExecuteTx is a paid mutator transaction binding the contract method 0xd450e04c.
//
// Solidity: function verifyHeaderAndExecuteTx(bytes proof, bytes rawHeader, bytes headerProof, bytes curRawHeader, bytes headerSig) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerSession) VerifyHeaderAndExecuteTx(proof []byte, rawHeader []byte, headerProof []byte, curRawHeader []byte, headerSig []byte) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.VerifyHeaderAndExecuteTx(&_EthCrossChainManager.TransactOpts, proof, rawHeader, headerProof, curRawHeader, headerSig)
}

// VerifyHeaderAndExecuteTx is a paid mutator transaction binding the contract method 0xd450e04c.
//
// Solidity: function verifyHeaderAndExecuteTx(bytes proof, bytes rawHeader, bytes headerProof, bytes curRawHeader, bytes headerSig) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactorSession) VerifyHeaderAndExecuteTx(proof []byte, rawHeader []byte, headerProof []byte, curRawHeader []byte, headerSig []byte) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.VerifyHeaderAndExecuteTx(&_EthCrossChainManager.TransactOpts, proof, rawHeader, headerProof, curRawHeader, headerSig)
}

// EthCrossChainManagerChangeBookKeeperEventIterator is returned from FilterChangeBookKeeperEvent and is used to iterate over the raw logs and unpacked data for ChangeBookKeeperEvent events raised by the EthCrossChainManager contract.
type EthCrossChainManagerChangeBookKeeperEventIterator struct {
	Event *EthCrossChainManagerChangeBookKeeperEvent // Event containing the contract specifics and raw log

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
func (it *EthCrossChainManagerChangeBookKeeperEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthCrossChainManagerChangeBookKeeperEvent)
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
		it.Event = new(EthCrossChainManagerChangeBookKeeperEvent)
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
func (it *EthCrossChainManagerChangeBookKeeperEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthCrossChainManagerChangeBookKeeperEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthCrossChainManagerChangeBookKeeperEvent represents a ChangeBookKeeperEvent event raised by the EthCrossChainManager contract.
type EthCrossChainManagerChangeBookKeeperEvent struct {
	Height    *big.Int
	RawHeader []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChangeBookKeeperEvent is a free log retrieval operation binding the contract event 0xe60d33488cba3977bf65766cd2f8ac9617f64bf3b3198aff6240ce5c7d43b690.
//
// Solidity: event ChangeBookKeeperEvent(uint256 height, bytes rawHeader)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) FilterChangeBookKeeperEvent(opts *bind.FilterOpts) (*EthCrossChainManagerChangeBookKeeperEventIterator, error) {

	logs, sub, err := _EthCrossChainManager.contract.FilterLogs(opts, "ChangeBookKeeperEvent")
	if err != nil {
		return nil, err
	}
	return &EthCrossChainManagerChangeBookKeeperEventIterator{contract: _EthCrossChainManager.contract, event: "ChangeBookKeeperEvent", logs: logs, sub: sub}, nil
}

// WatchChangeBookKeeperEvent is a free log subscription operation binding the contract event 0xe60d33488cba3977bf65766cd2f8ac9617f64bf3b3198aff6240ce5c7d43b690.
//
// Solidity: event ChangeBookKeeperEvent(uint256 height, bytes rawHeader)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) WatchChangeBookKeeperEvent(opts *bind.WatchOpts, sink chan<- *EthCrossChainManagerChangeBookKeeperEvent) (event.Subscription, error) {

	logs, sub, err := _EthCrossChainManager.contract.WatchLogs(opts, "ChangeBookKeeperEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthCrossChainManagerChangeBookKeeperEvent)
				if err := _EthCrossChainManager.contract.UnpackLog(event, "ChangeBookKeeperEvent", log); err != nil {
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

// ParseChangeBookKeeperEvent is a log parse operation binding the contract event 0xe60d33488cba3977bf65766cd2f8ac9617f64bf3b3198aff6240ce5c7d43b690.
//
// Solidity: event ChangeBookKeeperEvent(uint256 height, bytes rawHeader)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) ParseChangeBookKeeperEvent(log types.Log) (*EthCrossChainManagerChangeBookKeeperEvent, error) {
	event := new(EthCrossChainManagerChangeBookKeeperEvent)
	if err := _EthCrossChainManager.contract.UnpackLog(event, "ChangeBookKeeperEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EthCrossChainManagerCrossChainEventIterator is returned from FilterCrossChainEvent and is used to iterate over the raw logs and unpacked data for CrossChainEvent events raised by the EthCrossChainManager contract.
type EthCrossChainManagerCrossChainEventIterator struct {
	Event *EthCrossChainManagerCrossChainEvent // Event containing the contract specifics and raw log

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
func (it *EthCrossChainManagerCrossChainEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthCrossChainManagerCrossChainEvent)
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
		it.Event = new(EthCrossChainManagerCrossChainEvent)
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
func (it *EthCrossChainManagerCrossChainEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthCrossChainManagerCrossChainEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthCrossChainManagerCrossChainEvent represents a CrossChainEvent event raised by the EthCrossChainManager contract.
type EthCrossChainManagerCrossChainEvent struct {
	Sender               common.Address
	TxId                 []byte
	ProxyOrAssetContract common.Address
	ToChainId            uint64
	ToContract           []byte
	Rawdata              []byte
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterCrossChainEvent is a free log retrieval operation binding the contract event 0x6ad3bf15c1988bc04bc153490cab16db8efb9a3990215bf1c64ea6e28be88483.
//
// Solidity: event CrossChainEvent(address indexed sender, bytes txId, address proxyOrAssetContract, uint64 toChainId, bytes toContract, bytes rawdata)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) FilterCrossChainEvent(opts *bind.FilterOpts, sender []common.Address) (*EthCrossChainManagerCrossChainEventIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EthCrossChainManager.contract.FilterLogs(opts, "CrossChainEvent", senderRule)
	if err != nil {
		return nil, err
	}
	return &EthCrossChainManagerCrossChainEventIterator{contract: _EthCrossChainManager.contract, event: "CrossChainEvent", logs: logs, sub: sub}, nil
}

// WatchCrossChainEvent is a free log subscription operation binding the contract event 0x6ad3bf15c1988bc04bc153490cab16db8efb9a3990215bf1c64ea6e28be88483.
//
// Solidity: event CrossChainEvent(address indexed sender, bytes txId, address proxyOrAssetContract, uint64 toChainId, bytes toContract, bytes rawdata)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) WatchCrossChainEvent(opts *bind.WatchOpts, sink chan<- *EthCrossChainManagerCrossChainEvent, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EthCrossChainManager.contract.WatchLogs(opts, "CrossChainEvent", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthCrossChainManagerCrossChainEvent)
				if err := _EthCrossChainManager.contract.UnpackLog(event, "CrossChainEvent", log); err != nil {
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

// ParseCrossChainEvent is a log parse operation binding the contract event 0x6ad3bf15c1988bc04bc153490cab16db8efb9a3990215bf1c64ea6e28be88483.
//
// Solidity: event CrossChainEvent(address indexed sender, bytes txId, address proxyOrAssetContract, uint64 toChainId, bytes toContract, bytes rawdata)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) ParseCrossChainEvent(log types.Log) (*EthCrossChainManagerCrossChainEvent, error) {
	event := new(EthCrossChainManagerCrossChainEvent)
	if err := _EthCrossChainManager.contract.UnpackLog(event, "CrossChainEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EthCrossChainManagerInitGenesisBlockEventIterator is returned from FilterInitGenesisBlockEvent and is used to iterate over the raw logs and unpacked data for InitGenesisBlockEvent events raised by the EthCrossChainManager contract.
type EthCrossChainManagerInitGenesisBlockEventIterator struct {
	Event *EthCrossChainManagerInitGenesisBlockEvent // Event containing the contract specifics and raw log

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
func (it *EthCrossChainManagerInitGenesisBlockEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthCrossChainManagerInitGenesisBlockEvent)
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
		it.Event = new(EthCrossChainManagerInitGenesisBlockEvent)
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
func (it *EthCrossChainManagerInitGenesisBlockEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthCrossChainManagerInitGenesisBlockEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthCrossChainManagerInitGenesisBlockEvent represents a InitGenesisBlockEvent event raised by the EthCrossChainManager contract.
type EthCrossChainManagerInitGenesisBlockEvent struct {
	Height    *big.Int
	RawHeader []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterInitGenesisBlockEvent is a free log retrieval operation binding the contract event 0xf01968fc3a2655cf1b5144cb32de6dc898f91b9239c103744e8457152ab2fbde.
//
// Solidity: event InitGenesisBlockEvent(uint256 height, bytes rawHeader)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) FilterInitGenesisBlockEvent(opts *bind.FilterOpts) (*EthCrossChainManagerInitGenesisBlockEventIterator, error) {

	logs, sub, err := _EthCrossChainManager.contract.FilterLogs(opts, "InitGenesisBlockEvent")
	if err != nil {
		return nil, err
	}
	return &EthCrossChainManagerInitGenesisBlockEventIterator{contract: _EthCrossChainManager.contract, event: "InitGenesisBlockEvent", logs: logs, sub: sub}, nil
}

// WatchInitGenesisBlockEvent is a free log subscription operation binding the contract event 0xf01968fc3a2655cf1b5144cb32de6dc898f91b9239c103744e8457152ab2fbde.
//
// Solidity: event InitGenesisBlockEvent(uint256 height, bytes rawHeader)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) WatchInitGenesisBlockEvent(opts *bind.WatchOpts, sink chan<- *EthCrossChainManagerInitGenesisBlockEvent) (event.Subscription, error) {

	logs, sub, err := _EthCrossChainManager.contract.WatchLogs(opts, "InitGenesisBlockEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthCrossChainManagerInitGenesisBlockEvent)
				if err := _EthCrossChainManager.contract.UnpackLog(event, "InitGenesisBlockEvent", log); err != nil {
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

// ParseInitGenesisBlockEvent is a log parse operation binding the contract event 0xf01968fc3a2655cf1b5144cb32de6dc898f91b9239c103744e8457152ab2fbde.
//
// Solidity: event InitGenesisBlockEvent(uint256 height, bytes rawHeader)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) ParseInitGenesisBlockEvent(log types.Log) (*EthCrossChainManagerInitGenesisBlockEvent, error) {
	event := new(EthCrossChainManagerInitGenesisBlockEvent)
	if err := _EthCrossChainManager.contract.UnpackLog(event, "InitGenesisBlockEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EthCrossChainManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EthCrossChainManager contract.
type EthCrossChainManagerOwnershipTransferredIterator struct {
	Event *EthCrossChainManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EthCrossChainManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthCrossChainManagerOwnershipTransferred)
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
		it.Event = new(EthCrossChainManagerOwnershipTransferred)
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
func (it *EthCrossChainManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthCrossChainManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthCrossChainManagerOwnershipTransferred represents a OwnershipTransferred event raised by the EthCrossChainManager contract.
type EthCrossChainManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EthCrossChainManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EthCrossChainManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EthCrossChainManagerOwnershipTransferredIterator{contract: _EthCrossChainManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EthCrossChainManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EthCrossChainManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthCrossChainManagerOwnershipTransferred)
				if err := _EthCrossChainManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_EthCrossChainManager *EthCrossChainManagerFilterer) ParseOwnershipTransferred(log types.Log) (*EthCrossChainManagerOwnershipTransferred, error) {
	event := new(EthCrossChainManagerOwnershipTransferred)
	if err := _EthCrossChainManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EthCrossChainManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the EthCrossChainManager contract.
type EthCrossChainManagerPausedIterator struct {
	Event *EthCrossChainManagerPaused // Event containing the contract specifics and raw log

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
func (it *EthCrossChainManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthCrossChainManagerPaused)
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
		it.Event = new(EthCrossChainManagerPaused)
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
func (it *EthCrossChainManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthCrossChainManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthCrossChainManagerPaused represents a Paused event raised by the EthCrossChainManager contract.
type EthCrossChainManagerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) FilterPaused(opts *bind.FilterOpts) (*EthCrossChainManagerPausedIterator, error) {

	logs, sub, err := _EthCrossChainManager.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &EthCrossChainManagerPausedIterator{contract: _EthCrossChainManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EthCrossChainManagerPaused) (event.Subscription, error) {

	logs, sub, err := _EthCrossChainManager.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthCrossChainManagerPaused)
				if err := _EthCrossChainManager.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) ParsePaused(log types.Log) (*EthCrossChainManagerPaused, error) {
	event := new(EthCrossChainManagerPaused)
	if err := _EthCrossChainManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EthCrossChainManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the EthCrossChainManager contract.
type EthCrossChainManagerUnpausedIterator struct {
	Event *EthCrossChainManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *EthCrossChainManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthCrossChainManagerUnpaused)
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
		it.Event = new(EthCrossChainManagerUnpaused)
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
func (it *EthCrossChainManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthCrossChainManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthCrossChainManagerUnpaused represents a Unpaused event raised by the EthCrossChainManager contract.
type EthCrossChainManagerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*EthCrossChainManagerUnpausedIterator, error) {

	logs, sub, err := _EthCrossChainManager.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &EthCrossChainManagerUnpausedIterator{contract: _EthCrossChainManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EthCrossChainManagerUnpaused) (event.Subscription, error) {

	logs, sub, err := _EthCrossChainManager.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthCrossChainManagerUnpaused)
				if err := _EthCrossChainManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) ParseUnpaused(log types.Log) (*EthCrossChainManagerUnpaused, error) {
	event := new(EthCrossChainManagerUnpaused)
	if err := _EthCrossChainManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EthCrossChainManagerVerifyHeaderAndExecuteTxEventIterator is returned from FilterVerifyHeaderAndExecuteTxEvent and is used to iterate over the raw logs and unpacked data for VerifyHeaderAndExecuteTxEvent events raised by the EthCrossChainManager contract.
type EthCrossChainManagerVerifyHeaderAndExecuteTxEventIterator struct {
	Event *EthCrossChainManagerVerifyHeaderAndExecuteTxEvent // Event containing the contract specifics and raw log

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
func (it *EthCrossChainManagerVerifyHeaderAndExecuteTxEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthCrossChainManagerVerifyHeaderAndExecuteTxEvent)
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
		it.Event = new(EthCrossChainManagerVerifyHeaderAndExecuteTxEvent)
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
func (it *EthCrossChainManagerVerifyHeaderAndExecuteTxEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthCrossChainManagerVerifyHeaderAndExecuteTxEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthCrossChainManagerVerifyHeaderAndExecuteTxEvent represents a VerifyHeaderAndExecuteTxEvent event raised by the EthCrossChainManager contract.
type EthCrossChainManagerVerifyHeaderAndExecuteTxEvent struct {
	FromChainID      uint64
	ToContract       []byte
	CrossChainTxHash []byte
	FromChainTxHash  []byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterVerifyHeaderAndExecuteTxEvent is a free log retrieval operation binding the contract event 0x8a4a2663ce60ce4955c595da2894de0415240f1ace024cfbff85f513b656bdae.
//
// Solidity: event VerifyHeaderAndExecuteTxEvent(uint64 fromChainID, bytes toContract, bytes crossChainTxHash, bytes fromChainTxHash)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) FilterVerifyHeaderAndExecuteTxEvent(opts *bind.FilterOpts) (*EthCrossChainManagerVerifyHeaderAndExecuteTxEventIterator, error) {

	logs, sub, err := _EthCrossChainManager.contract.FilterLogs(opts, "VerifyHeaderAndExecuteTxEvent")
	if err != nil {
		return nil, err
	}
	return &EthCrossChainManagerVerifyHeaderAndExecuteTxEventIterator{contract: _EthCrossChainManager.contract, event: "VerifyHeaderAndExecuteTxEvent", logs: logs, sub: sub}, nil
}

// WatchVerifyHeaderAndExecuteTxEvent is a free log subscription operation binding the contract event 0x8a4a2663ce60ce4955c595da2894de0415240f1ace024cfbff85f513b656bdae.
//
// Solidity: event VerifyHeaderAndExecuteTxEvent(uint64 fromChainID, bytes toContract, bytes crossChainTxHash, bytes fromChainTxHash)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) WatchVerifyHeaderAndExecuteTxEvent(opts *bind.WatchOpts, sink chan<- *EthCrossChainManagerVerifyHeaderAndExecuteTxEvent) (event.Subscription, error) {

	logs, sub, err := _EthCrossChainManager.contract.WatchLogs(opts, "VerifyHeaderAndExecuteTxEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthCrossChainManagerVerifyHeaderAndExecuteTxEvent)
				if err := _EthCrossChainManager.contract.UnpackLog(event, "VerifyHeaderAndExecuteTxEvent", log); err != nil {
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

// ParseVerifyHeaderAndExecuteTxEvent is a log parse operation binding the contract event 0x8a4a2663ce60ce4955c595da2894de0415240f1ace024cfbff85f513b656bdae.
//
// Solidity: event VerifyHeaderAndExecuteTxEvent(uint64 fromChainID, bytes toContract, bytes crossChainTxHash, bytes fromChainTxHash)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) ParseVerifyHeaderAndExecuteTxEvent(log types.Log) (*EthCrossChainManagerVerifyHeaderAndExecuteTxEvent, error) {
	event := new(EthCrossChainManagerVerifyHeaderAndExecuteTxEvent)
	if err := _EthCrossChainManager.contract.UnpackLog(event, "VerifyHeaderAndExecuteTxEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IEthCrossChainDataABI is the input ABI used to generate the binding from.
const IEthCrossChainDataABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"fromChainId\",\"type\":\"uint64\"},{\"name\":\"fromChainTx\",\"type\":\"bytes32\"}],\"name\":\"checkIfFromChainTxExist\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"key1\",\"type\":\"bytes32\"},{\"name\":\"key2\",\"type\":\"bytes32\"},{\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"putExtraData\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"ethTxHashIndex\",\"type\":\"uint256\"}],\"name\":\"getEthTxHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key1\",\"type\":\"bytes32\"},{\"name\":\"key2\",\"type\":\"bytes32\"}],\"name\":\"getExtraData\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"curEpochPkBytes\",\"type\":\"bytes\"}],\"name\":\"putCurEpochConPubKeyBytes\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ethTxHash\",\"type\":\"bytes32\"}],\"name\":\"putEthTxHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurEpochStartHeight\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurEpochConPubKeyBytes\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"curEpochStartHeight\",\"type\":\"uint32\"}],\"name\":\"putCurEpochStartHeight\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"fromChainId\",\"type\":\"uint64\"},{\"name\":\"fromChainTx\",\"type\":\"bytes32\"}],\"name\":\"markFromChainTxExist\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEthTxHashIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IEthCrossChainDataFuncSigs maps the 4-byte function signature to its string representation.
var IEthCrossChainDataFuncSigs = map[string]string{
	"0586763c": "checkIfFromChainTxExist(uint64,bytes32)",
	"69d48074": "getCurEpochConPubKeyBytes()",
	"5ac40790": "getCurEpochStartHeight()",
	"29927875": "getEthTxHash(uint256)",
	"ff3d24a1": "getEthTxHashIndex()",
	"40602bb5": "getExtraData(bytes32,bytes32)",
	"e90bfdcf": "markFromChainTxExist(uint64,bytes32)",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"41973cd9": "putCurEpochConPubKeyBytes(bytes)",
	"8a8bd17f": "putCurEpochStartHeight(uint32)",
	"4c3ccf64": "putEthTxHash(bytes32)",
	"1afe374e": "putExtraData(bytes32,bytes32,bytes)",
	"f2fde38b": "transferOwnership(address)",
	"3f4ba83a": "unpause()",
}

// IEthCrossChainData is an auto generated Go binding around an Ethereum contract.
type IEthCrossChainData struct {
	IEthCrossChainDataCaller     // Read-only binding to the contract
	IEthCrossChainDataTransactor // Write-only binding to the contract
	IEthCrossChainDataFilterer   // Log filterer for contract events
}

// IEthCrossChainDataCaller is an auto generated read-only Go binding around an Ethereum contract.
type IEthCrossChainDataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthCrossChainDataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IEthCrossChainDataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthCrossChainDataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IEthCrossChainDataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthCrossChainDataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IEthCrossChainDataSession struct {
	Contract     *IEthCrossChainData // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IEthCrossChainDataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IEthCrossChainDataCallerSession struct {
	Contract *IEthCrossChainDataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IEthCrossChainDataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IEthCrossChainDataTransactorSession struct {
	Contract     *IEthCrossChainDataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IEthCrossChainDataRaw is an auto generated low-level Go binding around an Ethereum contract.
type IEthCrossChainDataRaw struct {
	Contract *IEthCrossChainData // Generic contract binding to access the raw methods on
}

// IEthCrossChainDataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IEthCrossChainDataCallerRaw struct {
	Contract *IEthCrossChainDataCaller // Generic read-only contract binding to access the raw methods on
}

// IEthCrossChainDataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IEthCrossChainDataTransactorRaw struct {
	Contract *IEthCrossChainDataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIEthCrossChainData creates a new instance of IEthCrossChainData, bound to a specific deployed contract.
func NewIEthCrossChainData(address common.Address, backend bind.ContractBackend) (*IEthCrossChainData, error) {
	contract, err := bindIEthCrossChainData(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainData{IEthCrossChainDataCaller: IEthCrossChainDataCaller{contract: contract}, IEthCrossChainDataTransactor: IEthCrossChainDataTransactor{contract: contract}, IEthCrossChainDataFilterer: IEthCrossChainDataFilterer{contract: contract}}, nil
}

// NewIEthCrossChainDataCaller creates a new read-only instance of IEthCrossChainData, bound to a specific deployed contract.
func NewIEthCrossChainDataCaller(address common.Address, caller bind.ContractCaller) (*IEthCrossChainDataCaller, error) {
	contract, err := bindIEthCrossChainData(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainDataCaller{contract: contract}, nil
}

// NewIEthCrossChainDataTransactor creates a new write-only instance of IEthCrossChainData, bound to a specific deployed contract.
func NewIEthCrossChainDataTransactor(address common.Address, transactor bind.ContractTransactor) (*IEthCrossChainDataTransactor, error) {
	contract, err := bindIEthCrossChainData(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainDataTransactor{contract: contract}, nil
}

// NewIEthCrossChainDataFilterer creates a new log filterer instance of IEthCrossChainData, bound to a specific deployed contract.
func NewIEthCrossChainDataFilterer(address common.Address, filterer bind.ContractFilterer) (*IEthCrossChainDataFilterer, error) {
	contract, err := bindIEthCrossChainData(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainDataFilterer{contract: contract}, nil
}

// bindIEthCrossChainData binds a generic wrapper to an already deployed contract.
func bindIEthCrossChainData(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IEthCrossChainDataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEthCrossChainData *IEthCrossChainDataRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IEthCrossChainData.Contract.IEthCrossChainDataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEthCrossChainData *IEthCrossChainDataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.IEthCrossChainDataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEthCrossChainData *IEthCrossChainDataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.IEthCrossChainDataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEthCrossChainData *IEthCrossChainDataCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IEthCrossChainData.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEthCrossChainData *IEthCrossChainDataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEthCrossChainData *IEthCrossChainDataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.contract.Transact(opts, method, params...)
}

// CheckIfFromChainTxExist is a free data retrieval call binding the contract method 0x0586763c.
//
// Solidity: function checkIfFromChainTxExist(uint64 fromChainId, bytes32 fromChainTx) view returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataCaller) CheckIfFromChainTxExist(opts *bind.CallOpts, fromChainId uint64, fromChainTx [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IEthCrossChainData.contract.Call(opts, out, "checkIfFromChainTxExist", fromChainId, fromChainTx)
	return *ret0, err
}

// CheckIfFromChainTxExist is a free data retrieval call binding the contract method 0x0586763c.
//
// Solidity: function checkIfFromChainTxExist(uint64 fromChainId, bytes32 fromChainTx) view returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) CheckIfFromChainTxExist(fromChainId uint64, fromChainTx [32]byte) (bool, error) {
	return _IEthCrossChainData.Contract.CheckIfFromChainTxExist(&_IEthCrossChainData.CallOpts, fromChainId, fromChainTx)
}

// CheckIfFromChainTxExist is a free data retrieval call binding the contract method 0x0586763c.
//
// Solidity: function checkIfFromChainTxExist(uint64 fromChainId, bytes32 fromChainTx) view returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataCallerSession) CheckIfFromChainTxExist(fromChainId uint64, fromChainTx [32]byte) (bool, error) {
	return _IEthCrossChainData.Contract.CheckIfFromChainTxExist(&_IEthCrossChainData.CallOpts, fromChainId, fromChainTx)
}

// GetCurEpochConPubKeyBytes is a free data retrieval call binding the contract method 0x69d48074.
//
// Solidity: function getCurEpochConPubKeyBytes() view returns(bytes)
func (_IEthCrossChainData *IEthCrossChainDataCaller) GetCurEpochConPubKeyBytes(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _IEthCrossChainData.contract.Call(opts, out, "getCurEpochConPubKeyBytes")
	return *ret0, err
}

// GetCurEpochConPubKeyBytes is a free data retrieval call binding the contract method 0x69d48074.
//
// Solidity: function getCurEpochConPubKeyBytes() view returns(bytes)
func (_IEthCrossChainData *IEthCrossChainDataSession) GetCurEpochConPubKeyBytes() ([]byte, error) {
	return _IEthCrossChainData.Contract.GetCurEpochConPubKeyBytes(&_IEthCrossChainData.CallOpts)
}

// GetCurEpochConPubKeyBytes is a free data retrieval call binding the contract method 0x69d48074.
//
// Solidity: function getCurEpochConPubKeyBytes() view returns(bytes)
func (_IEthCrossChainData *IEthCrossChainDataCallerSession) GetCurEpochConPubKeyBytes() ([]byte, error) {
	return _IEthCrossChainData.Contract.GetCurEpochConPubKeyBytes(&_IEthCrossChainData.CallOpts)
}

// GetCurEpochStartHeight is a free data retrieval call binding the contract method 0x5ac40790.
//
// Solidity: function getCurEpochStartHeight() view returns(uint32)
func (_IEthCrossChainData *IEthCrossChainDataCaller) GetCurEpochStartHeight(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _IEthCrossChainData.contract.Call(opts, out, "getCurEpochStartHeight")
	return *ret0, err
}

// GetCurEpochStartHeight is a free data retrieval call binding the contract method 0x5ac40790.
//
// Solidity: function getCurEpochStartHeight() view returns(uint32)
func (_IEthCrossChainData *IEthCrossChainDataSession) GetCurEpochStartHeight() (uint32, error) {
	return _IEthCrossChainData.Contract.GetCurEpochStartHeight(&_IEthCrossChainData.CallOpts)
}

// GetCurEpochStartHeight is a free data retrieval call binding the contract method 0x5ac40790.
//
// Solidity: function getCurEpochStartHeight() view returns(uint32)
func (_IEthCrossChainData *IEthCrossChainDataCallerSession) GetCurEpochStartHeight() (uint32, error) {
	return _IEthCrossChainData.Contract.GetCurEpochStartHeight(&_IEthCrossChainData.CallOpts)
}

// GetEthTxHash is a free data retrieval call binding the contract method 0x29927875.
//
// Solidity: function getEthTxHash(uint256 ethTxHashIndex) view returns(bytes32)
func (_IEthCrossChainData *IEthCrossChainDataCaller) GetEthTxHash(opts *bind.CallOpts, ethTxHashIndex *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IEthCrossChainData.contract.Call(opts, out, "getEthTxHash", ethTxHashIndex)
	return *ret0, err
}

// GetEthTxHash is a free data retrieval call binding the contract method 0x29927875.
//
// Solidity: function getEthTxHash(uint256 ethTxHashIndex) view returns(bytes32)
func (_IEthCrossChainData *IEthCrossChainDataSession) GetEthTxHash(ethTxHashIndex *big.Int) ([32]byte, error) {
	return _IEthCrossChainData.Contract.GetEthTxHash(&_IEthCrossChainData.CallOpts, ethTxHashIndex)
}

// GetEthTxHash is a free data retrieval call binding the contract method 0x29927875.
//
// Solidity: function getEthTxHash(uint256 ethTxHashIndex) view returns(bytes32)
func (_IEthCrossChainData *IEthCrossChainDataCallerSession) GetEthTxHash(ethTxHashIndex *big.Int) ([32]byte, error) {
	return _IEthCrossChainData.Contract.GetEthTxHash(&_IEthCrossChainData.CallOpts, ethTxHashIndex)
}

// GetEthTxHashIndex is a free data retrieval call binding the contract method 0xff3d24a1.
//
// Solidity: function getEthTxHashIndex() view returns(uint256)
func (_IEthCrossChainData *IEthCrossChainDataCaller) GetEthTxHashIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IEthCrossChainData.contract.Call(opts, out, "getEthTxHashIndex")
	return *ret0, err
}

// GetEthTxHashIndex is a free data retrieval call binding the contract method 0xff3d24a1.
//
// Solidity: function getEthTxHashIndex() view returns(uint256)
func (_IEthCrossChainData *IEthCrossChainDataSession) GetEthTxHashIndex() (*big.Int, error) {
	return _IEthCrossChainData.Contract.GetEthTxHashIndex(&_IEthCrossChainData.CallOpts)
}

// GetEthTxHashIndex is a free data retrieval call binding the contract method 0xff3d24a1.
//
// Solidity: function getEthTxHashIndex() view returns(uint256)
func (_IEthCrossChainData *IEthCrossChainDataCallerSession) GetEthTxHashIndex() (*big.Int, error) {
	return _IEthCrossChainData.Contract.GetEthTxHashIndex(&_IEthCrossChainData.CallOpts)
}

// GetExtraData is a free data retrieval call binding the contract method 0x40602bb5.
//
// Solidity: function getExtraData(bytes32 key1, bytes32 key2) view returns(bytes)
func (_IEthCrossChainData *IEthCrossChainDataCaller) GetExtraData(opts *bind.CallOpts, key1 [32]byte, key2 [32]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _IEthCrossChainData.contract.Call(opts, out, "getExtraData", key1, key2)
	return *ret0, err
}

// GetExtraData is a free data retrieval call binding the contract method 0x40602bb5.
//
// Solidity: function getExtraData(bytes32 key1, bytes32 key2) view returns(bytes)
func (_IEthCrossChainData *IEthCrossChainDataSession) GetExtraData(key1 [32]byte, key2 [32]byte) ([]byte, error) {
	return _IEthCrossChainData.Contract.GetExtraData(&_IEthCrossChainData.CallOpts, key1, key2)
}

// GetExtraData is a free data retrieval call binding the contract method 0x40602bb5.
//
// Solidity: function getExtraData(bytes32 key1, bytes32 key2) view returns(bytes)
func (_IEthCrossChainData *IEthCrossChainDataCallerSession) GetExtraData(key1 [32]byte, key2 [32]byte) ([]byte, error) {
	return _IEthCrossChainData.Contract.GetExtraData(&_IEthCrossChainData.CallOpts, key1, key2)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IEthCrossChainData.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) Paused() (bool, error) {
	return _IEthCrossChainData.Contract.Paused(&_IEthCrossChainData.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataCallerSession) Paused() (bool, error) {
	return _IEthCrossChainData.Contract.Paused(&_IEthCrossChainData.CallOpts)
}

// MarkFromChainTxExist is a paid mutator transaction binding the contract method 0xe90bfdcf.
//
// Solidity: function markFromChainTxExist(uint64 fromChainId, bytes32 fromChainTx) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactor) MarkFromChainTxExist(opts *bind.TransactOpts, fromChainId uint64, fromChainTx [32]byte) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "markFromChainTxExist", fromChainId, fromChainTx)
}

// MarkFromChainTxExist is a paid mutator transaction binding the contract method 0xe90bfdcf.
//
// Solidity: function markFromChainTxExist(uint64 fromChainId, bytes32 fromChainTx) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) MarkFromChainTxExist(fromChainId uint64, fromChainTx [32]byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.MarkFromChainTxExist(&_IEthCrossChainData.TransactOpts, fromChainId, fromChainTx)
}

// MarkFromChainTxExist is a paid mutator transaction binding the contract method 0xe90bfdcf.
//
// Solidity: function markFromChainTxExist(uint64 fromChainId, bytes32 fromChainTx) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) MarkFromChainTxExist(fromChainId uint64, fromChainTx [32]byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.MarkFromChainTxExist(&_IEthCrossChainData.TransactOpts, fromChainId, fromChainTx)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) Pause() (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.Pause(&_IEthCrossChainData.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) Pause() (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.Pause(&_IEthCrossChainData.TransactOpts)
}

// PutCurEpochConPubKeyBytes is a paid mutator transaction binding the contract method 0x41973cd9.
//
// Solidity: function putCurEpochConPubKeyBytes(bytes curEpochPkBytes) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactor) PutCurEpochConPubKeyBytes(opts *bind.TransactOpts, curEpochPkBytes []byte) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "putCurEpochConPubKeyBytes", curEpochPkBytes)
}

// PutCurEpochConPubKeyBytes is a paid mutator transaction binding the contract method 0x41973cd9.
//
// Solidity: function putCurEpochConPubKeyBytes(bytes curEpochPkBytes) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) PutCurEpochConPubKeyBytes(curEpochPkBytes []byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutCurEpochConPubKeyBytes(&_IEthCrossChainData.TransactOpts, curEpochPkBytes)
}

// PutCurEpochConPubKeyBytes is a paid mutator transaction binding the contract method 0x41973cd9.
//
// Solidity: function putCurEpochConPubKeyBytes(bytes curEpochPkBytes) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) PutCurEpochConPubKeyBytes(curEpochPkBytes []byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutCurEpochConPubKeyBytes(&_IEthCrossChainData.TransactOpts, curEpochPkBytes)
}

// PutCurEpochStartHeight is a paid mutator transaction binding the contract method 0x8a8bd17f.
//
// Solidity: function putCurEpochStartHeight(uint32 curEpochStartHeight) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactor) PutCurEpochStartHeight(opts *bind.TransactOpts, curEpochStartHeight uint32) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "putCurEpochStartHeight", curEpochStartHeight)
}

// PutCurEpochStartHeight is a paid mutator transaction binding the contract method 0x8a8bd17f.
//
// Solidity: function putCurEpochStartHeight(uint32 curEpochStartHeight) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) PutCurEpochStartHeight(curEpochStartHeight uint32) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutCurEpochStartHeight(&_IEthCrossChainData.TransactOpts, curEpochStartHeight)
}

// PutCurEpochStartHeight is a paid mutator transaction binding the contract method 0x8a8bd17f.
//
// Solidity: function putCurEpochStartHeight(uint32 curEpochStartHeight) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) PutCurEpochStartHeight(curEpochStartHeight uint32) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutCurEpochStartHeight(&_IEthCrossChainData.TransactOpts, curEpochStartHeight)
}

// PutEthTxHash is a paid mutator transaction binding the contract method 0x4c3ccf64.
//
// Solidity: function putEthTxHash(bytes32 ethTxHash) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactor) PutEthTxHash(opts *bind.TransactOpts, ethTxHash [32]byte) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "putEthTxHash", ethTxHash)
}

// PutEthTxHash is a paid mutator transaction binding the contract method 0x4c3ccf64.
//
// Solidity: function putEthTxHash(bytes32 ethTxHash) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) PutEthTxHash(ethTxHash [32]byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutEthTxHash(&_IEthCrossChainData.TransactOpts, ethTxHash)
}

// PutEthTxHash is a paid mutator transaction binding the contract method 0x4c3ccf64.
//
// Solidity: function putEthTxHash(bytes32 ethTxHash) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) PutEthTxHash(ethTxHash [32]byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutEthTxHash(&_IEthCrossChainData.TransactOpts, ethTxHash)
}

// PutExtraData is a paid mutator transaction binding the contract method 0x1afe374e.
//
// Solidity: function putExtraData(bytes32 key1, bytes32 key2, bytes value) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactor) PutExtraData(opts *bind.TransactOpts, key1 [32]byte, key2 [32]byte, value []byte) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "putExtraData", key1, key2, value)
}

// PutExtraData is a paid mutator transaction binding the contract method 0x1afe374e.
//
// Solidity: function putExtraData(bytes32 key1, bytes32 key2, bytes value) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) PutExtraData(key1 [32]byte, key2 [32]byte, value []byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutExtraData(&_IEthCrossChainData.TransactOpts, key1, key2, value)
}

// PutExtraData is a paid mutator transaction binding the contract method 0x1afe374e.
//
// Solidity: function putExtraData(bytes32 key1, bytes32 key2, bytes value) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) PutExtraData(key1 [32]byte, key2 [32]byte, value []byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutExtraData(&_IEthCrossChainData.TransactOpts, key1, key2, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IEthCrossChainData *IEthCrossChainDataTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IEthCrossChainData *IEthCrossChainDataSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.TransferOwnership(&_IEthCrossChainData.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.TransferOwnership(&_IEthCrossChainData.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) Unpause() (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.Unpause(&_IEthCrossChainData.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) Unpause() (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.Unpause(&_IEthCrossChainData.TransactOpts)
}

// IEthCrossChainManagerABI is the input ABI used to generate the binding from.
const IEthCrossChainManagerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_toChainId\",\"type\":\"uint64\"},{\"name\":\"_toContract\",\"type\":\"bytes\"},{\"name\":\"_method\",\"type\":\"bytes\"},{\"name\":\"_txData\",\"type\":\"bytes\"}],\"name\":\"crossChain\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IEthCrossChainManagerFuncSigs maps the 4-byte function signature to its string representation.
var IEthCrossChainManagerFuncSigs = map[string]string{
	"bd5cf625": "crossChain(uint64,bytes,bytes,bytes)",
}

// IEthCrossChainManager is an auto generated Go binding around an Ethereum contract.
type IEthCrossChainManager struct {
	IEthCrossChainManagerCaller     // Read-only binding to the contract
	IEthCrossChainManagerTransactor // Write-only binding to the contract
	IEthCrossChainManagerFilterer   // Log filterer for contract events
}

// IEthCrossChainManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IEthCrossChainManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthCrossChainManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IEthCrossChainManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthCrossChainManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IEthCrossChainManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthCrossChainManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IEthCrossChainManagerSession struct {
	Contract     *IEthCrossChainManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IEthCrossChainManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IEthCrossChainManagerCallerSession struct {
	Contract *IEthCrossChainManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// IEthCrossChainManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IEthCrossChainManagerTransactorSession struct {
	Contract     *IEthCrossChainManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// IEthCrossChainManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IEthCrossChainManagerRaw struct {
	Contract *IEthCrossChainManager // Generic contract binding to access the raw methods on
}

// IEthCrossChainManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IEthCrossChainManagerCallerRaw struct {
	Contract *IEthCrossChainManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IEthCrossChainManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IEthCrossChainManagerTransactorRaw struct {
	Contract *IEthCrossChainManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIEthCrossChainManager creates a new instance of IEthCrossChainManager, bound to a specific deployed contract.
func NewIEthCrossChainManager(address common.Address, backend bind.ContractBackend) (*IEthCrossChainManager, error) {
	contract, err := bindIEthCrossChainManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainManager{IEthCrossChainManagerCaller: IEthCrossChainManagerCaller{contract: contract}, IEthCrossChainManagerTransactor: IEthCrossChainManagerTransactor{contract: contract}, IEthCrossChainManagerFilterer: IEthCrossChainManagerFilterer{contract: contract}}, nil
}

// NewIEthCrossChainManagerCaller creates a new read-only instance of IEthCrossChainManager, bound to a specific deployed contract.
func NewIEthCrossChainManagerCaller(address common.Address, caller bind.ContractCaller) (*IEthCrossChainManagerCaller, error) {
	contract, err := bindIEthCrossChainManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainManagerCaller{contract: contract}, nil
}

// NewIEthCrossChainManagerTransactor creates a new write-only instance of IEthCrossChainManager, bound to a specific deployed contract.
func NewIEthCrossChainManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IEthCrossChainManagerTransactor, error) {
	contract, err := bindIEthCrossChainManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainManagerTransactor{contract: contract}, nil
}

// NewIEthCrossChainManagerFilterer creates a new log filterer instance of IEthCrossChainManager, bound to a specific deployed contract.
func NewIEthCrossChainManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IEthCrossChainManagerFilterer, error) {
	contract, err := bindIEthCrossChainManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainManagerFilterer{contract: contract}, nil
}

// bindIEthCrossChainManager binds a generic wrapper to an already deployed contract.
func bindIEthCrossChainManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IEthCrossChainManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEthCrossChainManager *IEthCrossChainManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IEthCrossChainManager.Contract.IEthCrossChainManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEthCrossChainManager *IEthCrossChainManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthCrossChainManager.Contract.IEthCrossChainManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEthCrossChainManager *IEthCrossChainManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEthCrossChainManager.Contract.IEthCrossChainManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEthCrossChainManager *IEthCrossChainManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IEthCrossChainManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEthCrossChainManager *IEthCrossChainManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthCrossChainManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEthCrossChainManager *IEthCrossChainManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEthCrossChainManager.Contract.contract.Transact(opts, method, params...)
}

// CrossChain is a paid mutator transaction binding the contract method 0xbd5cf625.
//
// Solidity: function crossChain(uint64 _toChainId, bytes _toContract, bytes _method, bytes _txData) returns(bool)
func (_IEthCrossChainManager *IEthCrossChainManagerTransactor) CrossChain(opts *bind.TransactOpts, _toChainId uint64, _toContract []byte, _method []byte, _txData []byte) (*types.Transaction, error) {
	return _IEthCrossChainManager.contract.Transact(opts, "crossChain", _toChainId, _toContract, _method, _txData)
}

// CrossChain is a paid mutator transaction binding the contract method 0xbd5cf625.
//
// Solidity: function crossChain(uint64 _toChainId, bytes _toContract, bytes _method, bytes _txData) returns(bool)
func (_IEthCrossChainManager *IEthCrossChainManagerSession) CrossChain(_toChainId uint64, _toContract []byte, _method []byte, _txData []byte) (*types.Transaction, error) {
	return _IEthCrossChainManager.Contract.CrossChain(&_IEthCrossChainManager.TransactOpts, _toChainId, _toContract, _method, _txData)
}

// CrossChain is a paid mutator transaction binding the contract method 0xbd5cf625.
//
// Solidity: function crossChain(uint64 _toChainId, bytes _toContract, bytes _method, bytes _txData) returns(bool)
func (_IEthCrossChainManager *IEthCrossChainManagerTransactorSession) CrossChain(_toChainId uint64, _toContract []byte, _method []byte, _txData []byte) (*types.Transaction, error) {
	return _IEthCrossChainManager.Contract.CrossChain(&_IEthCrossChainManager.TransactOpts, _toChainId, _toContract, _method, _txData)
}

// IUpgradableECCMABI is the input ABI used to generate the binding from.
const IUpgradableECCMABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newChainId\",\"type\":\"uint64\"}],\"name\":\"setChainId\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"upgradeToNew\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IUpgradableECCMFuncSigs maps the 4-byte function signature to its string representation.
var IUpgradableECCMFuncSigs = map[string]string{
	"8f32d59b": "isOwner()",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"6f31031d": "setChainId(uint64)",
	"3f4ba83a": "unpause()",
	"7e724ff3": "upgradeToNew(address)",
}

// IUpgradableECCM is an auto generated Go binding around an Ethereum contract.
type IUpgradableECCM struct {
	IUpgradableECCMCaller     // Read-only binding to the contract
	IUpgradableECCMTransactor // Write-only binding to the contract
	IUpgradableECCMFilterer   // Log filterer for contract events
}

// IUpgradableECCMCaller is an auto generated read-only Go binding around an Ethereum contract.
type IUpgradableECCMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUpgradableECCMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IUpgradableECCMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUpgradableECCMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IUpgradableECCMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUpgradableECCMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IUpgradableECCMSession struct {
	Contract     *IUpgradableECCM  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IUpgradableECCMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IUpgradableECCMCallerSession struct {
	Contract *IUpgradableECCMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IUpgradableECCMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IUpgradableECCMTransactorSession struct {
	Contract     *IUpgradableECCMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IUpgradableECCMRaw is an auto generated low-level Go binding around an Ethereum contract.
type IUpgradableECCMRaw struct {
	Contract *IUpgradableECCM // Generic contract binding to access the raw methods on
}

// IUpgradableECCMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IUpgradableECCMCallerRaw struct {
	Contract *IUpgradableECCMCaller // Generic read-only contract binding to access the raw methods on
}

// IUpgradableECCMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IUpgradableECCMTransactorRaw struct {
	Contract *IUpgradableECCMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIUpgradableECCM creates a new instance of IUpgradableECCM, bound to a specific deployed contract.
func NewIUpgradableECCM(address common.Address, backend bind.ContractBackend) (*IUpgradableECCM, error) {
	contract, err := bindIUpgradableECCM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IUpgradableECCM{IUpgradableECCMCaller: IUpgradableECCMCaller{contract: contract}, IUpgradableECCMTransactor: IUpgradableECCMTransactor{contract: contract}, IUpgradableECCMFilterer: IUpgradableECCMFilterer{contract: contract}}, nil
}

// NewIUpgradableECCMCaller creates a new read-only instance of IUpgradableECCM, bound to a specific deployed contract.
func NewIUpgradableECCMCaller(address common.Address, caller bind.ContractCaller) (*IUpgradableECCMCaller, error) {
	contract, err := bindIUpgradableECCM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IUpgradableECCMCaller{contract: contract}, nil
}

// NewIUpgradableECCMTransactor creates a new write-only instance of IUpgradableECCM, bound to a specific deployed contract.
func NewIUpgradableECCMTransactor(address common.Address, transactor bind.ContractTransactor) (*IUpgradableECCMTransactor, error) {
	contract, err := bindIUpgradableECCM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IUpgradableECCMTransactor{contract: contract}, nil
}

// NewIUpgradableECCMFilterer creates a new log filterer instance of IUpgradableECCM, bound to a specific deployed contract.
func NewIUpgradableECCMFilterer(address common.Address, filterer bind.ContractFilterer) (*IUpgradableECCMFilterer, error) {
	contract, err := bindIUpgradableECCM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IUpgradableECCMFilterer{contract: contract}, nil
}

// bindIUpgradableECCM binds a generic wrapper to an already deployed contract.
func bindIUpgradableECCM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IUpgradableECCMABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUpgradableECCM *IUpgradableECCMRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IUpgradableECCM.Contract.IUpgradableECCMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUpgradableECCM *IUpgradableECCMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUpgradableECCM.Contract.IUpgradableECCMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUpgradableECCM *IUpgradableECCMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUpgradableECCM.Contract.IUpgradableECCMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUpgradableECCM *IUpgradableECCMCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IUpgradableECCM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUpgradableECCM *IUpgradableECCMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUpgradableECCM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUpgradableECCM *IUpgradableECCMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUpgradableECCM.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_IUpgradableECCM *IUpgradableECCMCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IUpgradableECCM.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_IUpgradableECCM *IUpgradableECCMSession) IsOwner() (bool, error) {
	return _IUpgradableECCM.Contract.IsOwner(&_IUpgradableECCM.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_IUpgradableECCM *IUpgradableECCMCallerSession) IsOwner() (bool, error) {
	return _IUpgradableECCM.Contract.IsOwner(&_IUpgradableECCM.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IUpgradableECCM *IUpgradableECCMCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IUpgradableECCM.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IUpgradableECCM *IUpgradableECCMSession) Paused() (bool, error) {
	return _IUpgradableECCM.Contract.Paused(&_IUpgradableECCM.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IUpgradableECCM *IUpgradableECCMCallerSession) Paused() (bool, error) {
	return _IUpgradableECCM.Contract.Paused(&_IUpgradableECCM.CallOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_IUpgradableECCM *IUpgradableECCMTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUpgradableECCM.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_IUpgradableECCM *IUpgradableECCMSession) Pause() (*types.Transaction, error) {
	return _IUpgradableECCM.Contract.Pause(&_IUpgradableECCM.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_IUpgradableECCM *IUpgradableECCMTransactorSession) Pause() (*types.Transaction, error) {
	return _IUpgradableECCM.Contract.Pause(&_IUpgradableECCM.TransactOpts)
}

// SetChainId is a paid mutator transaction binding the contract method 0x6f31031d.
//
// Solidity: function setChainId(uint64 _newChainId) returns(bool)
func (_IUpgradableECCM *IUpgradableECCMTransactor) SetChainId(opts *bind.TransactOpts, _newChainId uint64) (*types.Transaction, error) {
	return _IUpgradableECCM.contract.Transact(opts, "setChainId", _newChainId)
}

// SetChainId is a paid mutator transaction binding the contract method 0x6f31031d.
//
// Solidity: function setChainId(uint64 _newChainId) returns(bool)
func (_IUpgradableECCM *IUpgradableECCMSession) SetChainId(_newChainId uint64) (*types.Transaction, error) {
	return _IUpgradableECCM.Contract.SetChainId(&_IUpgradableECCM.TransactOpts, _newChainId)
}

// SetChainId is a paid mutator transaction binding the contract method 0x6f31031d.
//
// Solidity: function setChainId(uint64 _newChainId) returns(bool)
func (_IUpgradableECCM *IUpgradableECCMTransactorSession) SetChainId(_newChainId uint64) (*types.Transaction, error) {
	return _IUpgradableECCM.Contract.SetChainId(&_IUpgradableECCM.TransactOpts, _newChainId)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_IUpgradableECCM *IUpgradableECCMTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUpgradableECCM.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_IUpgradableECCM *IUpgradableECCMSession) Unpause() (*types.Transaction, error) {
	return _IUpgradableECCM.Contract.Unpause(&_IUpgradableECCM.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_IUpgradableECCM *IUpgradableECCMTransactorSession) Unpause() (*types.Transaction, error) {
	return _IUpgradableECCM.Contract.Unpause(&_IUpgradableECCM.TransactOpts)
}

// UpgradeToNew is a paid mutator transaction binding the contract method 0x7e724ff3.
//
// Solidity: function upgradeToNew(address ) returns(bool)
func (_IUpgradableECCM *IUpgradableECCMTransactor) UpgradeToNew(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _IUpgradableECCM.contract.Transact(opts, "upgradeToNew", arg0)
}

// UpgradeToNew is a paid mutator transaction binding the contract method 0x7e724ff3.
//
// Solidity: function upgradeToNew(address ) returns(bool)
func (_IUpgradableECCM *IUpgradableECCMSession) UpgradeToNew(arg0 common.Address) (*types.Transaction, error) {
	return _IUpgradableECCM.Contract.UpgradeToNew(&_IUpgradableECCM.TransactOpts, arg0)
}

// UpgradeToNew is a paid mutator transaction binding the contract method 0x7e724ff3.
//
// Solidity: function upgradeToNew(address ) returns(bool)
func (_IUpgradableECCM *IUpgradableECCMTransactorSession) UpgradeToNew(arg0 common.Address) (*types.Transaction, error) {
	return _IUpgradableECCM.Contract.UpgradeToNew(&_IUpgradableECCM.TransactOpts, arg0)
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = map[string]string{
	"8f32d59b": "isOwner()",
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"f2fde38b": "transferOwnership(address)",
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Ownable *OwnableCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Ownable *OwnableSession) IsOwner() (bool, error) {
	return _Ownable.Contract.IsOwner(&_Ownable.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Ownable *OwnableCallerSession) IsOwner() (bool, error) {
	return _Ownable.Contract.IsOwner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PausableABI is the input ABI used to generate the binding from.
const PausableABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"}]"

// PausableFuncSigs maps the 4-byte function signature to its string representation.
var PausableFuncSigs = map[string]string{
	"5c975abb": "paused()",
}

// Pausable is an auto generated Go binding around an Ethereum contract.
type Pausable struct {
	PausableCaller     // Read-only binding to the contract
	PausableTransactor // Write-only binding to the contract
	PausableFilterer   // Log filterer for contract events
}

// PausableCaller is an auto generated read-only Go binding around an Ethereum contract.
type PausableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PausableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PausableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PausableSession struct {
	Contract     *Pausable         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PausableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PausableCallerSession struct {
	Contract *PausableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PausableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PausableTransactorSession struct {
	Contract     *PausableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PausableRaw is an auto generated low-level Go binding around an Ethereum contract.
type PausableRaw struct {
	Contract *Pausable // Generic contract binding to access the raw methods on
}

// PausableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PausableCallerRaw struct {
	Contract *PausableCaller // Generic read-only contract binding to access the raw methods on
}

// PausableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PausableTransactorRaw struct {
	Contract *PausableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPausable creates a new instance of Pausable, bound to a specific deployed contract.
func NewPausable(address common.Address, backend bind.ContractBackend) (*Pausable, error) {
	contract, err := bindPausable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pausable{PausableCaller: PausableCaller{contract: contract}, PausableTransactor: PausableTransactor{contract: contract}, PausableFilterer: PausableFilterer{contract: contract}}, nil
}

// NewPausableCaller creates a new read-only instance of Pausable, bound to a specific deployed contract.
func NewPausableCaller(address common.Address, caller bind.ContractCaller) (*PausableCaller, error) {
	contract, err := bindPausable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PausableCaller{contract: contract}, nil
}

// NewPausableTransactor creates a new write-only instance of Pausable, bound to a specific deployed contract.
func NewPausableTransactor(address common.Address, transactor bind.ContractTransactor) (*PausableTransactor, error) {
	contract, err := bindPausable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PausableTransactor{contract: contract}, nil
}

// NewPausableFilterer creates a new log filterer instance of Pausable, bound to a specific deployed contract.
func NewPausableFilterer(address common.Address, filterer bind.ContractFilterer) (*PausableFilterer, error) {
	contract, err := bindPausable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PausableFilterer{contract: contract}, nil
}

// bindPausable binds a generic wrapper to an already deployed contract.
func bindPausable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.PausableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transact(opts, method, params...)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Pausable.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableCallerSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// PausablePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Pausable contract.
type PausablePausedIterator struct {
	Event *PausablePaused // Event containing the contract specifics and raw log

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
func (it *PausablePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausablePaused)
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
		it.Event = new(PausablePaused)
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
func (it *PausablePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausablePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausablePaused represents a Paused event raised by the Pausable contract.
type PausablePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pausable *PausableFilterer) FilterPaused(opts *bind.FilterOpts) (*PausablePausedIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PausablePausedIterator{contract: _Pausable.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pausable *PausableFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PausablePaused) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausablePaused)
				if err := _Pausable.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pausable *PausableFilterer) ParsePaused(log types.Log) (*PausablePaused, error) {
	event := new(PausablePaused)
	if err := _Pausable.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PausableUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Pausable contract.
type PausableUnpausedIterator struct {
	Event *PausableUnpaused // Event containing the contract specifics and raw log

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
func (it *PausableUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableUnpaused)
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
		it.Event = new(PausableUnpaused)
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
func (it *PausableUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableUnpaused represents a Unpaused event raised by the Pausable contract.
type PausableUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pausable *PausableFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PausableUnpausedIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PausableUnpausedIterator{contract: _Pausable.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pausable *PausableFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PausableUnpaused) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableUnpaused)
				if err := _Pausable.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pausable *PausableFilterer) ParseUnpaused(log types.Log) (*PausableUnpaused, error) {
	event := new(PausableUnpaused)
	if err := _Pausable.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723058205bc3e87da8675f6ca37b1c6db827d34a569538235bcfb1108bb528f49c372cb464736f6c634300050a0032"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// UpgradableECCMABI is the input ABI used to generate the binding from.
const UpgradableECCMABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"EthCrossChainDataAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newChainId\",\"type\":\"uint64\"}],\"name\":\"setChainId\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newEthCrossChainManagerAddress\",\"type\":\"address\"}],\"name\":\"upgradeToNew\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"chainId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"ethCrossChainDataAddr\",\"type\":\"address\"},{\"name\":\"_chainId\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// UpgradableECCMFuncSigs maps the 4-byte function signature to its string representation.
var UpgradableECCMFuncSigs = map[string]string{
	"00ba1694": "EthCrossChainDataAddress()",
	"9a8a0592": "chainId()",
	"8f32d59b": "isOwner()",
	"8da5cb5b": "owner()",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"715018a6": "renounceOwnership()",
	"6f31031d": "setChainId(uint64)",
	"f2fde38b": "transferOwnership(address)",
	"3f4ba83a": "unpause()",
	"7e724ff3": "upgradeToNew(address)",
}

// UpgradableECCMBin is the compiled bytecode used for deploying new contracts.
var UpgradableECCMBin = "0x608060405234801561001057600080fd5b50604051610b7d380380610b7d8339818101604052604081101561003357600080fd5b508051602090910151600061004f6001600160e01b036100ff16565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3506000805460ff60a01b19169055600180546001600160401b039092167401000000000000000000000000000000000000000002600160a01b600160e01b03196001600160a01b039094166001600160a01b03199093169290921792909216179055610103565b3390565b610a6b806101126000396000f3fe608060405234801561001057600080fd5b50600436106100a85760003560e01c80637e724ff3116100715780637e724ff3146101265780638456cb591461014c5780638da5cb5b146101545780638f32d59b1461015c5780639a8a059214610164578063f2fde38b14610189576100a8565b8062ba1694146100ad5780633f4ba83a146100d15780635c975abb146100ed5780636f31031d146100f5578063715018a61461011c575b600080fd5b6100b56101af565b604080516001600160a01b039092168252519081900360200190f35b6100d96101be565b604080519115158252519081900360200190f35b6100d961033d565b6100d96004803603602081101561010b57600080fd5b503567ffffffffffffffff1661034d565b610124610418565b005b6100d96004803603602081101561013c57600080fd5b50356001600160a01b03166104a9565b6100d96105b8565b6100b561072d565b6100d961073c565b61016c610760565b6040805167ffffffffffffffff9092168252519081900360200190f35b6101246004803603602081101561019f57600080fd5b50356001600160a01b0316610777565b6001546001600160a01b031681565b60006101c861073c565b610207576040805162461bcd60e51b81526020600482018190526024820152600080516020610a17833981519152604482015290519081900360640190fd5b61020f61033d565b1561021c5761021c6107ca565b60015460408051635c975abb60e01b815290516001600160a01b03909216918291635c975abb916004808301926020929190829003018186803b15801561026257600080fd5b505afa158015610276573d6000803e3d6000fd5b505050506040513d602081101561028c57600080fd5b50511561033557806001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156102ce57600080fd5b505af11580156102e2573d6000803e3d6000fd5b505050506040513d60208110156102f857600080fd5b50516103355760405162461bcd60e51b81526004018080602001828103825260298152602001806109c86029913960400191505060405180910390fd5b600191505090565b600054600160a01b900460ff1690565b60008054600160a01b900460ff166103a3576040805162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b604482015290519081900360640190fd5b6103ab61073c565b6103ea576040805162461bcd60e51b81526020600482018190526024820152600080516020610a17833981519152604482015290519081900360640190fd5b506001805467ffffffffffffffff8316600160a01b0267ffffffffffffffff60a01b19909116178155919050565b61042061073c565b61045f576040805162461bcd60e51b81526020600482018190526024820152600080516020610a17833981519152604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b60008054600160a01b900460ff166104ff576040805162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b604482015290519081900360640190fd5b61050761073c565b610546576040805162461bcd60e51b81526020600482018190526024820152600080516020610a17833981519152604482015290519081900360640190fd5b6001546040805163f2fde38b60e01b81526001600160a01b03858116600483015291519190921691829163f2fde38b9160248082019260009290919082900301818387803b15801561059757600080fd5b505af11580156105ab573d6000803e3d6000fd5b5060019695505050505050565b60006105c261073c565b610601576040805162461bcd60e51b81526020600482018190526024820152600080516020610a17833981519152604482015290519081900360640190fd5b61060961033d565b61061557610615610872565b60015460408051635c975abb60e01b815290516001600160a01b03909216918291635c975abb916004808301926020929190829003018186803b15801561065b57600080fd5b505afa15801561066f573d6000803e3d6000fd5b505050506040513d602081101561068557600080fd5b505161033557806001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156106c657600080fd5b505af11580156106da573d6000803e3d6000fd5b505050506040513d60208110156106f057600080fd5b50516103355760405162461bcd60e51b81526004018080602001828103825260278152602001806109a16027913960400191505060405180910390fd5b6000546001600160a01b031690565b600080546001600160a01b03166107516108fc565b6001600160a01b031614905090565b600154600160a01b900467ffffffffffffffff1681565b61077f61073c565b6107be576040805162461bcd60e51b81526020600482018190526024820152600080516020610a17833981519152604482015290519081900360640190fd5b6107c781610900565b50565b600054600160a01b900460ff1661081f576040805162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b604482015290519081900360640190fd5b6000805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6108556108fc565b604080516001600160a01b039092168252519081900360200190a1565b600054600160a01b900460ff16156108c4576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b6000805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586108555b3390565b6001600160a01b0381166109455760405162461bcd60e51b81526004018080602001828103825260268152602001806109f16026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b039290921691909117905556fe70617573652045746843726f7373436861696e4461746120636f6e7472616374206661696c6564756e70617573652045746843726f7373436861696e4461746120636f6e7472616374206661696c65644f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573734f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572a265627a7a72305820eb87c3f2af252d01f13ec194b585e2faccf63fa5bc748197ed2cfc63952b2e5d64736f6c634300050a0032"

// DeployUpgradableECCM deploys a new Ethereum contract, binding an instance of UpgradableECCM to it.
func DeployUpgradableECCM(auth *bind.TransactOpts, backend bind.ContractBackend, ethCrossChainDataAddr common.Address, _chainId uint64) (common.Address, *types.Transaction, *UpgradableECCM, error) {
	parsed, err := abi.JSON(strings.NewReader(UpgradableECCMABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UpgradableECCMBin), backend, ethCrossChainDataAddr, _chainId)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UpgradableECCM{UpgradableECCMCaller: UpgradableECCMCaller{contract: contract}, UpgradableECCMTransactor: UpgradableECCMTransactor{contract: contract}, UpgradableECCMFilterer: UpgradableECCMFilterer{contract: contract}}, nil
}

// UpgradableECCM is an auto generated Go binding around an Ethereum contract.
type UpgradableECCM struct {
	UpgradableECCMCaller     // Read-only binding to the contract
	UpgradableECCMTransactor // Write-only binding to the contract
	UpgradableECCMFilterer   // Log filterer for contract events
}

// UpgradableECCMCaller is an auto generated read-only Go binding around an Ethereum contract.
type UpgradableECCMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradableECCMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UpgradableECCMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradableECCMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UpgradableECCMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradableECCMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UpgradableECCMSession struct {
	Contract     *UpgradableECCM   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UpgradableECCMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UpgradableECCMCallerSession struct {
	Contract *UpgradableECCMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// UpgradableECCMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UpgradableECCMTransactorSession struct {
	Contract     *UpgradableECCMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// UpgradableECCMRaw is an auto generated low-level Go binding around an Ethereum contract.
type UpgradableECCMRaw struct {
	Contract *UpgradableECCM // Generic contract binding to access the raw methods on
}

// UpgradableECCMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UpgradableECCMCallerRaw struct {
	Contract *UpgradableECCMCaller // Generic read-only contract binding to access the raw methods on
}

// UpgradableECCMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UpgradableECCMTransactorRaw struct {
	Contract *UpgradableECCMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUpgradableECCM creates a new instance of UpgradableECCM, bound to a specific deployed contract.
func NewUpgradableECCM(address common.Address, backend bind.ContractBackend) (*UpgradableECCM, error) {
	contract, err := bindUpgradableECCM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UpgradableECCM{UpgradableECCMCaller: UpgradableECCMCaller{contract: contract}, UpgradableECCMTransactor: UpgradableECCMTransactor{contract: contract}, UpgradableECCMFilterer: UpgradableECCMFilterer{contract: contract}}, nil
}

// NewUpgradableECCMCaller creates a new read-only instance of UpgradableECCM, bound to a specific deployed contract.
func NewUpgradableECCMCaller(address common.Address, caller bind.ContractCaller) (*UpgradableECCMCaller, error) {
	contract, err := bindUpgradableECCM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradableECCMCaller{contract: contract}, nil
}

// NewUpgradableECCMTransactor creates a new write-only instance of UpgradableECCM, bound to a specific deployed contract.
func NewUpgradableECCMTransactor(address common.Address, transactor bind.ContractTransactor) (*UpgradableECCMTransactor, error) {
	contract, err := bindUpgradableECCM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradableECCMTransactor{contract: contract}, nil
}

// NewUpgradableECCMFilterer creates a new log filterer instance of UpgradableECCM, bound to a specific deployed contract.
func NewUpgradableECCMFilterer(address common.Address, filterer bind.ContractFilterer) (*UpgradableECCMFilterer, error) {
	contract, err := bindUpgradableECCM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UpgradableECCMFilterer{contract: contract}, nil
}

// bindUpgradableECCM binds a generic wrapper to an already deployed contract.
func bindUpgradableECCM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UpgradableECCMABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpgradableECCM *UpgradableECCMRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UpgradableECCM.Contract.UpgradableECCMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpgradableECCM *UpgradableECCMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradableECCM.Contract.UpgradableECCMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpgradableECCM *UpgradableECCMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpgradableECCM.Contract.UpgradableECCMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpgradableECCM *UpgradableECCMCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UpgradableECCM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpgradableECCM *UpgradableECCMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradableECCM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpgradableECCM *UpgradableECCMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpgradableECCM.Contract.contract.Transact(opts, method, params...)
}

// EthCrossChainDataAddress is a free data retrieval call binding the contract method 0x00ba1694.
//
// Solidity: function EthCrossChainDataAddress() view returns(address)
func (_UpgradableECCM *UpgradableECCMCaller) EthCrossChainDataAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UpgradableECCM.contract.Call(opts, out, "EthCrossChainDataAddress")
	return *ret0, err
}

// EthCrossChainDataAddress is a free data retrieval call binding the contract method 0x00ba1694.
//
// Solidity: function EthCrossChainDataAddress() view returns(address)
func (_UpgradableECCM *UpgradableECCMSession) EthCrossChainDataAddress() (common.Address, error) {
	return _UpgradableECCM.Contract.EthCrossChainDataAddress(&_UpgradableECCM.CallOpts)
}

// EthCrossChainDataAddress is a free data retrieval call binding the contract method 0x00ba1694.
//
// Solidity: function EthCrossChainDataAddress() view returns(address)
func (_UpgradableECCM *UpgradableECCMCallerSession) EthCrossChainDataAddress() (common.Address, error) {
	return _UpgradableECCM.Contract.EthCrossChainDataAddress(&_UpgradableECCM.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint64)
func (_UpgradableECCM *UpgradableECCMCaller) ChainId(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _UpgradableECCM.contract.Call(opts, out, "chainId")
	return *ret0, err
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint64)
func (_UpgradableECCM *UpgradableECCMSession) ChainId() (uint64, error) {
	return _UpgradableECCM.Contract.ChainId(&_UpgradableECCM.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint64)
func (_UpgradableECCM *UpgradableECCMCallerSession) ChainId() (uint64, error) {
	return _UpgradableECCM.Contract.ChainId(&_UpgradableECCM.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_UpgradableECCM *UpgradableECCMCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _UpgradableECCM.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_UpgradableECCM *UpgradableECCMSession) IsOwner() (bool, error) {
	return _UpgradableECCM.Contract.IsOwner(&_UpgradableECCM.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_UpgradableECCM *UpgradableECCMCallerSession) IsOwner() (bool, error) {
	return _UpgradableECCM.Contract.IsOwner(&_UpgradableECCM.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UpgradableECCM *UpgradableECCMCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UpgradableECCM.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UpgradableECCM *UpgradableECCMSession) Owner() (common.Address, error) {
	return _UpgradableECCM.Contract.Owner(&_UpgradableECCM.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UpgradableECCM *UpgradableECCMCallerSession) Owner() (common.Address, error) {
	return _UpgradableECCM.Contract.Owner(&_UpgradableECCM.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_UpgradableECCM *UpgradableECCMCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _UpgradableECCM.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_UpgradableECCM *UpgradableECCMSession) Paused() (bool, error) {
	return _UpgradableECCM.Contract.Paused(&_UpgradableECCM.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_UpgradableECCM *UpgradableECCMCallerSession) Paused() (bool, error) {
	return _UpgradableECCM.Contract.Paused(&_UpgradableECCM.CallOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_UpgradableECCM *UpgradableECCMTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradableECCM.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_UpgradableECCM *UpgradableECCMSession) Pause() (*types.Transaction, error) {
	return _UpgradableECCM.Contract.Pause(&_UpgradableECCM.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_UpgradableECCM *UpgradableECCMTransactorSession) Pause() (*types.Transaction, error) {
	return _UpgradableECCM.Contract.Pause(&_UpgradableECCM.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UpgradableECCM *UpgradableECCMTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradableECCM.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UpgradableECCM *UpgradableECCMSession) RenounceOwnership() (*types.Transaction, error) {
	return _UpgradableECCM.Contract.RenounceOwnership(&_UpgradableECCM.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UpgradableECCM *UpgradableECCMTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _UpgradableECCM.Contract.RenounceOwnership(&_UpgradableECCM.TransactOpts)
}

// SetChainId is a paid mutator transaction binding the contract method 0x6f31031d.
//
// Solidity: function setChainId(uint64 _newChainId) returns(bool)
func (_UpgradableECCM *UpgradableECCMTransactor) SetChainId(opts *bind.TransactOpts, _newChainId uint64) (*types.Transaction, error) {
	return _UpgradableECCM.contract.Transact(opts, "setChainId", _newChainId)
}

// SetChainId is a paid mutator transaction binding the contract method 0x6f31031d.
//
// Solidity: function setChainId(uint64 _newChainId) returns(bool)
func (_UpgradableECCM *UpgradableECCMSession) SetChainId(_newChainId uint64) (*types.Transaction, error) {
	return _UpgradableECCM.Contract.SetChainId(&_UpgradableECCM.TransactOpts, _newChainId)
}

// SetChainId is a paid mutator transaction binding the contract method 0x6f31031d.
//
// Solidity: function setChainId(uint64 _newChainId) returns(bool)
func (_UpgradableECCM *UpgradableECCMTransactorSession) SetChainId(_newChainId uint64) (*types.Transaction, error) {
	return _UpgradableECCM.Contract.SetChainId(&_UpgradableECCM.TransactOpts, _newChainId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UpgradableECCM *UpgradableECCMTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _UpgradableECCM.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UpgradableECCM *UpgradableECCMSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UpgradableECCM.Contract.TransferOwnership(&_UpgradableECCM.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UpgradableECCM *UpgradableECCMTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UpgradableECCM.Contract.TransferOwnership(&_UpgradableECCM.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_UpgradableECCM *UpgradableECCMTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradableECCM.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_UpgradableECCM *UpgradableECCMSession) Unpause() (*types.Transaction, error) {
	return _UpgradableECCM.Contract.Unpause(&_UpgradableECCM.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_UpgradableECCM *UpgradableECCMTransactorSession) Unpause() (*types.Transaction, error) {
	return _UpgradableECCM.Contract.Unpause(&_UpgradableECCM.TransactOpts)
}

// UpgradeToNew is a paid mutator transaction binding the contract method 0x7e724ff3.
//
// Solidity: function upgradeToNew(address newEthCrossChainManagerAddress) returns(bool)
func (_UpgradableECCM *UpgradableECCMTransactor) UpgradeToNew(opts *bind.TransactOpts, newEthCrossChainManagerAddress common.Address) (*types.Transaction, error) {
	return _UpgradableECCM.contract.Transact(opts, "upgradeToNew", newEthCrossChainManagerAddress)
}

// UpgradeToNew is a paid mutator transaction binding the contract method 0x7e724ff3.
//
// Solidity: function upgradeToNew(address newEthCrossChainManagerAddress) returns(bool)
func (_UpgradableECCM *UpgradableECCMSession) UpgradeToNew(newEthCrossChainManagerAddress common.Address) (*types.Transaction, error) {
	return _UpgradableECCM.Contract.UpgradeToNew(&_UpgradableECCM.TransactOpts, newEthCrossChainManagerAddress)
}

// UpgradeToNew is a paid mutator transaction binding the contract method 0x7e724ff3.
//
// Solidity: function upgradeToNew(address newEthCrossChainManagerAddress) returns(bool)
func (_UpgradableECCM *UpgradableECCMTransactorSession) UpgradeToNew(newEthCrossChainManagerAddress common.Address) (*types.Transaction, error) {
	return _UpgradableECCM.Contract.UpgradeToNew(&_UpgradableECCM.TransactOpts, newEthCrossChainManagerAddress)
}

// UpgradableECCMOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the UpgradableECCM contract.
type UpgradableECCMOwnershipTransferredIterator struct {
	Event *UpgradableECCMOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *UpgradableECCMOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradableECCMOwnershipTransferred)
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
		it.Event = new(UpgradableECCMOwnershipTransferred)
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
func (it *UpgradableECCMOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradableECCMOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradableECCMOwnershipTransferred represents a OwnershipTransferred event raised by the UpgradableECCM contract.
type UpgradableECCMOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UpgradableECCM *UpgradableECCMFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*UpgradableECCMOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UpgradableECCM.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UpgradableECCMOwnershipTransferredIterator{contract: _UpgradableECCM.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UpgradableECCM *UpgradableECCMFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *UpgradableECCMOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UpgradableECCM.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradableECCMOwnershipTransferred)
				if err := _UpgradableECCM.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_UpgradableECCM *UpgradableECCMFilterer) ParseOwnershipTransferred(log types.Log) (*UpgradableECCMOwnershipTransferred, error) {
	event := new(UpgradableECCMOwnershipTransferred)
	if err := _UpgradableECCM.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// UpgradableECCMPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the UpgradableECCM contract.
type UpgradableECCMPausedIterator struct {
	Event *UpgradableECCMPaused // Event containing the contract specifics and raw log

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
func (it *UpgradableECCMPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradableECCMPaused)
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
		it.Event = new(UpgradableECCMPaused)
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
func (it *UpgradableECCMPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradableECCMPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradableECCMPaused represents a Paused event raised by the UpgradableECCM contract.
type UpgradableECCMPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_UpgradableECCM *UpgradableECCMFilterer) FilterPaused(opts *bind.FilterOpts) (*UpgradableECCMPausedIterator, error) {

	logs, sub, err := _UpgradableECCM.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &UpgradableECCMPausedIterator{contract: _UpgradableECCM.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_UpgradableECCM *UpgradableECCMFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *UpgradableECCMPaused) (event.Subscription, error) {

	logs, sub, err := _UpgradableECCM.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradableECCMPaused)
				if err := _UpgradableECCM.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_UpgradableECCM *UpgradableECCMFilterer) ParsePaused(log types.Log) (*UpgradableECCMPaused, error) {
	event := new(UpgradableECCMPaused)
	if err := _UpgradableECCM.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// UpgradableECCMUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the UpgradableECCM contract.
type UpgradableECCMUnpausedIterator struct {
	Event *UpgradableECCMUnpaused // Event containing the contract specifics and raw log

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
func (it *UpgradableECCMUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradableECCMUnpaused)
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
		it.Event = new(UpgradableECCMUnpaused)
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
func (it *UpgradableECCMUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradableECCMUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradableECCMUnpaused represents a Unpaused event raised by the UpgradableECCM contract.
type UpgradableECCMUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_UpgradableECCM *UpgradableECCMFilterer) FilterUnpaused(opts *bind.FilterOpts) (*UpgradableECCMUnpausedIterator, error) {

	logs, sub, err := _UpgradableECCM.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &UpgradableECCMUnpausedIterator{contract: _UpgradableECCM.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_UpgradableECCM *UpgradableECCMFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *UpgradableECCMUnpaused) (event.Subscription, error) {

	logs, sub, err := _UpgradableECCM.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradableECCMUnpaused)
				if err := _UpgradableECCM.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_UpgradableECCM *UpgradableECCMFilterer) ParseUnpaused(log types.Log) (*UpgradableECCMUnpaused, error) {
	event := new(UpgradableECCMUnpaused)
	if err := _UpgradableECCM.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// UtilsABI is the input ABI used to generate the binding from.
const UtilsABI = "[]"

// UtilsBin is the compiled bytecode used for deploying new contracts.
var UtilsBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7230582002d2aace566602b565cd74b2e6f14c147fb5b3cd9b76395e0fc2e1f9cea2443c64736f6c634300050a0032"

// DeployUtils deploys a new Ethereum contract, binding an instance of Utils to it.
func DeployUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Utils, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Utils{UtilsCaller: UtilsCaller{contract: contract}, UtilsTransactor: UtilsTransactor{contract: contract}, UtilsFilterer: UtilsFilterer{contract: contract}}, nil
}

// Utils is an auto generated Go binding around an Ethereum contract.
type Utils struct {
	UtilsCaller     // Read-only binding to the contract
	UtilsTransactor // Write-only binding to the contract
	UtilsFilterer   // Log filterer for contract events
}

// UtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type UtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UtilsSession struct {
	Contract     *Utils            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UtilsCallerSession struct {
	Contract *UtilsCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UtilsTransactorSession struct {
	Contract     *UtilsTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type UtilsRaw struct {
	Contract *Utils // Generic contract binding to access the raw methods on
}

// UtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UtilsCallerRaw struct {
	Contract *UtilsCaller // Generic read-only contract binding to access the raw methods on
}

// UtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UtilsTransactorRaw struct {
	Contract *UtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUtils creates a new instance of Utils, bound to a specific deployed contract.
func NewUtils(address common.Address, backend bind.ContractBackend) (*Utils, error) {
	contract, err := bindUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Utils{UtilsCaller: UtilsCaller{contract: contract}, UtilsTransactor: UtilsTransactor{contract: contract}, UtilsFilterer: UtilsFilterer{contract: contract}}, nil
}

// NewUtilsCaller creates a new read-only instance of Utils, bound to a specific deployed contract.
func NewUtilsCaller(address common.Address, caller bind.ContractCaller) (*UtilsCaller, error) {
	contract, err := bindUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UtilsCaller{contract: contract}, nil
}

// NewUtilsTransactor creates a new write-only instance of Utils, bound to a specific deployed contract.
func NewUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*UtilsTransactor, error) {
	contract, err := bindUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UtilsTransactor{contract: contract}, nil
}

// NewUtilsFilterer creates a new log filterer instance of Utils, bound to a specific deployed contract.
func NewUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*UtilsFilterer, error) {
	contract, err := bindUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UtilsFilterer{contract: contract}, nil
}

// bindUtils binds a generic wrapper to an already deployed contract.
func bindUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Utils *UtilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Utils.Contract.UtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Utils *UtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Utils.Contract.UtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Utils *UtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Utils.Contract.UtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Utils *UtilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Utils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Utils *UtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Utils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Utils *UtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Utils.Contract.contract.Transact(opts, method, params...)
}

// ZeroCopySinkABI is the input ABI used to generate the binding from.
const ZeroCopySinkABI = "[]"

// ZeroCopySinkBin is the compiled bytecode used for deploying new contracts.
var ZeroCopySinkBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72305820b888fe2faf222a1a6d1134fc03b43f795259f36656ea16f1d142834c844bfae764736f6c634300050a0032"

// DeployZeroCopySink deploys a new Ethereum contract, binding an instance of ZeroCopySink to it.
func DeployZeroCopySink(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZeroCopySink, error) {
	parsed, err := abi.JSON(strings.NewReader(ZeroCopySinkABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZeroCopySinkBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZeroCopySink{ZeroCopySinkCaller: ZeroCopySinkCaller{contract: contract}, ZeroCopySinkTransactor: ZeroCopySinkTransactor{contract: contract}, ZeroCopySinkFilterer: ZeroCopySinkFilterer{contract: contract}}, nil
}

// ZeroCopySink is an auto generated Go binding around an Ethereum contract.
type ZeroCopySink struct {
	ZeroCopySinkCaller     // Read-only binding to the contract
	ZeroCopySinkTransactor // Write-only binding to the contract
	ZeroCopySinkFilterer   // Log filterer for contract events
}

// ZeroCopySinkCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZeroCopySinkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroCopySinkTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZeroCopySinkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroCopySinkFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZeroCopySinkFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroCopySinkSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZeroCopySinkSession struct {
	Contract     *ZeroCopySink     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZeroCopySinkCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZeroCopySinkCallerSession struct {
	Contract *ZeroCopySinkCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ZeroCopySinkTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZeroCopySinkTransactorSession struct {
	Contract     *ZeroCopySinkTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ZeroCopySinkRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZeroCopySinkRaw struct {
	Contract *ZeroCopySink // Generic contract binding to access the raw methods on
}

// ZeroCopySinkCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZeroCopySinkCallerRaw struct {
	Contract *ZeroCopySinkCaller // Generic read-only contract binding to access the raw methods on
}

// ZeroCopySinkTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZeroCopySinkTransactorRaw struct {
	Contract *ZeroCopySinkTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZeroCopySink creates a new instance of ZeroCopySink, bound to a specific deployed contract.
func NewZeroCopySink(address common.Address, backend bind.ContractBackend) (*ZeroCopySink, error) {
	contract, err := bindZeroCopySink(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySink{ZeroCopySinkCaller: ZeroCopySinkCaller{contract: contract}, ZeroCopySinkTransactor: ZeroCopySinkTransactor{contract: contract}, ZeroCopySinkFilterer: ZeroCopySinkFilterer{contract: contract}}, nil
}

// NewZeroCopySinkCaller creates a new read-only instance of ZeroCopySink, bound to a specific deployed contract.
func NewZeroCopySinkCaller(address common.Address, caller bind.ContractCaller) (*ZeroCopySinkCaller, error) {
	contract, err := bindZeroCopySink(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySinkCaller{contract: contract}, nil
}

// NewZeroCopySinkTransactor creates a new write-only instance of ZeroCopySink, bound to a specific deployed contract.
func NewZeroCopySinkTransactor(address common.Address, transactor bind.ContractTransactor) (*ZeroCopySinkTransactor, error) {
	contract, err := bindZeroCopySink(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySinkTransactor{contract: contract}, nil
}

// NewZeroCopySinkFilterer creates a new log filterer instance of ZeroCopySink, bound to a specific deployed contract.
func NewZeroCopySinkFilterer(address common.Address, filterer bind.ContractFilterer) (*ZeroCopySinkFilterer, error) {
	contract, err := bindZeroCopySink(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySinkFilterer{contract: contract}, nil
}

// bindZeroCopySink binds a generic wrapper to an already deployed contract.
func bindZeroCopySink(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZeroCopySinkABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZeroCopySink *ZeroCopySinkRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZeroCopySink.Contract.ZeroCopySinkCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZeroCopySink *ZeroCopySinkRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZeroCopySink.Contract.ZeroCopySinkTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZeroCopySink *ZeroCopySinkRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZeroCopySink.Contract.ZeroCopySinkTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZeroCopySink *ZeroCopySinkCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZeroCopySink.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZeroCopySink *ZeroCopySinkTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZeroCopySink.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZeroCopySink *ZeroCopySinkTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZeroCopySink.Contract.contract.Transact(opts, method, params...)
}

// ZeroCopySourceABI is the input ABI used to generate the binding from.
const ZeroCopySourceABI = "[]"

// ZeroCopySourceBin is the compiled bytecode used for deploying new contracts.
var ZeroCopySourceBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723058203bd20457e88d7bf053774b7a7f897488c0a7dc52641f4f1e4a987dd592e2318b64736f6c634300050a0032"

// DeployZeroCopySource deploys a new Ethereum contract, binding an instance of ZeroCopySource to it.
func DeployZeroCopySource(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZeroCopySource, error) {
	parsed, err := abi.JSON(strings.NewReader(ZeroCopySourceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZeroCopySourceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZeroCopySource{ZeroCopySourceCaller: ZeroCopySourceCaller{contract: contract}, ZeroCopySourceTransactor: ZeroCopySourceTransactor{contract: contract}, ZeroCopySourceFilterer: ZeroCopySourceFilterer{contract: contract}}, nil
}

// ZeroCopySource is an auto generated Go binding around an Ethereum contract.
type ZeroCopySource struct {
	ZeroCopySourceCaller     // Read-only binding to the contract
	ZeroCopySourceTransactor // Write-only binding to the contract
	ZeroCopySourceFilterer   // Log filterer for contract events
}

// ZeroCopySourceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZeroCopySourceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroCopySourceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZeroCopySourceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroCopySourceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZeroCopySourceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroCopySourceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZeroCopySourceSession struct {
	Contract     *ZeroCopySource   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZeroCopySourceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZeroCopySourceCallerSession struct {
	Contract *ZeroCopySourceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ZeroCopySourceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZeroCopySourceTransactorSession struct {
	Contract     *ZeroCopySourceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ZeroCopySourceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZeroCopySourceRaw struct {
	Contract *ZeroCopySource // Generic contract binding to access the raw methods on
}

// ZeroCopySourceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZeroCopySourceCallerRaw struct {
	Contract *ZeroCopySourceCaller // Generic read-only contract binding to access the raw methods on
}

// ZeroCopySourceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZeroCopySourceTransactorRaw struct {
	Contract *ZeroCopySourceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZeroCopySource creates a new instance of ZeroCopySource, bound to a specific deployed contract.
func NewZeroCopySource(address common.Address, backend bind.ContractBackend) (*ZeroCopySource, error) {
	contract, err := bindZeroCopySource(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySource{ZeroCopySourceCaller: ZeroCopySourceCaller{contract: contract}, ZeroCopySourceTransactor: ZeroCopySourceTransactor{contract: contract}, ZeroCopySourceFilterer: ZeroCopySourceFilterer{contract: contract}}, nil
}

// NewZeroCopySourceCaller creates a new read-only instance of ZeroCopySource, bound to a specific deployed contract.
func NewZeroCopySourceCaller(address common.Address, caller bind.ContractCaller) (*ZeroCopySourceCaller, error) {
	contract, err := bindZeroCopySource(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySourceCaller{contract: contract}, nil
}

// NewZeroCopySourceTransactor creates a new write-only instance of ZeroCopySource, bound to a specific deployed contract.
func NewZeroCopySourceTransactor(address common.Address, transactor bind.ContractTransactor) (*ZeroCopySourceTransactor, error) {
	contract, err := bindZeroCopySource(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySourceTransactor{contract: contract}, nil
}

// NewZeroCopySourceFilterer creates a new log filterer instance of ZeroCopySource, bound to a specific deployed contract.
func NewZeroCopySourceFilterer(address common.Address, filterer bind.ContractFilterer) (*ZeroCopySourceFilterer, error) {
	contract, err := bindZeroCopySource(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySourceFilterer{contract: contract}, nil
}

// bindZeroCopySource binds a generic wrapper to an already deployed contract.
func bindZeroCopySource(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZeroCopySourceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZeroCopySource *ZeroCopySourceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZeroCopySource.Contract.ZeroCopySourceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZeroCopySource *ZeroCopySourceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZeroCopySource.Contract.ZeroCopySourceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZeroCopySource *ZeroCopySourceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZeroCopySource.Contract.ZeroCopySourceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZeroCopySource *ZeroCopySourceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZeroCopySource.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZeroCopySource *ZeroCopySourceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZeroCopySource.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZeroCopySource *ZeroCopySourceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZeroCopySource.Contract.contract.Transact(opts, method, params...)
}

