package main

import "fmt"

func main () {
	var x float32
	
	fmt.Println(" masukan angka untuk var x:")
	fmt.Scan(&x)

	var fx = ((x * x * x) + (3 * x)) / ((x * x * x * x) - (3 * x * x) + 4)
	fmt.Println(fx) // hasil sudah dibulatkan
}
