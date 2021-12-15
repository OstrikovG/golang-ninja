package main

import (
	"fmt"
	"reflect"
)

var a, b, k = 1, 2, 3

func main() {
	//sdfdsfdsfdsfdsf
	/*	dsfdsf
		dsddsf
		sdfdsf*/
	var message string
	var number int
	var float_number float32
	var b bool
	var message_byte = []byte("1а3a")
	var a rune = 'a'
	e, c, d := 1, 2, 3

	fmt.Println(reflect.TypeOf(message))
	fmt.Println(message)
	fmt.Println(number)
	fmt.Println(float_number)
	fmt.Println(reflect.TypeOf(float_number))
	fmt.Println(b)
	fmt.Println(message_byte)
	fmt.Println(reflect.TypeOf(message_byte))
	fmt.Println(a)
	fmt.Println(e, c, d)
	d, c = c, d
	fmt.Println(e, c, d)
	d, _ = 10, 12
	fmt.Println(e, c, d)
	print()
	fmt.Println(k)

}

func print() {
	e, c := 1, 2

	fmt.Println(e, c)
	fmt.Println(k)
	k = 12
}
