package chainname

import (
	"github.com/reiver/go-chainid"
)

const Zora = "Zora"

func init() {
	const value string =                   Zora
	var   key   string = createkey(chainid.Zora)

	chainNames[key] = value
}
