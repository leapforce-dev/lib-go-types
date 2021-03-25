package types

import (
	"encoding/json"
	"strconv"
	"strings"
)

type Float64String float64

func (f *Float64String) UnmarshalJSON(b []byte) error {
	var s string

	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	s = strings.Trim(s, " ")

	if s == "" {
		f = nil
		return nil
	}

	_f, err := strconv.ParseFloat(s, 64)
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
