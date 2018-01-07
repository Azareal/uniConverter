package gosora

import "../../common"

func Lookup(version string) common.Software {
	return NewPrealphaSoftware()
}
