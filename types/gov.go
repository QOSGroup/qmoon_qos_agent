package types

import (
	"time"
	"github.com/QOSGroup/qos/module/gov/types"
	btypes "github.com/QOSGroup/qbase/types"
)

type ResultProposal struct {

	ProposalID int64 `json:"proposal_id"` //  ID of the proposal
	Type string `json:"type"`
	Title string `json:"title"`
	Description string `json:"description"`
	Level string `json:"level"`

	Status           string `json:"proposal_status"`    //  Status of the Proposal
	FinalTallyResult types.TallyResult    `json:"final_tally_result"` //  Result of Tallys

	SubmitTime     time.Time     `json:"submit_time"`      //  Time of the block where TxGovSubmitProposal was included
	DepositEndTime time.Time     `json:"deposit_end_time"` // Time that the Proposal would expire if deposit amount isn't met
	TotalDeposit   btypes.BigInt `json:"total_deposit"`    //  Current deposit on this proposal. Initial value is set at InitialDeposit

	VotingStartTime   time.Time `json:"voting_start_time"` //  Time of the block where MinDeposit was reached. -1 if MinDeposit is not reached
	VotingStartHeight int64     `json:"voting_start_height"`
	VotingEndTime     time.Time `json:"voting_end_time"` // Time that the VotingPeriod for this proposal will end and votes will be tallied
}
