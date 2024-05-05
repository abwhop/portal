package gql

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/abwhop/portal_sync"
	"io"
	"net/http"
	"time"
)

type Gql struct {
	config     *portal_sync.PortalConfig
	httpClient *http.Client
}

type authedTransport struct {
	username string
	password string
	wrapped  http.RoundTripper
}

func (t *authedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(t.username, t.password)
	return t.wrapped.RoundTrip(req)
}

func NewGql(config *portal_sync.PortalConfig) *Gql {
	return &Gql{
		config: config,
		httpClient: &http.Client{
			Transport: &authedTransport{
				username: config.User,
				password: config.Password,
				wrapped:  http.DefaultTransport,
			},
		},
	}
}

func (g *Gql) Query(ctx context.Context, query string, model interface{}) error {
	//fmt.Println(query)
	req, err := http.NewRequest("POST", g.config.Server+`/graphql/`, bytes.NewBuffer([]byte(query)))
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(req.Context(), time.Duration(g.config.Timeout)*time.Millisecond)
	defer cancel()
	req = req.WithContext(ctx)
	resp, err := g.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var errorRespondGQL *ErrorRespondGQL
	var errorRespond *ErrorRespond

	b, err := io.ReadAll(resp.Body)
	fmt.Println(string(b))
	if err := json.Unmarshal(b, &errorRespond); err == nil && errorRespond != nil {
		fmt.Println("1")
		return errors.New(errorRespond.Error.Message)
	} else if err := json.Unmarshal(b, &errorRespondGQL); err == nil && len(errorRespondGQL.Errors) > 0 {
		fmt.Println("2")
		return errors.New(errorRespondGQL.Errors[0].Message)
	} else if err := json.Unmarshal(b, &model); err != nil {
		return err
	}
	return nil
}
