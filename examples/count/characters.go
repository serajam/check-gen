package count

type (
	Checker struct {
		CyrillicLen string `check:"len=70"`
		PolishLen   string `check:"len=111"`
		GermanLen   string `check:"len=120"`

		CyrillicMin string `check:"min=70"`
		PolishMin   string `check:"min=111"`
		GermanMin   string `check:"min=120"`

		CyrillicMax string `check:"max=70"`
		PolishMax   string `check:"max=111"`
		GermanMax   string `check:"max=120"`
	}

	CheckerRef struct {
		CyrillicLen *string `check:"len=70"`
		PolishLen   *string `check:"len=111"`
		GermanLen   *string `check:"len=120"`

		CyrillicMin *string `check:"min=70"`
		PolishMin   *string `check:"min=111"`
		GermanMin   *string `check:"min=120"`

		CyrillicMax *string `check:"max=70"`
		PolishMax   *string `check:"max=111"`
		GermanMax   *string `check:"max=120"`
	}

	CheckerSlice struct {
		Multiple     []string  `check:"deep,len=513"`
		MultipleRefs []*string `check:"deep,len=513"`
	}
)
