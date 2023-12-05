package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormatStrReturnCorrectValue(t *testing.T) {
	var text = "my name is ahmed , my age is 35"
	assert.Equal(t, text, FormatStr("my name is %s , my age is %d", "ahmed", 35))
}
