package rtda

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNilLocalVar(t *testing.T) {
	localVars := newLocalVars(0)
	assert.Nil(t, localVars, "")
}

func TestInt(t *testing.T) {
	localVars := newLocalVars(10)
	for i := range localVars {
		v := rand.Int31()
		localVars.SetInt(uint8(i), v)
		assert.Equal(t, v, localVars.GetInt(uint8(i)), "")
	}
}

func TestFloat(t *testing.T) {
	localVars := newLocalVars(10)
	for i := range localVars {
		v := rand.Float32()
		localVars.SetFloat(uint8(i), v)
		assert.Equal(t, v, localVars.GetFloat(uint8(i)), "")
	}
}

func TestLong(t *testing.T) {
	localVars := newLocalVars(10)
	for i := 0; i < 5; i++ {
		v := rand.Int63()
		localVars.SetLong(uint8(i*2), v)
		assert.Equal(t, v, localVars.GetLong(uint8(i*2)), "")
	}
}

func TestDouble(t *testing.T) {
	localVars := newLocalVars(10)
	for i := 0; i < 5; i++ {
		v := rand.Float64()
		localVars.SetDouble(uint8(i*2), v)
		assert.Equal(t, v, localVars.GetDouble(uint8(i*2)), "")
	}
}
