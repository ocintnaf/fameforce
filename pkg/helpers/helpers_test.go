package helpers

import (
	"errors"
	"testing"

	"github.com/ocintnaf/fameforce/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testExpectation struct {
	bearerToken string
	err         error
}

type test struct {
	desc       string
	authHeader string
	expected   testExpectation
}

func TestGetBearerToken(t *testing.T) {
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
