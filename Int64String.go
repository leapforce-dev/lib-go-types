package types

import (
	"strconv"
	"strings"
)

type Int64String int64

func (i *Int64String) UnmarshalJSON(b []byte) error {
	unquoted, err := strconv.Unquote(string(b))
	if err != nil {
		return nil
	}

	if strings.Trim(unquoted, " ") == "" {
		i = nil
		return nil
	}

	_i, err := strconv.ParseInt(unquoted, 10, 64)
	if err != nil {
		return err
	}

	*i = Int64String(_i)
	return nil
}

func (i *Int64String) ValuePtr() *int64 {
	if i == nil {
		return nil
	}

	_i := int64(*i)
	return &_i
}

func (i Int64String) Value() int64 {
	return int64(i)
}
