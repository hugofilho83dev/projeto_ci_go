package main

import "testing"

func TestSoma(t *testing.T) {
	total := soma(15, 25)

	if total != 30 {
		t.Errorf("Resultado da soma e inválido: Resultado %d. Esperado: %d", total, 30)
	}

}
