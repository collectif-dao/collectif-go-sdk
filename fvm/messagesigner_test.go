package fvm

import (
	"collective-go-sdk/config"
	"collective-go-sdk/keystore"
	"context"
	"fmt"
	"testing"

	"github.com/filecoin-project/lotus/chain/wallet"
	"github.com/stretchr/testify/assert"
)

func TestNewMessageSigner(t *testing.T) {
	keystore := keystore.NewMemKeystore()

	wallet, err := wallet.NewWallet(keystore)
	if err != nil {
		assert.Error(t, err)
	}

	msigner := NewMessageSigner(*wallet)
	assert.NotEmpty(t, msigner.Wallet.Get())

}

func TestNextNonce(t *testing.T) {
	config, err := config.LoadConfig("../config")
	if err != nil {
		assert.Error(t, err)
	}

	ctx := context.Background()
	client, err := NewLotusClient(ctx, config, MemoryKeyStore)
	if err != nil {
		assert.Error(t, err)
	}

	addr, err := client.MessageSigner.Wallet.WalletNew(ctx, "bls")

	nonce, err := client.NextNonce(addr)
	if err != nil {
		assert.Error(t, err)
	}

	fmt.Println(nonce)
	assert.Equal(t, nonce, uint64(0))

	err = client.MessageSigner.Wallet.WalletDelete(ctx, addr)
	if err != nil {
		assert.Error(t, err)
	}

	status, err := client.MessageSigner.Wallet.WalletHas(ctx, addr)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, status, false)
}
