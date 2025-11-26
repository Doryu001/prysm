package params

import (
	"fmt"
)

type Option func(*BeaconChainConfig)

func WithGenesisValidatorsRoot(gvr [32]byte) Option {
	return func(cfg *BeaconChainConfig) {
		cfg.GenesisValidatorsRoot = gvr
		log.WithField("genesis_validators_root", fmt.Sprintf("%#x", gvr)).Info("Setting genesis validators root")
	}
}
