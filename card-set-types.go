package nrdb

import (
	"fmt"
	"time"
)

type CardSetType struct {
	Document[CardSetTypeAttributes, CardSetTypeRelationships]
}

type CardSetTypeAttributes struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CardSetTypeRelationships struct {
	CardSets *Relationship `json:"card_sets"`
}

func (doc CardSetType) String() string {
	return fmt.Sprintf("%s (%s) - %s", doc.Attributes.Name, doc.ID, doc.Attributes.Description)
}

func (cl client) CardSetTypes() ([]*CardSetType, error) {
	var res Response[[]*CardSetType]

	if err := cl.nrdbReq("card_set_types", &res, nil); err != nil {
		return nil, err
	}

	return res.Data, nil
}
