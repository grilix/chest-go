package chest_ui

import (
	"github.com/grilix/chest-go/chest"
	"github.com/rivo/tview"
)

func Authenticate(
	app *tview.Application,
	client *chest.Client,
	done func(*chest.User),
) {
	host := tview.NewInputField().SetLabel("Host").SetFieldWidth(20)
	user := tview.NewInputField().SetLabel("Username").SetFieldWidth(20)
	pass := tview.NewInputField().SetLabel("Password").SetFieldWidth(20).
		SetMaskCharacter('*')

	form := tview.NewForm()
	form.AddFormItem(host)
	form.AddFormItem(user)
	form.AddFormItem(pass)

	pages := tview.NewPages().
		AddPage("login", form, true, true)

	form.
		AddButton("Login", func() {
			client.SetURL(host.GetText())

			credentials := &chest.Credentials{
				Username: user.GetText(),
				Password: pass.GetText(),
			}

			user, _, error := client.Authentication.Authenticate(credentials)
			if error == nil {
				done(user)
			} else {
				Alert(app, pages, "Login failed")
			}
		}).
		AddButton("Quit", func() {
			app.Stop()
		})

	form.SetBorder(true).SetTitle("Login").SetTitleAlign(tview.AlignLeft)
	app.SetRoot(pages, true).SetFocus(form)
}
