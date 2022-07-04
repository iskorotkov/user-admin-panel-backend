package entities

import (
	"go.uber.org/multierr"
)

type SingleError error

type CompositeError []SingleError

func (v CompositeError) Slice() []string {
	var res []string
	for _, err := range v {
		res = append(res, err.Error())
	}

	return res
}

func (v CompositeError) Combine() error {
	var errs []error
	for _, err := range v {
		errs = append(errs, err)
	}

	return multierr.Combine(errs...)
}
