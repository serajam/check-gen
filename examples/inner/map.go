package inner

type MapStruct struct {
	SpacesMapStruct   map[string]Inner  `check:"required,max=1,check,max=1"`
	SpacesSliceStruct []Inner           `check:"required,deep,check,min=1"`
	SpacesMapInnerRef map[string]*Inner `check:"required,deep,check,len=1"`
}
