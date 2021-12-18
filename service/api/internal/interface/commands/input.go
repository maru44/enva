package commands

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/maru44/enva/service/api/pkg/domain"
)

func inputEmailPassword() (*domain.CliUserValidateInput, error) {
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

			return &domain.CliUserValidateInput{
				EmailOrUsername: email,
				Password:        password,
			}, nil
		}
		return nil, errors.New("Email or Username must not be blank")
	}
}
