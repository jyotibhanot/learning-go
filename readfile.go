package main
import (
	"fmt"
	"bufio"
	"os"
)
func main() {
	file, err := os.Open("./read")
	if err!= nil {
		os.Exit(1)
	}
	scanner:= bufio.NewScanner(file)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Printf("Hello %s \n",m)
	}
}
