package distribution

import (
	"github.com/QOSGroup/qbase/client/types"
	"github.com/QOSGroup/qmoon_qos_agent/codec"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func Register(engine *gin.Engine) {
	engine.GET("/distribution/community/fee/pool", queryCommunityFeePool)
}

func queryCommunityFeePool(ctx *gin.Context) {
	nodeUrl := ctx.Query("node_url")
	viper.Set(types.FlagNode, nodeUrl)
	viper.Set(types.FlagNonceNode, nodeUrl)

	result, err := QueryCommunityFeePool(codec.Cdc)
	log.Printf("res:%+v, err:%+v", result, err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}
