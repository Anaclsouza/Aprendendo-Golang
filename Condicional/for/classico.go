package FOR

import "fmt"

//O For é a única estrutura de loop em GO e por isso ele é versátil e poderoso.

//Laço classico : é ideal quanfo você ja sabe o número de vezes que o LOOP deve ser executado
//ou quando deseja iterar com base em um contador

//Quando usar:

//*Iterações Baseadas em Contador*
//Quando você precisa iterar um número fixo de vezes ou deseja aumentar/diminuir uma variável de controle a cada iteração

func IteracoesContador() {

	//Enquanto 0 for menor que 10 faça:

	for i := 0; i < 10; i++ {
		fmt.Println("Enquanto 0 for menor que 10 faça",i)
	}

	
	//Imprima os multiplos de 5 de 1 a 50

	for i:=1; i<50;i++{
		if i%5 == 0{
			fmt.Println("Imprima os multiplos de 5 de 1 a  50", i)
		}
	}

	
	//Imprime a Tabuada do número 7

	numero :=7
	for i:=1; i<10;i++{
      
		resultado := numero * i
		fmt.Println("Imprime a Tabuada do número 7",resultado)

	}



}

//*Percorrendo estruturas com índice*
//Em arrays e slices, muitas vezes você quer usar os índices para acessar os elementos diretamente

func PercorreArrays() {

	arrayDeNumeros :=[]int{2,4,6,8,10}

	for i:=0;i<len(arrayDeNumeros);i++{

		arrayDeNumeros[i] *= 2

		fmt.Println("Percorre por indice e vai duplicando um a um",arrayDeNumeros)

	}
}

//*Controle de Fluxo por condições*
//Quando você quer iterar sobre uma sequência de valores mas também ter controle para adicionar condições extras na inicialização ou pós execução, como pular ou modificar a contagem entre iterações 

func ControleDeFluxo() {

	for i:=1;i<50;i++{

		if i%2 ==0{
			fmt.Println("imprime os números de 1 a 50 que são pares",i)
		}
	}
}

//*Iterações reversas*
//Quando você precisa percorrer uma lista de trás pra frente, por exemplo

func IteracoesReversas(){

	for i:=10;i>=1; i--{
		fmt.Println(i)
	}
}


