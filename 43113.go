package chainname

import (
	"github.com/reiver/go-chainid"
)

const AvalancheFujiTestnet = "Avalanche Fuji Testnet"

func init() {
	const value string =                   AvalancheFujiTestnet
	var   key   string = createkey(chainid.AvalancheFujiTestnet)

	chainNames[key] = value
}
