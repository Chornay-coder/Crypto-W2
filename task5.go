package main

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
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

func loadingEffect(msg string) {
	for _, c := range msg {
		fmt.Printf("%c", c)
		time.Sleep(25 * time.Millisecond)
	}
	fmt.Println()
}

func stringToBinary(s string) string {
	var binaryStrings []string
	for _, b := range []byte(s) {
		binaryStrings = append(binaryStrings, fmt.Sprintf("%08b", b))
	}
	return strings.Join(binaryStrings, " ")
}

func main() {
	var input string

	// Cool intro
	fmt.Println(Cyan + Bold + "\n╔══════════════════════════════════════════════╗")
	fmt.Println("║        BINARY, HEX, & BASE64 ENCODER 🔐      ║")
	fmt.Println("╚══════════════════════════════════════════════╝" + Reset)
	loadingEffect(Yellow + "Initializing encoder..." + Reset)
	time.Sleep(300 * time.Millisecond)

	fmt.Print(Yellow + "Enter a string to encode: " + Reset)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()

	fmt.Print(Bold + Cyan + "\n🔄 Encoding..." + Reset)
	time.Sleep(250 * time.Millisecond)
	fmt.Println()

	// Binary representation
	binaryRep := stringToBinary(input)
	fmt.Println(Green + Bold + "Binary Representation:" + Reset)
	fmt.Println(binaryRep)

	// Hexadecimal representation
	hexRep := hex.EncodeToString([]byte(input))
	fmt.Println(Green + Bold + "\nHexadecimal Representation:" + Reset)
	fmt.Println(hexRep)

	// Base64 representation
	base64Rep := base64.StdEncoding.EncodeToString([]byte(input))
	fmt.Println(Green + Bold + "\nBase64 Representation:" + Reset)
	fmt.Println(base64Rep)

	fmt.Println(Cyan + "\n──────────────────────────────" + Reset)
	fmt.Println(Green + Bold + "Encoding complete! 🎉" + Reset)
}
