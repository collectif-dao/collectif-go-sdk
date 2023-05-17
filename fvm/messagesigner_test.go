package fvm

import (
	"bytes"
	"collective-go-sdk/config"
	"collective-go-sdk/keystore"
	"context"
	"strconv"
	"testing"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	builtintypes "github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"
	"github.com/filecoin-project/lotus/chain/wallet"
	"github.com/stretchr/testify/assert"
	cbg "github.com/whyrusleeping/cbor-gen"
)

var (
	sampleCalldata = "b4699eb70000000000000000000000000000000000000000000000000000000000000080000000000000000000000000c740192cd9d6dbff71f59ad2ee25dfbee24c77b3000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000f424000000000000000000000000000000000000000000000000000000000000000020064000000000000000000000000000000000000000000000000000000000000"
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

	nonce, err := client.NextNonce(ctx, addr)
	if err != nil {
		assert.Error(t, err)
	}

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

func TestNextNonceForExistingAddress(t *testing.T) {
	config, err := config.LoadConfig("../config")
	if err != nil {
		assert.Error(t, err)
	}

	ctx := context.Background()
	client, err := NewLotusClient(ctx, config, MemoryKeyStore)
	if err != nil {
		assert.Error(t, err)
	}

	toAddr, err := address.NewFromString("t410fjtte454usorrzedcbgslfmiwg5yy6i2a2pxoyky")

	if err != nil {
		assert.Error(t, err)
	}

	nonce, err := client.NextNonce(ctx, toAddr)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, nonce, uint64(5))
}

func TestSignMessage(t *testing.T) {
	msgValue, err := strconv.ParseInt("1000", 10, 8)
	toAddress := "t2indjldhwmnwro4cqiujhnwhva74ndoeagpn4nhi"

	config, err := config.LoadConfig("../config")
	if err != nil {
		assert.Error(t, err)
	}

	ctx := context.Background()
	client, err := NewLotusClient(ctx, config, MemoryKeyStore)
	if err != nil {
		assert.Error(t, err)
	}

	wallet := client.MessageSigner.Wallet

	addr, err := wallet.GetDefault()
	if err != nil {
		assert.Error(t, err)
	}

	toAddr, err := address.NewFromString(toAddress)
	if err != nil {
		assert.Error(t, err)
	}

	var calldata []byte
	calldata, err = ethtypes.DecodeHexStringTrimSpace(sampleCalldata)
	if err != nil {
		assert.Error(t, err)
	}

	var buffer bytes.Buffer
	if err := cbg.WriteByteArray(&buffer, calldata); err != nil {
		assert.Error(t, err)
	}
	calldata = buffer.Bytes()

	val := abi.NewTokenAmount(msgValue)

	message := &types.Message{
		From:   addr,
		To:     toAddr,
		Value:  val,
		Method: builtintypes.MethodsEVM.InvokeContract,
		Params: calldata,
		Nonce:  5,
	}

	signedMessage, err := client.SignMessage(ctx, message, &api.MessageSendSpec{})
	if err != nil {
		assert.Error(t, err)
	}

	sb, err := SigningBytes(message, message.From.Protocol())
	if err != nil {
		assert.Error(t, err)
	}

	mb, err := message.ToStorageBlock()
	if err != nil {
		assert.Error(t, err)
	}

	sig, err := wallet.WalletSign(ctx, message.From, sb, api.MsgMeta{
		Type:  api.MTChainMsg,
		Extra: mb.RawData(),
	})

	assert.Equal(t, sig.Data, signedMessage.Signature.Data)
}
