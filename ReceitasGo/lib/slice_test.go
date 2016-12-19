package lib

import (
	"testing"
	"fmt"
	"Golang/ReceitasGo/tester"
)

func TestRemoveDaSlice(t *testing.T) {
	
	entrada := []Char {NUMERO, MAIUSCULA, ESPECIAL, MENORQUESETE}
	saida := []Char {NUMERO, MAIUSCULA}

	fmt.Printf("entrada original...%s\n", entrada)
	entrada = CharSliceRemoveFrom(entrada, ESPECIAL)
	fmt.Printf("entrada atualizada... %s\n", entrada)
	entrada = CharSliceRemoveFrom(entrada, MENORQUESETE)
	fmt.Printf("entrada atualizada... %s\n", entrada)
	fmt.Printf("saída... %s\n", entrada)

	for item := range entrada {
		if entrada[item] != saida[item] {
			t.Errorf("entrada: %s, saída: %s", entrada[item], saida[item])
			return
		}
	}

	tester.PassouEntradaValida("entrada", t)
}

func TestSlicesDevemSerIguais(t *testing.T) {

	slice1 := []string {"um", "dois", "três"}
	slice2 := []string {"um", "dois", "três"}

	if StringSliceIsEqual(slice1, slice2) {
		t.Logf("Passou! %q é igual a %q", slice1, slice2)
		return
	}

	t.Errorf("Falhou! %q não é igual a %q")
}

func TestSlicesDevemSerDiferentes(t *testing.T) {
	slice1 := []string {"um", "dois", "três"}
	slice2 := []string {"um", "dois", "trê"}

	if StringSliceIsEqual(slice1, slice2) {
		t.Errorf("Falhou! %q é igual a %q", slice1, slice2)
		return
	}

	t.Logf("Passou! %q não é igual a %q", slice1, slice2)
}

func TestSlicesDevemSerIguaisAposRemover(t *testing.T) {
	slice1 := []string {"um", "dois", "três", "quatro"}
	slice2 := []string {"um", "dois", "três"}

	slice1 = StringSliceRemoveFrom("quatro", slice1)

	if StringSliceIsEqual(slice1, slice2) {
		t.Logf("Passou! %q é igual a %q", slice1, slice2)
		return
	}

	t.Errorf("Falhou! %q não é igual a %q")
}