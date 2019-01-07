package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"github.com/jyotibhanot30/learning-go/armstrong"
)

func main() {
	file, err := os.Open("./input")
	if err != nil {
		fmt.Println("There's and error reading the file", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	m := scanner.Text()
	j, _ := strconv.Atoi(m)
	fmt.Println("Read number:", j)
	armstrong.IsArmstrong(j)
}
