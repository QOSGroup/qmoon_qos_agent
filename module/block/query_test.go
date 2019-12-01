package block

import (
	"github.com/QOSGroup/qmoon_qos_agent/codec"
	"strings"
	"testing"
)

func TestQueryTx(t *testing.T) {
	nodeUrl:="http://47.103.79.28:26657"

	ipString:= nodeUrl[:strings.LastIndex(nodeUrl, ":")]
	ipString = strings.ReplaceAll(ipString, "http://", "")

	Tout, err := QueryTx(codec.Cdc, ipString, "C4CE4728774B67C064DF3274063A9035C95D0E92AD3073C717047B7A48486CBD")
	// Tout, err := QueryStatus(codec.Cdc, ipString)
	if err != nil {
		t.Log(err)
		return
	}
	bytes, err := codec.Cdc.MarshalJSON(Tout)
	t.Log(string(bytes))
}
