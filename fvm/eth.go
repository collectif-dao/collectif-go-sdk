package fvm

import (
	"bytes"
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/filecoin-project/go-address"
	fBig "github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	lTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"
	"github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
)

type MessageResponse struct {
	Data    []byte
	Message cid.Cid
	Receipt types.MessageReceipt
}

func (c *LotusClient) performEthCall(ctx context.Context, from *ethtypes.EthAddress, target *ethtypes.EthAddress, method string, calldata []byte, abi *abi.ABI, res interface{}) error {
	msg := &ethtypes.EthCall{
		From: from,
		To:   target,
		Data: calldata,
	}

	data, err := c.messageCall(ctx, msg)
	if err != nil {
		return err
	}

	return abi.UnpackIntoInterface(res, method, data)
}

func (c *LotusClient) messageCall(ctx context.Context, msg *ethtypes.EthCall) ([]byte, error) {
	output, err := c.RPCClient.Call(ctx, msg)
	if err != nil {
		return nil, err
	}

	data, err := ethtypes.DecodeHexString(output)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *LotusClient) performLotusMessage(ctx context.Context, target *address.Address, method string, value *big.Int, calldata []byte, send bool) (*MessageResponse, error) {
	from := *c.Address
	var res MessageResponse
	res.Data = calldata

	nonce, err := c.MessageSigner.NextNonce(ctx, from)
	if err != nil {
		return &res, fmt.Errorf("failed to get the next nonce: %w", err)
	}

	v, err := fBig.FromString(value.String())
	if err != nil {
		return nil, err
	}

	msg := &lTypes.Message{
		To:     *target,
		From:   from,
		Nonce:  nonce,
		Value:  v,
		Method: builtin.MethodsEVM.InvokeContract,
		Params: calldata,
	}

	res, err = c.executeNativeMessage(ctx, msg, res, send)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

func (c *LotusClient) executeNativeMessage(ctx context.Context, msg *lTypes.Message, res MessageResponse, send bool) (MessageResponse, error) {
	fullMsg, err := c.RPCClient.EstimateMessageGas(ctx, msg, &api.MessageSendSpec{})
	if err != nil {
		return res, err
	}

	signedMessage, err := c.MessageSigner.SignMessage(ctx, fullMsg, &api.MessageSendSpec{})
	if err != nil {
		return res, err
	}

	if send {
		cid, err := c.RPCClient.PushToMpool(ctx, signedMessage)
		if err != nil {
			return res, err
		}

		l, err := c.RPCClient.WaitMessage(ctx, &cid)
		if err != nil {
			return res, err
		}

		res.Message = l.Message
		res.Receipt = l.Receipt
	}

	return res, nil

}

func (c *LotusClient) calculateCalldata(method string, abi *abi.ABI, args ...interface{}) ([]byte, error) {
	callData, err := c.calculateEthCalldata(method, abi, args...)
	if err != nil {
		return nil, err
	}

	var buffer bytes.Buffer
	if err := cbg.WriteByteArray(&buffer, callData); err != nil {
		return nil, err
	}
	callData = buffer.Bytes()

	return callData, nil
}

func (c *LotusClient) calculateEthCalldata(method string, abi *abi.ABI, args ...interface{}) ([]byte, error) {
	callData, err := abi.Pack(method, args...)
	if err != nil {
		return nil, err
	}

	return callData, nil
}
