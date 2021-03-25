package types

import (
	"strconv"
	"strings"
)

type String string

func (s *String) UnmarshalJSON(b []byte) error {
	unquoted, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}

	if strings.Trim(unquoted, " ") == "" {
		s = nil
		return nil
	}

	*s = String(unquoted)
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
