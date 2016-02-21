package controllers

import (
//"fmt"

	"github.com/revel/revel"
	"github.com/revel/samples/booking/app/models"

	"golang.org/x/oauth2"
)

type Wunderlist struct {
	*revel.Controller
}

type Url struct {
	url string
}

var WUNDERLIST = &oauth2.Config{
	ClientID:     "ec5ebac038a76c6ec3c4",
	ClientSecret: "4b55f1910a7299c2f26aec2e76488af8cdfea044f20ed92eaa9859a9be28",
	Scopes:       []string{},
	Endpoint:     oauth2.Endpoint{
		AuthURL:"https://www.wunderlist.com/oauth/authorize",
		TokenURL:"https://www.wunderlist.com/oauth/access_token",
	},
	RedirectURL:  "http://todoist.dev:9000/wunderlist/index/",
}

func (c Wunderlist) Auth() revel.Result {
	authUrl := WUNDERLIST.AuthCodeURL("RANDOM")
	return c.RenderJson(authUrl)
}

func (c Wunderlist) Index() revel.Result {
	values := c.Params.Query
	verificationCode := values.Get("code")
	return c.RenderJson(verificationCode)
}

func (c Wunderlist) connected() *models.User {
	return c.RenderArgs["user"].(*models.User)
}
