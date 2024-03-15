package nrdb

import (
	"fmt"
	"net/url"
	"time"
)

type Card struct {
	Document[CardPoolAttributes, CardPoolRelationships]
}

type CardAttributes struct {
	Name         string    `json:"name"`
	CardCycleIDs []string  `json:"card_cycles_ids"`
	CardSetIDs   []string  `json:"card_set_ids"`
	CardIDs      []string  `json:"card_ids"`
	NumCards     int       `json:"num_cards"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CardRelationships struct {
	CardCycles *Relationship `json:"card_cycles"`
	CardSets   *Relationship `json:"card_sets"`
	Cards      *Relationship `json:"cards"`
	Format     *Relationship `json:"format"`
	Snapshots  *Relationship `json:"snapshots"`
}

type CardFilter struct {
}

func (filter CardFilter) Encode() (url.Values, error) {
	query := url.Values{}
	return query, nil
}

func (doc Card) String() string {
	return fmt.Sprintf("%s (%s)", doc.Attributes.Name, doc.ID)
}

func (cl client) Cards(filter CardFilter) ([]*Card, error) {
	var res Response[Card]

	if err := cl.nrdbReq("cards", &res, nil); err != nil {
		return nil, err
	}

	return res.Data, nil
}
