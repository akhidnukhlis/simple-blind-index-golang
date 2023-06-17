package main

import (
	"fmt"
)

type User struct {
	Username string
	Password string
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
