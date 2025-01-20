package syncs_test

import (
	"testing"

	"github.com/Joaolfc0/goutil/syncs"
	"github.com/Joaolfc0/goutil/testutil/assert"
)

func TestGo(t *testing.T) {
	err := syncs.Go(func() error {
		return nil
	})
	assert.NoErr(t, err)
}
