package mint

import (
	"github.com/QOSGroup/qmoon_qos_agent/codec"
	"testing"
)

func TestQueryInflationPhrases(t *testing.T) {
	Tout, err := QueryInflationPhrases(codec.Cdc, "47.103.79.28")
	if err != nil {
		t.Log(err)
		return
	}
	bytes, err := codec.Cdc.MarshalJSON(Tout)
	t.Log(string(bytes))
}

func TestQueryApplied(t *testing.T) {
	Tout, err := QueryApplied(codec.Cdc, "47.103.79.28")
	if err != nil {
		t.Log(err)
		return
	}
	bytes, err := codec.Cdc.MarshalJSON(Tout)
	t.Log(string(bytes))
}
