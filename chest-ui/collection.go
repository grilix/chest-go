package chest_ui

import (
	"strconv"

	"github.com/gdamore/tcell"
	"github.com/grilix/chest-go/chest"
	"github.com/rivo/tview"
)

func ListCards(table *tview.Table, cards []*chest.Card) {
	color := tcell.ColorWhite

	table.SetCell(0, 0,
		tview.NewTableCell("Edition").
			SetTextColor(color).
			SetAlign(tview.AlignCenter))

	table.SetCell(0, 1,
		tview.NewTableCell("Edition").
			SetTextColor(color).
			SetAlign(tview.AlignCenter))

	table.SetCell(0, 2,
		tview.NewTableCell("Name").
			SetTextColor(color).
			SetAlign(tview.AlignCenter))

	for index, card := range cards {
		table.SetCell(index+1, 0,
			tview.NewTableCell(card.EditionCode).
				SetTextColor(color).
				SetAlign(tview.AlignCenter))

		table.SetCell(index+1, 1,
			tview.NewTableCell(strconv.Itoa(card.Count)).
				SetTextColor(color).
				SetAlign(tview.AlignCenter))

		table.SetCell(index+1, 2,
			tview.NewTableCell(card.Name).
				SetTextColor(color).
				SetAlign(tview.AlignCenter))
	}
}

func ShowCollection(
	app *tview.Application,
	client *chest.Client,
	done func(),
) {
	table := tview.NewTable().
		SetBorders(false)

	params := chest.TableParams{}
	collection, _, err := client.Collection.UserCollection(params)
	if err != nil {
		panic(err)
	}

	ListCards(table, collection.Cards)

	table.Select(0, 0).SetFixed(1, 0).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			done()
		}
	}).SetSelectedFunc(func(row int, column int) {
		// Navigation is done by pressing [ENTER], yup.
		// I have no idea how to use another keybinding,
		// will check again later.
		// Or not.
		table.Clear()
		params.Page++
		collection, _, err := client.Collection.UserCollection(params)

		if err != nil {
			panic(err)
		}

		ListCards(table, collection.Cards)
	})
	table.SetSelectable(true, false)

	pages := tview.NewPages().
		AddPage("test", table, true, true)

	app.SetRoot(table, true).SetFocus(pages)
}
