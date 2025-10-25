package main

import "fmt"

func myXor(a,b int){
	fmt.Println("\nXOR Operation:")
	fmt.Printf("%d ^ %d = %d\n", a, b, a^b)

    // XOR assignment: different → 1, same → 0;
    c := a
    c ^= b
    fmt.Printf("\nResult of XOR: %d ^ %d → %d\n", a, b, c)
	fmt.Printf("-------------------------------------------------------\n")
}

func myNot(a,b int){
	fmt.Println("\nNOT Operation:")

    // NOT assignment: flips all bits (0 → 1, 1 → 0)
    fmt.Printf("^%d = %d\n", a, ^a)
    fmt.Printf("^%d = %d\n", b, ^b)
	fmt.Printf("-------------------------------------------------------\n")
}

func myOr(a, b int){
	fmt.Println("\nOR Operation:")
	fmt.Printf("%d | %d = %d\n", a, b, a|b)

	// Or assignment: either bit 1 → 1, else 0
	c := a
	c |= b
	fmt.Printf("\nResult of OR: %d | %d → %d\n", a, b, c)
	fmt.Printf("-------------------------------------------------------\n")
}

func myAnd(a, b int){
	fmt.Println("\nAND Operation:")
	fmt.Printf("%d & %d = %d\n", a, b, a&b)

	// AND assignment: both bits 1 → 1, else 0
	c := a
	c &= b
	fmt.Printf("\nResult of AND: %d & %d → %d\n", a, b, c)
	fmt.Printf("-------------------------------------------------------\n")
} 

func myShiftleft(a, b int){
	fmt.Println("\nShift Left:")
	fmt.Printf("%d << %d = %d\n", a, b, a<<b)

	c := a
	c <<= b
	fmt.Printf("\nResult of shift left: %d << %d → %d\n", a, b, c)
	fmt.Printf("-------------------------------------------------------\n")
}

func myShiftright(a, b int){
	fmt.Println("\nShift right:")
	fmt.Printf("%d >> %d = %d\n", a, b, a>>b)

	c := a
	c >>= b
	fmt.Printf("\nResult of shift right: %d >> %d → %d\n", a, b, c)
	fmt.Printf("-------------------------------------------------------\n")
}

func main() {
    var a, b int

    fmt.Print("Enter first number (a): ")
    fmt.Scan(&a)

    fmt.Print("Enter second number (b): ")
    fmt.Scan(&b)

    myXor(a, b)
	myNot(a, b)
	myOr(a, b)
	myAnd(a, b)
	myShiftleft(a, b)
	myShiftright(a, b)

}