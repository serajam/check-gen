package inner

type Inner struct {
	Name  string `check:"required,min=1,max=100"`
	Value int    `check:"required,min=1,max=100"`
}

