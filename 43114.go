package chainname

import (
	"github.com/reiver/go-chainid"
)

const AvalancheCChain = "Avalanche C-Chain"

func init() {
	const value string =                   AvalancheCChain
	var   key   string = createkey(chainid.AvalancheCChain)

	chainNames[key] = value
}
