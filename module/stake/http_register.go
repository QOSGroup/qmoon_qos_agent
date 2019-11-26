package stake

import (
	"github.com/QOSGroup/qmoon_qos_agent/codec"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func Register(engine *gin.Engine) {
	engine.GET("/stake/validators", queryValidators)
	engine.GET("/stake/validators/total/bond/tokens", queryTotalValidatorBondToken)
	engine.GET("/stake/validator/delegations", queryDelegationsWithValidator)
	engine.GET("/stake/delegator/delegations", queryDelegationsWithDelegator)
}

func queryValidators(ctx *gin.Context) {
	nodeUrl := ctx.Query("node_url")
	ipString:= nodeUrl[:strings.LastIndex(nodeUrl, ":")]
	ipString = strings.ReplaceAll(ipString, "http://", "")
	//
	//height, err:= strconv.ParseInt(ctx.Query("height"), 10, 64)
	//if err!=nil {
	//	ctx.JSON(http.StatusInternalServerError, err)
	//	return
	//}

	result, err := QueryValidators(codec.Cdc, ipString, 0)
	log.Printf("res:%+v, err:%+v", result, err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func queryDelegationsWithDelegator(ctx *gin.Context) {
	nodeUrl := ctx.Query("node_url")
	ipString:= nodeUrl[:strings.LastIndex(nodeUrl, ":")]
	ipString = strings.ReplaceAll(ipString, "http://", "")
	delegatorAddr := ctx.Query("delegator")

	result, err := QueryDelegationsWithDelegator(codec.Cdc, ipString, delegatorAddr)
	log.Printf("res:%+v, err:%+v", result, err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func queryTotalValidatorBondToken(ctx *gin.Context) {
	nodeUrl := ctx.Query("node_url")
	ipString:= nodeUrl[:strings.LastIndex(nodeUrl, ":")]
	ipString = strings.ReplaceAll(ipString, "http://", "")
	//height, err:= strconv.ParseInt(ctx.Query("height"), 10, 64)
	//if err!=nil {
	//	ctx.JSON(http.StatusInternalServerError, err)
	//	return
	//}
	result, err := QueryTotalValidatorBondToken(codec.Cdc, ipString, 0)
	log.Printf("res:%+v, err:%+v", result, err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func queryDelegationsWithValidator(ctx *gin.Context) {
	nodeUrl := ctx.Query("node_url")
	ipString:= nodeUrl[:strings.LastIndex(nodeUrl, ":")]
	ipString = strings.ReplaceAll(ipString, "http://", "")
	validatorAddr := ctx.Query("validator")

	result, err := QueryDelegationsWithValidator(codec.Cdc, ipString, validatorAddr)
	log.Printf("res:%+v, err:%+v", result, err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}
