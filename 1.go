package chainname

import (
	"github.com/reiver/go-chainid"
)

const EthereumMainnet = "Ethereum Mainnet"

func init() {
	const value string =                   EthereumMainnet
	var   key   string = createkey(chainid.EthereumMainnet)

	chainNames[key] = value
}
