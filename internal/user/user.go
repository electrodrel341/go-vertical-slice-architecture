package user

import "time"

// User Domain
type User struct {
	Id             int64
	IdentityNumber string
	FirstName      string
	LastName       string
	Email          string
	DateOfBirth    time.Time
	CreatedAt      time.Time
}

// Create a new user instance
func New(in string, fn string, ln string, e string, dob time.Time) (*User, error) {
	return &User{
		IdentityNumber: in,
		FirstName:      fn,
		LastName:       ln,
		Email:          e,
		DateOfBirth:    dob,
	}, nil
}
