package utility

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (encrypt []byte, err error) {
	encrypt, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return
}
