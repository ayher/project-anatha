package client

import (
	govclient "github.com/ayher/project-anatha/x/governance/client"
	"github.com/ayher/project-anatha/x/upgrade/client/cli"
)

var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitSoftwareUpgradeProposal)
