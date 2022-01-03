package types

import (
	"encoding/json"
	"fmt"
	"strings"

	"cloud.google.com/go/civil"
	errortools "github.com/leapforce-libraries/go_errortools"
)

type TimeString civil.Time

func (ts *TimeString) UnmarshalJSON(b []byte) error {
	var returnError = func() error {
		errortools.CaptureError(fmt.Sprintf("Cannot parse '%s' to TimeString", string(b)))
		return nil
	}

	var s string

	err := json.Unmarshal(b, &s)
	if err != nil {
		return returnError()
	}

	s = strings.Trim(s, " ")

	if s == "" {
		ts = nil
		return nil
	}

	t, err := civil.ParseTime(s)
	if err != nil {
		return returnError()
	}

	*ts = TimeString(t)
	return nil
}

func (ts TimeString) MarshalJSON() ([]byte, error) {
	return json.Marshal(civil.Time(ts).String())
}

func (ts *TimeString) ValuePtr() *civil.Time {
	if ts == nil {
		return nil
	}

	_t := civil.Time(*ts)
	return &_t
}

func (ts TimeString) Value() civil.Time {
	return civil.Time(ts)
}
