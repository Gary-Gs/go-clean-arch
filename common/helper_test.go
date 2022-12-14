package common

import (
	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

// GoMonkey function mocking usage example
func TestContainsIgnoreCase(t *testing.T) {
	p := gomonkey.ApplyFunc(strings.ToUpper, func(s string) string {
		return strings.ToLower(s)
	})
	defer p.Reset()

	assert.Equal(t, true, ContainsIgnoreCase([]string{"a", "b", "c"}, "A"))
}

func TestMin(t *testing.T) {
	assert.Equal(t, -3.2, Min(-3.2, 0, 1.6))
}

func TestMax(t *testing.T) {
	assert.Equal(t, float64(8), Max(-3.2, 0, 8, 1.6))
}

func TestRemoveElementByIndex(t *testing.T) {
	// test string
	assert.Equal(t, []string{"a", "c"}, RemoveElementByIndex([]string{"a", "b", "c"}, 1))
	// test int
	assert.Equal(t, []int{1, 3}, RemoveElementByIndex([]int{1, 2, 3}, 1))
	// test float64
	assert.Equal(t, []float64{1.1, 3.3}, RemoveElementByIndex([]float64{1.1, 2.2, 3.3}, 1))
	// test out of bound index
	assert.Equal(t, []string{"a", "b", "c"}, RemoveElementByIndex([]string{"a", "b", "c"}, 3))
	// test negative range bound
	assert.Equal(t, []string{"a", "b", "c"}, RemoveElementByIndex([]string{"a", "b", "c"}, -1))
}

func TestInsertElementByIndex(t *testing.T) {
	// test string
	assert.Equal(t, []string{"a", "d", "b", "c"}, InsertElementByIndex([]string{"a", "b", "c"}, 1, "d"))
	// test int
	assert.Equal(t, []int{1, 4, 2, 3}, InsertElementByIndex([]int{1, 2, 3}, 1, 4))
	// test float64
	assert.Equal(t, []float64{1.1, 4.4, 2.2, 3.3}, InsertElementByIndex([]float64{1.1, 2.2, 3.3}, 1, 4.4))
	// test out of bound index
	assert.Equal(t, []string{"a", "b", "c"}, InsertElementByIndex([]string{"a", "b", "c"}, 4, "d"))
	// test negative range bound
	assert.Equal(t, []string{"a", "b", "c"}, InsertElementByIndex([]string{"a", "b", "c"}, -1, "d"))
}
