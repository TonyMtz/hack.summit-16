package services

import (
	"github.com/revel/revel"
)

type Provider interface {
	RedirectUrl() string
	Callback(params revel.Params) string
}

var (
	providers map[string]Provider
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

func Callback(provider string, params *revel.Params) string {
	if p, ok := providers[provider]; ok {
		return p.Callback(*params)
	}
	return "Unknown provider"
}
