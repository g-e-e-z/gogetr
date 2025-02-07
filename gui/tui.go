package gui

import (
	"github.com/rivo/tview"
	// "github.com/g-e-e-z/gogetr/requests"
)

func RunTUI() {
	app := tview.NewApplication()
	flex := tview.NewFlex().
        SetDirection(tview.FlexColumn).
        AddItem(NewRequestSelector(), 0, 1, true).
		// AddItem(list, 0, 1, true).
		// AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
		// 	AddItem(tview.NewBox().SetBorder(true).SetTitle("Top"), 0, 1, false).
		// 	AddItem(tview.NewBox().SetBorder(true).SetTitle("Middle (3 x height of Top)"), 0, 3, false).
		// 	AddItem(tview.NewBox().SetBorder(true).SetTitle("Bottom (5 rows)"), 5, 1, false), 0, 2, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Right (20 cols)"), 20, 1, false)
	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
