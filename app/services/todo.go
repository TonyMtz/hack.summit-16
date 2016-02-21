package services

import (
	"github.com/revel/revel"
	"github.com/TonyMtz/hack.summit-16.service/app/models"
	"log"
	"github.com/TonyMtz/hack.summit-16.service/app/utils"
)

type Provider interface {
	RedirectUrl() string
	Callback(params revel.Params) interface{} //TODO change params for map[string] string
}

var (
	providers map[string]Provider
	users map[string]models.User
)

func initProviders() {
	providers = make(map[string]Provider)
	providers["trello"] = NewTrello(
		revel.Config.StringDefault("trello.key", "empty"),
		revel.Config.StringDefault("trello.secret", "empty"),
	)
	providers["wunderlist"] = NewWunderlist(
		revel.Config.StringDefault("wunderlist.key", "empty"),
		revel.Config.StringDefault("wunderlist.key", "empty"),
	)
}

func init() {
	revel.OnAppStart(initProviders)
}

func Auth(provider string) string {
	if p, ok := providers[provider]; ok {
		return p.RedirectUrl()
	}
	return "Unknown provider"
}

func Callback(provider string, xtoken string, params *revel.Params) string {
	if p, ok := providers[provider]; ok {
		token := p.Callback(*params)
		if xtoken == "" {
			uid, err := utils.NewV4();
			if err != nil {
				log.Print(err)
				return "Error generating uuid"
			}
			xtoken = uid.String()
		}
		user, ok := users[xtoken]
		if !ok {
			user = models.User{Uid:xtoken, Token:make(map[string]interface{})}
			users[xtoken] = user
		}
		user.Token[provider] = token
		return xtoken //TODO change
	}
	return "Unknown provider"
}
