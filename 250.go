package chainname

import (
	"github.com/reiver/go-chainid"
)

const FantomOpera = "Fantom Opera"

func init() {
	const value string =                   FantomOpera
	var   key   string = createkey(chainid.FantomOpera)

	chainNames[key] = value
}
