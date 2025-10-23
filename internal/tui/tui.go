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

func MultiSelect[T comparable](title string, description string, values []T, convert func(t T) huh.Option[T]) ([]T, error) {
	var options []huh.Option[T]
	for _, value := range values {
		options = append(options, convert(value))
	}

	var selected []T
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[T]().
				Title(title).
				Description(description).
				Options(options...).
				Value(&selected),
		),
	)

	if err := form.Run(); err != nil {
		return selected, err
	}

	return selected, nil
}
