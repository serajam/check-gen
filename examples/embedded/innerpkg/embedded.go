package innerpkg

// Foo validates custom type
type Foo struct {
	Name string `check:"min=1,max=1000"`
}

// Bar embedded struct
type Bar struct {
	Surname string `check:"min=1,max=1000"`
	Foo     `check:"deep,check,min=1"`
}

// Bar embedded struct
type BarRef struct {
	Surname string `check:"min=1,max=1000"`
	*Foo    `check:"required,deep,check,min=1"`
}
