package client

import (
	"github.com/ayher/project-anatha/x/distribution/client/cli"
	govclient "github.com/ayher/project-anatha/x/governance/client"
)

var DevelopmentFundDistributionProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitDevelopmentFundDistributionProposal)
var SecurityTokenFundDistributionProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitSecurityTokenFundDistributionProposal)
