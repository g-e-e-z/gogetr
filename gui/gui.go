package gui

import (
	"github.com/rivo/tview"
)

func RunGUI() {
	app := tview.NewApplication()
	flex := tview.NewFlex().
        SetDirection(tview.FlexRow).
		// AddItem(tview.NewBox().SetBorder(true).SetTitle("Env and Build Info"), 4, 1, false).
        AddItem(NewEnvSelector(), 4, 1, false).
        AddItem(NewRequestSelector(), 0, 4, true)
	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
