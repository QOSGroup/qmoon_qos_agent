package distribution

import (
	"fmt"
	"github.com/QOSGroup/qbase/client/context"
	btypes "github.com/QOSGroup/qbase/types"
	"github.com/QOSGroup/qos/module/distribution/types"
	"github.com/tendermint/go-amino"
)

//func QueryValidatorPeriod(cdc *amino.Codec, validatorAddr string) (result mapper.ValidatorPeriodInfoQueryResult, err error) {
//	cliCtx := context.NewCLIContext().WithCodec(cdc)
//
//	var validator btypes.ValAddress
//	if o, err := qcliacc.GetValidatorAddrFromValue(cliCtx, validatorAddr); err == nil {
//		validator = o
//	}
//
//	path := types.BuildQueryValidatorPeriodInfoCustomQueryPath(validator)
//	res, err := cliCtx.Query(path, []byte(""))
//	if err != nil {
//		return
//	}
//
//	cliCtx.Codec.UnmarshalJSON(res, &result)
//	return
//}

func QueryCommunityFeePool(cdc *amino.Codec, ip string) (result btypes.BigInt, err error) {
	cliCtx := context.NewCLIContext().WithCodec(cdc).WithNodeIP(ip)

	res, err := cliCtx.Query(fmt.Sprintf("/store/%s/key", types.MapperName), types.BuildCommunityFeePoolKey())
	if err != nil {
		return
	}

	cdc.MustUnmarshalBinaryBare(res, &result)
	return
}
