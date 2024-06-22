package chainname

import (
	"github.com/reiver/go-chainid"
)

const ZoraSepoliaTestnet = "Zora Sepolia Testnet"

func init() {
	const value string =                   ZoraSepoliaTestnet
	var   key   string = createkey(chainid.ZoraSepoliaTestnet)

	chainNames[key] = value
}
