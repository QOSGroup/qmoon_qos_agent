package block

import (
	"github.com/QOSGroup/qbase/client/types"
	"github.com/QOSGroup/qmoon_qos_agent/codec"
	"github.com/spf13/viper"
	"testing"
)

func TestQueryTx(t *testing.T) {
	viper.Set(types.FlagNode, "47.103.79.28:26657")
	viper.Set(types.FlagNonceNode, "47.103.79.28:26657")
	viper.Set(types.FlagNonce, 0)

	//Tout, err := QueryTx(codec.Cdc, "443259A455F99566C901CD9A10F9541D26F8EED70B7480BD4F2312EC637A2875")
	Tout, err := QueryStatus(codec.Cdc)
	if err != nil {
		t.Log(err)
		return
	}
	bytes, err := codec.Cdc.MarshalJSON(Tout)
	t.Log(string(bytes))
}
