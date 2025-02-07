package reflects_test

import (
	"reflect"
	"testing"

	"github.com/Joaolfc0/goutil/dump"
	"github.com/Joaolfc0/goutil/reflects"
	"github.com/Joaolfc0/goutil/testutil/assert"
)

func TestMakeSliceByElem(t *testing.T) {
	slv := reflects.MakeSliceByElem(reflect.TypeOf("str"), 0, 2)
	slv = reflect.Append(slv, reflect.ValueOf("a"))
	slv = reflect.Append(slv, reflect.ValueOf("b"))

	sl := slv.Interface().([]string)
	dump.P(sl)
	assert.Len(t, sl, 2)
	assert.Eq(t, "a", sl[0])
}

func TestFlatSlice(t *testing.T) {

}
