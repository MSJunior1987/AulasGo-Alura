package contas

import "banco/clientes"

/*equivalente a classe*/
type ContaPoupanca struct {
	Titular                              clientes.Titular //instancia da classe clientes - o go nao possui heranca
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "O valor do deposito é menor que zero", c.saldo
	}

}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque //retorna o valor do saldo menos o valor do saque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
