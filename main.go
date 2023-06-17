package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type User struct {
	Username string
	Password string
}

// hashingSHA256 fungsi untuk melakukan hash SHA256
func hashingSHA256(data string) string {
	byteData := []byte(data)
	sha256Hash := sha256.New()
	sha256Hash.Write(byteData)
	hashResult := sha256Hash.Sum(nil)
	hashString := hex.EncodeToString(hashResult)
	return hashString
}

// login fungsi untuk mencocokkan hasil hash dari data yang terdaftar
func login(username, password string, users []User) bool {
	for _, user := range users {
		if user.Username == username {
			// Menghitung nilai hash dari password yang dimasukkan
			hashedPassword := hashingSHA256(password)

			// Membandingkan nilai hash password dengan nilai hash yang disimpan
			if user.Password == hashedPassword {
				return true // Login berhasil
			}
		}
	}
	return false // Login gagal
}

func main() {
	// Contoh data pengguna yang terdaftar
	users := []User{
		{Username: "my-username", Password: hashingSHA256("my-password")},
		{Username: "your-username", Password: hashingSHA256("your-password")},
	}

	// Contoh parameter login
	username := "my-username"
	password := "my-password"

	if login(username, password, users) {
		fmt.Println("Login berhasil.")
	} else {
		fmt.Println("Login gagal.")
	}
}
