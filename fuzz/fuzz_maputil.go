//go:build gofuzz
// +build gofuzz

package fuzz

import (
	"fmt"
	"reflect"

	"github.com/gookit/goutil/maputil"
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
			fmt.Printf("Encontrado valor string em key: %s\n", key)
		}
	})
	maputil.EachTypedMap(map[string]int{"key1": 1, "key2": 2}, func(key string, val int) {
		if val%2 == 0 {
			fmt.Printf("Encontrado valor par: %d\n", val)
		}
	})

	return 1
}
