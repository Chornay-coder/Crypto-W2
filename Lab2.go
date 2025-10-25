package main

import "fmt"

func main(){
	var a int
	var b int

	fmt.Print("Enter the value a: ")
	fmt.Scan(&a)

	fmt.Print("Enter the value b: ")
	fmt.Scan(&b)

    // 1. Both positive (&&)
    bothPositive := (a > 0) && (b > 0)
    fmt.Println("\nBoth numbers are positive (a > 0 && b > 0):", bothPositive)

    // 2. One greater than the other (||)
    oneGreater := (a > b) || (b > a)
    fmt.Println("One of those number is greater than the other (a > b || b > a):", oneGreater)

    // 3. Not equal (!)
    notEqual := (a != b)
    fmt.Println("Numbers are not equal (a != b):", notEqual)
}