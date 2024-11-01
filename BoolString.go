package types

import (
	"encoding/json"
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	"strconv"
	"strings"
)

type BoolString bool

func (bl *BoolString) UnmarshalJSON(b []byte) error {
	var returnError = func() error {
		errortools.CaptureError(fmt.Sprintf("Cannot parse '%s' to BoolString", string(b)))
		return nil
	}

	unquoted, err := strconv.Unquote(string(b))
	if err != nil {
		return returnError()
	}

	if strings.Trim(unquoted, " ") == "" {
		*bl = false
		return nil
	}

	if strings.ToLower(unquoted) == "false" {
		*bl = false
		return nil
	}

	if strings.ToLower(unquoted) == "true" {
		*bl = true
		return nil
	}

	i, err := strconv.ParseInt(unquoted, 10, 64)
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

func (bl BoolString) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprintf("%v", bl))
}

func (bl *BoolString) ValuePtr() *bool {
	if bl == nil {
		return nil
	}

	_b := bool(*bl)
	return &_b
}

func (bl BoolString) Value() bool {
	return bool(bl)
}
