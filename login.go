package main

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
