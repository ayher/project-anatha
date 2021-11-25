package types

import (
	"github.com/ayher/anatha/codec"
	gov "github.com/ayher/anatha/x/gov"
	"github.com/ayher/anatha/x/params"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterInterface((*gov.Content)(nil), nil)

	cdc.RegisterConcrete(MsgSubmitProposal{}, "governance/MsgSubmitProposal", nil)
	cdc.RegisterConcrete(MsgVote{}, "governance/MsgVote", nil)
	cdc.RegisterConcrete(MsgExpedite{}, "governance/MsgExpedite", nil)
	cdc.RegisterConcrete(TextProposal{}, "governance/TextProposal", nil)

	cdc.RegisterConcrete(AddGovernorProposal{}, "governance/AddGovernorProposal", nil)
	cdc.RegisterConcrete(RemoveGovernorProposal{}, "governance/RemoveGovernorProposal", nil)
}

var ModuleCdc = codec.New()

func RegisterProposalTypeCodec(o interface{}, name string) {
	ModuleCdc.RegisterConcrete(o, name, nil)
}

func init() {
	RegisterCodec(ModuleCdc)

	RegisterProposalTypeCodec(params.ParameterChangeProposal{}, "cosmos-sdk/ParameterChangeProposal")

	codec.RegisterCrypto(ModuleCdc)
}