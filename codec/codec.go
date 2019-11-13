package codec

import (
	bcli "github.com/QOSGroup/qbase/client"
	"github.com/QOSGroup/qbase/client/block"
	"github.com/QOSGroup/qos/app"
	"github.com/tendermint/go-amino"
)

var Cdc *amino.Codec

func init() {
	Cdc = RegisterCodec()
}

func RegisterCodec() *amino.Codec {
	cdc := app.MakeCodec()
	// query commands
	queryCommands := bcli.QueryCommand(cdc)
	app.ModuleBasics.AddQueryCommands(queryCommands, cdc)
	queryCommands.AddCommand(block.BlockCommand(cdc)...)

	// txs commands
	txsCommands := bcli.TxCommand()
	app.ModuleBasics.AddTxCommands(txsCommands, cdc)
	return cdc
}
