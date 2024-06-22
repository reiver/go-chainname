package chainname

import (
	"github.com/reiver/go-chainid"
)

const OPSepoliaTestnet = "OP Sepolia Testnet"

func init() {
	const value string =                   OPSepoliaTestnet
	var   key   string = createkey(chainid.OPSepoliaTestnet)

	chainNames[key] = value
}
