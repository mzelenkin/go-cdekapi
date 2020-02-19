package cdekapi

import (
	"context"
	"fmt"
	"time"
)

type authSuccess struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Expires     time.Time
	Scope       string `json:"score"`
	JTI         string `json:"jti"`
}

type authError struct {
	Error       string `json:"error"`
	Description string `json:"error_description"`
}

const oAuthTokenUrl = "oauth/token"

func (c *Client) getToken(ctx context.Context) error {

	body := map[string]string{
		"grant_type":    "client_credentials",
		"client_id":     string(c.auth.ClientID),
		"client_secret": string(c.auth.Secret),
	}

	resp, err := c.getClient().R().SetContext(ctx).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(body).
		SetResult(&authSuccess{}).
		SetError(&authError{}).
		Post(oAuthTokenUrl)

	if err != nil {
		return err
	}

	if resp.StatusCode() > 399 {
		if errStruct, ok := resp.Error().(*authError); ok {
			return fmt.Errorf("token get error: %s", errStruct.Description)
		}

		return fmt.Errorf("unknown http error: %s", resp.Status())
	}

	token, ok := resp.Result().(*authSuccess)
	if ok {
		c.token = token
		c.token.Expires = time.Now().Add(time.Duration(token.ExpiresIn) * time.Second)
		return nil
	}

	return fmt.Errorf("unknown http error: %s", resp.Status())
}
