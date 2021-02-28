package phone

type User struct {
	Phone    string  `check:"phone=7~15"`
	PhoneRef *string `check:"phone=7~15"`
}
