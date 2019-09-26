// Copyright 2018 The QOS Authors

package cmd

import (
	bcli "github.com/QOSGroup/qbase/client"
	"github.com/QOSGroup/qbase/client/block"
	"github.com/QOSGroup/qbase/client/types"
	"github.com/QOSGroup/qmoon_qos_agent/mudule/gov"
	"github.com/QOSGroup/qos/app"
	"github.com/spf13/viper"
	"github.com/tendermint/go-amino"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// ServerCmd qmoon http server
var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "http server",
	RunE:  server,
}

var (
	laddr string
	cdc   *amino.Codec
)

func init() {
	ServerCmd.PersistentFlags().StringVar(&laddr, "laddr", "0.0.0.0:19527", "listen addr")
	cdc = RegisterCodec()
}

func RegisterCodec() *amino.Codec {
	cdc = app.MakeCodec()
	// query commands
	queryCommands := bcli.QueryCommand(cdc)
	app.ModuleBasics.AddQueryCommands(queryCommands, cdc)
	queryCommands.AddCommand(block.BlockCommand(cdc)...)

	// txs commands
	txsCommands := bcli.TxCommand()
	app.ModuleBasics.AddTxCommands(txsCommands, cdc)
	return cdc
}

type TxQuery struct {
	Txs []string `json:"txs"`
}

func server(cmd *cobra.Command, args []string) error {
	r := gin.Default()

	//r.GET("/tx", nil)
	//r.POST("/txs", nil)
	r.GET("/gov/proposal", queryProposal)
	r.GET("/gov/proposals", queryProposals)
	r.GET("/gov/vote", queryVote)
	r.GET("/gov/votes", queryVotes)
	r.GET("/gov/deposit", queryDeposit)
	r.GET("/gov/deposit", queryDeposits)
	r.GET("/gov/tally", queryTally)
	return r.Run(laddr)
}

func queryProposal(ctx *gin.Context) {
	remote := ctx.Query("remote")
	viper.Set(types.FlagNode, remote)
	viper.Set(types.FlagNonceNode, remote)
	viper.Set(types.FlagNonce, 0)

	pId := ctx.GetInt("pId")
	result, err := gov.QueryProposal(cdc, int64(pId))
	log.Printf("res:%+v, err:%+v", result, err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func queryProposals(ctx *gin.Context) {
	remote := ctx.Query("remote")
	viper.Set(types.FlagNode, remote)
	viper.Set(types.FlagNonceNode, remote)
	viper.Set(types.FlagNonce, 0)

	limit := ctx.GetInt64("pId")
	depositor := ctx.Query("pId")
	voter := ctx.Query("pId")
	statusStr := ctx.Query("pId")
	result, err := gov.QueryProposals(cdc, limit, depositor, voter, statusStr)
	log.Printf("res:%+v, err:%+v", result, err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func queryVote(ctx *gin.Context) {
	remote := ctx.Query("remote")
	viper.Set(types.FlagNode, remote)
	viper.Set(types.FlagNonceNode, remote)
	viper.Set(types.FlagNonce, 0)

	pId := ctx.GetInt64("pId")
	addrStr := ctx.Query("pId")
	result, err := gov.QueryVote(cdc, pId, addrStr)
	log.Printf("res:%+v, err:%+v", result, err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func queryVotes(ctx *gin.Context) {
	remote := ctx.Query("remote")
	viper.Set(types.FlagNode, remote)
	viper.Set(types.FlagNonceNode, remote)
	viper.Set(types.FlagNonce, 0)

	pId := ctx.GetInt64("pId")
	result, err := gov.QueryVotes(cdc, pId)
	log.Printf("res:%+v, err:%+v", result, err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func queryDeposit(ctx *gin.Context) {
	remote := ctx.Query("remote")
	viper.Set(types.FlagNode, remote)
	viper.Set(types.FlagNonceNode, remote)
	viper.Set(types.FlagNonce, 0)

	pId := ctx.GetInt64("pId")
	addrStr := ctx.Query("pId")
	result, err := gov.QueryDeposit(cdc, pId, addrStr)
	log.Printf("res:%+v, err:%+v", result, err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func queryDeposits(ctx *gin.Context) {
	remote := ctx.Query("remote")
	viper.Set(types.FlagNode, remote)
	viper.Set(types.FlagNonceNode, remote)
	viper.Set(types.FlagNonce, 0)

	pId := ctx.GetInt64("pId")
	result, err := gov.QueryDeposits(cdc, pId)
	log.Printf("res:%+v, err:%+v", result, err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func queryTally(ctx *gin.Context) {
	remote := ctx.Query("remote")
	viper.Set(types.FlagNode, remote)
	viper.Set(types.FlagNonceNode, remote)
	viper.Set(types.FlagNonce, 0)

	pId := ctx.GetInt64("pId")
	addrStr := ctx.Query("pId")
	result, err := gov.QueryTally(cdc, pId, addrStr)
	log.Printf("res:%+v, err:%+v", result, err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}