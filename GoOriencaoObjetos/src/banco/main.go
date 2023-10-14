package main

import (
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

/*seria o contrato entre os metodos de saque entre CC e CP*/
type verificarConta interface {
	Sacar(valor float64) string
}

/*Lembrando que propriedades publicas comecam com maiuscula e privadas com minuscula*/
/*A linguagem Go não possui o conceito de classe ao invez disso o Go utiliza estruturas de dados em forma de structs.*/
/*A linguagem go utiliza ponteiros*/
/*O go possui interfaces - Uma interface é a definição de um conjunto de métodos comuns a um ou mais tipos. É o que permite a criação de tipos polimórficos em Go*/
/*Java possui um conceito muito parecido, também chamado de interface. A grande diferença é que, em Go, um tipo implementa uma interface implicitamente.*/
func main() {

	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	PagarBoleto(&contaDoDenis, 60)

	fmt.Println(contaDoDenis.ObterSaldo())

	contaDoLuisa := contas.ContaCorrente{}
	contaDoLuisa.Depositar(500)
	PagarBoleto(&contaDoLuisa, 500)

	fmt.Println(contaDoLuisa.ObterSaldo())
}
