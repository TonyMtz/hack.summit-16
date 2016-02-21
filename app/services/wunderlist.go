package services

import (
	"log"
	"fmt"
	"github.com/revel/revel"
	"golang.org/x/oauth2"
)

func init() {
	oauth2.RegisterBrokenAuthHeaderProvider("https://www.wunderlist.com/oauth/")
}

type Wunderlist struct {
	config *oauth2.Config
}

func NewWunderlist(key string, secret string) Wunderlist {
	w := Wunderlist{config:&oauth2.Config{
		ClientID:    key,
		ClientSecret: secret,
		Scopes:       []string{},
		Endpoint:     oauth2.Endpoint{
			AuthURL:"https://www.wunderlist.com/oauth/authorize",
			TokenURL:"https://www.wunderlist.com/oauth/access_token",
		},
	}}
	return w
}

func (w Wunderlist) RedirectUrl() string {
	callbackUrl := fmt.Sprintf("http://%s:%v/wunderlist/callback", revel.HttpAddr, revel.HttpPort) //TODO common code
	w.config.RedirectURL = callbackUrl
	return w.config.AuthCodeURL("RANDOM")
}

func (w Wunderlist) Callback(params revel.Params) interface {}{
	values := params.Query
	verificationCode := values.Get("code")

	token, err := w.config.Exchange(oauth2.NoContext, verificationCode)
	if err != nil {
		log.Fatal(err) //TODO throw?
		return "Error!"
	}
	return token
}

/*func (c Wunderlist) connected() *models.User {
	return c.RenderArgs["user"].(*models.User)
}*/
