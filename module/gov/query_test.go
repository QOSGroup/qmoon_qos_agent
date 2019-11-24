package gov

import (
	"github.com/QOSGroup/qmoon_qos_agent/codec"
	"testing"
)

func TestQueryProposal(t *testing.T) {

	Tout, err := QueryProposal(codec.Cdc,"47.103.79.28", 1)
	if err != nil {
		t.Log(err)
		return
	}
	bytes, err := codec.Cdc.MarshalJSON(Tout)
	t.Log(string(bytes))
}

func TestQueryProposals(t *testing.T) {

	Tout, err := QueryProposals(codec.Cdc, "47.103.79.28")
	if err != nil {
		t.Log(err)
		return
	}
	bytes, err := codec.Cdc.MarshalJSON(Tout)
	t.Log(string(bytes))
}

func TestQueryVotes(t *testing.T) {

	Tout, err := QueryVotes(codec.Cdc, "47.103.79.28", 1)
	if err != nil {
		t.Log(err)
		return
	}
	bytes, err := codec.Cdc.MarshalJSON(Tout)
	t.Log(string(bytes))
}

func TestQueryDeposits(t *testing.T) {
	Tout, err := QueryDeposits(codec.Cdc,  "47.103.79.28",1)
	if err != nil {
		t.Log(err)
		return
	}
	bytes, err := codec.Cdc.MarshalJSON(Tout)
	t.Log(string(bytes))
}

func TestQueryTally(t *testing.T) {

	Tout, err := QueryTally(codec.Cdc, "47.103.79.28", 1)
	if err != nil {
		t.Log(err)
		return
	}
	bytes, err := codec.Cdc.MarshalJSON(Tout)
	t.Log(string(bytes))
}
