package validators

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
)

type IntegerValidator struct {
	logger *logrus.Entry
}

func NewIntegerValidator(logger *logrus.Entry) *IntegerValidator {
	return &IntegerValidator{logger: logger}
}

func (i *IntegerValidator) Validate(source, target interface{}) (bool, error) {
	s := source.(int)
	t := target.(int)

	if s == t {
		i.logger.Debugf("%d is equal to %d", s, t)
		return true, nil
	}

	err := errors.New(fmt.Sprintf("%d is not equal to %d", s, t))

	return false, err
}
