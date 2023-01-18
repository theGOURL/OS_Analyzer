package analyst

import (
	"os/user"

	"github.com/theGOURL/warning"
)

// this function prints the username of the system,
// returning the username
func myOSUsername(errmessage string) string {
	whoami, err := user.Current()
	warning.PRINT_DEFAULT_ERRORS(err, errmessage)
	username := whoami.Username
	return username
}
