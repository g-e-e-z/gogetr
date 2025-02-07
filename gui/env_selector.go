package gui

import (
    "github.com/rivo/tview"
    // "github.com/g-e-e-z/gogetr/config"
)

func NewEnvSelector() *tview.InputField {
    input := tview.NewInputField().
        SetLabel("Select Environment: ").
        SetPlaceholder("dev.env, prod.env").
        SetAcceptanceFunc(func(text string, key rune) bool {
        // SetAcceptanceFunc(func(text string, key tview.Key) bool {
            // Logic to load the selected environment file
            // config.LoadEnvFile(text) // hypothetical function
            return true
        })

    return input
}

