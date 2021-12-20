package commands

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func inputEmailPassword() (string, string, error) {
	var email, password string

	fmt.Print("email or username: ")
	for {
		scan := bufio.NewScanner(os.Stdin)
		scan.Scan()
		email = scan.Text()

		if email != "" {
			fmt.Print("cli password: ")
			scan := bufio.NewScanner(os.Stdin)
			scan.Scan()
			password = scan.Text()

			return email, password, nil
		}
		return "", "", errors.New("Email or Username must not be blank")
	}
}
