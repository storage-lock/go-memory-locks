package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainRuns(t *testing.T) {
	main()
}

func TestMustPanics(t *testing.T) {
	assert.Panics(t, func() { must(errors.New("boom")) })
	assert.NotPanics(t, func() { must(nil) })
}
