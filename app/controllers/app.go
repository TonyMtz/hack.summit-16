package controllers

import (
	"github.com/revel/revel"
	"github.com/TonyMtz/hack.summit-16.service/app/services"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Auth(provider string) revel.Result {
	xtoken := c.Request.Header.Get("xtoken")
	return c.RenderJson(services.Auth(provider, xtoken))
}

func (c App) Callback(provider string) revel.Result {
	xtoken := c.Params.Get("xtoken")
	token := services.Callback(provider, xtoken, c.Params)
	return c.RenderJson(token)
}

func (c App) Options(provider string) revel.Result {
	return c.RenderText("GET,POST,PUT,DELETE,OPTIONS") // better than empty
}

func (c App) Cards(xtoken string) revel.Result {
	cards := services.Cards(xtoken)
	return c.RenderJson(cards)
}
