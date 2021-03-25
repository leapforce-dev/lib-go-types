package types

import (
	"encoding/json"
	"strings"
)

type String string

func (s *String) UnmarshalJSON(b []byte) error {
	var st string

	err := json.Unmarshal(b, &st)
	if err != nil {
		return err
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
