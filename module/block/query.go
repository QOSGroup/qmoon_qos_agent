package block

import (
	"encoding/hex"
	"fmt"
	"github.com/QOSGroup/qbase/client/context"
	btypes "github.com/QOSGroup/qbase/types"
	"github.com/tendermint/go-amino"
	go_amino "github.com/tendermint/go-amino"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"github.com/QOSGroup/qmoon_qos_agent/types"
	"time"
)

type LatestHeight struct {
	latest_height int64 `json:"latest_height"`
}

func QueryStatus(cdc *amino.Codec, ip string) (*ctypes.SyncInfo, error) {
	cliCtx := context.NewCLIContext().WithCodec(cdc).WithNodeIP(ip)
	node, err := cliCtx.GetNode()
	if err != nil {
		return nil, err
	}
	status, err := node.Status()
	if err != nil {
		return nil, err
	}
	return &status.SyncInfo, nil
	//if err != nil{
	//	return &LatestHeight{}, err
	//}
	//return &LatestHeight{latest_height:status.SyncInfo.LatestBlockHeight}, nil
}

func QueryTx(cdc *amino.Codec, ip string, tx string) (result types.TxResponseResult, err error) {
	cliCtx := context.NewCLIContext().WithCodec(cdc).WithNodeIP(ip)

	result1, err := QueryTxInner(cliCtx, tx)
	if err != nil {
		return
	}

	if result1.Empty() {
		err = fmt.Errorf("No transaction found with hash %s", tx)
		return
	}

	result = types.TxResponseResult{
		Height: result1.Height,
		TxHash: result1.TxHash,
		Code: result1.Code,
		Data: result1.Data,
		RawLog: result1.RawLog,
		Info: result1.Info,
		GasWanted: result1.GasWanted,
		GasUsed: result1.GasUsed,
		Codespace: result1.Codespace,
		Tx: result1.Tx.Type(),
		Timestamp: result1.Timestamp,
	}

	return

	//var path = types.BuildQueryDelegationsByDelegatorCustomQueryPath(delegator)
	//
	//res, err := cliCtx.Query(path, []byte(""))
	//if err != nil {
	//	return
	//}
	//
	////var result []mapper.DelegationQueryResult
	//cliCtx.Codec.UnmarshalJSON(res, &result)
	//return
}

func QueryBlock(cdc *amino.Codec, ip string, height int64) (*ctypes.ResultBlock, error){
	cliCtx := context.NewCLIContext().WithCodec(cdc).WithNodeIP(ip)
	node, err := cliCtx.GetNode()
	if err != nil {
		return nil, err
	}

	resBlock, err := node.Block(&height)

	if err != nil {
		return nil, err
	}

	return resBlock, nil
}

// QueryTx queries for a single transaction by a hash string in hex format. An
// error is returned if the transaction does not exist or cannot be queried.
func QueryTxInner(cliCtx context.CLIContext,hashHexStr string) (btypes.TxResponse, error) {
	hash, err := hex.DecodeString(hashHexStr)
	if err != nil {
		return btypes.TxResponse{}, err
	}

	node, err := cliCtx.GetNode()
	if err != nil {
		return btypes.TxResponse{}, err
	}

	resTx, err := node.Tx(hash, !cliCtx.TrustNode)
	if err != nil {
		return btypes.TxResponse{}, err
	}

	resBlocks, err := getBlocksForTxResults(cliCtx, []*ctypes.ResultTx{resTx})
	if err != nil {
		return btypes.TxResponse{}, err
	}

	out, err := formatTxResult(cliCtx.Codec, resTx, resBlocks[resTx.Height])
	if err != nil {
		return out, err
	}

	return out, nil
}

func getBlocksForTxResults(cliCtx context.CLIContext, resTxs []*ctypes.ResultTx) (map[int64]*ctypes.ResultBlock, error) {
	node, err := cliCtx.GetNode()
	if err != nil {
		return nil, err
	}

	resBlocks := make(map[int64]*ctypes.ResultBlock)

	for _, resTx := range resTxs {
		if _, ok := resBlocks[resTx.Height]; !ok {
			resBlock, err := node.Block(&resTx.Height)
			if err != nil {
				return nil, err
			}

			resBlocks[resTx.Height] = resBlock
		}
	}

	return resBlocks, nil
}

func formatTxResult(cdc *go_amino.Codec, resTx *ctypes.ResultTx, resBlock *ctypes.ResultBlock) (btypes.TxResponse, error) {
	tx, err := parseTx(cdc, resTx.Tx)
	if err != nil {
		return btypes.TxResponse{}, err
	}

	return btypes.NewResponseResultTx(resTx, tx, resBlock.Block.Time.Format(time.RFC3339)), nil
}

func parseTx(cdc *go_amino.Codec, txBytes []byte) (btypes.Tx, error) {
	var tx btypes.Tx

	err := cdc.UnmarshalBinaryBare(txBytes, &tx)
	if err != nil {
		return nil, err
	}

	return tx, nil
}
