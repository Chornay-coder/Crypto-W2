package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	"time"
)

// 🎨 ANSI escape codes for color and style
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Cyan   = "\033[36m"
	Bold   = "\033[1m"
)

// ⏳ Typing animation
func loadingEffect(msg string) {
	for _, c := range msg {
		fmt.Printf("%c", c)
		time.Sleep(25 * time.Millisecond)
	}
	fmt.Println()
}

// 🔐 XOR encryption/decryption function
// Supports single-byte key or repeating string key
func xorEncrypt(text string, key string) string {
	if len(key) == 0 {
		return text
	}
	result := make([]byte, len(text))
	keyBytes := []byte(key)
	for i, b := range []byte(text) {
		result[i] = b ^ keyBytes[i%len(keyBytes)]
	}
	return string(result)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// ✨ Intro banner
	fmt.Println(Cyan + Bold + "\n╔══════════════════════════════════════════════╗")
	fmt.Println("║        XOR ENCRYPTION / DECRYPTION 🔐        ║")
	fmt.Println("╚══════════════════════════════════════════════╝" + Reset)
	loadingEffect(Yellow + "Initializing XOR tool..." + Reset)
	time.Sleep(300 * time.Millisecond)

	// 📝 Input: plaintext
	fmt.Print(Yellow + "Enter plaintext message: " + Reset)
	scanner.Scan()
	plaintext := strings.TrimSpace(scanner.Text())

	// 📝 Input: key
	fmt.Print(Yellow + "Key" + Reset + " (single char or string): ")
	scanner.Scan()
	key := strings.TrimSpace(scanner.Text())

	if plaintext == "" || key == "" {
		fmt.Println(Red + "Error: Plaintext and key cannot be empty!" + Reset)
		return
	}

	fmt.Print(Bold + Cyan + "\n🔄 Encrypting..." + Reset)
	time.Sleep(250 * time.Millisecond)
	fmt.Println()

	// ⚙️ Encrypt
	cipherText := xorEncrypt(plaintext, key)

	// 🧩 Show ciphertext as HEX
	fmt.Println(Green + Bold + "Ciphertext (Hex):" + Reset)
	for _, b := range []byte(cipherText) {
		fmt.Printf("%02x ", b)
	}
	fmt.Println("\n")

	// 🧬 Show Base64 version
	base64Cipher := base64.StdEncoding.EncodeToString([]byte(cipherText))
	fmt.Println(Green + Bold + "Ciphertext (Base64):" + Reset)
	fmt.Println(base64Cipher)

	fmt.Print(Bold + Cyan + "\n🔄 Decrypting..." + Reset)
	time.Sleep(250 * time.Millisecond)
	fmt.Println()

	// 🔓 Decrypt (same function)
	decryptedText := xorEncrypt(cipherText, key)

	fmt.Println(Green + Bold + "Decrypted Text:" + Reset)
	fmt.Println(decryptedText)

	fmt.Println(Cyan + "\n──────────────────────────────" + Reset)
	fmt.Println(Green + Bold + "XOR operation complete! 🎉" + Reset)
}
