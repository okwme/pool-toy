package faucet

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/interchainberlin/pooltoy/regex"
	"github.com/interchainberlin/pooltoy/x/faucet/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto"
)

// TODO: rewrite test
func TestEmoji(t *testing.T) {

	sdk.SetCoinDenomRegex(func() string {
		return regex.NewDnmRegex
	})

	moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte("foo")))
	moduleAcct2 := sdk.AccAddress(crypto.AddressHash([]byte("bar")))
	denom := "🥵"
	msg := types.NewMsgMint(moduleAcct, moduleAcct2, denom)

	err := msg.ValidateBasic()
	require.NoError(t, err)

	_, err = sdk.ParseCoinsNormalized("1" + msg.Denom)
	if err != nil {
		fmt.Println("Not correct interface for Emoji")
	}

	fmt.Println("final msg.Denom", msg.Denom)
	// require.True(t, false)
}
