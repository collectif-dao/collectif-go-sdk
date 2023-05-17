package fvm

import (
	"collective-go-sdk/fvm/types"
	"context"
	"encoding/json"
	"fmt"

	"github.com/filecoin-project/go-address"
	ltypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/valyala/fasthttp"
	"golang.org/x/xerrors"
)

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

func (c *LotusClient) GetNonce(ctx context.Context, address address.Address) (uint64, error) {
	i := make([]interface{}, 0)
	i = append(i, address)

	response, _ := c.rpcClient2.Call(ctx, "Filecoin.MpoolGetNonce", i)

	value, err := response.GetInt()
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return uint64(value), nil
}

func (c *LotusClient) PushToMpool(msg *ltypes.SignedMessage) (cid.Cid, error) {
	i := make([]interface{}, 0)
	i = append(i, msg)

	res, err := c.HandleRequest("Filecoin.MpoolPush", i)
	if err != nil {
		return cid.Undef, xerrors.Errorf("Failed to marshal request body %v", err)
	}

	mcid, err := cid.Parse(res)
	if err != nil {
		return cid.Cid{}, xerrors.Errorf("Failed to parse cid of message %v", err)
	}

	return mcid, nil
}

func (c *LotusClient) VerifyCid(ctx context.Context, cid string) (bool, error) {
	i := newCidParam(cid)

	response, _ := c.rpcClient2.Call(ctx, "Filecoin.ChainHasObj", i)

	value, err := response.GetBool()
	if err != nil {
		return false, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return value, nil
}

func (c *LotusClient) HandleRequest(method string, params []interface{}) ([]byte, error) {
	requestBody := types.NewRPCRequestBody(method, params)

	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, xerrors.Errorf("Failed to marshal request body %v", err)
	}

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(c.rpcClient.Addr)
	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/json")
	req.SetBody(body)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err = fasthttp.Do(req, resp)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}

	return resp.Body(), nil
}
