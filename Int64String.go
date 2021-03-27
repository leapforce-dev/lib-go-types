package types

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	errortools "github.com/leapforce-libraries/go_errortools"
)

type Int64String int64
type Int64Strings []Int64String

func (i *Int64String) UnmarshalJSON(b []byte) error {
	var returnError = func() error {
		errortools.CaptureError(fmt.Sprintf("Cannot parse '%s' to Int64String", string(b)))
		return nil
	}

	var s string

	err := json.Unmarshal(b, &s)
	if err != nil {
		return returnError()
	}

	s = strings.Trim(s, " ")

	if s == "" {
		i = nil
		return nil
	}

	_i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return returnError()
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
