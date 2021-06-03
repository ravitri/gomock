package user

import "github.com/ravitri/gomock/doer"

type User struct {
	Doer doer.Doer
}

func (u *User) Use() error {
	return u.Doer.DoSomething(1989, "Hello RBT")
}
