package block

import (
	"github.com/QOSGroup/qmoon_qos_agent/codec"
	"strings"
	"testing"
)

func TestQueryTx(t *testing.T) {
	nodeUrl:="47.103.79.28:26657"

	ipString:= nodeUrl[:strings.LastIndex(nodeUrl, ":")]
	strings.ReplaceAll(ipString, "http://", "")

	//Tout, err := QueryTx(codec.Cdc, "443259A455F99566C901CD9A10F9541D26F8EED70B7480BD4F2312EC637A2875")
	Tout, err := QueryStatus(codec.Cdc, ipString)
	if err != nil {
		t.Log(err)
		return
	}
	bytes, err := codec.Cdc.MarshalJSON(Tout)
	t.Log(string(bytes))
}
