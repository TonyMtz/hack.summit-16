package services

import (
	"log"
	"github.com/mrjones/oauth"
	"github.com/revel/revel"
	"fmt"
	"github.com/TonyMtz/hack.summit-16.service/app/models"
	"encoding/json"
)

type Trello struct {
	tokens   map[string]*oauth.RequestToken
	consumer *oauth.Consumer
}

func NewTrello(key string, secret string) Trello {
	t := Trello{tokens:make(map[string]*oauth.RequestToken), consumer:oauth.NewConsumer(
		key,
		secret,
		oauth.ServiceProvider{
			RequestTokenUrl:   "https://trello.com/1/OAuthGetRequestToken",
			AuthorizeTokenUrl: "https://trello.com/1/OAuthAuthorizeToken",
			AccessTokenUrl:    "https://trello.com/1/OAuthGetAccessToken",
		},
	)}
	t.consumer.AdditionalAuthorizationUrlParams["name"] = "Trello OAuth"
	// Token Expiration - Default 30 days
	t.consumer.AdditionalAuthorizationUrlParams["expiration"] = "never"
	// Authorization Scope
	t.consumer.AdditionalAuthorizationUrlParams["scope"] = "read"
	t.consumer.Debug(revel.DevMode)
	return t
}

func (t Trello) RedirectUrl(xtoken string) string {
	callbackUrl := fmt.Sprintf("http://%s:%v/trello/callback", revel.HttpAddr, revel.HttpPort)
	if xtoken != "" {
		callbackUrl = callbackUrl + "?xtoken=" + xtoken
	}
	token, url, err := t.consumer.GetRequestTokenAndUrl(callbackUrl)
	if err != nil {
		log.Fatal(err)
	}
	t.tokens[token.Token] = token
	return url
}

func (t Trello) Callback(params revel.Params) interface{} {
	values := params.Query
	verificationCode := values.Get("oauth_verifier")
	tokenKey := values.Get("oauth_token")

	accessToken, err := t.consumer.AuthorizeToken(t.tokens[tokenKey], verificationCode)
	if err != nil {
		log.Fatal(err) //TODO throw?
	}
	return *accessToken
}

type TrelloCard struct {
	Id   string        `json:id`
	Name string        `json:name`
	Desc string        `json:desc`
}

func (t Trello) Cards(token interface{}) []models.Card {
	tkn, ok := token.(oauth.AccessToken)
	if !ok {
		log.Fatal("Castig error")
	}
	client, err := t.consumer.MakeHttpClient(&tkn)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Get("https://api.trello.com/1/members/me/cards?filter=visible&fields=name,desc")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var cards []models.Card

	decoder := json.NewDecoder(resp.Body)

	/*//TODO
	content, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(content))
	return nil
	//TODO*/

	var tr []TrelloCard
	if err := decoder.Decode(&tr); err != nil {
		log.Fatal(err)
	}

	for _, tc := range tr {
		if tc.Desc != "" {
			cards = append(cards, models.Card{Id:tc.Id, Title:tc.Name, Desc:tc.Desc, Provider:"trello" })
		}
	}

	return cards
}
