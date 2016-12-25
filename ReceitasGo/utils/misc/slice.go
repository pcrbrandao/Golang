package misc

func StringSliceIsEqual(a []string, b []string) bool {

	if a == nil && b == nil {
		return true;
	}

	if a == nil || b == nil {
		return false;
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func StringSliceRemoveFrom(item string, slice []string) []string {
	novaSlice := []string{}

	for _, itemValor := range slice {
		if itemValor != item {
			novaSlice = append(novaSlice, itemValor)
		}
	}
	return novaSlice
}

// Retorna um novo slice sem o item informado
func CharSliceRemoveFrom(slice []Char, val Char) []Char {

	novaSlice := []Char{}

	for _, itemValor := range slice {
		if itemValor != val {
			novaSlice = append(novaSlice, itemValor)
		}
	}
	return novaSlice
}
