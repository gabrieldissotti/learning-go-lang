/*
- Utilizando o operador curto de declaração, atribua estes valores às variáveis com os identificadores "x", "y", e "z".
    1. 42
    2. "James Bond"
    3. true
- Agora demonstre os valores nestas variáveis, com:
    1. Uma única declaração print
    2. Múltiplas declarações print
- Solução: https://play.golang.org/p/yYXnWXIQNa
*/


package main 

import "fmt"

func main () {

	x := 41
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}