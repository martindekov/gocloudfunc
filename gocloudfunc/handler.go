package function

import (
	"fmt"
)

// Handle a serverless request
func Handle(req []byte) string {
	return fmt.Sprintf("He llo, G o. Yo u said: %s", string(req))
}
