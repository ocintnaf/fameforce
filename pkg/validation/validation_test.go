package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	desc     string
	arg      toValidate
	expected []string
}

type toValidate struct {
	Id  string `validate:"uuid,required"`
	Age int    `validate:"numeric,min=0,required"`
}

func TestValidate(t *testing.T) {
	tests := []test{
		{
			desc:     "should return a list of validation errors [invalid id]",
			arg:      toValidate{"invalid-id", 5},
			expected: []string{"Id must be a valid UUID"},
		},
		{
			desc:     "should return a list of validation errors [missing age]",
			arg:      toValidate{"945671c2-d2a3-42b3-9e8f-3295d7f86702", 0},
			expected: []string{"Age is a required field"},
		},
		{
			desc:     "should return a list of validation errors [invalid id, invalid age]",
			arg:      toValidate{"invalid-id", -1},
			expected: []string{"Id must be a valid UUID", "Age must be 0 or greater"},
		},
		{
			desc:     "should return an empty list of validation errors [valid fields]",
			arg:      toValidate{"945671c2-d2a3-42b3-9e8f-3295d7f86702", 5},
			expected: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			actual := Validate(test.arg)

			assert.Equal(t, test.expected, actual)
		})
	}
}
