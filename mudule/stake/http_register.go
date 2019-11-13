package stake

import (
	"github.com/QOSGroup/qbase/client/types"
	"github.com/QOSGroup/qmoon_qos_agent/codec"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func Register(engine *gin.Engine) {
	engine.GET("/stake/validators", queryValidators)
	engine.GET("/stake/validators/total/bond/tokens", queryTotalValidatorBondToken)
	engine.GET("/stake/validator/delegations", queryDelegationsWithValidator)
}

func queryValidators(ctx *gin.Context) {
	nodeUrl := ctx.Query("node_url")
	viper.Set(types.FlagNode, nodeUrl)
	viper.Set(types.FlagNonceNode, nodeUrl)

	height := ctx.Query("height")
	viper.Set(types.FlagHeight, height)
	viper.Set(types.FlagTrustNode, true)

	result, err := QueryValidators(codec.Cdc)
	log.Printf("res:%+v, err:%+v", result, err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func queryTotalValidatorBondToken(ctx *gin.Context) {
	nodeUrl := ctx.Query("node_url")
	viper.Set(types.FlagNode, nodeUrl)
	viper.Set(types.FlagNonceNode, nodeUrl)

	height := ctx.Query("height")
	viper.Set(types.FlagHeight, height)
	viper.Set(types.FlagTrustNode, true)

	result, err := QueryTotalValidatorBondToken(codec.Cdc)
	log.Printf("res:%+v, err:%+v", result, err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func queryDelegationsWithValidator(ctx *gin.Context) {
	nodeUrl := ctx.Query("node_url")
	viper.Set(types.FlagNode, nodeUrl)
	viper.Set(types.FlagNonceNode, nodeUrl)

	validatorAddr := ctx.Query("validator")

	result, err := QueryDelegationsWithValidator(codec.Cdc, validatorAddr)
	log.Printf("res:%+v, err:%+v", result, err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}
