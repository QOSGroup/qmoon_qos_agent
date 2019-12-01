package block

import (
	"github.com/QOSGroup/qmoon_qos_agent/codec"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func Register(engine *gin.Engine) {
	engine.GET("/tx", queryTx)
	engine.GET("/block", queryBlock)
	engine.GET("/status", queryStatus)
}

func queryStatus(ctx *gin.Context) {
	nodeUrl := ctx.Query("node_url")
	ipString:= nodeUrl[:strings.LastIndex(nodeUrl, ":")]
	ipString = strings.ReplaceAll(ipString, "http://", "")
	result, err := QueryStatus(codec.Cdc, ipString)
	log.Printf("res:%+v, err:%+v", result, err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func queryTx(ctx *gin.Context) {
	nodeUrl := ctx.Query("node_url")
	ipString:= nodeUrl[:strings.LastIndex(nodeUrl, ":")]
	ipString = strings.ReplaceAll(ipString, "http://", "")
	tx := ctx.Query("hash")
	result, err := QueryTx(codec.Cdc, ipString, tx)
	log.Printf("res:%+v, err:%+v", result, err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func queryBlock(ctx *gin.Context) {
	nodeUrl := ctx.Query("node_url")
	ipString:= nodeUrl[:strings.LastIndex(nodeUrl, ":")]
	ipString = strings.ReplaceAll(ipString, "http://", "")
	if height, err := strconv.ParseInt(ctx.Query("height"), 10, 64); err == nil {
		result, err := QueryBlock(codec.Cdc, ipString, height)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, result)
	} else {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
}