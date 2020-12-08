package main

import "fmt"

// Retorna se o CPF é um CPF valido ou não.
func validaCpf(cpf string) bool {
	bVal := todosNumerosIguais(cpf)
	return bVal
}

// Certifica-se de que o CPF não é um CPF com todos os números iguais.
func todosNumerosIguais(cpf string) bool {
	for i := 0; i < len(cpf); i++ {
		fmt.Printf("%c\n", cpf[i])
	}
	return false
}

func main() {
	resultado := validaCpf("46494695892")
	print(resultado)
}
