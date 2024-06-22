package chainname

import (
	"github.com/reiver/go-chainid"
)

const ArbitrumNova = "Arbitrum Nova"

func init() {
	const value string =                   ArbitrumNova
	var   key   string = createkey(chainid.ArbitrumNova)

	chainNames[key] = value
}
