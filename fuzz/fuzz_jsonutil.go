//go:build gofuzz
// +build gofuzz

package fuzz

import (
	"fmt"

	"github.com/Joaolfc0/goutil/jsonutil"
)

// FuzzJSONUtil testa as funções do pacote jsonutil.
func FuzzJSONUtil(data []byte) int {
	input := string(data)

	// Testa validação de JSON
	isJSON := jsonutil.IsJSON(input)
	_ = jsonutil.IsArray(input)
	_ = jsonutil.IsObject(input)

	// Consistência entre IsJSONFast e IsJSON
	isFast := jsonutil.IsJSONFast(input)
	if isFast && !isJSON {
		panic(fmt.Sprintf("Inconsistência: IsJSONFast retornou true, mas IsJSON retornou false para: %q", input))
	}

	// Testa remoção de comentários
	withoutComments := jsonutil.StripComments(input)
	if jsonutil.IsJSON(withoutComments) && !isJSON {
		panic(fmt.Sprintf("Inconsistência: StripComments transformou um JSON inválido em válido: %q", input))
	}

	// Testa transformação e formatação
	var v interface{}
	err := jsonutil.Mapping(input, &v)
	if err == nil {
		formatted, _ := jsonutil.Pretty(v)
		if !jsonutil.IsJSON(formatted) {
			panic(fmt.Sprintf("Pretty produziu JSON inválido: %q", formatted))
		}
	}

	return 1
}
