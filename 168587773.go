package chainname

import (
	"github.com/reiver/go-chainid"
)

const BlastSepoliaTestnet = "Blast Sepolia Testnet"

func init() {
	const value string =                   BlastSepoliaTestnet
	var   key   string = createkey(chainid.BlastSepoliaTestnet)

	chainNames[key] = value
}
