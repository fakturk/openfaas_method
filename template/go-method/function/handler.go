package function

import (
	"fmt"
)

type Parent struct{
	name string
}

// Handle a serverless request
func (p *Parent) Handle(req []byte) string {
	return fmt.Sprintf("Hello, Go. You said: %s and name of the parent is :%s", string(req),p.name)
}
