package testutil_test

import (
	"testing"

	"github.com/Joaolfc0/goutil/testutil"
	"github.com/Joaolfc0/goutil/testutil/assert"
)

func TestNewBuffer(t *testing.T) {
	buf := testutil.NewBuffer()

	buf.WriteStr("ab", "-", "cd")
	assert.Eq(t, "ab-cd", buf.ResetGet())

	buf.WriteAny(23, "abc")
	assert.Eq(t, "23abc", buf.ResetAndGet())

	buf.Writeln("abc")
	assert.Eq(t, "abc\n", buf.ResetAndGet())
}
