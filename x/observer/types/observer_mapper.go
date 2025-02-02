package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zeta-chain/zetacore/common"
)

// Validate observer mapper contains an existing chain
func (m *ObserverSet) Validate() error {
	for _, observerAddress := range m.ObserverList {
		_, err := sdk.AccAddressFromBech32(observerAddress)
		if err != nil {
			return err
		}
	}
	return nil
}

func CheckReceiveStatus(status common.ReceiveStatus) error {
	switch status {
	case common.ReceiveStatus_Success:
		return nil
	case common.ReceiveStatus_Failed:
		return nil
	default:
		return ErrInvalidStatus
	}
}
