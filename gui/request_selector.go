package gui

import (
	// "github.com/gdamore/tcell/v2"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/g-e-e-z/gogetr/requests"
	"github.com/rivo/tview"
)

func NewRequestSelector() *tview.Flex {
	// list := tview.NewList().
	// 	AddItem("List item 1", "Some explanatory text", 'a', nil).
	// 	AddItem("List item 2", "Some explanatory text", 'b', nil).
	// 	AddItem("List item 3", "Some explanatory text", 'c', nil).
	// 	AddItem("List item 4", "Some explanatory text", 'd', nil)
		// AddItem("Quit", "Press to exit", 'q', func() {
		// 	app.Stop()
		// })
	// groups := map[string][]string{
	// 	"Group 1": {"Request 1", "Request 2", "Request 3"},
	// 	"Group 2": {"Request 4", "Request 5"},
	// 	"Group 3": {"Request 6", "Request 7", "Request 8", "Request 9"},
	// }

    pwd, _ := os.Getwd()
	requestsDir := filepath.Join(pwd, "requests_dir")
    groups, err := requests.LoadAllRequests(requestsDir)
    if err != nil {
        log.Panic(err)
    }
	// Create a new Flex layout to contain groups
	requestSelector := tview.NewList()
    requestSelector.SetBorder(true).SetTitle("Requests")

    requestViewer := tview.NewTextView().
		SetWrap(false).
		SetDynamicColors(true)
	requestViewer.SetBorderPadding(1, 1, 2, 0)
    requestViewer.SetBorder(true).SetTitle("Request Viewer")

	// Loop over each group and add requests to a collapsible section
	for groupName, requests := range groups {
		// group := tview.NewFlex().SetDirection(tview.FlexRow)

		// Create a simple header with the group name
		// groupHeader := tview.NewTextView().SetText(groupName)

		// Create a list to hold requests in this group
		// list := tview.NewList()
		for _, request := range requests.Requests {
			// Add request to list (no subtext or callback for now)
			requestSelector.AddItem(request.Name, groupName, 0, func() {
                requestViewer.Clear()
                fmt.Fprint(requestViewer, request.ViewerFormat())
            })


		// Add the header and list to the group
		// group.AddItem(groupHeader, 0, 1, false) // Add group header
		// group.AddItem(list, 0, 5, false)       // Add request list to group

		// Add the group to the main request view
		// requestView.AddItem(group, 0, 1, false)
	}}

    return tview.NewFlex().
		AddItem(tview.NewFlex().
			SetDirection(tview.FlexColumn).
			AddItem(requestSelector, 0, 1, true), 0, 1, true).
		AddItem(requestViewer, 0, 1, false)
}
