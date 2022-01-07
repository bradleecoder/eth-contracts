abigen --sol ./contracts/core/cross_chain_manager/data/EthCrossChainData.sol --pkg eccd_abi > ./go_abi/eccd_abi/eccd_abi.go

abigen --sol ./contracts/core/cross_chain_manager/logic/EthCrossChainManager.sol --pkg eccm_abi > ./go_abi/eccm_abi/eccm_abi.go

abigen --sol ./contracts/core/cross_chain_manager/upgrade/EthCrossChainManagerProxy.sol --pkg eccmp_abi > ./go_abi/eccmp_abi/eccmp_abi.go

abigen --sol ./contracts/core/lock_proxy/LockProxy.sol --pkg lock_proxy_abi > ./go_abi/lock_proxy_abi/lock_proxy.go

abigen --sol ./contracts/core/assets/btc/BTCX.sol --pkg btcx_abi > ./go_abi/btcx_abi/btcx_abi.go

abigen --sol ./contracts/core/assets/ont/xONT.sol --pkg ontx_abi > ./go_abi/ontx_abi/ontx_abi.go

abigen --sol ./contracts/core/assets/ong/ONGX.sol --pkg ongx_abi > ./go_abi/ongx_abi/ongx_abi.go

abigen --sol ./contracts/core/assets/oep4_template/OEP4Template.sol --pkg oep4_abi > ./go_abi/oep4_abi/oep4_abi.go

abigen --sol ./contracts/core/assets/erc20_template/ERC20Template.sol --pkg erc20_abi > ./go_abi/erc20_abi/erc20_abi.go
