package main

// Retorna se o CPF é um CPF valido ou não.
func validaCpf(cpf string) bool {
	bVal := todosNumerosIguais(cpf)
	return bVal
}

// Certifica-se de que o CPF não é um CPF com todos os números iguais.
func todosNumerosIguais(cpf string) bool {
	for i := 0; i < len(cpf); i++ {
		if cpf[i] != cpf[0] {
			return false
		}
	}
	return true
}

// Remove caracteres especiais do CPF.
func removedorDeCaracterEspecial(cpf string) string {
	//for i := 0; i < len(cpf); i++ {
	//	if strconv.ParseInt(cpf[i])
	//}
	return "46494698888"
}

func main() {
	resultado := validaCpf("1111111111111")
	print(resultado)
}
