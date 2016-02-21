package services

import (
	"log"
	"github.com/mrjones/oauth"
	"github.com/revel/revel"
	"fmt"
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
	t.consumer.Debug(true)
	return t
}

func (t Trello) RedirectUrl() string {
	callbackUrl := fmt.Sprintf("http://%s:%v/trello/callback", revel.HttpAddr, revel.HttpPort)
	token, url, err := t.consumer.GetRequestTokenAndUrl(callbackUrl)
	if err != nil {
		log.Fatal(err)
	}
	t.tokens[token.Token] = token
	return url
}

/*func (t Trello) GetTrelloToken() revel.Result {
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
}*/
