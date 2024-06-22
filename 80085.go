package chainname

import (
	"github.com/reiver/go-chainid"
)

const BerachainArtio = "Berachain Artio"

func init() {
	const value string =                   BerachainArtio
	var   key   string = createkey(chainid.BerachainArtio)

	chainNames[key] = value
}
