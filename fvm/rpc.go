package fvm

import (
	"collective-go-sdk/fvm/types"
	"encoding/binary"
	"encoding/json"

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

func (c *LotusClient) GetNonce(address address.Address) (uint64, error) {
	i := make([]interface{}, 0)
	i = append(i, address)

	res, err := c.HandleRequest("Filecoin.MpoolGetNonce", i)
	if err != nil {
		return 0, xerrors.Errorf("Failed to marshal request body %v", err)
	}

	return binary.BigEndian.Uint64(res), nil
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

func (c *LotusClient) VerifyCid(cid string) ([]byte, error) {
	i := newCidParam(cid)

	res, err := c.HandleRequest("Filecoin.ChainHasObj", i)
	if err != nil {
		return nil, xerrors.Errorf("Failed to marshal request body %v", err)
	}

	return res, nil
}

func (c *LotusClient) HandleRequest(method string, params []interface{}) ([]byte, error) {
	req := fasthttp.AcquireRequest()

	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/json")

	requestBody := types.NewRPCRequestBody(method)
	requestBody.Params = append(requestBody.Params, params...)

	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, xerrors.Errorf("Failed to marshal request body %v", err)
	}

	req.AppendBody(body)

	resp := fasthttp.AcquireResponse()
	if err = c.rpcClient.Do(req, resp); err != nil {
		return nil, err
	}

	bodyBytes := resp.Body()

	return bodyBytes, nil
}
