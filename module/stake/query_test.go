package stake

import (
	"github.com/QOSGroup/qmoon_qos_agent/codec"
	"testing"
)

func TestQueryValidators(t *testing.T) {
	Tout, err := QueryValidators(codec.Cdc, "47.103.79.28", 0)
	if err != nil {
		t.Log(err)
		return
	}
	bytes, err := codec.Cdc.MarshalJSON(Tout)
	t.Log(string(bytes))
}

func TestQueryTotalValidatorBondToken(t *testing.T) {

	Tout, err := QueryTotalValidatorBondToken(codec.Cdc, "47.103.79.28", 1000)
	if err != nil {
		t.Log(err)
		return
	}
	bytes, err := codec.Cdc.MarshalJSON(Tout)
	t.Log(string(bytes))
}

func TestQueryDelegationsWithValidator(t *testing.T) {
	Tout, err := QueryDelegationsWithValidator(codec.Cdc, "47.103.79.28", "qosval19hrl38w5lm6sklw2hzrzrjtsxudpy8hyfaea3e")
	if err != nil {
		t.Log(err)
		return
	}
	bytes, err := codec.Cdc.MarshalJSON(Tout)
	t.Log(string(bytes))
}
