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

	return 1
}
