package examples

type IntCheck struct {
	Number    int  `check:"required,min=1,max=1000"`
	NumberRef *int `check:"required,min=1,max=100"`
}

type Int64Check struct {
	Number    int64  `check:"required,min=1,max=1000"`
	NumberRef *int64 `check:"required,min=1,max=100"`
}

type Int32Check struct {
	Number    int32  `check:"required,min=1,max=1000"`
	NumberRef *int32 `check:"required,min=1,max=100"`
}

type Int16Check struct {
	Number    int16  `check:"required,min=1,max=1000"`
	NumberRef *int16 `check:"required,min=1,max=100"`
}

type Int8Check struct {
	Number    int8  `check:"required,min=1,max=120"`
	NumberRef *int8 `check:"required,min=1,max=120"`
}
