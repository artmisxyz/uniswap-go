package weth

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"gotest.tools/assert"
	"math/big"
	"testing"
)

func TestDeployWeth(t *testing.T) {
	key, _ := crypto.GenerateKey()
	auth, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
	if err != nil {
		t.Fatal(err)
	}
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1_000_000_000_000_000_000)}
	blockchain := backends.NewSimulatedBackend(alloc, 30_000_000)

	address, _, weth, err := DeployWeth(auth, blockchain)
	if err != nil {
		panic(err)
	}
	blockchain.Commit()

	if len(address.Bytes()) == 0 {
		t.Error("Expected a valid deployment address. Received empty address byte array instead")
	}
	callOpts := &bind.CallOpts{
		From:    common.Address{},
		Context: context.Background(),
	}
	symbol, err := weth.Symbol(callOpts)
	assert.Equal(t, symbol, "WETH")
}
