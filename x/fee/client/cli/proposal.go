package cli

import (
	"bufio"
	"github.com/ayher/anatha/client/context"
	"github.com/ayher/anatha/codec"
	sdk "github.com/ayher/anatha/types"
	"github.com/ayher/anatha/x/auth"
	"github.com/ayher/anatha/x/auth/client/utils"
	govutils "github.com/ayher/project-anatha/x/fee/client/utils"
	"github.com/ayher/project-anatha/x/fee/internal/types"
	govtypes "github.com/ayher/project-anatha/x/governance"
	"github.com/spf13/cobra"
)

func GetCmdSubmitAddFeeExcludedMessageProposal(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-fee-excluded-message [proposal-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a proposal to exclude a message type from fees",
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			proposal, err := govutils.ParseFeeExclusionProposalJSON(cdc, args[0])
			if err != nil {
				return err
			}

			from := cliCtx.GetFromAddress()

			content := types.NewAddFeeExcludedMessageProposal(proposal.Title, proposal.Description, proposal.MessageType)

			msg := govtypes.NewMsgSubmitProposal(content, from)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}

func GetCmdSubmitRemoveFeeExcludedMessageProposal(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-fee-excluded-message [proposal-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a proposal to remove a message from fee exclusion",
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			proposal, err := govutils.ParseFeeExclusionProposalJSON(cdc, args[0])
			if err != nil {
				return err
			}

			from := cliCtx.GetFromAddress()

			content := types.NewRemoveFeeExcludedMessageProposal(proposal.Title, proposal.Description, proposal.MessageType)

			msg := govtypes.NewMsgSubmitProposal(content, from)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}