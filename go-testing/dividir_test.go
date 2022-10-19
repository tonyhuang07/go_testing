package testing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividir(t *testing.T) {
	result, err := dividir(10, 0)
	assert.Equal(t, 0, result)
	assert.Equal(t, "el dominador no puede ser 0", err.Error())

}
