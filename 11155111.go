package chainname

import (
	"github.com/reiver/go-chainid"
)

const Sepolia = "Sepolia"

func init() {
	const value string =                   Sepolia
	var   key   string = createkey(chainid.Sepolia)

	chainNames[key] = value
}
