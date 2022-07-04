package entities

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

const (
	GenderMale        Gender = "male"
	GenderFemale      Gender = "female"
	GenderOther       Gender = "other"
	GenderUnspecified Gender = "unspecified"
)

type Gender string

type User struct {
	gorm.Model
	Name   string
	Phone  string
	Email  string
	Gender Gender
}

func (u User) Trim() User {
	return User{
		Model:  u.Model,
		Name:   strings.TrimSpace(u.Name),
		Phone:  strings.TrimSpace(u.Phone),
		Email:  strings.TrimSpace(u.Email),
		Gender: u.Gender,
	}
}

func (u User) Validate() CompositeError {
	var errs CompositeError

	if u.Name == "" {
		errs = append(errs, fmt.Errorf("name can't be empty"))
	}

	if u.Phone == "" {
		errs = append(errs, fmt.Errorf("phone can't be empty"))
	}

	// We don't validate email field correctness with regular expressions as it is too complex.
	// See https://davidcel.is/2012/09/06/stop-validating-email.html for examples.
	if u.Email == "" {
		errs = append(errs, fmt.Errorf("email can't be empty"))
	}

	usernameAndHost := strings.SplitN(u.Email, "@", 2)
	if len(usernameAndHost) < 2 {
		errs = append(errs, fmt.Errorf("email must contain @"))
	}

	if len(usernameAndHost) >= 2 && strings.Count(usernameAndHost[1], ".") == 0 {
		errs = append(errs, fmt.Errorf("email must contain ."))
	}

	if u.Gender != GenderMale &&
		u.Gender != GenderFemale &&
		u.Gender != GenderOther &&
		u.Gender != GenderUnspecified {
		errs = append(errs, fmt.Errorf("unsupported gender"))
	}

	if u != u.Trim() {
		errs = append(errs, fmt.Errorf("fields have leading or trailing whitespace"))
	}

	return errs
}

func (u User) BeforeSave(db *gorm.DB) error {
	return u.Validate().Combine()
}
