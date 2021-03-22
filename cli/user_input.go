package cli

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
)

// An example of validate function:
//
// validate := func(input string) error {
//	 _, err := strconv.ParseFloat(input, 64)
//	 if err != nil {
// 		 return errors.New("Invalid number")
//	 }
//	 return nil
// }

// UserInput allow to get the user input with optional validate function.
func UserInput(label string, validate ...promptui.ValidateFunc) (string, error) {

	var validation promptui.ValidateFunc

	if len(validate) == 0 {
		validation = promptui.ValidateFunc(func(s string) error {
			return nil
		})
	} else if len(validate) == 1 {
		validation = validate[0]
	} else if len(validate) > 1 {
		return "", errors.New("it's permited only one validation function parameter.")
	}

	prompt := promptui.Prompt{
		Label:    label,
		Validate: validation,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}

	return result, nil
}
