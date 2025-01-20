//go:build gofuzz
// +build gofuzz

package fuzz

import (
	"reflect"

	"github.com/Joaolfc0/goutil/reflects"
)

// FuzzReflects testa as funções do pacote reflects.
func FuzzReflects(data []byte) int {
	// Converte os dados de entrada em uma string
	input := string(data)

	// Cria valores de teste usando reflection
	testValue := reflect.ValueOf(input)

	// Testa HasChild
	_ = reflects.HasChild(testValue)

	// Testa IsArrayOrSlice
	_ = reflects.IsArrayOrSlice(testValue.Kind())

	// Testa IsSimpleKind
	_ = reflects.IsSimpleKind(testValue.Kind())

	// Testa IsAnyInt e IsIntLike
	_ = reflects.IsAnyInt(testValue.Kind())
	_ = reflects.IsIntLike(testValue.Kind())

	// Testa IsIntx e IsUintX
	_ = reflects.IsIntx(testValue.Kind())
	_ = reflects.IsUintX(testValue.Kind())

	// Testa IsNil e IsValidPtr
	_ = reflects.IsNil(testValue)
	_ = reflects.IsValidPtr(testValue)

	// Testa CanBeNil
	if testType := testValue.Type(); testType != nil {
		_ = reflects.CanBeNil(testType)
	}

	// Testa IsFunc
	_ = reflects.IsFunc(func() {})

	// Testa IsEqual
	_ = reflects.IsEqual(input, string(data))

	// Testa IsEmpty, IsEmptyValue e IsEmptyReal
	_ = reflects.IsEmpty(testValue)
	_ = reflects.IsEmptyValue(testValue)
	_ = reflects.IsEmptyReal(testValue)

	return 1
}

func FuzzReflectsConv(data []byte) int {
	if len(data) == 0 {
		return 0
	}

	// Converte os dados de entrada em uma string
	input := string(data)

	// Testa BaseTypeVal e ToBaseVal
	value := reflect.ValueOf(input)
	_, _ = reflects.BaseTypeVal(value)
	_, _ = reflects.ToBaseVal(value)

	// Testa ConvToType e ValueByType
	_, _ = reflects.ConvToType(input, reflect.TypeOf(input))
	_, _ = reflects.ValueByType(input, reflect.TypeOf(input))

	// Testa ConvToKind e ValueByKind
	_, _ = reflects.ConvToKind(input, reflect.String)
	_, _ = reflects.ValueByKind(input, reflect.String)

	// Testa ConvSlice com valores diversos
	oldSlice := reflect.ValueOf([]string{input})
	_, _ = reflects.ConvSlice(oldSlice, reflect.TypeOf(input))

	// Testa String e ToString
	reflectVal := reflect.ValueOf(input)
	_ = reflects.String(reflectVal)
	_, _ = reflects.ToString(reflectVal)

	// Testa ValToString com diferentes configurações
	_, _ = reflects.ValToString(reflectVal, false)
	_, _ = reflects.ValToString(reflectVal, true)

	return 1
}
