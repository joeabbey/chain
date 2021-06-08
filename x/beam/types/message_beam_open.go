package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgOpenBeam{}

// NewMsgOpenBeam Build a open beam message based on parameters
func NewMsgOpenBeam(id string, creator string, owner string, amount sdk.Coin, secret string, schema string, data *BeamData) *MsgOpenBeam {
	return &MsgOpenBeam{
		Id:             id,
		CreatorAddress: creator,
		Amount:         amount,
		Secret:         secret,
		Schema:         schema,
		Data:           data,
		ClaimAddress:   owner,
	}
}

// Route dunno
func (msg MsgOpenBeam) Route() string {
	return RouterKey
}

// Type Return the message type
func (msg MsgOpenBeam) Type() string {
	return "OpenBeam"
}

// GetSigners Return the list of signers for the given message
func (msg *MsgOpenBeam) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.GetCreatorAddress())
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes Return the generated bytes from the signature
func (msg *MsgOpenBeam) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic Validate the message payload before dispatching to the local kv store
func (msg *MsgOpenBeam) ValidateBasic() error {
	if len(msg.Id) <= 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "Invalid id supplied (%d)", len(msg.Id))
	}

	// Ensure the address is correct and that we are able to acquire it
	_, err := sdk.AccAddressFromBech32(msg.GetCreatorAddress())
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid creator address (%s)", err)
	}

	// Validate the secret
	if len(msg.Secret) <= 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "Invalid secret supplied")
	}

	// Validate the schema
	if msg.GetSchema() != BEAM_SCHEMA_REVIEW && msg.GetSchema() != BEAM_SCHEMA_REWARD {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "Invalid schema must be review or reward")
	}

	// If we have an amount, make sure it is not negative nor zero
	if msg.Amount.IsNegative() || msg.Amount.IsZero() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "Invalid amount: must be greater than 0")
	}
	return nil
}