package examples

type NestedSliceStructCheck struct {
	InnerStructSlice    []Inner  `check:"required,min=1,max=100,deep,check"`
	InnerStructSliceRef []*Inner `check:"required,min=1,max=100,deep,check"`
}

type Inner struct {
	Name  string `check:"required,min=1,max=100"`
	Value int    `check:"required,min=1,max=100"`
}

type SliceCheck struct {
	SpacesStr     []string  `check:"required,deep,word"`
	SpacesInt     []int     `check:"required,len=100,min=1,max=100,deep,min=1,max=100"`
	SpacesInt8    []int8    `check:"required,len=100,min=1,max=100,deep,max=100,min=1,max=100"`
	SpacesInt16   []int16   `check:"required,len=100,min=1,max=100,deep,max=1000,min=1,max=100"`
	SpacesInt32   []int32   `check:"required,len=100,min=1,max=100,deep,max=1000,min=1,max=100"`
	SpacesInt64   []int64   `check:"required,len=100,min=1,max=100,deep,max=1000,min=1,max=100"`
	SpacesFloat64 []float64 `check:"required,len=100,min=1,max=100,deep,max=1000,min=1,max=100"`
	SpacesFloat32 []float32 `check:"required,len=100,min=1,max=100,deep,max=1000,min=1,max=100"`
	SpacesByte    []byte    `check:"required,len=100,min=1,max=100,deep,min=1,max=100"`
}

type SliceRefCheck struct {
	SpacesStr     []*string  `check:"required,deep,digit"`
	SpacesInt     []*int     `check:"required,len=100,min=1,max=100,deep,max=1000,min=1,max=100"`
	SpacesInt8    []*int8    `check:"required,len=100,min=1,max=100,deep,max=100,min=1,max=100"`
	SpacesInt16   []*int16   `check:"required,len=100,min=1,max=100,deep,max=1000,min=1,max=100"`
	SpacesInt32   []*int32   `check:"required,len=100,min=1,max=100,deep,max=1000,min=1,max=100"`
	SpacesInt64   []*int64   `check:"required,len=100,min=1,max=100,deep,max=1000,min=1,max=100"`
	SpacesFloat64 []*float64 `check:"required,len=100,min=1,max=100,deep,max=1000,min=1,max=100"`
	SpacesFloat32 []*float32 `check:"required,len=100,min=1,max=100,deep,max=1000,min=1,max=100"`
	SpacesByte    []*byte    `check:"required,len=100,min=1,max=100,deep,max=1000,min=1,max=100"`
}
