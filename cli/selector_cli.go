package cli

import (
	"github.com/manifoldco/promptui"
	"log"
)

// SelectorCli helps to get user input showing a selector with the passed options and label.
func SelectorCli(label string, options ...string) (string, error) {
	s := promptui.Select{
		Label: label,
		Items: options,
	}

	_, result, err := s.Run()
	if err != nil {
		log.Fatal(err)
	}

	return result, nil
}
