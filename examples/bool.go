package examples

type BoolCheck struct {
	Yes bool  // cant be checked
	No  *bool `check:"required"`
}
