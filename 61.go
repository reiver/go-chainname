package chainname

import (
	"github.com/reiver/go-chainid"
)

const EthereumClassic = "Ethereum Classic"

func init() {
	const value string =                   EthereumClassic
	var   key   string = createkey(chainid.EthereumClassic)

	chainNames[key] = value
}
