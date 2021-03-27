package types

import (
	"encoding/json"
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
)

type BoolInt bool

func (bl *BoolInt) UnmarshalJSON(b []byte) error {
	var returnError = func() error {
		errortools.CaptureError(fmt.Sprintf("Cannot parse '%s' to BoolInt", string(b)))
		return nil
	}

	var i int

	err := json.Unmarshal(b, &i)
	if err != nil {
		return returnError()
	}

	if i == 0 {
		*bl = false
		return nil
	}

	if i == 1 {
		*bl = true
		return nil
	}

	return returnError()
}

func (bl *BoolInt) ValuePtr() *bool {
	if bl == nil {
		return nil
	}

	_b := bool(*bl)
	return &_b
}

func (bl BoolInt) Value() bool {
	return bool(bl)
}
