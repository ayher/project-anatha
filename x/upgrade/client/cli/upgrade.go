package cli

import (
	"bufio"
	"github.com/ayher/project-anatha/x/governance"
	"github.com/ayher/project-anatha/x/upgrade/internal/types"
	"github.com/spf13/cobra"

	"github.com/ayher/anatha/client/context"
	"github.com/ayher/anatha/codec"
	sdk "github.com/ayher/anatha/types"
	"github.com/ayher/anatha/x/auth"
	"github.com/ayher/anatha/x/auth/client/utils"
	upgradeutils "github.com/ayher/project-anatha/x/upgrade/client/utils"
)

func GetCmdSubmitSoftwareUpgradeProposal(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "software-upgrade [proposal-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a software upgrade change proposal",
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			proposal, err := upgradeutils.ParseSoftwareUpgradeProposalJSON(cdc, args[0])
			if err != nil {
				return err
			}

			from := cliCtx.GetFromAddress()
			content := types.NewSoftwareUpgradeProposal(proposal.Title, proposal.Description, types.Plan {
				Name: proposal.Plan.Name,
				Height: proposal.Plan.Height,
				Info: proposal.Plan.Info,
			})

			msg := governance.NewMsgSubmitProposal(content, from)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}