package mathutil

import "github.com/Joaolfc0/goutil/comdef"

// Abs get absolute value of given value
func Abs[T comdef.Int](val T) T {
	if val >= 0 {
		return val
	}
	return -val
}
