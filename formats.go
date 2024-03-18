package nrdb

import (
	"fmt"
	"time"
)

func (cl client) Format(formatID string) (*Format, error) {
	var res Response[Format]

	if err := cl.nrdbReq(fmt.Sprintf("formats/%s", formatID), &res, nil); err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (cl client) Formats() ([]*Format, error) {
	var res Response[[]*Format]

	if err := cl.nrdbReq("formats", &res, nil); err != nil {
		return nil, err
	}

	return res.Data, nil
}

type Format struct {
	Document[FormatAttributes, FormatRelationships]
}

func (doc Format) String() string {
	return fmt.Sprintf("%s (%s)", doc.Attributes.Name, doc.ID)
}

func (doc Format) Name() string {
	if doc.Attributes == nil {
		return ""
	}
	return doc.Attributes.Name
}

func (doc Format) ActiveCardPoolID() string {
	if doc.Attributes == nil {
		return ""
	}
	return doc.Attributes.ActiveCardPoolID
}

func (doc Format) ActiveRestrictionID() string {
	if doc.Attributes == nil {
		return ""
	}
	return doc.Attributes.ActiveRestrictionID
}

func (doc Format) ActiveSnapshotID() string {
	if doc.Attributes == nil {
		return ""
	}
	return doc.Attributes.ActiveSnapshotID
}

func (doc Format) RestrictionIDs() []string {
	if doc.Attributes == nil {
		return nil
	}
	return doc.Attributes.RestrictionIDs
}

func (doc Format) SnapshotIDs() []string {
	if doc.Attributes == nil {
		return nil
	}
	return doc.Attributes.SnapshotIDs
}

type FormatAttributes struct {
	Name                string    `json:"name"`
	ActiveCardPoolID    string    `json:"active_card_pool_id"`
	ActiveRestrictionID string    `json:"active_restriction_id"`
	ActiveSnapshotID    string    `json:"active_snapshot_id"`
	RestrictionIDs      []string  `json:"restriction_ids"`
	SnapshotIDs         []string  `json:"snapshot_ids"`
	UpdatedAt           time.Time `json:"updated_at"`
}

type FormatRelationships struct {
	CardPools    *Relationship `json:"card_pools"`
	Restrictions *Relationship `json:"restrictions"`
	Snapshots    *Relationship `json:"snapshots"`
}
