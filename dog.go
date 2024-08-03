package dog

import (
	"strings"
)

func WhenGrownUp(string_to_pass_in string) string {
	return "When the puppy grows up it says: " + strings.ToUpper(string_to_pass_in)
}
