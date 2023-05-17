package types

type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type RPCRequestBody struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      string        `json:"id"`
}

type RPCResponseBody struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  []byte `json:"result"`
	Error   Error  `json:"error"`
	ID      int64  `json:"id"`
}

func NewRPCRequestBody(method string, i []interface{}) *RPCRequestBody {
	return &RPCRequestBody{
		Jsonrpc: "2.0",
		Method:  method,
		Params:  i,
		ID:      "0",
	}
}
