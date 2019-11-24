package distribution

import (
	"github.com/QOSGroup/qmoon_qos_agent/codec"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func Register(engine *gin.Engine) {
	engine.GET("/distribution/community/fee/pool", queryCommunityFeePool)
}

func queryCommunityFeePool(ctx *gin.Context) {
	nodeUrl := ctx.Query("node_url")
	ipString:= nodeUrl[:strings.LastIndex(nodeUrl, ":")]
	ipString = strings.ReplaceAll(ipString, "http://", "")

	result, err := QueryCommunityFeePool(codec.Cdc, ipString)
	log.Printf("res:%+v, err:%+v", result, err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}
