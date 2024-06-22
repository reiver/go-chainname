package chainname

import (
	"github.com/reiver/go-chainid"
)

const Holesky = "Holesky"

func init() {
	const value string =                   Holesky
	var   key   string = createkey(chainid.Holesky)

	chainNames[key] = value
}
