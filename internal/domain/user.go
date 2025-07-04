package domain

type User struct {
	ID   int64
	Name string
}

func NewUser(name string) *User {
	return &User{Name: name}
}
