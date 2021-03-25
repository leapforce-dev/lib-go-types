package types

import (
	"strconv"
	"strings"
)

type Int64String int64
type Int64Strings []Int64String

func (i *Int64String) UnmarshalJSON(b []byte) error {
	unquoted, err := strconv.Unquote(string(b))
	if err != nil {
		return err
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

func (is *Int64Strings) ToInt64() []int64 {
	var _i []int64

	if is != nil {
		for _, i := range *is {
			_i = append(_i, int64(i))
		}
	}

	return _i
}
