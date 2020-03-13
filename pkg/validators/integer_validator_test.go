package validators

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntegerValidator_Validate(t *testing.T) {
	testCases := []struct {
		desc     string
		source   int
		target   int
		expected bool
	}{
		{desc: "Source and Target are equal", source: 1, target: 1, expected: true},
		{desc: "Source and Target are not equal", source: 1, target: 3, expected: false},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			validator := new(IntegerValidator)
			valid, _ := validator.Validate(tc.source, tc.target)
			assert.Equal(t, tc.expected, valid)
		})
	}
}
