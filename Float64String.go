package types

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	errortools "github.com/leapforce-libraries/go_errortools"
)

type Float64String float64

func (f *Float64String) UnmarshalJSON(b []byte) error {
	var returnError = func() error {
		errortools.CaptureError(fmt.Sprintf("Cannot parse '%s' to Float64String", string(b)))
		return nil
	}

	// first try to parse to float64 directly
	var ii float64

	err := json.Unmarshal(b, &ii)
	if err == nil {
		*f = Float64String(ii)
		return nil
	}

	var s string

	err = json.Unmarshal(b, &s)
	if err != nil {
		return returnError()
	}

	s = strings.Trim(s, " ")

	if s == "" {
		return nil
	}

	/* handle exponential number without 'E' */
	re := regexp.MustCompile(`\d(-|\+)\d`)

	f1 := re.Find([]byte(s))
	if f1 != nil {
		s1 := string(f1)
		s = strings.Replace(s, s1, s1[:1]+"E"+s1[len(s1)-2:], 1)
	}
	/* */

	_f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return returnError()
	}

	*f = Float64String(_f)
	return nil
}

func (f Float64String) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprintf("%f", f))
}

func (f *Float64String) ValuePtr() *float64 {
	if f == nil {
		return nil
	}

	_f := float64(*f)
	return &_f
}

func (f Float64String) Value() float64 {
	return float64(f)
}
