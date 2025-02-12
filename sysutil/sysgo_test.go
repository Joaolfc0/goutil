package sysutil_test

import (
	"testing"

	"github.com/Joaolfc0/goutil/dump"
	"github.com/Joaolfc0/goutil/sysutil"
	"github.com/Joaolfc0/goutil/testutil/assert"
)

func TestGoVersion(t *testing.T) {
	assert.NotEmpty(t, sysutil.GoVersion())

	info, err := sysutil.ParseGoVersion("go version go1.19.2 darwin/amd64")
	assert.NoErr(t, err)
	assert.NotEmpty(t, info)
	assert.Eq(t, "1.19.2", info.Version)
	assert.Eq(t, "darwin", info.GoOS)
	assert.Eq(t, "amd64", info.Arch)

	info, err = sysutil.OsGoInfo()
	assert.NoErr(t, err)
	assert.NotEmpty(t, info)
	dump.P(info)
}
