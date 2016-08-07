package controllers

import "github.com/revel/revel"


type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	greeting := "Aloha from Sunday App"
	return c.Render(greeting)
}
