package gui

import (
    "github.com/rivo/tview"
)

func NewHistoryModal() *tview.Modal {
    modal := tview.NewModal().
        SetText("Request History").
        AddButtons([]string{"OK"}).
        SetDoneFunc(func(buttonIndex int, buttonLabel string) {
            // Close the modal or show history details
        })

    return modal
}

