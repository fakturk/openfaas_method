# openfaas_method
golang openfaas method example instead of function

Golang method is nothing but a function with a special receiver argument.

If you want to use use Golang methods in OpenFaaS you can find a template in the template/go-method folder that use method of Parent type struct isntead of a function.

We can create a new function by the following code

`faas-cli new --lang go-method <function-name>`

Now we have a handler.go file in the 'function-name' folder  

<code>
package function

import (
	"fmt"
)

//Parent struct
type Parent struct {
	Name string
}

// Handle a serverless request
func (p *Parent) Handle(req []byte) string {
	return fmt.Sprintf("Hello, Go. You said: %s and name of the parent is :%s", string(req), p.Name)
}

</code>

Now we can build, push and deploy our code

`faas-cli build -f <function-name>.yml `

`faas-cli push -f <function-name>.yml `

`faas-cli deploy -f <function-name>.yml `

Or we can replace build/push/deploy with `faas-cli up`