package array

import "testing"

// Dado que eu tenho um Array vazio, quando eu adiciono valores 1, 2 e 3,
// então o Array deve conter esses valores
func TestArray_DadoQueEuTenhoUmArrayVazio_QuandoEuAdicionoValores_DeveConterValores(t *testing.T) {
	arr := New()

	// Quando eu adiciono valores 1, 2 e 3
	arr.Add(1)
	arr.Add(2)
	arr.Add(3)

	// Então o Array deve conter os valores 1, 2 e 3
	expected := []int{1, 2, 3}
	result := arr.GetAll()

	if len(result) != len(expected) {
		t.Errorf("Esperado %v, mas obteve %v", expected, result)
	}

	if !arr.Find(2) {
		t.Errorf("Esperado encontrar o valor 2")
	}

	if arr.Find(4) {
		t.Errorf("Não esperava encontrar o valor 4")
	}
}

// Dado que eu tenho um Array, quando ele está vazio,
// então deve estar vazio
func TestArray_QuandoVazio_DeveEstarVazio(t *testing.T) {
	arr := New()

	// Então deve estar vazio
	if len(arr.GetAll()) != 0 {
		t.Errorf("Esperado Array vazio")
	}
}
