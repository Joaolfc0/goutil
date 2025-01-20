//go:build gofuzz
// +build gofuzz

package fuzz

import (
	"reflect"
	"strings"

	"github.com/Joaolfc0/goutil/maputil"
)

// FuzzMapUtil testa as funções do pacote maputil.
func FuzzMapUtil(data []byte) int {
	// Converte os dados de entrada em uma string
	input := string(data)

	// Cria mapas para teste
	testMap := map[string]any{
		"key1": "value1",
		"key2": map[string]any{
			"subkey1": "subvalue1",
			"subkey2": []map[string]any{
				{"nestedkey1": "nestedvalue1"},
			},
		},
		"key3": []string{"value3_1", "value3_2"},
	}

	// Testa DeepGet e QuietGet
	_ = maputil.DeepGet(testMap, input)
	_ = maputil.QuietGet(testMap, input)

	// Testa GetFromAny
	_, _ = maputil.GetFromAny(input, testMap)

	// Testa GetByPath e GetByPathKeys
	_, _ = maputil.GetByPath(input, testMap)
	_, _ = maputil.GetByPathKeys(testMap, []string{"key2", "subkey2", "0", "nestedkey1"})

	// Testa Keys, TypedKeys, Values, TypedValues
	_ = maputil.Keys(testMap)
	_ = maputil.TypedKeys(map[string]int{"key1": 1, "key2": 2})
	_ = maputil.Values(testMap)
	_ = maputil.TypedValues(map[string]int{"key1": 1, "key2": 2})

	// Testa EachAnyMap e EachTypedMap
	maputil.EachAnyMap(testMap, func(key string, val any) {
		if reflect.ValueOf(val).Kind() == reflect.String && key == input {
		}
	})
	maputil.EachTypedMap(map[string]int{"key1": 1, "key2": 2}, func(key string, val int) {
		if val%2 == 0 {
		}
	})

	return 1
}

func FuzzMapUtilConv(data []byte) int {
	// Converte os dados de entrada em uma string
	input := string(data)

	// Teste para KeyToLower
	testMapStr := map[string]string{
		input:  "Value1",
		"KEY2": input,
	}
	_ = maputil.KeyToLower(testMapStr)

	// Teste para ToStringMap
	testMapAny := map[string]any{
		"key1": input,
		"key2": len(input),
		"key3": strings.Contains(input, "test"),
	}
	_ = maputil.ToStringMap(testMapAny)

	// Teste para ToAnyMap e TryAnyMap
	_ = maputil.ToAnyMap(testMapStr)
	_, _ = maputil.TryAnyMap(testMapStr)

	// Teste para CombineToSMap
	keys := []string{input, "key2", "key3"}
	values := []string{"value1", input}
	_ = maputil.CombineToSMap(keys, values)

	// Teste para HTTPQueryString
	queryMap := map[string]any{
		"param1": input,
		"param2": len(input),
		"param3": strings.Contains(input, "query"),
	}
	_ = maputil.HTTPQueryString(queryMap)

	// Teste para StringsMapToAnyMap
	strMap := map[string][]string{
		input:  {"value1", "value2"},
		"key2": {input},
	}
	_ = maputil.StringsMapToAnyMap(strMap)

	// Teste para ToString
	_ = maputil.ToString(testMapAny)

	// Teste para Flatten
	treeMap := map[string]any{
		input: map[string]any{
			"level2": input,
			"level3": map[string]any{
				"level4": len(input),
			},
		},
	}
	_ = maputil.Flatten(treeMap)

	return 1
}
