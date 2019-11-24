package mint

import (
	"github.com/QOSGroup/qmoon_qos_agent/codec"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func Register(engine *gin.Engine) {
	engine.GET("/mint/inflation/phrases", queryInflationPhrases)
	engine.GET("/mint/total", queryTotal)
	engine.GET("/mint/applied", queryApplied)
}

func queryInflationPhrases(ctx *gin.Context) {
	nodeUrl := ctx.Query("node_url")

	ipString:= nodeUrl[:strings.LastIndex(nodeUrl, ":")]
	ipString = strings.ReplaceAll(ipString, "http://", "")

	result, err := QueryInflationPhrases(codec.Cdc, ipString)
	log.Printf("res:%+v, err:%+v", result, err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func queryTotal(ctx *gin.Context) {
	nodeUrl := ctx.Query("node_url")
	ipString:= nodeUrl[:strings.LastIndex(nodeUrl, ":")]
	ipString = strings.ReplaceAll(ipString, "http://", "")
	result, err := QueryTotal(codec.Cdc, ipString)
	log.Printf("res:%+v, err:%+v", result, err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func queryApplied(ctx *gin.Context) {
	nodeUrl := ctx.Query("node_url")
	ipString:= nodeUrl[:strings.LastIndex(nodeUrl, ":")]
	ipString = strings.ReplaceAll(ipString, "http://", "")
	result, err := QueryApplied(codec.Cdc, ipString)
	log.Printf("res:%+v, err:%+v", result, err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}
