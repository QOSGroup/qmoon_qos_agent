package codec

import (
	bcli "github.com/QOSGroup/qbase/client"
	"github.com/QOSGroup/qbase/client/block"

	atype "github.com/QOSGroup/qmoon_qos_agent/types"
	"github.com/QOSGroup/qos/app"
	"github.com/tendermint/go-amino"
)

var Cdc *amino.Codec

func init() {
	Cdc = RegisterCodec()
}

func RegisterCodec() *amino.Codec {
	cdc := app.MakeCodec()
	cdc.RegisterConcrete(&atype.ResultProposal{}, "gov/Proposal", nil)

	//app.RegisterCodec(cdc)
	//qos_cdc.RegisterCrypto(cdc)
	//qos_cdc.RegisterEvidences(cdc)
	//Cdc = cdc.Seal()

	// query commands
	queryCommands := bcli.QueryCommand(cdc)
	app.ModuleBasics.AddQueryCommands(queryCommands, cdc)
	queryCommands.AddCommand(block.BlockCommand(cdc)...)

	// txs commands
	txsCommands := bcli.TxCommand()
	app.ModuleBasics.AddTxCommands(txsCommands, cdc)
	return cdc
}
