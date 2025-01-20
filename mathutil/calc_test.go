package mathutil_test

import (
	"testing"

	"github.com/Joaolfc0/goutil/mathutil"
	"github.com/Joaolfc0/goutil/testutil/assert"
)

func TestAbs(t *testing.T) {
	assert.Eq(t, 1, mathutil.Abs(1))
	assert.Eq(t, 1, mathutil.Abs(-1))
}
