package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func calculator() {
    fmt.Println("Welcome to the Julia Calculator")

    for {
        fmt.Println("Please enter the operation (+, -, *, /) or 'quit' to exit:")
        reader := bufio.NewReader(os.Stdin)
        op, _ := reader.ReadString('\n')
        op = op[:len(op)-1] // Remove newline character

        if op == "quit" {
            break
        } else if op == "+" || op == "-" || op == "*" || op == "/" {
            fmt.Println("Enter the first number:")
            num1Str, _ := reader.ReadString('\n')
            num1, _ := strconv.ParseFloat(num1Str[:len(num1Str)-1], 64) // Convert string to float

            fmt.Println("Enter the second number:")
            num2Str, _ := reader.ReadString('\n')
            num2, _ := strconv.ParseFloat(num2Str[:len(num2Str)-1], 64) // Convert string to float

            var result float64
            switch op {
            case "+":
                result = num1 + num2
            case "-":
                result = num1 - num2
            case "*":
                result = num1 * num2
            case "/":
                if num2 == 0 {
                    fmt.Println("Error: Division by zero is not allowed.")
                    continue
                } else {
                    result = num1 / num2
                }
            }

            fmt.Printf("The result is %v\n", result)
        } else {
            fmt.Println("Invalid operation. Please enter +, -, *, / or 'quit'.")
        }
    }
}

func main() {
    calculator()
}
