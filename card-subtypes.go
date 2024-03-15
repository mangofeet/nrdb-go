package nrdb

import (
	"fmt"
	"time"
)

func (cl client) CardSubtypes() ([]*CardSubtype, error) {
	var res Response[[]*CardSubtype]

	if err := cl.nrdbReq("card_subtypes", &res, nil); err != nil {
		return nil, err
	}

	return res.Data, nil
}

type CardSubtype struct {
	Document[CardSubtypeAttributes, CardSubtypeRelationships]
}

type CardSubtypeAttributes struct {
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CardSubtypeRelationships struct {
	Cards     *Relationship `json:"cards"`
	Printings *Relationship `json:"printings"`
}

func (doc CardSubtype) String() string {
	return fmt.Sprintf("%s (%s)", doc.Attributes.Name, doc.ID)
}

func (doc CardSubtype) Cards() ([]*Card, error) {
	return nil, nil
}
