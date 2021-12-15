
code-gen:
	~/go/bin/abigen --abi=factory/factory.abi --bin=factory/factory.bin --pkg=factory --out=factory/factory.go
	~/go/bin/abigen --abi=nftpositionmanager/nftpositionmanager.abi --bin=nftpositionmanager/nftpositionmanager.bin --pkg=nftpositionmanager --out=nftpositionmanager/nftpositionmanager.go
	~/go/bin/abigen --abi=router/router.abi --bin=router/router.bin --pkg=router --out=router/router.go
	~/go/bin/abigen --abi=weth/weth.abi --bin=weth/weth.bin --pkg=weth --out=weth/weth.go
	~/go/bin/abigen --abi=erc20/erc20.abi --pkg=erc20 --out=erc20/erc20.go