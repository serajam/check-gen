package innerpkg

import (
	"github.com/serajam/check-gen/examples/embedded/innerpkg"
)

// Foo validates custom type
type Foo struct {
	Name         string `check:"min=1,max=1000"`
	innerpkg.Bar `check:"deep,check,min=1"`
}

// Foo validates custom type
type FooRef struct {
	Name          string `check:"min=1,max=1000"`
	*innerpkg.Bar `check:"deep,check,min=1"`
}
