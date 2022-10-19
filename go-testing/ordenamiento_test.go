package testing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	arr := []int{10, 7, 8, 9, 1, 5}
	quickSort(arr, 0, len(arr)-1)
	assert.Equal(t, []int{1, 5, 7, 8, 9, 10}, arr)
}
