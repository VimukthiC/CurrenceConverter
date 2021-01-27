package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var dollarPrice int64 = 180

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	answer := input * dollarPrice
	fmt.Printf("dollars : %d", answer)
}
