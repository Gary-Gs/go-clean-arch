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
