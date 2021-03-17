package types

import (
	"strconv"
	"strings"
)

type Float64String float64

func (f *Float64String) UnmarshalJSON(b []byte) error {
	unquoted, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}

	if strings.Trim(unquoted, " ") == "" {
		f = nil
		return nil
	}

	_f, err := strconv.ParseFloat(unquoted, 64)
	if err != nil {
		return err
	}

	*f = Float64String(_f)
	return nil
}

func (f *Float64String) ValuePtr() *float64 {
	if f == nil {
		return nil
	}

	_f := float64(*f)
	return &_f
}

func (f Float64String) Value() float64 {
	return float64(f)
}
