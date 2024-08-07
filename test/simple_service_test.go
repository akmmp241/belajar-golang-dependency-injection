package test

import (
	"akmmp241/belajar-golang-restful-api/simple"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleService(t *testing.T) {
	t.Run("Error", func(t *testing.T) {
		simpleService, err := simple.InitializeService(true)
		assert.Nil(t, simpleService)
		assert.NotNil(t, err)
	})

	t.Run("Not Error", func(t *testing.T) {
		simpleService, err := simple.InitializeService(false)
		assert.Nil(t, err)
		assert.NotNil(t, simpleService)
	})
}
