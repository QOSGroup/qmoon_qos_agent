package stake

import (
	"github.com/QOSGroup/qbase/client/types"
	"github.com/QOSGroup/qmoon_qos_agent/codec"
	"github.com/spf13/viper"
	"testing"
)

func TestQueryValidators(t *testing.T) {
	viper.Set(types.FlagNode, "47.103.79.28:26657")
	viper.Set(types.FlagNonceNode, "47.103.79.28:26657")
	viper.Set(types.FlagHeight, 0)
	viper.Set(types.FlagTrustNode, true)

	Tout, err := QueryValidators(codec.Cdc)
	if err != nil {
		t.Log(err)
		return
	}
	bytes, err := codec.Cdc.MarshalJSON(Tout)
	t.Log(string(bytes))
}

func TestQueryTotalValidatorBondToken(t *testing.T) {
	viper.Set(types.FlagNode, "39.97.234.227:26657")
	viper.Set(types.FlagNonceNode, "39.97.234.227:26657")
	viper.Set(types.FlagHeight, 0)
	viper.Set(types.FlagTrustNode, true)

	Tout, err := QueryTotalValidatorBondToken(codec.Cdc)
	if err != nil {
		t.Log(err)
		return
	}
	bytes, err := codec.Cdc.MarshalJSON(Tout)
	t.Log(string(bytes))
}

func TestQueryDelegationsWithValidator(t *testing.T) {
	viper.Set(types.FlagNode, "47.103.79.28:26657")
	viper.Set(types.FlagNonceNode, "47.103.79.28:26657")
	viper.Set(types.FlagHeight, 0)
	viper.Set(types.FlagTrustNode, true)

	Tout, err := QueryDelegationsWithValidator(codec.Cdc, "qosval19hrl38w5lm6sklw2hzrzrjtsxudpy8hyfaea3e")
	if err != nil {
		t.Log(err)
		return
	}
	bytes, err := codec.Cdc.MarshalJSON(Tout)
	t.Log(string(bytes))
}
