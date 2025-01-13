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

	// Testa funções de conversão de bases
	for _, base := range []int{2, 8, 10, 16, 32, 36, 62, 64} {
		// Tenta converter input como uma string em base 10 para outra base
		_ = strutil.Base10Conv(input, base)

		// Converte entre diferentes bases usando BaseConv
		for _, fromBase := range []int{2, 10, 16} {
			_ = strutil.BaseConv(input, fromBase, base)
		}
	}

	// Testa BaseConvByTpl diretamente com diferentes templates
	_ = strutil.BaseConvByTpl(input, strutil.Base10Chars, strutil.Base16Chars)
	_ = strutil.BaseConvByTpl(input, strutil.Base62Chars, strutil.Base36Chars)
	_ = strutil.BaseConvByTpl(input, strutil.Base64Chars, strutil.Base32Chars)

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
