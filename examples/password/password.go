package password

type User struct {
	Name     string `check:"word"`
	Password string `check:"min=8,password"`
	Desc     string `check:"word"`
}

type User2 struct {
	Name        string  `check:"word"`
	PasswordRef *string `check:"min=8,password"`
	Desc        string  `check:"word"`
}
