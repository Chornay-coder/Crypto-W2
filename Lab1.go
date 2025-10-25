package main

import "fmt"

func main(){
	var a int
	var b int

	fmt.Print("Enter value a: ")
	fmt.Scan(&a)

	fmt.Print("Enter value b: ")
	fmt.Scan(&b)

	// 1. Operation =
    result1 := b
    fmt.Println("\nThe result for a = b →", result1)

    // 2. Operation +=
    result2 := a + b
    fmt.Println("The sum of a and b =", result2)

    // 3. Operation -=
    result3 := a - b
    fmt.Println("The sub of a and b =", result3)

    // 4. Operation *=
    result4 := a * b
    fmt.Println("The multi of a and b =", result4)

    // 5. Operation /= (avoid division by zero)
    if b != 0 {
        result5 := a / b
        fmt.Println("The division of a and b =", result5)
    } else {
        fmt.Println("Cannot divide by zero!")
    }

    // 6. Operation %= (avoid modulo by zero)
    if b != 0 {
        result6 := a % b
        fmt.Println("The modulus of a and b =", result6)
    } else {
        fmt.Println("Cannot perform modulo by zero!")
    }
}


