package function

import (
	"fmt"
)

//Parent struct
type Parent struct {
	Name string
}

func Start(req []byte) string {
	parent := Parent{Name: "Fatih"}
	return parent.Handle(req)
}

// Handle a serverless request
func (p *Parent) Handle(req []byte) string {
	return fmt.Sprintf("Hello, Go. You said: %s and name of the parent is :%s", string(req), p.Name)
}
