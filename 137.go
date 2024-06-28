package chainname

import (
	"github.com/reiver/go-chainid"
)

const PolygonMainnet = "Polygon Mainnet"

func init() {
	const value string =                   PolygonMainnet
	var   key   string = createkey(chainid.PolygonMainnet)

	chainNames[key] = value
}
