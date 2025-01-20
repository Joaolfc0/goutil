package sysutil_test

import (
	"os"
	"runtime"
	"testing"

	"github.com/Joaolfc0/goutil/sysutil"
	"github.com/Joaolfc0/goutil/testutil/assert"
)

func TestBasic_usage(t *testing.T) {
	assert.NotEmpty(t, sysutil.BinDir())
	assert.NotEmpty(t, sysutil.BinName())
	assert.NotEmpty(t, sysutil.BinFile())
}

func TestProcessExists(t *testing.T) {
	if runtime.GOOS != "windows" {
		pid := os.Getpid()
		assert.True(t, sysutil.ProcessExists(pid))
	} else {
		t.Skip("on Windows")
	}
}
