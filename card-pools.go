package nrdb

import (
	"fmt"
	"time"
)

type CardPool struct {
	Document[CardPoolAttributes, CardPoolRelationships]
}

type CardPoolAttributes struct {
	Name         string    `json:"name"`
	CardCycleIDs []string  `json:"card_cycles_ids"`
	CardSetIDs   []string  `json:"card_set_ids"`
	CardIDs      []string  `json:"card_ids"`
	NumCards     int       `json:"num_cards"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CardPoolRelationships struct {
	CardCycles *Relationship `json:"card_cycles"`
	CardSets   *Relationship `json:"card_sets"`
	Cards      *Relationship `json:"cards"`
	Format     *Relationship `json:"format"`
	Snapshots  *Relationship `json:"snapshots"`
}

func (doc CardPool) String() string {
	return fmt.Sprintf("%s (%s)", doc.Attributes.Name, doc.ID)
}

func (cl client) CardPools() ([]*CardPool, error) {
	var res Response[[]*CardPool]

	if err := cl.nrdbReq("card_pools", &res, nil); err != nil {
		return nil, err
	}

	return res.Data, nil
}
