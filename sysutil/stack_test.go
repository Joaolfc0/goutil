package sysutil_test

import (
	"testing"

	"github.com/Joaolfc0/goutil/sysutil"
	"github.com/Joaolfc0/goutil/testutil/assert"
)

func TestCallersInfo(t *testing.T) {
	cs := sysutil.CallersInfos(0, 2)
	// dump.P(cs)
	assert.NotEmpty(t, cs)
	assert.Len(t, cs, 2)
	assert.StrContains(t, cs[0].String(), "goutil/sysutil/stack.go")
}
