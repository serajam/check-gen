package custom

type StructInCustom struct {
	Name string `check:"required"`
}

type Structs []StructInCustom
type Strings []string
type Map map[string]string
