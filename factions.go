package nrdb

import (
	"fmt"
	"net/url"
	"time"
)

func (cl client) Faction(factionID string) (*Faction, error) {
	var res Response[Faction]

	if err := cl.nrdbReq(fmt.Sprintf("factions/%s", factionID), &res, nil); err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (cl client) Factions(filter *FactionFilter) ([]*Faction, error) {
	var res Response[[]*Faction]

	var query url.Values
	if filter != nil {
		q, err := filter.Query()
		if err != nil {
			return nil, fmt.Errorf("encoding filter: %w", err)
		}
		query = q
	}

	if err := cl.nrdbReq("factions", &res, query); err != nil {
		return nil, err
	}

	return res.Data, nil
}

type Faction struct {
	Document[FactionAttributes, FactionRelationships]
}

func (doc Faction) Name() string {
	if doc.Attributes == nil {
		return ""
	}
	return doc.Attributes.Name
}

func (doc Faction) Description() string {
	if doc.Attributes == nil {
		return ""
	}
	return doc.Attributes.Description
}

func (doc Faction) IsMini() bool {
	if doc.Attributes == nil {
		return false
	}
	return doc.Attributes.IsMini
}

func (doc Faction) SideID() string {
	if doc.Attributes == nil {
		return ""
	}
	return doc.Attributes.SideID
}

type FactionAttributes struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsMini      bool      `json:"is_mini"`
	SideID      string    `json:"side_id"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type FactionRelationships struct {
	Side      *Relationship `json:"side"`
	Cards     *Relationship `json:"cards"`
	Printings *Relationship `json:"printings"`
}

type FactionFilter struct {
	SideID *string
	IsMini *bool
}

func (filter FactionFilter) Query() (url.Values, error) {
	query := url.Values{}
	if filter.SideID != nil {
		query.Set("filter[side_id]", *filter.SideID)
	}

	if filter.IsMini != nil {
		query.Set("filter[is_mini]", fmt.Sprintf("%t", *filter.IsMini))
	}

	return query, nil

}

func (doc Faction) String() string {
	return fmt.Sprintf("%s (%s)", doc.Attributes.Name, doc.ID)
}
