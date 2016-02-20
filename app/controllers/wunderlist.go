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
	ClientID:     "943076975742162",
	ClientSecret: "d3229ebe3501771344bb0f2db2324014",
	Scopes:       []string{},
	Endpoint:     oauth2.Endpoint{
		AuthURL:"https://www.wunderlist.com/oauth/authorize",
		TokenURL:"https://www.wunderlist.com/oauth/access_token",
	},
	RedirectURL:  "http://todoist.dev:9000/wunderlist/index",
}

func (c Wunderlist) Auth() revel.Result {
	authUrl := WUNDERLIST.AuthCodeURL("RANDOM")
	return c.RenderJson(authUrl)
}

func (c Wunderlist) connected() *models.User {
	return c.RenderArgs["user"].(*models.User)
}

// https://www.wunderlist.com/oauth/authorize?client_id=ID&redirect_uri=URL&state=RANDOM
// https://www.wunderlist.com/oauth/authorize?client_id=ID&redirect_uri=URL&state=state

// https://www.wunderlist.com/oauth/authorize?client_id=943076975742162&redirect_uri=http://todoist.com/wunderlist/index&state=RANDOM

