package messagesigner

import (
	"bytes"
	"collective-go-sdk/config"
	"strconv"

	// "collective-go-sdk/fvm"
	"collective-go-sdk/keystore"
	"collective-go-sdk/rpc"
	"context"
	"testing"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin"
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

	rpc := rpc.RPCClient{}

	msigner := NewMessageSigner(*wallet, &rpc)
	assert.NotEmpty(t, msigner.Wallet.Get())
}

func TestNextNonce(t *testing.T) {
	config, err := config.LoadConfig("../")
	if err != nil {
		assert.Error(t, err)
	}

	cache := keystore.MemoryKeyStore
	ctx := context.Background()
	rpcClient := rpc.NewRPCClient(config.RPCConfig[config.DefaultNetwork].Address, config.RPCConfig[config.DefaultNetwork].APIToken)

	ks, err := cache.PrepareKeystore(config.FSKeyStoreDir)
	if err != nil {
		assert.Error(t, err)
	}

	w, err := wallet.NewWallet(ks)
	if err != nil {
		assert.Error(t, err)
	}

	ms := NewMessageSigner(*w, rpcClient)

	addr, err := ms.Wallet.WalletNew(ctx, "bls")
	if err != nil {
		assert.Error(t, err)
	}

	nonce, err := ms.NextNonce(ctx, addr)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, nonce, uint64(0))

	err = ms.Wallet.WalletDelete(ctx, addr)
	if err != nil {
		assert.Error(t, err)
	}

	status, err := ms.Wallet.WalletHas(ctx, addr)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, status, false)
}

func TestNextNonceForExistingAddress(t *testing.T) {
	config, err := config.LoadConfig("../")
	if err != nil {
		assert.Error(t, err)
	}

	cache := keystore.MemoryKeyStore
	ctx := context.Background()
	rpcClient := rpc.NewRPCClient(config.RPCConfig[config.DefaultNetwork].Address, config.RPCConfig[config.DefaultNetwork].APIToken)

	ks, err := cache.PrepareKeystore(config.FSKeyStoreDir)
	if err != nil {
		assert.Error(t, err)
	}

	w, err := wallet.NewWallet(ks)
	if err != nil {
		assert.Error(t, err)
	}

	ms := NewMessageSigner(*w, rpcClient)

	toAddr, err := address.NewFromString("t410fjtte454usorrzedcbgslfmiwg5yy6i2a2pxoyky")

	if err != nil {
		assert.Error(t, err)
	}

	nonce, err := ms.NextNonce(ctx, toAddr)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, nonce, uint64(0))
}

func TestSignMessage(t *testing.T) {
	msgValue, err := strconv.ParseInt("1000", 10, 8)
	toAddress := "t2indjldhwmnwro4cqiujhnwhva74ndoeagpn4nhi"

	config, err := config.LoadConfig("../")
	if err != nil {
		assert.Error(t, err)
	}

	ctx := context.Background()
	rpcClient := rpc.NewRPCClient(config.RPCConfig[config.DefaultNetwork].Address, config.RPCConfig[config.DefaultNetwork].APIToken)

	w, err := wallet.NewWallet(keystore.NewMemKeystore())
	if err != nil {
		assert.Error(t, err)
	}

	ms := NewMessageSigner(*w, rpcClient)

	addr, err := ms.Wallet.WalletNew(ctx, types.KTSecp256k1)
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
		Method: builtin.MethodsEVM.InvokeContract,
		Params: calldata,
		Nonce:  5,
	}

	signedMessage, err := ms.SignMessage(ctx, message, &api.MessageSendSpec{})
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

	sig, err := ms.Wallet.WalletSign(ctx, message.From, sb, api.MsgMeta{
		Type:  api.MTChainMsg,
		Extra: mb.RawData(),
	})

	assert.Equal(t, sig.Data, signedMessage.Signature.Data)
}

func TestWallet(t *testing.T) {
	ctx := context.Background()

	w1, err := wallet.NewWallet(keystore.NewMemKeystore())
	if err != nil {
		t.Fatal(err)
	}

	a1, err := w1.WalletNew(ctx, types.KTSecp256k1)
	if err != nil {
		t.Fatal(err)
	}

	exists, err := w1.WalletHas(ctx, a1)
	if err != nil {
		t.Fatal(err)
	}

	if !exists {
		t.Fatalf("address doesn't exist in wallet")
	}

	w2, err := wallet.NewWallet(keystore.NewMemKeystore())
	if err != nil {
		t.Fatal(err)
	}

	a2, err := w2.WalletNew(ctx, types.KTSecp256k1)
	if err != nil {
		t.Fatal(err)
	}

	a3, err := w2.WalletNew(ctx, types.KTSecp256k1)
	if err != nil {
		t.Fatal(err)
	}

	addrs, err := w2.WalletList(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if len(addrs) != 2 {
		t.Fatalf("wrong number of addresses in wallet")
	}

	err = w2.WalletDelete(ctx, a2)
	if err != nil {
		t.Fatal(err)
	}

	exists, err = w2.WalletHas(ctx, a2)
	if err != nil {
		t.Fatal(err)
	}
	if exists {
		t.Fatalf("failed to delete wallet address")
	}

	err = w2.SetDefault(a3)
	if err != nil {
		t.Fatal(err)
	}

	def, err := w2.GetDefault()
	if !assert.Equal(t, a3, def) {
		t.Fatal(err)
	}
}
