package cli

import (
	"bufio"
	"github.com/ayher/anatha/x/params"
	"github.com/spf13/cobra"

	"github.com/ayher/anatha/client/context"
	"github.com/ayher/anatha/codec"
	sdk "github.com/ayher/anatha/types"
	"github.com/ayher/anatha/x/auth"
	"github.com/ayher/anatha/x/auth/client/utils"
	govutils "github.com/ayher/project-anatha/x/governance/client/utils"
	"github.com/ayher/project-anatha/x/governance/internal/types"
)

func GetCmdSubmitParamChangeProposal(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "param-change [proposal-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a parameter change proposal",
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			proposal, err := govutils.ParseParamChangeProposalJSON(cdc, args[0])
			if err != nil {
				return err
			}

			from := cliCtx.GetFromAddress()
			content := params.NewParameterChangeProposal(proposal.Title, proposal.Description, proposal.Changes.ToParamChanges())

			msg := types.NewMsgSubmitProposal(content, from)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}