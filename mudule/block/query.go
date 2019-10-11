package block

import (
	"encoding/hex"
	"fmt"
	"github.com/QOSGroup/qbase/client/context"
	btypes "github.com/QOSGroup/qbase/types"
	"github.com/tendermint/go-amino"
	go_amino "github.com/tendermint/go-amino"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"time"
)

func QueryTx(cdc *amino.Codec, tx string) (result btypes.TxResponse, err error) {
	cliCtx := context.NewCLIContext().WithCodec(cdc)
	result, err = QueryTxInner(cliCtx, tx)
	if err != nil {
		return
	}

	if result.Empty() {
		err = fmt.Errorf("No transaction found with hash %s", tx)
		return
	}
	return
}

// QueryTx queries for a single transaction by a hash string in hex format. An
// error is returned if the transaction does not exist or cannot be queried.
func QueryTxInner(cliCtx context.CLIContext, hashHexStr string) (btypes.TxResponse, error) {
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
