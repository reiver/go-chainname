package chainname

import (
	"github.com/reiver/go-chainid"
)

const RootstockTestnet = "Rootstock Testnet"

func init() {
	const value string =                   RootstockTestnet
	var   key   string = createkey(chainid.RootstockTestnet)

	chainNames[key] = value
}
