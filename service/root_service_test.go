package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRootService(t *testing.T) {
	assert.NotPanics(t, func() {
		_ = NewRootService(1).(*rootServiceImpl)
	})
}
