package chainname

import (
	"github.com/reiver/go-chainid"
)

const RootstockMainnet = "Rootstock Mainnet"

func init() {
	const value string =                   RootstockMainnet
	var   key   string = createkey(chainid.RootstockMainnet)

	chainNames[key] = value
}
