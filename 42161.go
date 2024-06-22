package chainname

import (
	"github.com/reiver/go-chainid"
)

const ArbitrumOne = "Arbitrum One"

func init() {
	const value string =                   ArbitrumOne
	var   key   string = createkey(chainid.ArbitrumOne)

	chainNames[key] = value
}
