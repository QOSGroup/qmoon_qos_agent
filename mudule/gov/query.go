package gov

import (
	"errors"
	"fmt"
	qcliacc "github.com/QOSGroup/qbase/client/account"
	"github.com/QOSGroup/qbase/client/context"
	btypes "github.com/QOSGroup/qbase/types"
	"github.com/QOSGroup/qos/module/gov/mapper"
	"github.com/QOSGroup/qos/module/gov/types"
	go_amino "github.com/tendermint/go-amino"
)

func QueryProposal(cdc *go_amino.Codec, pID int64) (types.Proposal, error) {
	cliCtx := context.NewCLIContext().WithCodec(cdc)
	//pID, err := strconv.ParseUint(args[0], 10, 64)
	//if err != nil {
	//	return fmt.Errorf("proposal id %s is not a valid uint value", args[0])
	//}

	path := mapper.BuildQueryProposalPath(pID)
	res, err := cliCtx.Query(path, []byte{})

	if err != nil {
		return types.Proposal{}, nil
	}

	if len(res) == 0 {
		return types.Proposal{}, errors.New("no result found")
	}

	var result types.Proposal
	err = cliCtx.Codec.UnmarshalJSON(res, &result)
	return result, err
}

func QueryProposals(cdc *go_amino.Codec, limit int64, depositor, voter, statusStr string) ([]types.Proposal, error) {
	cliCtx := context.NewCLIContext().WithCodec(cdc)

	var depositorAddr btypes.AccAddress
	var voterAddr btypes.AccAddress
	var status types.ProposalStatus

	if d, err := qcliacc.GetAddrFromValue(cliCtx, depositor); err == nil {
		depositorAddr = d
	}

	if d, err := qcliacc.GetAddrFromValue(cliCtx, voter); err == nil {
		voterAddr = d
	}

	status = toProposalStatus(statusStr)

	queryParam := mapper.QueryProposalsParam{
		Depositor: depositorAddr,
		Voter:     voterAddr,
		Status:    status,
		Limit:     limit,
	}

	data, err := cliCtx.Codec.MarshalJSON(queryParam)
	if err != nil {
		return nil, err
	}

	path := mapper.BuildQueryProposalsPath()
	res, err := cliCtx.Query(path, data)

	if len(res) == 0 {
		return nil, errors.New("no result found")
	}

	var result []types.Proposal
	err = cliCtx.Codec.UnmarshalJSON(res, &result)
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("no matching proposals found")
	}

	return result, err
}

func toProposalStatus(statusStr string) types.ProposalStatus {
	switch statusStr {
	case "DepositPeriod", "deposit_period":
		return types.StatusDepositPeriod
	case "VotingPeriod", "voting_period":
		return types.StatusVotingPeriod
	case "Passed", "passed":
		return types.StatusPassed
	case "Rejected", "rejected":
		return types.StatusRejected
	default:
		return types.StatusNil
	}
}

func QueryVote(cdc *go_amino.Codec, pID int64, addrStr string) (types.Vote, error) {
	cliCtx := context.NewCLIContext().WithCodec(cdc)

	addr, err := qcliacc.GetAddrFromValue(cliCtx, addrStr)
	if err != nil {
		return types.Vote{}, fmt.Errorf("voter %s is not a valid address value", addrStr)
	}

	path := mapper.BuildQueryVotePath(pID, addr.String())
	res, err := cliCtx.Query(path, []byte{})
	if err != nil {
		return types.Vote{}, err
	}

	if len(res) == 0 {
		return types.Vote{}, errors.New("no result found")
	}

	var vote types.Vote
	if err := cliCtx.Codec.UnmarshalJSON(res, &vote); err != nil {
		return types.Vote{}, err
	}

	return vote, err
}

func QueryVotes(cdc *go_amino.Codec, pID int64) ([]types.Vote, error) {
	cliCtx := context.NewCLIContext().WithCodec(cdc)

	path := mapper.BuildQueryVotesPath(pID)
	res, err := cliCtx.Query(path, []byte{})
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, errors.New("no result found")
	}

	var votes []types.Vote
	if err := cliCtx.Codec.UnmarshalJSON(res, &votes); err != nil {
		return nil, err
	}

	if len(votes) == 0 {
		return nil, errors.New("no votes found")
	}

	return votes, err
}

func QueryDeposit(cdc *go_amino.Codec, pID int64, addrStr string) (types.Deposit, error) {
	cliCtx := context.NewCLIContext().WithCodec(cdc)

	addr, err := qcliacc.GetAddrFromValue(cliCtx, addrStr)
	if err != nil {
		return types.Deposit{}, fmt.Errorf("voter %s is not a valid address value", addrStr)
	}

	path := mapper.BuildQueryVotePath(pID, addr.String())
	res, err := cliCtx.Query(path, []byte{})
	if err != nil {
		return types.Deposit{}, err
	}

	if len(res) == 0 {
		return types.Deposit{}, errors.New("no result found")
	}

	var deposit types.Deposit
	if err := cliCtx.Codec.UnmarshalJSON(res, &deposit); err != nil {
		return types.Deposit{}, nil
	}

	return deposit, err
}

func QueryDeposits(cdc *go_amino.Codec, pID int64) ([]types.Deposit, error) {
	cliCtx := context.NewCLIContext().WithCodec(cdc)

	path := mapper.BuildQueryVotesPath(pID)
	res, err := cliCtx.Query(path, []byte{})
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, errors.New("no result found")
	}

	var deposits []types.Deposit
	if err := cliCtx.Codec.UnmarshalJSON(res, &deposits); err != nil {
		return nil, err
	}

	return deposits, err
}

func QueryTally(cdc *go_amino.Codec, pID int64, addrStr string) (types.TallyResult, error) {
	cliCtx := context.NewCLIContext().WithCodec(cdc)

	addr, err := qcliacc.GetAddrFromValue(cliCtx, addrStr)
	if err != nil {
		return types.TallyResult{}, fmt.Errorf("voter %s is not a valid address value", addrStr)
	}

	path := mapper.BuildQueryVotePath(pID, addr.String())
	res, err := cliCtx.Query(path, []byte{})
	if err != nil {
		return types.TallyResult{}, err
	}

	if len(res) == 0 {
		return types.TallyResult{}, errors.New("no result found")
	}

	var result types.TallyResult
	if err := cliCtx.Codec.UnmarshalJSON(res, &result); err != nil {
		return types.TallyResult{}, nil
	}

	return result, err
}
