package chest_ui

import (
	//"strconv"

	"github.com/gdamore/tcell"
	"github.com/grilix/chest-go/chest"
	"github.com/rivo/tview"
)

func listDecks(table *tview.Table, decks []*chest.Deck) {
	color := tcell.ColorWhite

	table.SetCell(0, 0,
		tview.NewTableCell("Name").
			SetTextColor(color).
			SetAlign(tview.AlignCenter))

	for index, deck := range decks {
		table.SetCell(index+1, 0,
			tview.NewTableCell(deck.Name).
				SetTextColor(color).
				SetAlign(tview.AlignCenter))
	}
}

func showDeckCards(
	app *tview.Application,
	deck *chest.Deck,
	done func(),
) {
	table := tview.NewTable().
		SetBorders(false)

	ListCards(table, deck.Cards)

	table.Select(0, 0).SetFixed(1, 0).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			done()
		}
	}).SetSelectedFunc(func(row int, column int) {
		// TODO: ?
	})
	table.SetSelectable(true, false)

	pages := tview.NewPages().
		AddPage("test", table, true, true)

	app.SetRoot(table, true).SetFocus(pages)
}

func ShowDecks(
	app *tview.Application,
	client *chest.Client,
	done func(),
) {
	table := tview.NewTable().
		SetBorders(false)

	decks, _, err := client.Decks.UserDecks()
	if err != nil {
		panic(err)
	}

	listDecks(table, decks.Decks)
	pages := tview.NewPages().
		AddPage("test", table, true, true)

	table.Select(0, 0).SetFixed(1, 0).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			done()
		}
	}).SetSelectedFunc(func(row int, column int) {
		deck, _, err := client.Decks.UserDeck(decks.Decks[row-1].Id)
		if err != nil {
			panic(err)
		}

		showDeckCards(app, deck, func() {
			app.SetRoot(table, true).SetFocus(pages)
		})
	})
	table.SetSelectable(true, false)

	app.SetRoot(table, true).SetFocus(pages)
}
