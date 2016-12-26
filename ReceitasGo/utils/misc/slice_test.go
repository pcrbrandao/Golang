package misc_test

import (
	"testing"
	"fmt"
	"Golang/ReceitasGo/utils/misc"
	"Golang/ReceitasGo/tests"
)

func TestRemoveDaSlice(t *testing.T) {
	
	entrada := []misc.Char {misc.NUMERO, misc.MAIUSCULA, misc.ESPECIAL, misc.MENORQUESETE}
	saida := []misc.Char {misc.NUMERO, misc.MAIUSCULA}

	fmt.Printf("entrada original...%s\n", entrada)
	entrada = misc.CharSliceRemoveFrom(entrada, misc.ESPECIAL)
	fmt.Printf("entrada atualizada... %s\n", entrada)
	entrada = misc.CharSliceRemoveFrom(entrada, misc.MENORQUESETE)
	fmt.Printf("entrada atualizada... %s\n", entrada)
	fmt.Printf("saída... %s\n", entrada)

	for item := range entrada {
		if entrada[item] != saida[item] {
			t.Errorf("entrada: %s, saída: %s", entrada[item], saida[item])
			return
		}
	}

	tests.PassouEntradaValida("entrada", t)
}

func TestSlicesDevemSerIguais(t *testing.T) {

	slice1 := []string {"um", "dois", "três"}
	slice2 := []string {"um", "dois", "três"}

	if misc.StringSliceIsEqual(slice1, slice2) {
		t.Logf("Passou! %q é igual a %q", slice1, slice2)
		return
	}

	t.Errorf("Falhou! %q não é igual a %q")
}

func TestSlicesDevemSerDiferentes(t *testing.T) {
	slice1 := []string {"um", "dois", "três"}
	slice2 := []string {"um", "dois", "trê"}

	if misc.StringSliceIsEqual(slice1, slice2) {
		t.Errorf("Falhou! %q é igual a %q", slice1, slice2)
		return
	}

	t.Logf("Passou! %q não é igual a %q", slice1, slice2)
}

func TestSlicesDevemSerIguaisAposRemover(t *testing.T) {
	slice1 := []string {"um", "dois", "três", "quatro"}
	slice2 := []string {"um", "dois", "três"}

	slice1 = misc.StringSliceRemoveFrom("quatro", slice1)

	if misc.StringSliceIsEqual(slice1, slice2) {
		t.Logf("Passou! %q é igual a %q", slice1, slice2)
		return
	}

	t.Errorf("Falhou! %q não é igual a %q")
}