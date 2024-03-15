package nrdb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type Client interface {
	WithHTTPCLient(http.Client) Client

	CardCycles() ([]*CardCycle, error)
	CardPools() ([]*CardPool, error)
	CardSetTypes() ([]*CardSetType, error)
	CardSets(*CardSetFilter) ([]*CardSet, error)
	CardSubtypes() ([]*CardSubtype, error)
	CardTypes(*CardTypeFilter) ([]*CardType, error)
	Cards(*CardFilter) ([]*Card, error)
}

type Filter interface {
	Encode() (url.Values, error)
}

type client struct {
	http http.Client
}

var defaultHTTPClient = http.Client{Timeout: time.Second * 15}

func NewClient() Client {
	return client{
		http: defaultHTTPClient,
	}
}

func (cl client) WithHTTPCLient(httpClient http.Client) Client {
	cl.http = httpClient
	return cl
}

func (cl client) nrdbReq(path string, out any, query url.Values) error {

	reqURL := url.URL{
		Scheme:   "https",
		Host:     "api-preview.netrunnerdb.com",
		Path:     fmt.Sprintf("/api/v3/public/%s", path),
		RawQuery: query.Encode(),
	}

	req, err := http.NewRequest(http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return fmt.Errorf("making request: %w", err)
	}

	res, err := cl.http.Do(req)
	if err != nil {
		return fmt.Errorf("making request: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected %d status", res.StatusCode)
	}

	if err := json.NewDecoder(res.Body).Decode(out); err != nil {
		return fmt.Errorf("parsing payload: %w", err)
	}

	return nil

}
