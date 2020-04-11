package internal

import (
	util "gosh/pkg"
)

func validateBooleanVariable(b string) bool {
	yesOpts := [3]string{"yes", "1", "true"}

	return util.ItemExists(yesOpts, b)
}
