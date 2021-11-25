package client

import (
	"github.com/ayher/project-anatha/x/fee/client/cli"
	govclient "github.com/ayher/project-anatha/x/governance/client"
)

var (
	AddFeeExcludedMessageProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitAddFeeExcludedMessageProposal)
	RemoveFeeExcludedMessageProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitRemoveFeeExcludedMessageProposal)
)