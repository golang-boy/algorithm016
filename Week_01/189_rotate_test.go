package Week_01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotate1(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	k := 6
	rotate1(nums, k)
	assert.Equal(t, nums, []int{4, 5, 6, 7, 8, 9, 1, 2, 3})
}

func TestRotate2(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	k := 6
	rotate2(nums, k)
	assert.Equal(t, nums, []int{4, 5, 6, 7, 8, 9, 1, 2, 3})
}
