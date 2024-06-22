package chainname

import (
	"github.com/reiver/go-chainid"
)

const Blast = "Blast"

func init() {
	const value string =                   Blast
	var   key   string = createkey(chainid.Blast)

	chainNames[key] = value
}
