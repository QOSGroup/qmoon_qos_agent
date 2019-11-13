package mint

import (
	"errors"
	"github.com/QOSGroup/qbase/client/context"
	btypes "github.com/QOSGroup/qbase/types"
	"github.com/QOSGroup/qos/module/mint/mapper"
	"github.com/QOSGroup/qos/module/mint/types"
	"github.com/tendermint/go-amino"
)

func QueryInflationPhrases(cdc *amino.Codec) (result []types.InflationPhrase, err error) {
	cliCtx := context.NewCLIContext().WithCodec(cdc)

	path := mapper.BuildQueryPhrasesPath()
	res, err := cliCtx.Query(path, []byte{})
	if err != nil {
		return
	}
	if len(res) == 0 {
		err = errors.New("no result found")
		return
	}

	err = cdc.UnmarshalJSON(res, &result)
	return
}

func QueryTotal(cdc *amino.Codec) (result btypes.BigInt, err error) {
	cliCtx := context.NewCLIContext().WithCodec(cdc)

	path := mapper.BuildQueryTotalPath()
	res, err := cliCtx.Query(path, []byte{})
	if err != nil {
		return
	}
	if len(res) == 0 {
		err = errors.New("no result found")
		return
	}

	err = cdc.UnmarshalJSON(res, &result)
	return
}

func QueryApplied(cdc *amino.Codec) (result btypes.BigInt, err error) {
	cliCtx := context.NewCLIContext().WithCodec(cdc)

	path := mapper.BuildQueryAppliedPath()
	res, err := cliCtx.Query(path, []byte{})
	if err != nil {
		return
	}
	if len(res) == 0 {
		err = errors.New("no result found")
		return
	}

	err = cdc.UnmarshalJSON(res, &result)
	return
}
