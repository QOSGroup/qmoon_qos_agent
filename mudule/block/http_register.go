package block

import (
	"github.com/QOSGroup/qbase/client/types"
	"github.com/QOSGroup/qmoon_qos_agent/codec"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"strconv"
)

func Register(engine *gin.Engine) {
	engine.GET("/block/tx", queryTx)
	engine.GET("/block", queryBlock)
	engine.GET("/status", queryStatus)
}

func queryStatus(ctx *gin.Context) {
	nodeUrl := ctx.Query("node_url")
	viper.Set(types.FlagNode, nodeUrl)
	viper.Set(types.FlagNonceNode, nodeUrl)

	result, err := QueryStatus(codec.Cdc)
	log.Printf("res:%+v, err:%+v", result, err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func queryTx(ctx *gin.Context) {
	nodeUrl := ctx.Query("node_url")
	viper.Set(types.FlagNode, nodeUrl)
	viper.Set(types.FlagNonceNode, nodeUrl)

	tx := ctx.Query("tx")
	result, err := QueryTx(codec.Cdc, tx)
	log.Printf("res:%+v, err:%+v", result, err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func queryBlock(ctx *gin.Context) {
	nodeUrl := ctx.Query("node_url")
	viper.Set(types.FlagNode, nodeUrl)
	viper.Set(types.FlagNonceNode, nodeUrl)

	if height, err := strconv.ParseInt(ctx.Query("height"), 10, 64); err == nil {
		result, err := QueryBlock(codec.Cdc, height)
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