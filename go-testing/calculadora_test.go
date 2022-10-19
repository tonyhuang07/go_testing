package testing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumar(t *testing.T) {
	assert.Equal(t, 10, Sumar(3, 7))
}

func TestRestar(t *testing.T) {
	assert.Equal(t, 10, Restar(17, 7))
}
