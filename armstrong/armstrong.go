package armstrong
import (
	"fmt"
//	"os"
//	"bufio"
//	"strconv"
)
func main() {
//	scanner := bufio.NewScanner(os.Stdin)
//	fmt.Print("Enter a number:")
//	scanner.Scan()
//	n := scanner.Text()
	for i:=0; i<200; i++ {
//	m, _ := strconv.Atoi(n)
		IsArmstrong(i)
	}
}

func IsArmstrong(a int) {
	j := a
	sum := 0
	for j != 0 {
		n1 := j % 10
		sum += n1 * n1 * n1
		j = j / 10
	}
	if sum == a {
		fmt.Printf("%d is Armstrong! \n",a)
	} else {
		fmt.Printf("%d is Not Armstrong! \n",a)
	}
}
