package types

import (
	"encoding/json"
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	"strings"
)

type String string

func (s *String) UnmarshalJSON(b []byte) error {
	var returnError = func() error {
		errortools.CaptureError(fmt.Sprintf("Cannot parse '%s' to String", string(b)))
		return nil
	}

	var st string

	err := json.Unmarshal(b, &st)
	if err != nil {
		return returnError()
	}

	st = strings.Trim(st, " ")

	if st == "" {
		s = nil
		return nil
	}

	*s = String(st)
	return nil
}

func (s *String) ValuePtr() *string {
	if s == nil {
		return nil
	}

	_s := string(*s)
	return &_s
}

func (s String) Value() string {
	return string(s)
}
