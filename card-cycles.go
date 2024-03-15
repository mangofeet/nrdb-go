package nrdb

import (
	"fmt"
	"time"
)

type CardCycle struct {
	Document[CardCycleAttributes, CardCycleRelationships]
}

type CardCycleAttributes struct {
	Name            string    `json:"name"`
	DateRelease     string    `json:"date_release"`
	CardSetIDs      []string  `json:"card_set_ids"`
	FirstPrintingID string    `json:"first_printing_id"`
	Position        int       `json:"position"`
	ReleasedBy      string    `json:"released_by"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type CardCycleRelationships struct {
	CardSets  *Relationship `json:"card_sets"`
	Cards     *Relationship `json:"cards"`
	Printings *Relationship `json:"printings"`
}

func (doc CardCycle) String() string {
	return fmt.Sprintf("%s (%s)", doc.Attributes.Name, doc.ID)
}

func (cl client) CardCycles() ([]*CardCycle, error) {
	var res Response[[]*CardCycle]

	if err := cl.nrdbReq("card_cycles", &res, nil); err != nil {
		return nil, err
	}

	return res.Data, nil
}
