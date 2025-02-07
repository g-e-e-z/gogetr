package gui

import (
    "github.com/rivo/tview"
)

func NewGroupViewer() *tview.Flex {
    groupViewer := tview.NewFlex().
        SetDirection(tview.FlexRow)

    // Example of adding groups
    groupViewer.AddItem(tview.NewButton("Group 1").SetSelectedFunc(func() {
        // Load and display requests for Group 1
    }), 0, 1, false)

    groupViewer.AddItem(tview.NewButton("Group 2").SetSelectedFunc(func() {
        // Load and display requests for Group 2
    }), 0, 1, false)

    return groupViewer
}

