package chainname

var chainNames map[string]string = map[string]string{}

func ChainName(chainid uint64) string {

	var key string = createkey(chainid)

	name, found := chainNames[key]
	if !found {
		return ""
	}

	return name
}
