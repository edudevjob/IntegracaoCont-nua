package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 15)
	if total != 30 {
		t.Errorf("Resultado da Soma inválido: ZZZ Resultado %d. Esperado: %d", total, 30)
	}
}
