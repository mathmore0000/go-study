package calcCli

import "fmt"

var validOperations = []string{"+", "-", "/", "*"}

func main() {
	var operation int
	var n1, n2 float32
	for operation != 1 {
		fmt.Print("\033[H\033[2J")
		fmt.Println("Choose one operation:")
		fmt.Println("1) +\n2) -\n3) /\n4) *")
		fmt.Scan(&operation)
		if !isValidOperation(operation) {
			fmt.Printf("\nInvalid operation, please choose a valid one!!\n\n")
			continue
		}

		fmt.Println("Type in the first number:")
		fmt.Scan(&n1)

		fmt.Println("Type in the second number:")
		fmt.Scan(&n2)

		fmt.Printf("Result: %v", calc(operation-1, n1, n2))

		fmt.Println("\n\nDo you wish to leave?")
		fmt.Println("1) yes\n2) no")
		fmt.Scan(&operation)
	}
}

func isValidOperation(chosenOperation int) (isValid bool) {
	return chosenOperation > 0 && chosenOperation < 5

}

func calc(operation int, n1 float32, n2 float32) (result float32) {
	switch validOperations[operation] {
	case "+":
		return n1 + n2
	case "-":
		return n1 - n2
	case "/":
		return n1 / n2
	case "*":
		return n1 * n2
	}
	return 0
}
