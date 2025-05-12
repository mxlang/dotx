package tui

import (
	"github.com/charmbracelet/huh"
)

func Confirm(title string, description string) bool {
	confirm := false
	form := huh.NewConfirm().
		Title(title).
		Description(description).
		Value(&confirm)

	err := form.Run()
	if err != nil {
		return false
	}

	return confirm
}
