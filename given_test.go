package fizzbuzz_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Given struct {
	t      *testing.T
	actual interface{}
}

func (g Given) ShouldBe(expected interface{}) {
	assert.Equal(g.t, expected, g.actual)
}
