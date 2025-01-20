package panics_test

import (
	"testing"

	"github.com/Joaolfc0/goutil/errorx/panics"
	"github.com/Joaolfc0/goutil/testutil/assert"
)

func TestIsTrue(t *testing.T) {
	assert.Panics(t, func() {
		panics.IsTrue(false)
	})
}
