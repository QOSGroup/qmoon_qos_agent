package mint

import (
	"github.com/QOSGroup/qbase/client/types"
	"github.com/QOSGroup/qmoon_qos_agent/codec"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func Register(engine *gin.Engine) {
	engine.GET("/mint/inflation/phrases", queryInflationPhrases)
	engine.GET("/mint/total", queryTotal)
	engine.GET("/mint/Aapplied", queryApplied)
}

func queryInflationPhrases(ctx *gin.Context) {
	nodeUrl := ctx.Query("node_url")
	viper.Set(types.FlagNode, nodeUrl)
	viper.Set(types.FlagNonceNode, nodeUrl)

	result, err := QueryInflationPhrases(codec.Cdc)
	log.Printf("res:%+v, err:%+v", result, err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func queryTotal(ctx *gin.Context) {
	nodeUrl := ctx.Query("node_url")
	viper.Set(types.FlagNode, nodeUrl)
	viper.Set(types.FlagNonceNode, nodeUrl)

	result, err := QueryTotal(codec.Cdc)
	log.Printf("res:%+v, err:%+v", result, err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func queryApplied(ctx *gin.Context) {
	nodeUrl := ctx.Query("node_url")
	viper.Set(types.FlagNode, nodeUrl)
	viper.Set(types.FlagNonceNode, nodeUrl)

	result, err := QueryApplied(codec.Cdc)
	log.Printf("res:%+v, err:%+v", result, err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}
