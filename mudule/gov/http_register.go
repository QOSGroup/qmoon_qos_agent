package gov

import (
	"fmt"
	"github.com/QOSGroup/qbase/client/types"
	"github.com/QOSGroup/qmoon_qos_agent/codec"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"strconv"
)

func Register(engine *gin.Engine) {
	engine.GET("/gov/proposal", queryProposal)
	engine.GET("/gov/proposals", queryProposals)
	engine.GET("/gov/votes", queryVotes)
	engine.GET("/gov/deposits", queryDeposits)
	engine.GET("/gov/tally", queryTally)
}

func queryProposal(ctx *gin.Context) {
	nodeUrl := ctx.Query("node_url")
	viper.Set(types.FlagNode, nodeUrl)
	viper.Set(types.FlagNonceNode, nodeUrl)

	pId, err := strconv.ParseInt(ctx.Query("pId"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	result, err := QueryProposal(codec.Cdc, pId)
	log.Printf("res:%+v, err:%+v", result, err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func queryProposals(ctx *gin.Context) {
	nodeUrl := ctx.Query("node_url")
	viper.Set(types.FlagNode, nodeUrl)
	viper.Set(types.FlagNonceNode, nodeUrl)

	//statusStr := ctx.Query("status")
	result, err := QueryProposals(codec.Cdc)
	log.Printf("res:%+v, err:%+v", result, err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	bytes, err := codec.Cdc.MarshalJSON(result)
	fmt.Println(string(bytes))

	ctx.JSON(http.StatusOK, result)
}

func queryVotes(ctx *gin.Context) {
	nodeUrl := ctx.Query("node_url")
	viper.Set(types.FlagNode, nodeUrl)
	viper.Set(types.FlagNonceNode, nodeUrl)

	pId, err := strconv.ParseInt(ctx.Query("pId"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	result, err := QueryVotes(codec.Cdc, pId)
	log.Printf("res:%+v, err:%+v", result, err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func queryDeposits(ctx *gin.Context) {
	nodeUrl := ctx.Query("node_url")
	viper.Set(types.FlagNode, nodeUrl)
	viper.Set(types.FlagNonceNode, nodeUrl)

	pId, err := strconv.ParseInt(ctx.Query("pId"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	result, err := QueryDeposits(codec.Cdc, pId)
	log.Printf("res:%+v, err:%+v", result, err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func queryTally(ctx *gin.Context) {
	nodeUrl := ctx.Query("node_url")
	viper.Set(types.FlagNode, nodeUrl)
	viper.Set(types.FlagNonceNode, nodeUrl)

	pId, err := strconv.ParseInt(ctx.Query("pId"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	result, err := QueryTally(codec.Cdc, pId)
	log.Printf("res:%+v, err:%+v", result, err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}
