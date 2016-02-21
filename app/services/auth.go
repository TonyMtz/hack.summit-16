package services

import (
	"github.com/revel/revel"
	//"log"
)

type Provider interface {
	RedirectUrl() string
	Callback(params revel.Params) string
}

var (
	providers map[string]Provider
)

func init() {
	//log.Print("trello" +  revel.Config.StringDefault("trello.key", "default"))
	providers = make(map[string]Provider)
	providers["trello"] = NewTrello("afb6671d5446eb923f98a0111aa8227d", "4508a1f0f51d4e77ec3f32f87bfdd3b63048fffa659040952f012a9e02986ad5") //TODO
	providers["wunderlist"] = NewWunderlist("3db038e285c0dfc22875", "ab16496c9830bcf716fd7a6770dc745e7a9d70a8829c0801a6d27258b35e") //TODO
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