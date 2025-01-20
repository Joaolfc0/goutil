//go:build gofuzz
// +build gofuzz

package fuzz

import (
	"fmt"

	"github.com/Joaolfc0/goutil/byteutil"
)

// FuzzByteUtil testa as funções no pacote byteutil.
func FuzzByteUtil(data []byte) int {
	if len(data) == 0 {
		return 0
	}

	// Cria uma instância do Buffer
	buffer := byteutil.NewBuffer()

	// Testa WriteByte e PrintByte
	if len(data) > 0 {
		buffer.PrintByte(data[0])
		if len(buffer.Bytes()) > 0 && buffer.Bytes()[0] != data[0] {
			panic(fmt.Sprintf("PrintByte falhou: esperado %q, obtido %q", string(data[0]), buffer.Bytes()))
		}
		buffer.Reset()
	}

	// Testa WriteStr1 e WriteStr1Nl
	input := string(data)
	buffer.WriteStr1(input)
	if buffer.String() != input {
		panic(fmt.Sprintf("WriteStr1 falhou: esperado %q, obtido %q", input, buffer.String()))
	}
	buffer.Reset()

	buffer.WriteStr1Nl(input)
	expected := input + "\n"
	if buffer.String() != expected {
		panic(fmt.Sprintf("WriteStr1Nl falhou: esperado %q, obtido %q", expected, buffer.String()))
	}
	buffer.Reset()

	// Testa WriteAny e WriteAnyNl
	buffer.WriteAny(input)
	if buffer.String() != input {
		panic(fmt.Sprintf("WriteAny falhou: esperado %q, obtido %q", input, buffer.String()))
	}
	buffer.Reset()

	buffer.WriteAnyNl(input)
	expected = input + "\n"
	if buffer.String() != expected {
		panic(fmt.Sprintf("WriteAnyNl falhou: esperado %q, obtido %q", expected, buffer.String()))
	}
	buffer.Reset()

	// Testa ResetAndGet
	buffer.WriteStr1(input)
	retrieved := buffer.ResetAndGet()
	if retrieved != input {
		panic(fmt.Sprintf("ResetAndGet falhou: esperado %q, obtido %q", input, retrieved))
	}
	if buffer.Len() != 0 {
		panic("ResetAndGet falhou: buffer não foi limpo")
	}

	// Testa Close, Flush e Sync
	err := buffer.Close()
	if err != nil {
		panic(fmt.Sprintf("Close falhou: %v", err))
	}

	err = buffer.Flush()
	if err != nil {
		panic(fmt.Sprintf("Flush falhou: %v", err))
	}

	err = buffer.Sync()
	if err != nil {
		panic(fmt.Sprintf("Sync falhou: %v", err))
	}

	return 1
}
