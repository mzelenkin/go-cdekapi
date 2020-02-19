package cdekapi

import (
	"context"
	"testing"
)

func TestAuth(t *testing.T) {
	c := New("")
	c.auth = authData{
		ClientID: "z9GRRu7FxmO53CQ9cFfI6qiy32wpfTkd",
		Secret:   "w24JTCv4MnAcuRTx0oHjHLDtyt3I6IBq",
	}

	err := c.getToken(context.Background())
	if err != nil {
		t.Error(err)
	}

	if c.token == nil {
		t.Fail()
	}

	//fmt.Printf("Client: %#v\n", c.token)
}
