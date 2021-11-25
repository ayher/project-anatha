package client


import (
	govclient "github.com/ayher/project-anatha/x/governance/client"
	"github.com/ayher/project-anatha/x/hra/client/cli"
)

var (
	RegisterBlockchainIdProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitRegisterBlockchainIdProposal)
	RemoveBlockchainIdProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitRemoveBlockchainIdProposal)
)