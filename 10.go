package chainname

import (
	"github.com/reiver/go-chainid"
)

const OPMainnet = "OP Mainnet"

func init() {
	const value string =                   OPMainnet
	var   key   string = createkey(chainid.OPMainnet)

	chainNames[key] = value
}
