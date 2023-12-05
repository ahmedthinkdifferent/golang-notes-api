package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToInt(t *testing.T) {
	var num string = "1234"
	assert.Equal(t, 1234, ToInt(num))
}

func TestToIntReturnZeroForNonIntVal(t *testing.T) {
	var num string = "1234x"
	assert.Equal(t, 0, ToInt(num))
}
