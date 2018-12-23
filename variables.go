// Variables in Go.
package main

import "fmt"
import "reflect"

func main() {
	var a int
	fmt.Println(a)
	
	var b int = 1
	fmt.Println(b)
	
	var c, d int = 1, 2
	fmt.Println(c, d)

	f := 1
	fmt.Println(f)
	f = 2
	fmt.Println(reflect.TypeOf(f))

	fmt.Printf("Printing - %d\n", 23)
}
