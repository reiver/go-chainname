package chainname

import (
	"github.com/reiver/go-chainid"
)

const FantomTestnet = "Fantom Testnet"

func init() {
	const value string =                   FantomTestnet
	var   key   string = createkey(chainid.FantomTestnet)

	chainNames[key] = value
}
