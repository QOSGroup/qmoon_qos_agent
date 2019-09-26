package gov

import (
	bcli "github.com/QOSGroup/qbase/client"
	"github.com/QOSGroup/qbase/client/block"
	"github.com/QOSGroup/qbase/client/types"
	"github.com/QOSGroup/qos/app"
	"github.com/spf13/viper"
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
	txsCommands := bcli.TxCommand()
	app.ModuleBasics.AddTxCommands(txsCommands, cdc)
}

func TestQueryProposal(t *testing.T) {
	viper.Set(types.FlagNode, "39.97.234.227:26657")
	viper.Set(types.FlagNonceNode, "39.97.234.227:26657")
	viper.Set(types.FlagNonce, 0)

	Tout, err := QueryProposal(cdc, 1)
	if err != nil {
		t.Log(err)
		return
	}
	bytes, err := cdc.MarshalJSON(Tout)
	t.Log(string(bytes))
}
