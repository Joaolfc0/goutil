//go:build gofuzz
// +build gofuzz

package fuzz

import (
	"fmt"
	"regexp"

	"github.com/gookit/goutil/strutil"
)

// FuzzStrutil é um fuzz target para a biblioteca `strutil`.
func FuzzStrutil(data []byte) int {

	input := string(data)

	// Testa funções relacionadas a parsing de strings
	_ = strutil.IsNumeric(input)
	_ = strutil.IsVersion(input)

	// Testa funções de prefixo/sufixo
	_ = strutil.HasPrefix(input, "test")
	_ = strutil.HasSuffix(input, "end")

	// Testa funções de manipulação de substrings
	_ = strutil.SimpleMatch(input, []string{"test", "sample", "check$"})
	_ = strutil.LikeMatch("prefix%", input)

	// Testa funções relacionadas a padrões globais
	_ = strutil.GlobMatch("a*/b", input)
	_ = strutil.PathMatch("a*/b", input)

	// Testa consistência com expressões regulares
	verRegex := regexp.MustCompile(`^[0-9][\d.]+(-\w+)?$`)
	matched := verRegex.MatchString(input)
	if matched && !strutil.IsVersion(input) {
		panic(fmt.Sprintf("Mismatch between verRegex and IsVersion for input: %q", input))
	}

	// // Testa funções de conversão de bases
	// for _, base := range []int{2, 8, 10, 16, 32, 36, 62, 64} {
	// 	// Tenta converter input como uma string em base 10 para outra base
	// 	_ = strutil.Base10Conv(input, base)

	// 	// Converte entre diferentes bases usando BaseConv
	// 	for _, fromBase := range []int{2, 10, 16} {
	// 		_ = strutil.BaseConv(input, fromBase, base)
	// 	}
	// }

	// Testa BaseConvByTpl diretamente com diferentes templates
	// _ = strutil.BaseConvByTpl(input, strutil.Base10Chars, strutil.Base16Chars)
	// _ = strutil.BaseConvByTpl(input, strutil.Base62Chars, strutil.Base36Chars)
	// _ = strutil.BaseConvByTpl(input, strutil.Base64Chars, strutil.Base32Chars)

	// Testa funções de similaridade
	comp := strutil.NewComparator(input, "example")
	rate, ok := comp.Similar(0.5)
	if !ok {
		fmt.Printf("Similarity too low: %.2f\n", rate)
	}

	rate, ok = strutil.Similarity(input, "example", 0.5)
	if !ok {
		fmt.Printf("Global Similarity too low: %.2f\n", rate)
	}

	return 1
}

// FuzzRandomChars testa as funções de geração de strings aleatórias no arquivo strutil.
func FuzzRandomChars(data []byte) int {
	if len(data) == 0 {
		return 0
	}

	input := len(data) % 100 // Limita o tamanho para evitar inputs muito grandes

	// Testa RandomChars
	generated := strutil.RandomChars(input)
	if len(generated) != input {
		panic(fmt.Sprintf("RandomChars generated string has wrong length: expected %d, got %d", input, len(generated)))
	}

	// Testa RandomCharsV2
	generatedV2 := strutil.RandomCharsV2(input)
	if len(generatedV2) != input {
		panic(fmt.Sprintf("RandomCharsV2 generated string has wrong length: expected %d, got %d", input, len(generatedV2)))
	}

	// Testa RandomCharsV3
	generatedV3 := strutil.RandomCharsV3(input)
	if len(generatedV3) != input {
		panic(fmt.Sprintf("RandomCharsV3 generated string has wrong length: expected %d, got %d", input, len(generatedV3)))
	}

	// Testa RandWithTpl
	template := "0123456789abcdef"
	randTpl := strutil.RandWithTpl(input, template)
	if len(randTpl) != input {
		panic(fmt.Sprintf("RandWithTpl generated string has wrong length: expected %d, got %d", input, len(randTpl)))
	}
	for _, char := range randTpl {
		if !containsRune(template, char) {
			panic(fmt.Sprintf("RandWithTpl generated string contains invalid character: %c", char))
		}
	}

	// Testa RandomString
	_, err := strutil.RandomString(input)
	if err != nil {
		panic(fmt.Sprintf("RandomString returned an error: %v", err))
	}

	return 1
}

func containsRune(s string, r rune) bool {
	for _, char := range s {
		if char == r {
			return true
		}
	}
	return false
}
