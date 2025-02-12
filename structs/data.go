package structs

import (
	"sync"

	"github.com/Joaolfc0/goutil/internal/comfunc"
	"github.com/Joaolfc0/goutil/maputil"
	"github.com/Joaolfc0/goutil/mathutil"
	"github.com/Joaolfc0/goutil/strutil"
)

// LiteData simple map[string]any struct. no lock
type LiteData = Data

// NewLiteData create, not locked
func NewLiteData(data map[string]any) *Data {
	if data == nil {
		data = make(map[string]any)
	}
	return &LiteData{data: data}
}

/*************************************************************
 * data struct and allow enable lock
 *************************************************************/

// Data struct, allow enable lock
type Data struct {
	sync.RWMutex
	lock bool
	data map[string]any
}

// NewData create new data instance
func NewData() *Data {
	return &Data{
		lock: true,
		data: make(map[string]any),
	}
}

// WithLock for operate data
func (d *Data) WithLock() *Data {
	d.lock = true
	return d
}

// EnableLock for operate data
func (d *Data) EnableLock() *Data {
	return d.WithLock()
}

// Data get all
func (d *Data) Data() map[string]any {
	return d.data
}

// SetData set all data
func (d *Data) SetData(data map[string]any) {
	if !d.lock {
		d.data = data
		return
	}

	d.RLock()
	d.data = data
	d.RUnlock()
}

// DataLen of data
func (d *Data) DataLen() int {
	return len(d.data)
}

// ResetData all data
func (d *Data) ResetData() {
	d.data = make(map[string]any)
}

// Merge load new data
func (d *Data) Merge(mp map[string]any) {
	d.data = maputil.SimpleMerge(mp, d.data)
}

// Set value to data
func (d *Data) Set(key string, val any) {
	d.SetValue(key, val)
}

// SetValue to data
func (d *Data) SetValue(key string, val any) {
	if d.lock {
		d.Lock()
		defer d.Unlock()
	}

	d.data[key] = val
}

// Value get from data
func (d *Data) Value(key string) (val any, ok bool) {
	if d.lock {
		d.RLock()
		defer d.RUnlock()
	}

	val, ok = maputil.GetByPath(key, d.data)
	return
}

// Get val from data
func (d *Data) Get(key string) any {
	return d.GetVal(key)
}

// GetVal get from data
func (d *Data) GetVal(key string) any {
	if d.lock {
		d.RLock()
		defer d.RUnlock()
	}

	val, _ := maputil.GetByPath(key, d.data)
	return val
}

// StrVal get from data
func (d *Data) StrVal(key string) string {
	return strutil.QuietString(d.GetVal(key))
}

// IntVal get from data
func (d *Data) IntVal(key string) int {
	return mathutil.QuietInt(d.GetVal(key))
}

// BoolVal get from data
func (d *Data) BoolVal(key string) bool {
	val, ok := d.Value(key)
	if !ok {
		return false
	}
	return comfunc.Bool(val)
}

// String format data
func (d *Data) String() string {
	return maputil.ToString(d.data)
}

// OrderedData data TODO
type OrderedData struct {
	maputil.Data
	cap  int
	keys []string
	// vals []any
}

// NewOrderedData instance.
func NewOrderedData(cap int) *OrderedData {
	return &OrderedData{cap: cap, Data: make(maputil.Data, cap)}
}

// Load data
func (om *OrderedData) Load(data map[string]any) {
	om.Data.Load(data)
	om.keys = om.Data.Keys()
}

// Set key and value to map
func (om *OrderedData) Set(key string, val any) {
	om.keys = append(om.keys, key)
	om.Data.Set(key, val)
}
