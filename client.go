package userstream

import (
	"bufio"
	"github.com/mrjones/oauth"
	"log"
	"net/http"
)

var (
	userStreamsEndPoint = "https://userstream.twitter.com/1.1/user.json"
	requestTokenUrl     = "https://api.twitter.com/oauth/request_token"
	authorizeTokenUrl   = "https://api.twitter.com/oauth/authorize"
	accessTokenUrl      = "https://api.twitter.com/oauth/access_token"
)

type Client struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

func (c *Client) UserStream(callback func(interface{})) {
	c.connect(userStreamsEndPoint, callback)
}

func (c *Client) connect(endPoint string, callback func(interface{})) {
	consumer := oauth.NewConsumer(
		c.ConsumerKey,
		c.ConsumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   requestTokenUrl,
			AuthorizeTokenUrl: authorizeTokenUrl,
			AccessTokenUrl:    accessTokenUrl,
		},
	)

	response, err := consumer.Post(endPoint, nil, c.accessToken())
	if err != nil {
		log.Fatal(err)
	}

	c.readStream(response, callback)
}

func (c *Client) readStream(response *http.Response, callback func(interface{})) {
	reader := bufio.NewReader(response.Body)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			log.Fatal(err)
		}

		if len(line) > 0 && string(line) != "\r\n" {
			object := ParseJson(string(line))
			callback(object)
		}
	}
}

func (c *Client) accessToken() *oauth.AccessToken {
	return &oauth.AccessToken{
		Token:  c.AccessToken,
		Secret: c.AccessTokenSecret,
	}
}
