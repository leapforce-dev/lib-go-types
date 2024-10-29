package types

import (
	"encoding/json"
	"fmt"
	errortools "github.com/leapforce-dev/lib-go-errortools"
	"regexp"
	"strconv"
	"strings"
)

type Float64String struct {
	f      float64
	format string
}

func (f *Float64String) UnmarshalJSON(b []byte) error {
	var returnError = func() error {
		errortools.CaptureError(fmt.Sprintf("Cannot parse '%s' to Float64String", string(b)))
		return nil
	}

	// first try to parse to float64 directly
	var ii float64

	err := json.Unmarshal(b, &ii)
	if err == nil {
		*f = Float64String{ii, extractFormat(fmt.Sprintf("%v", ii))}
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

	*f = Float64String{_f, extractFormat(s)}
	return nil
}

func extractFormat(s string) string {
	s_ := strings.Split(s, ".")
	if len(s_) == 1 {
		return "%.0f"
	} else if len(s_) == 2 {
		return "%" + fmt.Sprintf(".%vf", len(s_[1]))
	}

	return "%f"
}

func (f Float64String) MarshalJSON() ([]byte, error) {
	format := "%f"
	if f.format != "" {
		format = f.format
	}
	return json.Marshal(fmt.Sprintf(format, f.f))
}

func (f *Float64String) ValuePtr() *float64 {
	if f == nil {
		return nil
	}

	return &f.f
}

func (f Float64String) Value() float64 {
	return f.f
}

func NewFloat64String(f float64) Float64String {
	return Float64String{
		f:      f,
		format: extractFormat(fmt.Sprintf("%v", f)),
	}
}

func NewFloat64StringFromFloat32(f32 float32) Float64String {
	s := fmt.Sprintf("%v", f32)
	f, _ := strconv.ParseFloat(s, 64)
	return Float64String{
		f:      f,
		format: extractFormat(s),
	}
}

func (f *Float64String) SetFormat(format string) {
	if f != nil {
		f.format = format
	}
}
