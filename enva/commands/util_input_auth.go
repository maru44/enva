package commands

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"syscall"

	"golang.org/x/term"
)

func inputEmailPassword() (string, string, error) {
	var email string

	fmt.Print("email or username: ")
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	email = scan.Text()

	if email == "" {
		return "", "", errors.New("Email or Username must not be blank")
	}

	fmt.Print("cli password: ")
	// should be casted to int in GOOS=windows
	password, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", "", err
	}
	fmt.Print("\n")

	return email, string(password), nil
}
