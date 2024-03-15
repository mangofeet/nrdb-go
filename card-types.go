package nrdb

import (
	"fmt"
	"net/url"
	"time"
)

type CardType struct {
	Document[CardTypeAttributes, CardTypeRelationships]
}

type CardTypeAttributes struct {
	Name      string    `json:"name"`
	SideID    string    `json:"side_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CardTypeRelationships struct {
	Cards *Relationship `json:"cards"`
	Side  *Relationship `json:"side"`
}

type CardTypeFilter struct {
	SideID string
}

func (filter CardTypeFilter) Encode() (url.Values, error) {
	query := url.Values{}
	if filter.SideID != "" {
		query.Set("filter[side_id]", filter.SideID)
	}

	return query, nil
}

func (doc CardType) String() string {
	return fmt.Sprintf("%s - %s (%s)", doc.Attributes.Name, doc.Attributes.SideID, doc.ID)
}

func (cl client) CardTypes(filter *CardTypeFilter) ([]*CardType, error) {
	var res Response[CardType]

	var query url.Values
	if filter != nil {
		q, err := filter.Encode()
		if err != nil {
			return nil, fmt.Errorf("encoding filter: %w", err)
		}
		query = q
	}

	if err := cl.nrdbReq("card_types", &res, query); err != nil {
		return nil, err
	}

	return res.Data, nil
}
