package gui

import (
	// "github.com/gdamore/tcell/v2"
	// "log"
	// "os"
	// "path/filepath"
	//
	// "github.com/g-e-e-z/gogetr/requests"
	"github.com/rivo/tview"
)

func NewRequestEditor() *tview.TextView {
    code := tview.NewTextView().
		SetWrap(false).
		SetDynamicColors(true)
	code.SetBorderPadding(1, 1, 2, 0)

	// basic := func() {
	// 	table.SetBorders(false).
	// 		SetSelectable(false, false).
	// 		SetSeparator(' ')
	// 	code.Clear()
	// 	fmt.Fprint(code, tableBasic)
	// }

	return code
}
