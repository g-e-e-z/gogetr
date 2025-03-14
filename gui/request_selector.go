package gui

import (
	// "github.com/gdamore/tcell/v2"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/g-e-e-z/gogetr/requests"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewRequestSelector() *tview.Flex {
	// TODO: This is bad
	pwd, _ := os.Getwd()
	requestsDir := filepath.Join(pwd, "requests_dir")
	groups, err := requests.LoadAllRequests(requestsDir)
	if err != nil {
		log.Panic(err)
	}

	// Create a new Flex layout to contain groups
	requestSelector := tview.NewList()
	requestSelector.SetBorder(true).SetTitle("Collection")

	// Create the viewer to display the request info
	requestViewer := tview.NewTextView().
		SetWrap(true).
		SetDynamicColors(true)
	requestViewer.SetBorderPadding(1, 1, 2, 0)
	requestViewer.SetBorder(true).SetTitle("Response")

	// The selected request's data
	var selectedRequest *requests.Request

	// TODO: Unify custom input logic
	requestSelector.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'k':
			curr_idx := requestSelector.GetCurrentItem()
			requestSelector.SetCurrentItem(curr_idx - 1)
			return nil
		case 'j':
			curr_idx := requestSelector.GetCurrentItem()
			if curr_idx + 1 == requestSelector.GetItemCount() {
				requestSelector.SetCurrentItem(0)
			} else {
				requestSelector.SetCurrentItem(curr_idx + 1)
			}
			return nil
		case rune(tcell.KeyCtrlR):
			// Execute the selected request
			if selectedRequest != nil {
				response, err := selectedRequest.Execute()
				if err != nil {
					log.Println("Error executing request:", err)
				} else {
					// TODO: Move this to a dedicated window, maybe?
					requestViewer.Clear()
					fmt.Fprint(requestViewer, selectedRequest.ParseResponse(response))
				}
			}
			return nil
		}
		return event
	})

	for groupName, groupRequests := range groups {
		for _, req := range groupRequests.Requests {
			// Associate each request with a closure that sets selectedRequest and displays it
			// NOTE: Want to take a moment, this next line has a function(arg) that returns function() - Nice
			// Should learn more about closures
			requestSelector.AddItem(req.Name, groupName, 0, func(req requests.Request) func() {
				return func() {
					selectedRequest = &req
					requestViewer.Clear()
					// Format and display the request details in the viewer
					fmt.Fprint(requestViewer, req.ViewerFormat()) // Format request as per ViewerFormat method
				}
			}(req))
		}
	}

	return tview.NewFlex().
		AddItem(tview.NewFlex().
			SetDirection(tview.FlexColumn).
			AddItem(requestSelector, 0, 1, true), 0, 1, true).
		AddItem(requestViewer, 0, 2, false)
}
