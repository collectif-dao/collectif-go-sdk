package rpc

import (
	"collective-go-sdk/rpc/types"
	"context"
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/api"
	ltypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"
	"github.com/ipfs/go-cid"
	"github.com/ybbus/jsonrpc/v3"
)

type RPCClient struct {
	Client jsonrpc.RPCClient
}

func NewRPCClient(endpoint string, api_token string) *RPCClient {
	var cHeader map[string]string

	if api_token != "" {
		cHeader = map[string]string{
			"Authorization": "Bearer " + api_token,
		}
	}

	rpc := jsonrpc.NewClientWithOpts(endpoint, &jsonrpc.RPCClientOpts{
		CustomHeaders: cHeader,
	})

	return &RPCClient{
		Client: rpc,
	}
}

type filParams struct {
	Cid interface{} `json:"/"`
}

func newCidParam(cid interface{}) []interface{} {
	i := make([]interface{}, 0)

	p := &filParams{
		Cid: cid,
	}

	i = append(i, p)

	return i
}

func (r *RPCClient) Version(ctx context.Context) (*types.APIVersion, error) {
	i := make([]interface{}, 0)

	response, err := r.Client.Call(ctx, "Filecoin.Version", i)
	if err != nil {
		return nil, fmt.Errorf("failed to perform Filecoin.Version call: %w", err)
	}

	version := types.APIVersion{}

	err = response.GetObject(&version)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &version, nil
}

func (r *RPCClient) GetNonce(ctx context.Context, address address.Address) (uint64, error) {
	i := make([]interface{}, 0)
	i = append(i, address)

	response, err := r.Client.Call(ctx, "Filecoin.MpoolGetNonce", i)
	if err != nil {
		return 0, fmt.Errorf("failed to perform Filecoin.MpoolGetNonce call: %w", err)
	}

	value, err := response.GetInt()
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return uint64(value), nil
}

func (r *RPCClient) Call(ctx context.Context, msg *ethtypes.EthCall) (string, error) {
	i := make([]interface{}, 0)
	i = append(i, msg, "latest")

	response, err := r.Client.Call(ctx, "eth_call", i)
	if err != nil {
		return "", fmt.Errorf("failed perform eth_call: %w", err)
	}

	res, err := response.GetString()
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return res, nil
}

func (r *RPCClient) EstimateMessageGas(ctx context.Context, msg *ltypes.Message, spec *api.MessageSendSpec) (*ltypes.Message, error) {
	i := make([]interface{}, 0)

	tsk, err := r.GetTipSetKey(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to perform GetTipSetKey: %w", err)
	}

	i = append(i, msg, spec, tsk)

	response, err := r.Client.Call(ctx, "Filecoin.GasEstimateMessageGas", i)
	if err != nil {
		return nil, fmt.Errorf("failed to perform Filecoin.GasEstimateMessageGas call: %w", err)
	}

	if response.Error != nil {
		return nil, fmt.Errorf("Filecoin.GasEstimateMessageGas call failed with: %s", response.Error.Message)
	}

	rMsg := ltypes.Message{}
	err = response.GetObject(&rMsg)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &rMsg, nil
}

func (r *RPCClient) GasEstimateGasLimit(ctx context.Context, msg *ltypes.Message) (uint64, error) {
	i := make([]interface{}, 0)

	tsk, err := r.GetTipSetKey(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to perform GetTipSetKey: %w", err)
	}

	i = append(i, msg, tsk)

	response, err := r.Client.Call(ctx, "Filecoin.GasEstimateGasLimit", i)
	if err != nil {
		return 0, fmt.Errorf("failed to perform Filecoin.GasEstimateGasLimit call: %w", err)
	}

	gasLimit, err := response.GetInt()
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return uint64(gasLimit), nil
}

func (r *RPCClient) PushToMpool(ctx context.Context, sMsg *ltypes.SignedMessage) (cid.Cid, error) {
	i := make([]interface{}, 0)
	i = append(i, sMsg)

	response, err := r.Client.Call(ctx, "Filecoin.MpoolPush", i)
	if err != nil {
		return cid.Cid{}, fmt.Errorf("failed to perform Filecoin.MpoolPush call: %w", err)
	}

	c := cid.Cid{}
	err = response.GetObject(&c)
	if err != nil {
		return cid.Cid{}, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return c, nil
}

func (r *RPCClient) WaitMessage(ctx context.Context, cid *cid.Cid) (api.MsgLookup, error) {
	i := make([]interface{}, 0)
	i = append(i, cid, 0, nil, true)

	response, err := r.Client.Call(ctx, "Filecoin.StateWaitMsg", i)
	if err != nil {
		return api.MsgLookup{}, fmt.Errorf("failed to perform Filecoin.StateWaitMsg call: %w", err)
	}

	l := api.MsgLookup{}
	err = response.GetObject(&l)
	if err != nil {
		return api.MsgLookup{}, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return l, nil
}

func (r *RPCClient) VerifyCid(ctx context.Context, cid string) (bool, error) {
	i := newCidParam(cid)

	response, err := r.Client.Call(ctx, "Filecoin.ChainHasObj", i)
	if err != nil {
		return false, fmt.Errorf("failed to perform Filecoin.ChainHasObj call: %w", err)
	}

	value, err := response.GetBool()
	if err != nil {
		return false, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return value, nil
}

func (r *RPCClient) GetChainHead(ctx context.Context) (*ltypes.TipSet, error) {
	i := make([]interface{}, 0)

	response, err := r.Client.Call(ctx, "Filecoin.ChainHead", i)
	if err != nil {
		return nil, fmt.Errorf("failed to perform a ChainHead call: %w", err)
	}

	tipSet := ltypes.TipSet{}

	err = response.GetObject(&tipSet)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &tipSet, nil
}

func (r *RPCClient) GetTipSetKey(ctx context.Context) (*ltypes.TipSetKey, error) {
	tipSet, err := r.GetChainHead(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to perform a ChainHead call: %w", err)
	}

	key := tipSet.Key()

	return &key, nil
}

func (r *RPCClient) LookupId(ctx context.Context, addr address.Address, key *ltypes.TipSetKey) (*address.Address, error) {
	i := make([]interface{}, 0)
	i = append(i, addr, key)

	response, err := r.Client.Call(ctx, "Filecoin.StateLookupID", i)
	if err != nil {
		return nil, fmt.Errorf("failed to perform a Filecoin.StateLookupID call: %w", err)
	}

	fAddr := address.Address{}

	err = response.GetObject(&fAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &fAddr, nil
}

func (r *RPCClient) StateMinerInfo(ctx context.Context, minerAddress address.Address, key *ltypes.TipSetKey) (*api.MinerInfo, error) {
	i := make([]interface{}, 0)
	i = append(i, minerAddress, key)

	response, err := r.Client.Call(ctx, "Filecoin.StateMinerInfo", i)
	if err != nil {
		return nil, fmt.Errorf("failed to perform a Filecoin.StateMinerInfo call: %w", err)
	}

	info := api.MinerInfo{}
	err = response.GetObject(&info)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &info, nil
}
