package nrdb

import (
	"fmt"
	"net/url"
	"time"
)

type CardSet struct {
	Document[CardSetAttributes, CardSetRelationships]
}

type CardSetAttributes struct {
	Name            string    `json:"name"`
	CardCycleID     string    `json:"card_cycle_id"`
	CardSetTypeID   string    `json:"card_set_type_id"`
	DateRelease     string    `json:"date_release"`
	FirstPrintingID string    `json:"first_printing_id"`
	LegacyCode      string    `json:"legacy_code"`
	ReleasedBy      string    `json:"released_by"`
	Size            int       `json:"size"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type CardSetRelationships struct {
	CardCycle   *Relationship `json:"card_cycle"`
	CardSetType *Relationship `json:"card_set_type"`
	Cards       *Relationship `json:"cards"`
	Printings   *Relationship `json:"printings"`
}

type CardSetFilter struct {
	CardCycleID   string
	CardSetTypeID string
}

func (filter CardSetFilter) Encode() (url.Values, error) {
	query := url.Values{}
	if filter.CardCycleID != "" {
		query.Set("filter[card_cycle_id]", filter.CardCycleID)
	}

	if filter.CardSetTypeID != "" {
		query.Set("filter[card_set_type_id]", filter.CardSetTypeID)
	}

	return query, nil

}

func (doc CardSet) String() string {
	return fmt.Sprintf("%s (%s)", doc.Attributes.Name, doc.ID)
}

func (cl client) CardSets(filter *CardSetFilter) ([]*CardSet, error) {
	var res Response[CardSet]

	var query url.Values
	if filter != nil {
		q, err := filter.Encode()
		if err != nil {
			return nil, fmt.Errorf("encoding filter: %w", err)
		}
		query = q
	}

	if err := cl.nrdbReq("card_sets", &res, query); err != nil {
		return nil, err
	}

	return res.Data, nil
}
