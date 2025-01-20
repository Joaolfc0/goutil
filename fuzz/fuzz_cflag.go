//go:build gofuzz
// +build gofuzz

package fuzz

import (
	"bytes"
	"fmt"

	"github.com/gookit/goutil/cflag"
)

// FuzzCFlag testa a estrutura de App e comandos.
func FuzzCFlag(data []byte) int {
	args := bytes.Fields(data)
	if len(args) == 0 {
		return 0
	}

	// Cria um aplicativo CLI com comandos de teste
	app := cflag.NewApp(func(a *cflag.App) {
		a.Name = "TestApp"
		a.Desc = "Um aplicativo CLI de exemplo"
		a.Version = "1.0.0"
	})

	// Adiciona comandos fictícios
	app.Add(cflag.NewCmd("test", "Um comando de teste").Config(func(c *cflag.Cmd) {
		c.Func = func(cmd *cflag.Cmd) error {
			return nil // Simula sucesso
		}
	}))

	app.Add(cflag.NewCmd("error", "Um comando que retorna erro").Config(func(c *cflag.Cmd) {
		c.Func = func(cmd *cflag.Cmd) error {
			return fmt.Errorf("erro simulado")
		}
	}))

	// Simula execução com argumentos fornecidos
	err := app.RunWithArgs(toStrings(args))
	if err != nil && err.Error() == "input not exists command" {
		return 0
	}

	return 1
}

// toStrings converte bytes em strings para usar como argumentos.
func toStrings(data [][]byte) []string {
	args := make([]string, len(data))
	for i, d := range data {
		args[i] = string(d)
	}
	return args
}
