package gosora

import "../../common"

func Lookup(version string) (soft common.Software, exists bool) {
	return NewPrealphaSoftware()
}
