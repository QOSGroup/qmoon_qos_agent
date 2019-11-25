package stake

import (
	"errors"
	qcliacc "github.com/QOSGroup/qbase/client/account"
	"github.com/QOSGroup/qbase/client/context"
	bctypes "github.com/QOSGroup/qbase/client/types"
	"github.com/QOSGroup/qbase/store"
	btypes "github.com/QOSGroup/qbase/types"
	"github.com/QOSGroup/qos/module/stake/mapper"
	"github.com/QOSGroup/qos/module/stake/types"
	"github.com/spf13/viper"
	"github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/rpc/client"
	"time"
)

const (
	activeDesc   = "active"
	inactiveDesc = "inactive"

	inactiveRevokeDesc        = "Revoked"
	inactiveMissVoteBlockDesc = "Kicked"
	inactiveMaxValidatorDesc  = "Replaced"
	inactiveDoubleDesc        = "DoubleSign"
)

type validatorDisplayInfo struct {
	OperatorAddress btypes.ValAddress  `json:"validator"`
	Owner           btypes.AccAddress  `json:"owner"`
	ConsAddress     btypes.ConsAddress `json:"consensusAddress"`
	ConsPubKey      string             `json:"consensusPubKey"`
	BondTokens      btypes.BigInt      `json:"bondTokens"`
	Description     types.Description  `json:"description"`
	Commission      types.Commission   `json:"commission"`

	Status         string    `json:"status"`
	InactiveDesc   string    `json:"InactiveDesc"`
	InactiveTime   time.Time `json:"inactiveTime"`
	InactiveHeight int64     `json:"inactiveHeight"`

	MinPeriod  int64         `json:"minPeriod"`
	BondHeight int64         `json:"bondHeight"`
	SelfBond   btypes.BigInt `json:"selfBond"`
}

func toValidatorDisplayInfo(validator types.Validator) validatorDisplayInfo {

	consPubKey, _ := btypes.ConsensusPubKeyString(validator.ConsPubKey)

	info := validatorDisplayInfo{
		OperatorAddress: validator.OperatorAddress,
		Owner:           validator.Owner,
		ConsAddress:     validator.ConsAddress(),
		ConsPubKey:      consPubKey,
		BondTokens:      validator.BondTokens,
		Description:     validator.Description,
		InactiveTime:    validator.InactiveTime,
		InactiveHeight:  validator.InactiveHeight,
		MinPeriod:       validator.MinPeriod,
		BondHeight:      validator.BondHeight,
		Commission:      validator.Commission,
	}

	if validator.Status == types.Active {
		info.Status = activeDesc
	} else {
		info.Status = inactiveDesc
	}

	if validator.InactiveCode == types.Revoke {
		info.InactiveDesc = inactiveRevokeDesc
	} else if validator.InactiveCode == types.MissVoteBlock {
		info.InactiveDesc = inactiveMissVoteBlockDesc
	} else if validator.InactiveCode == types.MaxValidator {
		info.InactiveDesc = inactiveMaxValidatorDesc
	} else if validator.InactiveCode == types.DoubleSign {
		info.InactiveDesc = inactiveDoubleDesc
	}

	return info
}

func QueryValidators(cdc *amino.Codec, ip string, height int64) (validators []validatorDisplayInfo, err error) {
	cliCtx := context.NewCLIContext().WithCodec(cdc).WithNodeIP(ip)

	node, err := cliCtx.GetNode()
	if err != nil {
		return
	}

	opts := buildQueryOptions(height)

	subspace := "/store/validator/subspace"
	result, err := node.ABCIQueryWithOptions(subspace, types.BuildValidatorPrefixKey(), opts)

	if err != nil {
		return
	}

	valueBz := result.Response.GetValue()
	if len(valueBz) == 0 {
		err = errors.New("response empty value")
		return
	}

	//var validators []validatorDisplayInfo

	var vKVPair []store.KVPair
	cdc.UnmarshalBinaryLengthPrefixed(valueBz, &vKVPair)
	for _, kv := range vKVPair {
		var validator types.Validator
		cdc.UnmarshalBinaryBare(kv.Value, &validator)
		validators = append(validators, toValidatorDisplayInfo(validator))
	}
	for index, v := range validators {
		result, err := QueryDelegationInfo(cdc, ip, v.OperatorAddress, v.Owner)
		if err == nil {
			validators[index].SelfBond = result.Amount
		}
	}
	return
}

func buildQueryOptions(height int64) client.ABCIQueryOptions {
	if height <= 0 {
		height = 0
	}

	trust := viper.GetBool(bctypes.FlagTrustNode)

	return client.ABCIQueryOptions{
		Height: height,
		Prove:  trust,
	}
}

func QueryTotalValidatorBondToken(cdc *amino.Codec, ip string, height int64) (result btypes.BigInt, err error) {
	result = btypes.ZeroInt()
	validators, err := QueryValidators(cdc, ip, height)
	if err != nil {
		return
	}
	if len(validators) == 0 {
		return
	}
	for _, v := range validators {
		result = result.Add(v.BondTokens)
	}
	return
}

func QueryDelegationInfo(cdc *amino.Codec, ip string, validator btypes.ValAddress, delegator btypes.AccAddress) (result mapper.DelegationQueryResult, err error) {
	cliCtx := context.NewCLIContext().WithCodec(cdc).WithNodeIP(ip)

	//var validator btypes.ValAddress
	//var delegator btypes.AccAddress
	//
	//if o, err := qcliacc.GetValidatorAddrFromFlag(cliCtx, validatorAddr); err == nil {
	//	validator = o
	//}
	//
	//if d, err := qcliacc.GetAddrFromFlag(cliCtx, delegatorAddr); err == nil {
	//	delegator = d
	//}

	var path = types.BuildGetDelegationCustomQueryPath(delegator, validator)

	res, err := cliCtx.Query(path, []byte(""))
	if err != nil {
		return
	}

	//var result mapper.DelegationQueryResult
	cliCtx.Codec.UnmarshalJSON(res, &result)
	return
}

func QueryDelegationsWithValidator(cdc *amino.Codec, ip string, validatorAddr string) (result []mapper.DelegationQueryResult, err error) {
	cliCtx := context.NewCLIContext().WithCodec(cdc).WithNodeIP(ip)

	var validator btypes.ValAddress

	if o, err := qcliacc.GetValidatorAddrFromValue(cliCtx, validatorAddr); err == nil {
		validator = o
	}

	var path = types.BuildQueryDelegationsByOwnerCustomQueryPath(validator)

	res, err := cliCtx.Query(path, []byte(""))
	if err != nil {
		return
	}

	//var result []mapper.DelegationQueryResult
	cliCtx.Codec.UnmarshalJSON(res, &result)
	return
}

func QueryDelegationsWithDelegator(cdc *amino.Codec, ip string, delegatorAddr string) (result []mapper.DelegationQueryResult, err error) {
	cliCtx := context.NewCLIContext().WithCodec(cdc).WithNodeIP(ip)

	var delegator btypes.AccAddress

	if o, err := qcliacc.GetAddrFromValue(cliCtx, delegatorAddr); err == nil {
		delegator = o
	}
	var path = types.BuildQueryDelegationsByDelegatorCustomQueryPath(delegator)

	res, err := cliCtx.Query(path, []byte(""))
	if err != nil {
		return
	}

	//var result []mapper.DelegationQueryResult
	cliCtx.Codec.UnmarshalJSON(res, &result)
	return
}
