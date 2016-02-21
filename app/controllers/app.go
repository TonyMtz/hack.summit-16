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
	//xtoken := c.Request.Header.Get("xtoken")
	//a := services.Auth{}
	//a.Some()

	//tokenUrl := fmt.Sprintf("http://%s/maketoken", c.Request.Host)

	return c.RenderText(services.Auth(provider))
}

func (c App) Callback(provider string) revel.Result {
	xtoken := c.Params.Get("xtoken")
	return c.RenderText(services.Callback(provider, xtoken, c.Params))
}