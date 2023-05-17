package fvm

import (
	"context"
	"sync"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	ltypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"
	lwallet "github.com/filecoin-project/lotus/chain/wallet"
	"golang.org/x/xerrors"
)

type MessageSigner struct {
	Wallet lwallet.LocalWallet
	lock   sync.Mutex
}

func NewMessageSigner(wallet lwallet.LocalWallet) *MessageSigner {
	return &MessageSigner{
		Wallet: wallet,
	}
}

func (c *LotusClient) NextNonce(ctx context.Context, addr address.Address) (uint64, error) {
	nonce, err := c.GetNonce(ctx, addr)
	if err != nil {
		return 0, xerrors.Errorf("failed to get nonce from mempool: %w", err)
	}

	return nonce, nil
}

func (c *LotusClient) SignMessage(ctx context.Context, msg *types.Message, spec *api.MessageSendSpec) (*types.SignedMessage, error) {
	c.MessageSigner.lock.Lock()
	defer c.MessageSigner.lock.Unlock()

	if msg.Nonce == 0 {
		nonce, err := c.NextNonce(ctx, msg.From)
		if err != nil {
			return nil, xerrors.Errorf("failed to create nonce: %w", err)
		}

		msg.Nonce = nonce
	}

	sb, err := SigningBytes(msg, msg.From.Protocol())
	if err != nil {
		return nil, err
	}

	mb, err := msg.ToStorageBlock()
	if err != nil {
		return nil, xerrors.Errorf("serializing message: %w", err)
	}

	sig, err := c.MessageSigner.Wallet.WalletSign(ctx, msg.From, sb, api.MsgMeta{
		Type:  api.MTChainMsg,
		Extra: mb.RawData(),
	})
	if err != nil {
		return nil, xerrors.Errorf("failed to sign message: %w, addr=%s", err, msg.From)
	}

	smsg := &types.SignedMessage{
		Message:   *msg,
		Signature: *sig,
	}

	return smsg, nil
}

func SigningBytes(msg *ltypes.Message, sigType address.Protocol) ([]byte, error) {
	if sigType == address.Delegated {
		txArgs, err := ethtypes.EthTxArgsFromUnsignedEthMessage(msg)
		if err != nil {
			return nil, xerrors.Errorf("failed to reconstruct eth transaction: %w", err)
		}
		rlpEncodedMsg, err := txArgs.ToRlpUnsignedMsg()
		if err != nil {
			return nil, xerrors.Errorf("failed to repack eth rlp message: %w", err)
		}
		return rlpEncodedMsg, nil
	}

	return msg.Cid().Bytes(), nil
}
