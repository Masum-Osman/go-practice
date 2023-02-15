package crypto

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

var secretkey = "4234kxzjcjj3nxnxbcvsjfj"

func generateSalt() string {
	randomBytes := make([]byte, 16)
	fmt.Println(randomBytes)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return ""
	}

	return base64.URLEncoding.EncodeToString(randomBytes)
}

func Hmac() {
	message := "Today web engineering has modern apps adhere to what is known as a single-page app (SPA) model."
	salt := generateSalt()

	fmt.Println("Message: " + message)
	fmt.Println("\nSalt: " + salt)

	hash := hmac.New(sha256.New, []byte(secretkey))
	io.WriteString(hash, message+salt)
	fmt.Printf("\nHMAC-Sha256: %x", hash.Sum(nil))

	hash = hmac.New(sha256.New, []byte(secretkey))
	io.WriteString(hash, message+salt)
	fmt.Printf("\n\nHMAC-sha512: %x", hash.Sum(nil))
}
