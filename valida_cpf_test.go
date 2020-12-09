package main

import "testing"

func TestValidaCpf(t *testing.T) {
	resultado := validaCpf("46494695888")
	esperado := false
	if resultado != esperado {
		t.Errorf("resultado '%t', esperando '%t'", resultado, esperado)
	}
}

func TestTodosNumerosIguais(t *testing.T) {
	resultado := todosNumerosIguais("46494695888")
	esperado := false
	if resultado != esperado {
		t.Errorf("resultado '%t', esperando '%t'", resultado, esperado)
	}
}

func TestRemovedorDeCaracterEspecial(t *testing.T) {
	resultado := removedorDeCaracterEspecial("464-946-988.88")
	esperado := "46494698888"
	if resultado != esperado {
		t.Errorf("resultado '%s', esperando '%s'", resultado, esperado)
	}
}
