package gov

import (
	"fmt"
	"github.com/QOSGroup/qmoon_qos_agent/codec"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
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
	ipString:= nodeUrl[:strings.LastIndex(nodeUrl, ":")]
	ipString = strings.ReplaceAll(ipString, "http://", "")

	pId, err := strconv.ParseInt(ctx.Query("pId"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	result, err := QueryProposal(codec.Cdc, ipString, pId)
	log.Printf("res:%+v, err:%+v", result, err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func queryProposals(ctx *gin.Context) {
	nodeUrl := ctx.Query("node_url")
	ipString:= nodeUrl[:strings.LastIndex(nodeUrl, ":")]
	ipString = strings.ReplaceAll(ipString, "http://", "")
	//statusStr := ctx.Query("status")
	result, err := QueryProposals(codec.Cdc, ipString)
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
	ipString:= nodeUrl[:strings.LastIndex(nodeUrl, ":")]
	ipString = strings.ReplaceAll(ipString, "http://", "")
	pId, err := strconv.ParseInt(ctx.Query("pId"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	result, err := QueryVotes(codec.Cdc, ipString, pId)
	log.Printf("res:%+v, err:%+v", result, err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func queryDeposits(ctx *gin.Context) {
	nodeUrl := ctx.Query("node_url")

	ipString:= nodeUrl[:strings.LastIndex(nodeUrl, ":")]
	ipString = strings.ReplaceAll(ipString, "http://", "")

	pId, err := strconv.ParseInt(ctx.Query("pId"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	result, err := QueryDeposits(codec.Cdc, ipString, pId)
	log.Printf("res:%+v, err:%+v", result, err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func queryTally(ctx *gin.Context) {
	nodeUrl := ctx.Query("node_url")
	ipString:= nodeUrl[:strings.LastIndex(nodeUrl, ":")]
	ipString = strings.ReplaceAll(ipString, "http://", "")

	pId, err := strconv.ParseInt(ctx.Query("pId"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	result, err := QueryTally(codec.Cdc, ipString, pId)
	log.Printf("res:%+v, err:%+v", result, err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}
