package tui

import (
	"github.com/charmbracelet/huh"
)

func Confirm(title string, description string) (bool, error) {
	confirm := false
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title(title).
				Description(description).
				Value(&confirm),
		),
	)

	if err := form.Run(); err != nil {
		return false, err
	}

	return confirm, nil
}

func Text(title string) (string, error) {
	var text string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewText().
				Title(title).
				Value(&text),
		),
	)

	if err := form.Run(); err != nil {
		return "", err
	}

	return text, nil
}
