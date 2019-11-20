package gov

import (
	"github.com/QOSGroup/qbase/client/types"
	"github.com/QOSGroup/qmoon_qos_agent/codec"
	"github.com/spf13/viper"
	"testing"
)

func TestQueryProposal(t *testing.T) {
	//time := time.Time{}
	//bytes, _ := codec.Cdc.MarshalJSON(time)
	//t.Log(string(bytes))

	viper.Set(types.FlagNode, "39.97.234.227:26657")
	viper.Set(types.FlagNonceNode, "39.97.234.227:26657")
	viper.Set(types.FlagNonce, 0)

	Tout, err := QueryProposal(codec.Cdc, 1)
	if err != nil {
		t.Log(err)
		return
	}
	bytes, err := codec.Cdc.MarshalJSON(Tout)
	t.Log(string(bytes))
}

func TestQueryProposals(t *testing.T) {
	viper.Set(types.FlagNode, "47.103.79.28:26657")
	viper.Set(types.FlagNonceNode, "47.103.79.28:26657")
	viper.Set(types.FlagNonce, 0)

	Tout, err := QueryProposals(codec.Cdc)
	if err != nil {
		t.Log(err)
		return
	}
	bytes, err := codec.Cdc.MarshalJSON(Tout)
	t.Log(string(bytes))
}

func TestQueryVotes(t *testing.T) {
	viper.Set(types.FlagNode, "39.97.234.227:26657")
	viper.Set(types.FlagNonceNode, "39.97.234.227:26657")
	viper.Set(types.FlagNonce, 0)

	Tout, err := QueryVotes(codec.Cdc, 1)
	if err != nil {
		t.Log(err)
		return
	}
	bytes, err := codec.Cdc.MarshalJSON(Tout)
	t.Log(string(bytes))
}

func TestQueryDeposits(t *testing.T) {
	viper.Set(types.FlagNode, "39.97.234.227:26657")
	viper.Set(types.FlagNonceNode, "39.97.234.227:26657")
	viper.Set(types.FlagNonce, 0)

	Tout, err := QueryDeposits(codec.Cdc, 1)
	if err != nil {
		t.Log(err)
		return
	}
	bytes, err := codec.Cdc.MarshalJSON(Tout)
	t.Log(string(bytes))
}

func TestQueryTally(t *testing.T) {
	viper.Set(types.FlagNode, "39.97.234.227:26657")
	viper.Set(types.FlagNonceNode, "39.97.234.227:26657")
	viper.Set(types.FlagNonce, 0)

	Tout, err := QueryTally(codec.Cdc, 1)
	if err != nil {
		t.Log(err)
		return
	}
	bytes, err := codec.Cdc.MarshalJSON(Tout)
	t.Log(string(bytes))
}
