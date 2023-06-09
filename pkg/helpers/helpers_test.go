package helpers

import (
	"errors"
	"testing"

	"github.com/ocintnaf/fameforce/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetBearerToken(t *testing.T) {
	type testExpectation struct {
		bearerToken string
		err         error
	}

	type test struct {
		desc       string
		authHeader string
		expected   testExpectation
	}

	tests := []test{
		{
			desc:       "should return an error if Authorization header is missing",
			authHeader: "",
			expected: testExpectation{
				bearerToken: "",
				err:         errors.New("invalid bearer token format"),
			},
		},
		{
			desc:       "should return an error if Authorization header is malformed",
			authHeader: "invalid_token",
			expected: testExpectation{
				bearerToken: "",
				err:         errors.New("invalid bearer token format"),
			},
		},
		{
			desc:       "should return a bearer token",
			authHeader: "Bearer my_token",
			expected: testExpectation{
				bearerToken: "my_token",
				err:         nil,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			headerGetterMock := mocks.NewHeaderGetterMock()

			headerGetterMock.On("Get", "Authorization", mock.Anything).Return(test.authHeader)

			actualBearerToken, actualErr := GetBearerToken(headerGetterMock)

			headerGetterMock.AssertCalled(t, "Get", "Authorization", mock.Anything)

			assert.Equal(t, test.expected.bearerToken, actualBearerToken)
			assert.Equal(t, actualErr, test.expected.err)
		})
	}
}

func TestValidate(t *testing.T) {
	type toValidate struct {
		ID  string `json:"id" validate:"uuid,required"`
		Age int    `json:"age" validate:"numeric,min=0,required"`
	}

	type test struct {
		desc     string
		arg      toValidate
		expected ValidationErrors
	}

	tests := []test{
		{
			desc: "should return a list of validation errors [invalid id]",
			arg:  toValidate{"invalid-id", 5},
			expected: ValidationErrors{
				"id": "id must be a valid UUID",
			},
		},
		{
			desc: "should return a list of validation errors [missing age]",
			arg:  toValidate{"945671c2-d2a3-42b3-9e8f-3295d7f86702", 0},
			expected: ValidationErrors{
				"age": "age is a required field",
			},
		},
		{
			desc: "should return a list of validation errors [invalid id, invalid age]",
			arg:  toValidate{"invalid-id", -1},
			expected: ValidationErrors{
				"id":  "id must be a valid UUID",
				"age": "age must be 0 or greater",
			},
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

func TestGetCtxLocal(t *testing.T) {
	t.Run("should return an error if the given key is not present", func(t *testing.T) {
		ctxLocalerMock := mocks.NewCtxLocalerMock()

		ctxLocalerMock.On("Locals", mock.Anything, mock.Anything).Return(nil)

		actualValue, actualErr := GetCtxLocal[string](ctxLocalerMock, "MyKey")

		expectedError := "key: MyKey not found in context"

		assert.Equal(t, "", actualValue)
		assert.EqualError(t, actualErr, expectedError)
		ctxLocalerMock.AssertCalled(t, "Locals", "MyKey", mock.Anything)
	})
}
