package entity

import "time"

type Product struct {
	ID        int32
	Email     string
	Password  string
	FirstName string
	LastName  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Books     []int32
}

func NewUser(email, password, firstName, lastName string) (*Product, error) {
	u := &Product{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		CreatedAt: time.Now(),
	}
	u.Password = password
	return u, nil
}
