package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter num: ")
	scanner.Scan()
	text := scanner.Text()
	num, _ := strconv.Atoi(text)
	fmt.Printf("%d - Facorial %d\n", num, fact(num))
}

func fact(n int) int {
	if n == 0 { return 1 }
	return n * fact(n - 1)
}
