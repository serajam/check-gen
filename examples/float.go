package examples

type Float64Check struct {
	Number    float64  `check:"required,min=1.1,max=1000.15"`
	NumberRef *float64 `check:"required,min=1,max=100"`
}

type Float32Check struct {
	Number    float32  `check:"required,min=1,max=1000"`
	NumberRef *float32 `check:"required,min=1,max=100"`
}
