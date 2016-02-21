package services

type Provider interface {
	RedirectUrl() string
}

var (
	providers map[string]Provider
)

func init() {
	providers = make(map[string]Provider)
	providers["trello"] = NewTrello("afb6671d5446eb923f98a0111aa8227d", "4508a1f0f51d4e77ec3f32f87bfdd3b63048fffa659040952f012a9e02986ad5") //TODO
}

func Auth(provider string) string {
	if p, ok := providers[provider]; ok {
		return p.RedirectUrl()
	}
	return "Unkonw provider"
}