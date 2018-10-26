package function

import (
	"fmt"
)

// Handle a serverless request
func Handle(req []byte) string {
	return fmt.Sprintf("He o, G o. Yo u said: %s", string(req))
}
