package utils

import "golang.org/x/crypto/bcrypt"

func HashPasscode(passcode string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(passcode), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedBytes), nil
}

func ComparePasscode(hashedPasscode, passcode string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPasscode), []byte(passcode))
}
