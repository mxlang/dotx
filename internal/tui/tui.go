package tui

import (
	"github.com/charmbracelet/huh"
)

func Confirm(title string, description string) (bool, error) {
	confirm := false
	form := huh.NewConfirm().
		Title(title).
		Description(description).
		Value(&confirm)

	if err := form.Run(); err != nil {
		return false, err
	}

	return confirm, nil
}
