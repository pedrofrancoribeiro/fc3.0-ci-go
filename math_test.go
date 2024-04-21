package main

import "testing"

func TestSoma(t *testing.T) {
	total := soma(20, 20)

	if total != 30 {
		t.Errorf("Resultado de soma é inválido: Resultado %d. Esperado: %d", total, 40)
	}
}