package validator

import "fmt"

func (v *Validator) MinLength(field, value string, high int) bool {

	if _, ok := v.Errors[field]; ok {
		return false
	}
	if len(value) < high {
		v.Errors[field] = fmt.Sprint("%s must be atleast (%%D) characters long", field, high)
		return false
	}
	return true
}
