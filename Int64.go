package types

import (
	"strconv"
	"strings"
)

type Int64 int64

func (i *Int64) UnmarshalJSON(b []byte) error {
	s := string(b)
	if strings.Trim(s, " ") == "" {
		i = nil
		return nil
	}

	_i, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		return err
	}

	*i = Int64(_i)
	return nil
}

func (i *Int64) Value() *int64 {
	if i == nil {
		return nil
	}

	_i := int64(*i)
	return &_i
}
