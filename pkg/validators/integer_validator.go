package validators

import "fmt"

type IntegerValidator struct {
}

func (i *IntegerValidator) Validate(source, target interface{}) (bool, error) {
	s := source.(int)
	t := target.(int)

	if s == t {
		return true, nil
	}

	return false, fmt.Errorf("%d is not equal to %d", s, t)
}
