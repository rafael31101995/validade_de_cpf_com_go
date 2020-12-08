package main

import "fmt"

func validaCpf(cpf string) bool {
	bVal := false
	return bVal
}

func todosNumerosIguais(cpf string) bool {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	return true
}

func main() {
	resultado := validaCpf("46494695892")
	print(resultado)
}
