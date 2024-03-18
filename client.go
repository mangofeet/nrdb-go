package nrdb

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

type Client interface {
	WithHTTPClient(http.Client) Client

	CardCycles() ([]*CardCycle, error)
	CardPools() ([]*CardPool, error)
	CardSetTypes() ([]*CardSetType, error)
	CardSets(*CardSetFilter) ([]*CardSet, error)
	CardSubtypes() ([]*CardSubtype, error)
	CardTypes(*CardTypeFilter) ([]*CardType, error)
	Cards(*CardFilter) ([]*Card, error)
	AllCards(*CardFilter) ([]*Card, error)
	Card(cardID string) (*Card, error)
	Factions(*FactionFilter) ([]*Faction, error)
	Faction(factionID string) (*Faction, error)
	Formats() ([]*Format, error)
	Format(formatID string) (*Format, error)
	Illustrators() ([]*Illustrator, error)
	Illustrator(illustratorID string) (*Illustrator, error)
	Printings(*PrintingFilter) ([]*Printing, error)
	AllPrintings(*PrintingFilter) ([]*Printing, error)
	Printing(printingID string) (*Printing, error)
}

type Filter interface {
	Query() (url.Values, error)
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

func (cl client) WithHTTPClient(httpClient http.Client) Client {
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

	log.Println(reqURL.String())

	return cl.doNRDBReq(reqURL.String(), out)
}

func (cl client) doNRDBReq(reqURL string, out any) error {
	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
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
