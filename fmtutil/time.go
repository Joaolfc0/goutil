package fmtutil

import (
	"github.com/Joaolfc0/goutil/mathutil"
)

// HowLongAgo format a seconds, get how lang ago
func HowLongAgo(sec int64) string {
	return mathutil.HowLongAgo(sec)
}
