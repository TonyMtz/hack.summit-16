package services

import (
	"log"
	"fmt"
	"github.com/revel/revel"
	"golang.org/x/oauth2"
	"github.com/TonyMtz/hack.summit-16.service/app/models"
	"encoding/json"
	"io/ioutil"
)

func init() {
	oauth2.RegisterBrokenAuthHeaderProvider("https://www.wunderlist.com/oauth/")
}

type Wunderlist struct {
	config *oauth2.Config
}

func NewWunderlist(key string, secret string) Wunderlist {
	w := Wunderlist{config:&oauth2.Config{
		ClientID:     key,
		ClientSecret: secret,
		Scopes:       []string{},
		Endpoint:     oauth2.Endpoint{
			AuthURL:  "https://www.wunderlist.com/oauth/authorize",
			TokenURL: "https://www.wunderlist.com/oauth/access_token",
		},
	}}
	return w
}

func (w Wunderlist) RedirectUrl(xtoken string) string {
	callbackUrl := fmt.Sprintf("http://%s:%v/wunderlist/callback", revel.HttpAddr, revel.HttpPort) //TODO common code
	if xtoken != "" {
		callbackUrl = callbackUrl + "?xtoken=" + xtoken
	}
	w.config.RedirectURL = callbackUrl
	return w.config.AuthCodeURL("RANDOM")
}

func (w Wunderlist) Callback(params revel.Params) interface{} {
	values := params.Query
	verificationCode := values.Get("code")

	token, err := w.config.Exchange(oauth2.NoContext, verificationCode)
	if err != nil {
		log.Fatal(err) //TODO throw?
		return "Error!"
	}
	return *token
}

type WunderlistCard struct {
	Id        int        `json:id`
	TaskID    int        `json:task_id`
	Content   string     `json:content`
	CreatedAt string     `json:created_at`
	UpdatedAt string     `json:updated_at`
	Revision  string     `json:revision`
}

func (w Wunderlist) Cards(token interface{}) []models.Card {
	tkn, ok := token.(oauth2.Token)
	if !ok {
		log.Fatal("Castig error")
	}
	client := w.config.Client(oauth2.NoContext, &tkn)
	resp, err := client.Get("http://a.wunderlist.com/api/v1/notes")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var cards []models.Card

	decoder := json.NewDecoder(resp.Body)

	//TODO comment
	content, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(content))
	return nil
	//TODO comment

	var wcs []WunderlistCard
	if err := decoder.Decode(&wcs); err != nil {
		log.Fatal(err)
	}

	for _, wc := range wcs {
		if wc.Content != "" {
			cards = append(cards, models.Card{Id:string(wc.Id), Desc:wc.Content, Provider:"wunderlist" })
		}
	}

	return cards
}