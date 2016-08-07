package controllers

import "github.com/revel/revel"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	greeting := "Aloha from Sunday App"
	return c.Render(greeting)
}

func (c App) Hello(myName string) revel.Result {
	c.Validation.Required(myName).Message("sorry but your name is required.")
	c.Validation.MinSize(myName, 3).Message("c'mon your name is longer than at least 2 letters...")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}

	return c.Render(myName)
}
