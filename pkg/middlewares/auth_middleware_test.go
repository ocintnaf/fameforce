package middlewares

import (
	"reflect"
	"testing"

	"github.com/ocintnaf/fameforce/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNewAuthMiddleware(t *testing.T) {
	t.Run("should return a new authentication middleware", func(t *testing.T) {
		idTokenVerifierMock := mocks.NewIDTokenVerifierMock()

		actual := NewAuthMiddleware(idTokenVerifierMock)

		assert.NotNil(t, actual)
		assert.True(t, reflect.TypeOf(actual).Kind() == reflect.Func)
	})
}
