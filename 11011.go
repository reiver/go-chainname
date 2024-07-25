package chainname

import (
	"github.com/reiver/go-chainid"
)

const ShapeSepolia = "Shape Sepolia"

func init() {
	const value string =                   ShapeSepolia
	var   key   string = createkey(chainid.ShapeSepolia)

	chainNames[key] = value
}
