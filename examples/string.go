package examples

type StringCheck struct {
	Name       string  `check:"required,max=20,min=1,len=10"`
	NameRef    *string `check:"required,max=20,min=1"`
	SurnameRef *string `check:"max=20,min=1,len=100"`
	IdRef      *string `check:"len=100,min=1,max=1000"`

	NumberStringRef *string `check:"required,digit"`
	NumberString    string  `check:"required,len=5"`

	StringAllnumRef *string `check:"required,word"`
	StringAllnum    string  `check:"required,word"`

	Test     string    `check:"uuid"`
	TestRef  *string   `check:"uuid"`
	UTest    []string  `check:"required,deep,uuid"`
	UTestRef []*string `check:"required,deep,uuid"`

	WTest    []string  `check:"required,deep,word"`
	WTestRef []*string `check:"required,deep,word"`

	DTest    []string  `check:"required,deep,digit"`
	DTestRef []*string `check:"required,deep,digit"`

	SpacesMapString   map[string]string `check:"required"`
}
