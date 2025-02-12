package testutil

import (
	"github.com/Joaolfc0/goutil/byteutil"
)

// Buffer wrap and extends the bytes.Buffer
type Buffer = byteutil.Buffer

// NewBuffer instance
func NewBuffer() *byteutil.Buffer {
	return byteutil.NewBuffer()
}
