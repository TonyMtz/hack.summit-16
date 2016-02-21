package controllers

import (
	"fmt"
	"log"
	"io/ioutil"

	"github.com/mrjones/oauth"
	"github.com/revel/revel"
)

type Trello struct {
	*revel.Controller
}

var (
	tokens map[string]*oauth.RequestToken
	consumer      *oauth.Consumer
)

func init() {
	tokens = make(map[string]*oauth.RequestToken)

	consumer = oauth.NewConsumer(
		"afb6671d5446eb923f98a0111aa8227d",
		"4508a1f0f51d4e77ec3f32f87bfdd3b63048fffa659040952f012a9e02986ad5",
		oauth.ServiceProvider{
			RequestTokenUrl:   "https://trello.com/1/OAuthGetRequestToken",
			AuthorizeTokenUrl: "https://trello.com/1/OAuthAuthorizeToken",
			AccessTokenUrl:    "https://trello.com/1/OAuthGetAccessToken",
		},
	)
	// App Name
	consumer.AdditionalAuthorizationUrlParams["name"] = "Trello OAuth"
	// Token Expiration - Default 30 days
	consumer.AdditionalAuthorizationUrlParams["expiration"] = "never"
	// Authorization Scope
	consumer.AdditionalAuthorizationUrlParams["scope"] = "read"
	consumer.Debug(true)
}

func (c Trello) RedirectUserToTrello() revel.Result {
	tokenUrl := fmt.Sprintf("http://%s/maketoken", c.Request.Host)
	token, requestUrl, err := consumer.GetRequestTokenAndUrl(tokenUrl)
	if err != nil {
		log.Fatal(err)
	}
	tokens[token.Token] = token
	return c.Redirect(requestUrl)
}

func (c Trello) GetTrelloToken() revel.Result {
	values := c.Params.Query
	verificationCode := values.Get("oauth_verifier")
	tokenKey := values.Get("oauth_token")

	accessToken, err := consumer.AuthorizeToken(tokens[tokenKey], verificationCode)
	if err != nil {
		log.Fatal(err)
	}

	client, err := consumer.MakeHttpClient(accessToken)
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Get("https://trello.com/1/members/me")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	return c.RenderText(string(content))
}
