package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var hashedPassword string

func TestHashPasswordReturnValidHash(t *testing.T) {
	hash, err := HashPassword("ahmed1")
	assert.NotEmpty(t, hash)
	assert.Nil(t, err)
	hashedPassword = hash
}

func TestIsPasswordCorrectReturnTrue(t *testing.T) {
	samePassword := IsPasswordCorrect("ahmed1", hashedPassword)
	assert.True(t, samePassword)
}

func TestIsPasswordCorrectReturnFalseIfPasswordNotCorrect(t *testing.T) {
	samePassword := IsPasswordCorrect("ahmed2", hashedPassword)
	assert.False(t, samePassword)
}
