package chest_ui

import (
	//"strconv"

	"github.com/grilix/chest-go/chest"
	"github.com/rivo/tview"
)

func ShowHome(app *tview.Application, client *chest.Client) {
	list := tview.NewList().
		// Will probably change to false
		ShowSecondaryText(true)
	pages := tview.NewPages().
		AddPage("home", list, true, true)

	list.
		AddItem("Collection", "See your collection", 'a', func() {
			ShowCollection(app, client, func() {
				app.SetRoot(list, true).SetFocus(pages)
			})
		}).
		AddItem("Decks", "See your decks", 'b', func() {
			ShowDecks(app, client, func() {
				app.SetRoot(list, true).SetFocus(pages)
			})
		}).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})

	app.SetRoot(list, true).SetFocus(pages)
}
