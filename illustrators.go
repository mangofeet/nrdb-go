package nrdb

import (
	"fmt"
	"time"
)

func (cl client) Illustrator(illustratorID string) (*Illustrator, error) {
	var res Response[Illustrator]

	if err := cl.nrdbReq(fmt.Sprintf("illustrators/%s", illustratorID), &res, nil); err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (cl client) Illustrators() ([]*Illustrator, error) {
	var res Response[[]*Illustrator]

	if err := cl.nrdbReq("illustrators", &res, nil); err != nil {
		return nil, err
	}

	return res.Data, nil
}

type Illustrator struct {
	Document[IllustratorAttributes, IllustratorRelationships]
}

func (doc Illustrator) String() string {
	return fmt.Sprintf("%s (%s)", doc.Attributes.Name, doc.ID)
}

func (doc Illustrator) Name() string {
	if doc.Attributes == nil {
		return ""
	}
	return doc.Attributes.Name
}

func (doc Illustrator) NumPrintings() int {
	if doc.Attributes == nil {
		return 0
	}
	return doc.Attributes.NumPrintings
}

type IllustratorAttributes struct {
	Name         string    `json:"name"`
	NumPrintings int       `json:"num_printings"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type IllustratorRelationships struct {
	Printings *Relationship `json:"printings"`
}
