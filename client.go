package cdekapi

import (
	"context"
	"github.com/go-resty/resty/v2"
	"strings"
	"time"
)

const (
	APITestBaseURL   = "http://api.edu.cdek.ru/v2"
	APIBaseURL       = "http://api.cdek.ru/v2"
	APICalculatorURL = "http://api.cdek.ru/calculator/calculate_price_by_json.php"

	apiOfficesList  = "/deliverypoints"
	apiListOfCities = "/location/cities"
)

type authData struct {
	ClientID string
	Secret   string
}

type Client struct {
	baseUrl string
	auth    authData
	token   *authSuccess
}

func New(baseUrl string) *Client {
	if baseUrl == "" {
		baseUrl = APITestBaseURL
	}

	return &Client{baseUrl: baseUrl}
}

func (c *Client) SetAuth(clientId string, secret string) *Client {
	c.auth = authData{
		ClientID: clientId,
		Secret:   secret,
	}
	return c
}

func (c Client) getClient() *resty.Client {
	client := resty.New()
	client.SetHostURL(c.baseUrl)
	client.SetDebug(true)
	client.SetHeader("Accept", "application/json")
	//client.SetHeader("User-Agent", "go-cdeckapi")
	return client
}

func (c Client) getRequest(ctx context.Context) *resty.Request {
	return c.getClient().
		OnBeforeRequest(func(client *resty.Client, req *resty.Request) error {
			if req.URL != APICalculatorURL {
				if c.token == nil || (c.token != nil && time.Now().Sub(c.token.Expires) > 0) {
					if err := c.getToken(req.Context()); err != nil {
						return err
					}
				}

				req.SetHeader("Authorization", strings.Title(c.token.TokenType)+" "+c.token.AccessToken)
			}

			return nil
		}).
		R().SetContext(ctx)
}
