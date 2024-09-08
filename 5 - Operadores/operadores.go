package main

import "fmt"

func main() {
	// Aritméticos
	// + - * / % ++ --
	// Relacionais
	// == != > < >= <=
	// Lógicos
	// && || !
	// Bitwise
	// & | ^ << >>
	// Atribuição
	// = += -= *= /= %= <<= >>= &= |= ^=
	// Outros
	// & * <- <-chan <-chan<- chan chan<- <-

	// Operadores aritméticos
	// Soma
	var a = 1 + 2
	println(a) // 3

	// Subtração
	var b = 2 - 1
	println(b) // 1

	// Multiplicação
	var c = 2 * 3
	println(c) // 6

	// Divisão
	var d = 6 / 2
	println(d) // 3

	// Modulo
	var e = 7 % 3
	println(e) // 1

	// Incremento
	var f = 1
	f++
	println(f) // 2

	// Decremento
	var g = 2
	g--
	println(g) // 1

	// Operadores relacionais
	// Igual
	var h = 1 == 1
	println(h) // true

	// Diferente
	var i = 1 != 1
	println(i) // false

	// Maior
	var j = 2 > 1
	println(j) // true

	// Menor
	var k = 1 < 2
	println(k) // true

	// Maior ou igual
	var l = 2 >= 1
	println(l) // true

	// Menor ou igual
	var m = 1 <= 2
	println(m) // true

	// Operadores lógicos
	// E (simplificado)
	var n = true
	println(n) // true

	// OU
	var o = true || false
	println(o) // true

	// Negação
	var p = !true
	println(p) // false

	// Operadores bitwise
	// E
	var q = 1 & 1
	println(q) // 1

	// OU
	var r = 1 | 1
	println(r) // 1

	// XOR (OU exclusivo)
	// O operador ^ em Go realiza uma operação XOR bit a bit (ou "exclusive OR") entre dois números inteiros.
	// Explicação de XOR:
	// A operação XOR compara os bits correspondentes de dois números:
	// Se os bits forem diferentes, o resultado é 1.
	// Se os bits forem iguais, o resultado é 0.
	var s = 1 ^ 1
	println(s) // 0

	var variavel1 string = "String1" // Declarando variável
	variavel2 := "String2"           // Declarando variável com inferência de tipo
	fmt.Println(variavel1, variavel2)

	// operador ternario em go não existe mas podemos fazer algo parecido
	var x = 10
	var y = 20
	var resultado int
	if x > y {
		resultado = x
	} else {
		resultado = y
	}
	fmt.Println(resultado)

}
