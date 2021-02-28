package examples

import "github.com/sk-telemed/show2doc-dev-tools/validation-generator/examples/custom"

type Custom struct {
	Structs custom.Structs `check:"required,min=1,max=100,deep,check"`
	Strings custom.Strings `check:"required,min=1,max=100,deep,len=15"`
	Map     custom.Map     `check:"required,min=1,max=100,deep,len=15"`
}
