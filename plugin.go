package main

import (
	"context"
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type Handler struct {
	lastBlockHeight int64
}

func (h *Handler) Handle(ctx sdk.Context, req *abci.RequestPrepareProposal) (*abci.ResponsePrepareProposal, error) {
	ctx.Logger().Info(fmt.Sprintf("Handling MEV for block %d with %d transactions. Last height is %d", req.Height, len(req.Txs), h.lastBlockHeight))
	h.lastBlockHeight = req.Height
	txRecords := make([]*abci.TxRecord, 0, len(req.Txs))
	for _, tx := range req.Txs {
		txRecords = append(txRecords, &abci.TxRecord{Action: abci.TxRecord_UNMODIFIED, Tx: tx})
	}
	return &abci.ResponsePrepareProposal{
		TxRecords: txRecords,
	}, nil
}

func (h *Handler) RPCSubmission(ctx context.Context, req json.RawMessage) (res json.RawMessage, err error) {
	mapReq := map[string]interface{}{}
	if err := json.Unmarshal(req, &mapReq); err != nil {
		return nil, err
	}
	fmt.Printf("MEV RPC submission received %s\n", mapReq)
	mapRes := map[string]interface{}{"success": true}
	return json.Marshal(mapRes)
}

var HandlerInstance = Handler{}
