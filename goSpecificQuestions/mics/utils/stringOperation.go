package utils

import (
	"fmt"
	"reflect"
)

func stringOperation() {
	str := "aakash"

	fmt.Printf("str type = %T\n", str)
	// str[1] is byte  (unit8)
	fmt.Printf("literal type in when indexed eg str[1]=%T\n", str[1])
	for _, v := range str {
		// v is of rune type which is of 4 bytes
		fmt.Printf("literal type in when used range= %T\n", v)
		fmt.Println(reflect.TypeOf(v))
		break
	}

	lit := 'a' // is of rune type which is of 4 bytes

	fmt.Printf("type of lit = %T\n", lit)

	var bb byte
	bb = str[1] + 2 // operation on byte   types
	fmt.Printf("str[1] value = %v\n", str[1])
	fmt.Printf("bb val = %v and type = %T\n", bb, bb)

	if bb == 'c' {
		fmt.Println("successful rune and byte comparision")
	}

}
