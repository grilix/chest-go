package chest_ui

import (
	"encoding/json"
	"io/ioutil"

	"github.com/grilix/chest-go/chest"
	"github.com/rivo/tview"
)

type Settings struct {
	Host  string `json:"host"`
	Token string `json:"token"`
}

func Alert(app *tview.Application, container *tview.Pages, text string) {
	modal := tview.NewModal().
		SetText(text).
		AddButtons([]string{"Ok"})

	previous := app.GetFocus()

	modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		app.SetFocus(previous)

		container.RemovePage("app-modal")
	})

	container.AddPage("app-modal", modal, true, true)
}

func LoadSettings(client *chest.Client) error {
	file, _ := ioutil.ReadFile(".chest.json")
	var settings Settings
	err := json.Unmarshal(file, &settings)

	if err != nil {
		return err
	}

	client.Token = settings.Token
	client.SetURL(settings.Host)

	return nil
}

func StoreSettings(client *chest.Client) error {
	settings := Settings{Host: client.GetURL(), Token: client.Token}
	b, err := json.Marshal(settings)

	if err != nil {
		return err
	}

	ioutil.WriteFile(".chest.json", b, 0664)

	return nil
}

func Loop(client *chest.Client) {
	app := tview.NewApplication()

	LoadSettings(client)

	if client.IsAuthenticated() {
		ShowCollection(app, client)
	} else {
		Authenticate(app, client, func(_ *chest.User) {
			StoreSettings(client)

			ShowCollection(app, client)
		})
	}

	if err := app.Run(); err != nil {
		panic(err)
	}
}
