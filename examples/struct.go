package examples

type Nested struct {
	Name  string `check:"required,min=1,max=100"`
	Value int    `check:"required,min=1,max=100"`
}

type NestedStructCheck struct {
	NestedStruct         Nested  `check:"required,check"`
	NestedStructRefCheck *Nested `check:"required,check"`
}
