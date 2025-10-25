package main

import (
    "fmt"
    "time"
)

const (
    Reset  = "\033[0m"
    Red    = "\033[31m"
    Green  = "\033[32m"
    Yellow = "\033[33m"
    Cyan   = "\033[36m"
    Bold   = "\033[1m"
)

func add(a, b int) int { return a + b }
func sub(a, b int) int { return a - b }
func mul(a, b int) int { return a * b }

func div(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("\n⚠️  You cannot divide a number by 0. Please enter a different value.")
    }
    return a / b, nil
}

func mod(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("\n⚠️  You cannot use 0 as the divisor for modulo. Please enter another number.")
    }
    return a % b, nil
}

func loadingEffect(msg string) {
    for _, c := range msg {
        fmt.Printf("%c", c)
        time.Sleep(25 * time.Millisecond)
    }
    fmt.Println()
}

func main() {
    var choice, a, b int

    // Cool intro
    fmt.Println(Cyan + Bold + "\n╔══════════════════════════════════╗")
    fmt.Println("║         MINI CALCULATOR 🧮        ║")
    fmt.Println("╚══════════════════════════════════╝" + Reset)
    loadingEffect(Yellow + "Initializing..." + Reset)
    time.Sleep(300 * time.Millisecond)

    for {
        fmt.Println(Yellow + Bold + "\n🔹 MENU OPTIONS 🔹\n" + Reset)
        fmt.Println("1  Add")
        fmt.Println("2  Subtract")
        fmt.Println("3  Multiply")
        fmt.Println("4  Divide")
        fmt.Println("5  Modulo")
        fmt.Println("6  Exit")
        fmt.Print(Cyan + "\n👉 Choose an option: " + Reset)
        fmt.Scan(&choice)

        if choice == 6 {
            fmt.Println(Green + Bold + "\nThanks for using the Mini Calculator! 👋" + Reset)
            break
        }

        if choice < 1 || choice > 6 {
            fmt.Println(Red + Bold + "\n Invalid option! Try again. " + Reset)
            continue
        }

        fmt.Print(Yellow + "Enter value a: " + Reset)
        fmt.Scan(&a)
        fmt.Print(Yellow + "Enter value b: " + Reset)
        fmt.Scan(&b)

        fmt.Print(Bold + Cyan + "\n🧮 Calculating..." + Reset)
        time.Sleep(250 * time.Millisecond)
        fmt.Println()

        switch choice {
        case 1:
            fmt.Println(Green + Bold + "\n Result:", add(a, b), Reset)
        case 2:
            fmt.Println(Green + Bold + "\n Result:", sub(a, b), Reset)
        case 3:
            fmt.Println(Green + Bold + "\n Result:", mul(a, b), Reset)
        case 4:
            result, err := div(a, b)
            if err != nil {
                fmt.Println(Red + Bold + err.Error() + Reset)
            } else {
                fmt.Println(Green + Bold + "\n Result:", result, Reset)
            }
        case 5:
            result, err := mod(a, b)
            if err != nil {
                fmt.Println(Red + Bold + err.Error() + Reset)
            } else {
                fmt.Println(Green + Bold + "\n Result:", result, Reset)
            }
        }

        fmt.Println(Cyan + "──────────────────────────────" + Reset)
        time.Sleep(300 * time.Millisecond)
    }
}
