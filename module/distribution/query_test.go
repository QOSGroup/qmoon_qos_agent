package distribution

import (
	bcli "github.com/QOSGroup/qbase/client"
	"github.com/QOSGroup/qbase/client/block"
	"github.com/QOSGroup/qos/app"
	"github.com/tendermint/go-amino"
	"testing"
)

var (
	cdc *amino.Codec
)

func init() {
	cdc = app.MakeCodec()
	// query commands
	queryCommands := bcli.QueryCommand(cdc)
	app.ModuleBasics.AddQueryCommands(queryCommands, cdc)
	queryCommands.AddCommand(block.BlockCommand(cdc)...)

	// txs commands
	//txsCommands := bcli.TxCommand()
	//app.ModuleBasics.AddTxCommands(txsCommands, cdc)
}

func TestQueryCommunityFeePool(t *testing.T) {
	Tout, err := QueryCommunityFeePool(cdc, "47.103.79.28")
	if err != nil {
		t.Log(err)
		return
	}
	bytes, err := cdc.MarshalJSON(Tout)
	t.Log(string(bytes))
}
