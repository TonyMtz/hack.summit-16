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
	return c.RenderText(services.Auth(provider, xtoken))
}

func (c App) Callback(provider string) revel.Result {
	xtoken := c.Params.Get("xtoken")
	token := services.Callback(provider, xtoken, c.Params)
	return c.Redirect("/?xtoken=" + token)
}